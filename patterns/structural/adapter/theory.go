package adapter

/*
Адаптер
Также известен как: Wrapper, Обёртка, Adapter

 Суть паттерна
Адаптер — это структурный паттерн проектирования, который позволяет объектам с несовместимыми интерфейсами
работать вместе.

Адаптер выступает прослойкой между двумя объектами, превращая вызовы одного в вызовы понятные другому.
*/

/*
Проблема
Представьте, что вы делаете приложение для торговли на бирже. Ваше приложение скачивает биржевые котировки
из нескольких источников в XML, а затем рисует красивые графики.

В какой-то момент вы решаете улучшить приложение, применив стороннюю библиотеку аналитики. Но вот беда — библиотека
поддерживает только формат данных JSON, несовместимый с вашим приложением.

image_1.png

Вы смогли бы переписать библиотеку, чтобы та поддерживала формат XML. Но, во-первых, это может нарушить работу
существующего кода, который уже зависит от библиотеки. А во-вторых, у вас может просто не быть доступа к
её исходному коду.
*/

/*
Решение

- Вы можете создать адаптер. Это объект-переводчик, который трансформирует интерфейс или данные одного объекта в такой вид,
чтобы он стал понятен другому объекту.

- При этом адаптер оборачивает один из объектов, так что другой объект даже не знает о наличии первого.
Например, вы можете обернуть объект, работающий в метрах, адаптером, который бы конвертировал данные в футы.

- Адаптеры могут не только переводить данные из одного формата в другой, но и помогать объектам с разными интерфейсами
работать сообща. Это работает так:

1. Адаптер имеет интерфейс, который совместим с одним из объектов.
2. Поэтому этот объект может свободно вызывать методы адаптера.
3. Адаптер получает эти вызовы и перенаправляет их второму объекту, но уже в том формате и последовательности,
которые понятны второму объекту.

Иногда возможно создать даже двухсторонний адаптер, который работал бы в обе стороны.

image_2.png

Таким образом, в приложении биржевых котировок вы могли бы создать класс XML_To_JSON_Adapter,
который бы оборачивал объект того или иного класса библиотеки аналитики. Ваш код посылал бы адаптеру запросы
в формате XML, а адаптер сначала транслировал входящие данные в формат JSON, а затем передавал бы их методам
обёрнутого объекта аналитики.
*/

/*
 Аналогия из жизни

Когда вы в первый раз летите за границу, вас может ждать сюрприз при попытке зарядить ноутбук.
Стандарты розеток в разных странах отличаются. Ваша европейская зарядка будет бесполезна в США без специального
адаптера, позволяющего подключиться к розетке другого типа.
*/

/*
Структура

image_3.png

image_4.png
*/

/*
Применимость

1. Когда вы хотите использовать сторонний класс, но его интерфейс не соответствует остальному коду приложения.

- Адаптер позволяет создать объект-прокладку, который будет превращать вызовы приложения в формат, понятный
стороннему классу.

2. Когда вам нужно использовать несколько существующих подклассов, но в них не хватает какой-то общей функциональности,
причём расширить суперкласс вы не можете.

- Вы могли бы создать ещё один уровень подклассов и добавить в них недостающую функциональность.
Но при этом придётся дублировать один и тот же код в обеих ветках подклассов.

Более элегантным решением было бы поместить недостающую функциональность в адаптер и приспособить его для работы
с суперклассом. Такой адаптер сможет работать со всеми подклассами иерархии. Это решение будет сильно напоминать
паттерн Декоратор.
*/

/*
Шаги реализации

1. Убедитесь, что у вас есть два класса с несовместимыми интерфейсами:

- полезный сервис — служебный класс, который вы не можете изменять (он либо сторонний, либо от него зависит другой код);
- один или несколько клиентов — существующих классов приложения, несовместимых с сервисом из-за неудобного или
несовпадающего интерфейса.

2. Опишите клиентский интерфейс, через который классы приложения смогли бы использовать класс сервиса.

3. Создайте класс адаптера, реализовав этот интерфейс.

4. Поместите в адаптер поле, которое будет хранить ссылку на объект сервиса. Обычно это поле заполняют объектом, переданным в конструктор адаптера. В случае простой адаптации этот объект можно передавать через параметры методов адаптера.

5. Реализуйте все методы клиентского интерфейса в адаптере. Адаптер должен делегировать основную работу сервису.

6. Приложение должно использовать адаптер только через клиентский интерфейс. Это позволит легко изменять и добавлять
адаптеры в будущем.
*/

/*
Преимущества
- Отделяет и скрывает от клиента подробности преобразования различных интерфейсов.
*/

/*
Недостатки
- Усложняет код программы из-за введения дополнительных классов.
*/

/*
 Отношения с другими паттернами

- Мост проектируют загодя, чтобы развивать большие части приложения отдельно друг от друга.
Адаптер применяется постфактум, чтобы заставить несовместимые классы работать вместе.

- Адаптер предоставляет совершенно другой интерфейс для доступа к существующему объекту.
С другой стороны, при использовании паттерна Декоратор интерфейс либо остается прежним, либо расширяется.
Причём Декоратор поддерживает рекурсивную вложенность, чего не скажешь об Адаптере.

- С Адаптером вы получаете доступ к существующему объекту через другой интерфейс. Используя Заместитель,
интерфейс остается неизменным. Используя Декоратор, вы получаете доступ к объекту через расширенный интерфейс.

- Фасад задаёт новый интерфейс, тогда как Адаптер повторно использует старый. Адаптер оборачивает только один класс,
а Фасад оборачивает целую подсистему. Кроме того, Адаптер позволяет двум существующим интерфейсам работать сообща,
вместо того, чтобы задать полностью новый.

- Мост, Стратегия и Состояние (а также слегка и Адаптер) имеют схожие структуры классов — все они построены
на принципе «композиции», то есть делегирования работы другим объектам. Тем не менее, они отличаются тем, что решают
разные проблемы. Помните, что паттерны — это не только рецепт построения кода определённым образом, но и описание
проблем, которые привели к данному решению.
*/