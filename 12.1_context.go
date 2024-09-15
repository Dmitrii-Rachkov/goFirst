//package main
//
//import (
//	"context"
//	"fmt"
//	"io"
//	"log"
//	"math/rand"
//	"net/http"
//	"time"
//)

/*
Что такое контекст?
Если мы взглянем на документацию к пакету context, то первый абзац будет таким:

Package context defines the Context type, which carries deadlines, cancellation signals,
and other request-scoped values across API boundaries and between processes.

Пакет context определяет тип Context, который позволяет управлять дедлайнами, сигналами отмены и другими
значениями области действия запросов между границами API и процессами.

Что это, чёрт побери, значит?

А значит это примерно следующее:

Контекст — это объект, который предназначен в первую очередь для того, чтобы иметь возможность отменить извне
выполнение потенциально долгой операции. Кроме того, с помощью контекста можно хранить и передавать информацию
между функциями и методами внутри вашей программы.

Отменять долгие операции с помощью контекста можно несколькими способами:

По явному сигналу отмены (context.WithCancel)
По истечению промежутка времени (context.WithTimeout)
По наступлению временной отметки или дедлайна (context.WithDeadline)

Пример. Столик в ресторане

Вы хотите забронировать столик в ресторане. Для этого вы набираете ресторан, и ждете, пока на той стороне возьмут
трубку. Далее происходит одно из двух:

Сотрудник ресторана берёт трубку. В таком случае вы начинаете диалог — всё хорошо;
На той стороне никто не берёт трубку в течение минуты, двух, трёх…

Во втором случае вы не будете ждать вечно, и положите трубку, когда вам надоест ждать, либо вы поймёте,
что это не имеет смысла.

Похожим образом работает контекст с таймаутом: как только проходит определённый промежуток времени, исполнение
операции останавливается.

context.WithTimeout()
Ваше приложение отправляет запрос во внешнюю систему, например, в API другого сервиса, который владеет
интересующими вас данными. Так как мы не контролируем  внешние системы, мы не можем быть на 100% уверены,
что API ответит за приемлемое время, или вообще ответит когда-либо. Чтобы не зависнуть навечно в ожидании
ответа от API, в запрос можно передать контекст:
*/

//ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
//defer cancel()

//req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://example.com", nil)
//if err != nil {
//return nil, fmt.Errorf("failed to create request with ctx: %w", err)
//}

//res, err := http.DefaultClient.Do(req)
//if err != nil {
//return nil, fmt.Errorf("failed to perform http request: %w", err)
//}
//
//return res, nil
//
//// Давайте разбираться, что здесь написано.
//
//ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
//defer cancel()

/*
Первым делом происходит инициализация контекста с таймаутом в 15 секунд. Конструкция defer cancel() гарантирует,
что после выхода из функции или горутины контекст будёт отменён, и таким образом вы избежите утекания
горутины — явления, когда горутина продолжает выполняться и существовать в памяти, но результат её работы
больше никого не интересует.
*/

/*
В следующем блоке ничего особенного, создаётся объект *http.Request, куда встраивается созданный нами контекст:
*/

//req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://example.com", nil)
//if err != nil {
//return nil, fmt.Errorf("failed to create request with ctx: %w", err)
//}

/*
Ну, и непосредственно исполнение запроса:
*/

//res, err := http.DefaultClient.Do(req)
//if err != nil {
//return nil, fmt.Errorf("failed to perform http request: %w", err)
//}

