package proxy

/*
Proxy (Заместитель)

Суть паттерна

Заместитель — это структурный паттерн проектирования, который позволяет подставлять вместо реальных объектов
специальные объекты-заменители. Эти объекты перехватывают вызовы к оригинальному объекту, позволяя сделать
что-то до или после передачи вызова оригиналу.

Заместитель — это объект, который выступает прослойкой между клиентом и реальным сервисным объектом.
Заместитель получает вызовы от клиента, выполняет свою функцию (контроль доступа, кеширование, изменение запроса и прочее),
а затем передаёт вызов сервисному объекту.

Заместитель имеет тот же интерфейс, что и реальный объект, поэтому для клиента нет разницы — работать через заместителя
или напрямую.
*/

/*
Проблема

Для чего вообще контролировать доступ к объектам? Рассмотрим такой пример: у вас есть внешний ресурсоёмкий объект,
который нужен не все время, а изредка.

image_1.png

Мы могли бы создавать этот объект не в самом начале программы, а только тогда, когда он кому-то реально понадобится.
Каждый клиент объекта получил бы некий код отложенной инициализации. Но, вероятно, это привело бы к множественному
дублированию кода.

В идеале, этот код хотелось бы поместить прямо в служебный класс, но это не всегда возможно.
Например, код класса может находиться в закрытой сторонней библиотеке.
*/

/*
Решение

Паттерн Заместитель предлагает создать новый класс-дублёр, имеющий тот же интерфейс, что и оригинальный служебный объект.
При получении запроса от клиента объект-заместитель сам бы создавал экземпляр служебного объекта и переадресовывал бы ему
всю реальную работу.

image_2.png

Но в чём же здесь польза?
Вы могли бы поместить в класс заместителя какую-то промежуточную логику, которая выполнялась бы до (или после)
вызовов этих же методов в настоящем объекте. А благодаря одинаковому интерфейсу, объект-заместитель можно передать
в любой код, ожидающий сервисный объект.
*/

/*
 Аналогия из жизни

image_3.png

Платёжная карточка — это заместитель пачки наличных. И карточка, и наличные имеют общий интерфейс — ими можно оплачивать товары.
Для покупателя польза в том, что не надо таскать с собой тонны наличных, а владелец магазина рад, что ему не нужно делать
дорогостоящую инкассацию наличности в банк — деньги поступают к нему на счёт напрямую.
*/

/*
Структура

image_4.png
*/

/*
Применимость

1. Ленивая инициализация (виртуальный прокси).
Когда у вас есть тяжёлый объект, грузящий данные из файловой системы или базы данных.

- Вместо того, чтобы грузить данные сразу после старта программы, можно сэкономить ресурсы и создать объект тогда,
когда он действительно понадобится.

2. Защита доступа (защищающий прокси). Когда в программе есть разные типы пользователей, и вам хочется защищать объект
от неавторизованного доступа. Например, если ваши объекты — это важная часть операционной системы,
а пользователи — сторонние программы (хорошие или вредоносные).

- Прокси может проверять доступ при каждом вызове и передавать выполнение служебному объекту, если доступ разрешён.

3. Локальный запуск сервиса (удалённый прокси). Когда настоящий сервисный объект находится на удалённом сервере.

- В этом случае заместитель транслирует запросы клиента в вызовы по сети в протоколе, понятном удалённому сервису.

4. Логирование запросов (логирующий прокси). Когда требуется хранить историю обращений к сервисному объекту.

- Заместитель может сохранять историю обращения клиента к сервисному объекту.

5. Кеширование объектов («умная» ссылка). Когда нужно кешировать результаты запросов клиентов и управлять их
жизненным циклом.

- Заместитель может подсчитывать количество ссылок на сервисный объект, которые были отданы клиенту и остаются активными.
Когда все ссылки освобождаются, можно будет освободить и сам сервисный объект (например, закрыть подключение к базе данных).

- Кроме того, Заместитель может отслеживать, не менял ли клиент сервисный объект. Это позволит использовать объекты
повторно и здóрово экономить ресурсы, особенно если речь идёт о больших прожорливых сервисах.
*/

/*
Шаги реализации

1. Определите интерфейс, который бы сделал заместитель и оригинальный объект взаимозаменяемыми.

2. Создайте класс заместителя. Он должен содержать ссылку на сервисный объект.
Чаще всего, сервисный объект создаётся самим заместителем. В редких случаях заместитель получает готовый сервисный
объект от клиента через конструктор.

3. Реализуйте методы заместителя в зависимости от его предназначения. В большинстве случаев, проделав какую-то полезную работу,
методы заместителя должны передать запрос сервисному объекту.

4. Подумайте о введении фабрики, которая решала бы, какой из объектов создавать — заместитель или реальный сервисный объект.
Но, с другой стороны, эта логика может быть помещена в создающий метод самого заместителя.

5. Подумайте, не реализовать ли вам ленивую инициализацию сервисного объекта при первом обращении клиента к методам
заместителя.
*/

/*
Преимущества

- Позволяет контролировать сервисный объект незаметно для клиента.
- Может работать, даже если сервисный объект ещё не создан.
- Может контролировать жизненный цикл служебного объекта.
*/

/*
Недостатки

- Усложняет код программы из-за введения дополнительных классов.
- Увеличивает время отклика от сервиса.
*/

/*
Отношения с другими паттернами

- С Адаптером вы получаете доступ к существующему объекту через другой интерфейс.
Используя Заместитель, интерфейс остается неизменным. Используя Декоратор, вы получаете доступ к объекту через
расширенный интерфейс.

- Фасад похож на Заместитель тем, что замещает сложную подсистему и может сам её инициализировать.
Но в отличие от Фасада, Заместитель имеет тот же интерфейс, что его служебный объект, благодаря чему их можно
взаимозаменять.

- Декоратор и Заместитель имеют схожие структуры, но разные назначения. Они похожи тем, что оба построены на принципе
композиции и делегируют работу другим объектам. Паттерны отличаются тем, что Заместитель сам управляет жизнью
сервисного объекта, а обёртывание Декораторов контролируется клиентом.
*/
