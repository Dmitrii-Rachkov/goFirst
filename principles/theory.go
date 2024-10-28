package principles

/*
1. YAGNI
You Aren’t Gonna Need It / Вам это не понадобится

Этот принцип прост и очевиден, но ему далеко не все следуют. Если пишете код, то будьте уверены, что он вам понадобится.
Не пишите код, если думаете, что он пригодится позже.

Этот принцип применим при рефакторинге. Если вы занимаетесь рефакторингом метода, класса или файла, не бойтесь удалять лишние методы.
Даже если раньше они были полезны – теперь они не нужны.

Может наступить день, когда они снова понадобятся – тогда вы сможете воспользоваться git-репозиторием, чтобы воскресить
их из мертвых.
*/

/*
2. DRY
Don’t Repeat Yourself / Не повторяйтесь

Дублирование кода – пустая трата времени и ресурсов. Вам придется поддерживать одну и ту же логику и тестировать код
сразу в двух местах, причем если вы измените код в одном месте, его нужно будет изменить и в другом.

В большинстве случаев дублирование кода происходит из-за незнания системы. Прежде чем что-либо писать, проявите
прагматизм: осмотритесь. Возможно, эта функция где-то реализована. Возможно, эта бизнес-логика существует в другом месте.
Повторное использование кода – всегда разумное решение.
*/

/*
3. KISS
Keep It Simple, Stupid / Будь проще

Этот принцип был разработан ВМС США в 1960 году. Этот принцип гласит, что простые системы будут работать лучше и надежнее.

У этого принципа много общего с переизобретением колеса, которым занимались в 1970-х. Тогда он звучал как деловая
и рекламная метафора.

Применительно к разработке ПО он значит следующее – не придумывайте к задаче более сложного решения, чем ей требуется.

Иногда самое разумное решение оказывается и самым простым. Написание производительного, эффективного и
простого кода – это прекрасно.

Одна из самых распространенных ошибок нашего времени – использование новых инструментов исключительно из-за того, что
они блестят. Разработчиков следует мотивировать использовать новейшие технологии не потому, что они новые, а потому
что они подходят для работы.
*/

/*
4. Big Design Up Front
Глобальное проектирование прежде всего

Этот подход к разработке программного обеспечения очень важен, и его часто игнорируют. Прежде чем переходить к реализации,
убедитесь, что все хорошо продумано.

Зачастую продумывание решений избавляло нас от проблем при разработке… Внесение изменений в спецификации занимало час или два.
Если бы мы вносили эти изменения в код, на это уходили бы недели. Я даже не могу выразить, насколько сильно я верю в важность
проектирования перед реализацией, хотя адепты экстремального программирования предали эту практику анафеме.
Я экономил время и делал свои продукты лучше, используя BDUF, и я горжусь этим фактом, чтобы там ни говорили фанатики
экстремального программирования. Они просто ошибаются, иначе сказать не могу.

Многие разработчики считают, что если они не пишут код, то они не добиваются прогресса. Это неверный подход.
Составив план, вы избавите себя от необходимости раз за разом начинать с нуля.

Иногда в недостатках и процессах разработки архитектуры должны быть замешаны и другие люди. Чем раньше вы все это обсудите,
тем лучше будет для всех.

Очень распространенный контраргумент заключается в том, что стоимость решения проблем зачастую ниже стоимости времени планирования.
Чем с меньшим количеством ошибок столкнется пользователь, тем лучше будет его опыт. У вас может не быть другого шанса
справиться с этими ошибками.
*/

/*
5. Avoid Premature Optimization
Избегайте преждевременной оптимизации

Эта практика побуждает разработчиков оптимизировать код до того, как необходимость этой оптимизации будет доказана.
Думаю, что если вы следуете KISS или YAGNI, вы не попадетесь на этот крючок.

Поймите правильно, предвидеть, что произойдет что-то плохое – это хорошо. Но прежде чем вы погрузитесь в детали реализации,
убедитесь, что эти оптимизации действительно полезны.

Очень простой пример – масштабирование. Вы не станете покупать 40 серверов из предположения, что ваше новое приложение
станет очень популярным. Вы будете добавлять серверы по мере необходимости.

Преждевременная оптимизация может привести к задержкам в коде и, следовательно, увеличит затраты времени на вывод
функций на рынок.

Многие считают преждевременную оптимизацию корнем всех зол.
*/

/*
6. Бритва Оккама

Бри́тва О́ккама (иногда ле́звие О́ккама) — методологический принцип, в кратком виде гласящий: «Не следует множить сущее без
необходимости» (либо «Не следует привлекать новые сущности без крайней на то необходимости»).

Что это значит в мире программирования? Не создавайте ненужных сущностей без необходимости.
Будьте прагматичны — подумайте, нужны ли они, поскольку они могут в конечном итоге усложнить вашу кодовую базу.
*/