/*
Дерево контекстов

Вы можете спросить: а что за context.Background()?

Дело в том, что любой контекст должен наследоваться от какого-то другого, родительского контекста.
Исключения: Background и TODO. Background — это контекст-заглушка, используемый как правило как самый
верхний родитель для всех дочерних контекстов в иерархии. TODO — это тоже заглушка, но используется в тех случаях,
когда мы ещё не определились, какой тип контекста мы хотим использовать. Эти два типа контекста по сути одно и
тоже, и разница исключительно семантическая.

Окей, зачем нужна схема с родительскими и дочерними контекстами? Это сделано для того, чтобы внутри функции,
куда был проброшен контекст, не было возможности повлиять на условия отмены сверху. Таким образом мы имеем
гарантию (с некоторым оговорками), что контекст с дедлайном отменится не позже данного дедлайна.
Кроме того, это даёт возможность дополнять родительский контекст и передавать дальше по цепочке новый контекст,
обогащённый новыми данными.
*/

/*
Для наглядности рассмотрим ещё один пример. Если мы запустим программу ниже, мы увидим: несмотря на то,
что внутри функции doWork таймаут переопределяется на больший, отмена контекста все равно наступит
через 10 секунд:
*/

//package main
//
//import (
//"context"
//	"fmt"
//	"io"
//	"log"
//	"math/rand"
//	"net/http"
//	"time"
//)

//func main() {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//
//	doWork(ctx)
//}

//func doWork(ctx context.Context) {
//	newCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
//	defer cancel()
//
//	log.Println("starting working...")
//
//	for {
//		select {
//		case <-newCtx.Done():
//			log.Printf("ctx done: %v", ctx.Err())
//			return
//		default:
//			log.Println("working...")
//			time.Sleep(1 * time.Second)
//		}
//	}
//}

/*
context.WithDeadline()
Контекст с таймаутом по сути является удобной обёрткой над контекстом с дедлайном.
Программу из предыдущего примера можно выразить немного по-другому:
*/

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

//func main() {
//	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
//	defer cancel()
//
//	doWork(ctx)
//}

func doWork(ctx context.Context) {
	newCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	log.Println("starting working...")

	for {
		select {
		case <-newCtx.Done():
			log.Printf("ctx done: %v", ctx.Err())
			return
		default:
			log.Println("working...")
			time.Sleep(1 * time.Second)
		}
	}
}

/*
time.Now().Add(10*time.Second) — это ровно то, что делает функция context.WithTimeout(), вызывая внутри себя
context.WithDeadline().
*/

/*
context.WithCancel()
Представьте, что вы очень торопитесь на важную встречу. Чтобы добраться побыстрее, вы решаете вызвать такси.
Где быстрее найдётся машина, вы не знаете, поэтому решаете начать поиск в нескольких сервисах одновременно.
Когда в одном из них найдётся машина, вы отмените поиск в остальных. Для похожих и других задач можно
использовать контекст с функцией отмены.

Давайте представим, как описанная ситуация могла бы выглядеть в виде кода:
*/

//package main
//
//import (
//"context"
//"log"
//"math/rand"
//"sync"
//"time"
//)
//
//func main() {
//	var (
//		resultCh    = make(chan string)
//		ctx, cancel = context.WithCancel(context.Background())
//		services    = []string{"Super", "Villagemobil", "Sett Taxi", "Index Go"}
//		wg          sync.WaitGroup
//		winner      string
//	)
//
//	defer cancel()
//
//	for i := range services {
//		svc := services[i]
//
//		wg.Add(1)
//		go func() {
//			requestRide(ctx, svc, resultCh)
//			wg.Done()
//		}()
//	}
//
//	go func() {
//		winner = <-resultCh
//		cancel()
//	}()
//
//	wg.Wait()
//	log.Printf("found car in %q", winner)
//}

func requestRide(ctx context.Context, serviceName string, resultCh chan string) {
	time.Sleep(3 * time.Second)

	for {
		select {
		case <-ctx.Done():
			log.Printf("stopped the search in %q (%v)", serviceName, ctx.Err())
			return
		default:
			if rand.Float64() > 0.75 {
				resultCh <- serviceName
				return
			}

			continue
		}
	}
}

