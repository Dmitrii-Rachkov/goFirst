package main

/*
Модули и зависимости

Что такое модули
В Go модулем принято называть любое приложение, которое можно опубликовать, версионировать,
импортировать или скачать. С помощью модулей мы можем управлять зависимостями.

Добавим в пример с Hello, Hexlet пакет для логгирования. В коде это будет выглядеть так:

package main

import "github.com/sirupsen/logrus" // Указываем путь до нужного пакета внутри репозитория

func main() {
    logrus.Println("Hello, Hexlet!")
}

Чтобы код запустился, надо этот пакет установить. Для этого можно использовать команду go get:
GO111MODULE=off go get "github.com/sirupsen/logrus" # О назначении GO111MODULE чуть позже

Далее можно запускать пакет командой go run:
GO111MODULE=off go run .

INFO[0000] Hello, Hexlet!

Зависимость установилась, и код заработал. Обратите внимание, что здесь мы использовали переменную
окружения GO111MODULE=off — именно эта переменная отключает систему модулей.
*/

/*
Как создать модуль
Чтобы превратить папку с кодом в Go-модуль, можно использовать команду go mod init:

go mod init github.com/hexlet/hello-hexlet # Сразу указываем имя модуля

go: creating new go.mod: module github.com/hexlet/hello-hexlet

Команда сгенерировала go.mod файл со следующим содержимым:
module github.com/hexlet/hello-hexlet
go 1.17

Файл начинается с объявления имени модуля (module github.com/hexlet/hello-hexlet).
Это уникальный идентификатор, по которому модуль хранится в индексе модулей Go.
Обычно в качестве имени модуля указывают его адрес в репозитории.

Далее указывается минимальная совместимая версия языка (go 1.17) и список зависимостей.
Пока список зависимостей пуст. Чтобы обновить его, можно воспользоваться командой go get,
убрав переменную GO111MODULE=off. Но есть и другой способ — команда go mod tidy:

go mod tidy

go: finding module for package github.com/sirupsen/logrus
go: downloading github.com/sirupsen/logrus v1.8.1
go: found github.com/sirupsen/logrus in github.com/sirupsen/logrus v1.8.1
go: downloading golang.org/x/sys v0.0.0-20191026070338-33540a1f6037
go: downloading github.com/stretchr/testify v1.2.2
go: downloading github.com/pmezard/go-difflib v1.0.0
go: downloading github.com/davecgh/go-spew v1.1.1

Команда go mod tidy проверяет импорты в коде, загружает недостающие зависимости и удаляет лишние).
Файл go.mod обновился и теперь включает в себя раздел с зависимостями:

module github.com/hexlet/hello-hexlet

go 1.17

require github.com/sirupsen/logrus v1.8.1

require golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect

Также появился файл go.sum:

github.com/sirupsen/logrus v1.8.1 h1:dJKuHgqk1NNQlqoA6BTlM1Wf9DOH3NBjQyu0h9+AZZE=
github.com/sirupsen/logrus v1.8.1/go.mod h1:yWOB1SBYBC5VeMP7gHvWumXLIWorT60ONWic61uBYv0=
...
golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 h1:YyJpGZS1sBuBCzLAR1VEpK193GlqGZbnPFnPV/5Rsb4=
golang.org/x/sys v0.0.0-20191026070338-33540a1f6037/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=

Оба файла обновляются при каждом добавлении или удалении зависимости:

- Файл go.mod включает путь до модуля и его версию
- Файл go.sum добавляет по две записи на каждую зависимость:
	Первая запись с названием модуля, его версией и хэш-суммой
	Вторая запись с хэш-суммой go.mod файла модуля
*/

/*
Как менять версии зависимостей
Модули предоставляют инструменты для работы с разными версиями пакетов. По умолчанию Go добавляет
последнюю доступную версию пакета. В нашем примере это v1.8.1. Чтобы проверить, какие еще версии
пакета доступны, используем команду go list:

go list -m --versions github.com/sirupsen/logrus

github.com/sirupsen/logrus v0.1.0 v0.1.1 v0.2.0 ... v1.7.0 v1.7.1 v1.8.0 v1.8.1

По умолчанию команда go list выдает адрес текущего пакета, по которому его можно импортировать.
В примере выше мы использовали два флага:

-m указывает, что нас интересует только модуль, а не его пакеты
--versions указывает все возможные для скачивания версии пакета

Чтобы изменить версию, используем команду go get:

go get github.com/sirupsen/logrus@v1.5.0

go: downloading github.com/sirupsen/logrus v1.5.0
go: downloading github.com/konsorten/go-windows-terminal-sequences v1.0.1
go get: added github.com/konsorten/go-windows-terminal-sequences v1.0.1
go get: downgraded github.com/sirupsen/logrus v1.8.1 => v1.5.0
*/

/*
Как удалять зависимости
Удалить зависимость из проекта довольно просто — достаточно удалить импорт этой зависимости из кода
и запустить go mod tidy. В обновленном файле go.mod этого модуля больше не будет.
*/

/*
Выводы
Модуль — это само приложение, в корне которого находится файл go.mod со списком
зависимостей текущего модуля
*/
