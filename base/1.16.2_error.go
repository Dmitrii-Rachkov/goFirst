package main

/*
Обработка ошибок

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)
// При объявлении функции мы указываем, что она возвращает значение типа string
// и может вернуть error - ошибку:
// Если переменная a == 43 то возвращаем 'ok' и пустоту
// Иначе возвращаем пустую строку и новую ошибку с текстом 'some error
func foo(a int) (string, error) {
	if a == 42 {
		return "ok", nil
	}
	return "", errors.New("some error")
}

func main() {
	s, err := foo(42)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(s)
}
*/

/*
Механизм обработки ошибок в Go отличается от обработки исключений в большинстве языков программирования,
ведь в Golang ошибки исключениями не являются. Если говорить в целом, то ошибка в Go — это возвращаемое
значение с типомerror, которое демонстрирует сбой. А с точки зрения кода — интерфейс.
В качестве ошибки может выступать любой объект, который этому интерфейсу удовлетворяет.

Выглядит это так:
type error interface {
    Error() string
}

Как обрабатывать ошибки в Go?
Чтобы обработать ошибку в Golang, необходимо сперва вернуть из функции переменную с объявленным
типом error и проверить её на nil:
if err != nil {
  return err
}

Если метод возвращает ошибку, значит, потенциально в его работе может возникнуть проблема,
которую нужно обработать. В качестве реализации обработчика может выступать логирование
ошибки или более сложные сценарии. Например, переоткрытие установленного сетевого соединения,
повторный вызов метода и тому подобные операции.

Если метод возвращает разные типы ошибок, то их нужно проверять отдельно.
То есть сначала происходит определение ошибки, а потом для каждого типа пишется свой обработчик.

В Go ошибки возвращаются и проверяются явно. Разработчик сам определяет, какие ошибки метод может вернуть,
и реализовать их обработку на вызывающей стороне.
*/

/*
Создание ошибок
Перед тем как обработать ошибку, нужно её создать. В стандартной библиотеке для этого есть две
встроенные функции — обе позволяют указывать и отображать сообщение об ошибке:
errors.New
fmt.Errorf

Метод errors.New() создаёт ошибку, принимая в качестве параметра текстовое сообщение.
package main

import (
  "errors"
  "fmt"
)

func main() {
  err := errors.New("emit macho dwarf: elf header corrupted")
  fmt.Print(err)
}

С помощью метода fmt.Errorf можно добавить дополнительную информацию об ошибке.
Данные будут храниться внутри одной конкретной строки.
package main

import (
  "fmt"
)

func main() {
  const name, id = "bueller", 17
  err := fmt.Errorf("user %q (id %d) not found", name, id)
  fmt.Print(err)
}

Такой способ подходит, если эта дополнительная информация нужна только для логирования на вызывающей
стороне. Если же с ней предстоит работать, можно воспользоваться другими механизмами.
*/

/*
Оборачивание ошибок
Поскольку Error — это интерфейс, можно создать удовлетворяющую ему структуру с собственными полями.
Тогда на вызывающей стороне этими самыми полями можно будет оперировать.
package main

import (
  "fmt"
)

type NotFoundError struct {
  UserId int
}

func (err NotFoundError) Error() string {
  return fmt.Sprintf("user with id %d not found", err.UserId)
}

func SearchUser(id int) error {
  // some logic for search
  // ...
  // if not found
  var err NotFoundError
  err.UserId = id
  return err
}

func main() {
  const id = 17
  err := SearchUser(id)
  if err != nil {
     fmt.Println(err)
     //type error checking
     notFoundErr, ok := err.(NotFoundError)
     if ok {
        fmt.Println(notFoundErr.UserId)
     }
  }
}
*/

/*
Представим другую ситуацию. У нас есть метод, который вызывает внутри себя ещё один метод.
В каждом из них проверяется своя ошибка. Иногда требуется в метод верхнего уровня передать
сразу обе эти ошибки.

В Go есть соглашение о том, что ошибка, которая содержит внутри себя другую ошибку, может реализовать
Unwrap, который будет возвращать исходную ошибку.

Также для оборачивания ошибок в fmt.Errorf есть плейсхолдер %w, который и позволяет произвести такую
упаковку.:
package main

import (
  "errors"
  "fmt"
  "os"
)

func main() {
  err := openFile("non-existing")
  if err != nil {
    fmt.Println(err.Error())
    // get internal error
    fmt.Println(errors.Unwrap(err))
  }
}

func openFile(filename string) error {
  if _, err := os.Open(filename); err != nil {
    return fmt.Errorf("error opening %s: %w", filename, err)
  }
  return nil
}
*/

/*
Проверка типов с Is и As
В Go 1.13 в пакете Errors появились две функции, которые позволяют определить
тип ошибки — чтобы написать тот или иной обработчик:

errors.Is
errors.As

Метод errors.Is, по сути, сравнивает текущую ошибку с заранее заданным значением ошибки:
package main

import (
  "errors"
  "fmt"
  "io/fs"
  "os"
)

func main() {
  if _, err := os.Open("non-existing"); err != nil {
    if errors.Is(err, fs.ErrNotExist) {
      fmt.Println("file does not exist")
    } else {
      fmt.Println(err)
    }
  }
}
Если это будет та же самая ошибка, то функция вернёт true, если нет — false.

errors.As проверяет, относится ли ошибка к конкретному типу
(раньше надо было явно приводить тип ошибки к тому типу, который хотим проверить):
package main

  import (
  "errors"
  "fmt"
  "io/fs"
  "os"
)

func main() {
  if _, err := os.Open("non-existing"); err != nil {
    var pathError *fs.PathError
    if errors.As(err, &pathError) {
      fmt.Println("Failed at path:", pathError.Path)
    } else {
      fmt.Println(err)
    }
  }
}

Помимо прочего, эти методы удобны тем, что упрощают работу с упакованными ошибками,
позволяя проверить каждую из них за один вызов.
*/

/*
Recover
Эта функция нужна, чтобы вернуть контроль при панике. В таком случае работа приложения не прекращается,
а восстанавливается и продолжается в нормальном режиме.

Recover всегда должна вызываться в функции defer. Чтобы сообщить об ошибке как возвращаемом значении,
вы должны вызвать функцию recover в той же горутине, что и паника, получить структуру ошибки
из функции восстановления и передать её в переменную:

package main

import (
  "errors"
  "fmt"
)

func A() {
  defer fmt.Println("Then we can't save the earth!")
  defer func() {
    if x := recover(); x != nil {
      fmt.Printf("Panic: %+v\n", x)
    }
  }()
  B()
}

func B() {
  defer fmt.Println("And if it keeps getting hotter...")
  C()
}

func C() {
  defer fmt.Println("Turn on the air conditioner...")
  Break()
}

func Break() {
  defer fmt.Println("If it's more than 30 degrees...")
  panic(errors.New("Global Warming!!!"))
}

func main() {
  A()
}
*/