/*
context.WithValue()
Окей, а что насчёт передачи значений через контекст? Для этого в пакете существует функция WithValue.
Давайте взглянем, как это работает:
*/
//package main
//
//import (
//"context"
//"log"
//)
//
//func main() {
//	ctx := context.WithValue(context.Background(), "name", "Joe")
//
//	log.Printf("name = %v", ctx.Value("name"))
//	log.Printf("age = %v", ctx.Value("age"))
//}

/*
Обратите внимание, что метод Value возвращает значение типа interface{}, поэтому скорее всего вам будет
необходимо привести его к нужному типу. Кроме того, если ключ не представлен в контексте, метод вернёт nil.
*/

/*
Когда стоит передавать данные через контекст?
Короткий ответ — никогда. Передача данных через контекст является антипаттерном, поскольку это порождает
неявный контракт между компонентами вашего приложения, к тому же ещё и ненадёжный. Исключение составляют случаи,
когда вам нужно предоставить компоненту из внешней библиотеку вашу реализацию интерфейса, который вы не можете
менять. Например, middleware в HTTP сервере.
*/

/*
Пример. HTTP Middleware
Представьте, что вы хотите, чтобы ваш API принимал запросы только от аутентифицированных клиентов.
Однако вызывать методы для аутентификации в каждом обработчике не кажется удачной идеей.
Но вы можете сделать так, чтобы перед тем как вызовется обработчик запроса, вызвался метод, который проведёт
аутентификацию, и либо вызовет следующий метод в цепочке (в данном случае обработчик), либо вернёт HTTP
с ошибкой аутентификации. Это и есть пример классического middleware.

Вот как это может выглядеть:
*/
//package main
//
//import (
//"context"
//"fmt"
//"io"
//"log"
//"net/http"
//)

type ctxKey string

const keyUserID ctxKey = "user_id"

func main() {
	mux := http.NewServeMux()

	mux.Handle("/restricted", authMiddleware(handleRestricted()))

	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Token")

		if token != "very-secret-token" {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, "invalid token")
			return
		}

		ctx := context.WithValue(r.Context(), keyUserID, 42)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handleRestricted() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(keyUserID).(int)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "internal error, try again later please")
			return
		}

		io.WriteString(w, fmt.Sprintf("hello, user #%d!", userID))
	})
}

/*
Когда использовать контекст?
Метод ходит куда-то по сети;
Горутина исполняется потенциально «долго».
Если у вас есть сомнения насчёт того, соответствует ли функция одному из этих критериев,
то лучше всё-таки добавить контекст. Это не усложнит вам жизнь, но потенциально упростит её в будущем.
Особенно это касается объявляемых вами интерфейсов — внутри реализации может происходить всё что угодно,
в том числе сетевые вызовы и долгие операции.
*/

/*
Советы и лучшие практики

Передавайте контекст всегда первым аргументом — это общепринятое соглашение;

Передавайте контекст только в функции и методы, не храните в состоянии (внутри структуры).

Контексты спроектированы так, чтобы их использовали как одноразовые и неизменяемые объекты.
Например, если вы сохраните контекст с таймутом в 15 секунд в поле структуры, а спустя 15 секунд
попробуете выполнить операцию с данным контекстом, у вас ничего не получится.
Обнулить счётчик таймаута вы тоже не сможете;

Используйте context.WithValue только в крайних случаях. В 99,(9)% случаев вы сможете передать данные через
аргументы функции;

context.Background должен использоваться только как самый верхний родительский контекст,
поскольку он является заглушкой и не предоставляет средств контроля;

Используйте context.TODO, если пока не уверены, какой контекст нужно использовать;

Не забывайте вызывать функцию отмены контекста, т.к. функции, принимающей контекст может потребоваться
время на завершение перед выходом;

Передавайте только контекст, без функции отмены. Контроль за завершением контекста должен оставаться
на вызывающей стороне, иначе логика приложения может стать очень запутанной.

https://pkg.go.dev/context
https://go.dev/blog/context
https://gobyexample.com/context
*/
