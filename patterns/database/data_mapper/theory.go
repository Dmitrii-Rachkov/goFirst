package data_mapper

/*
Определение

Data Mapper — это шаблон проектирования, который отделяет представление данных в памяти от хранилища данных.
Он позволяет бизнес-логике и логике доступа к данным работать независимо, помогая в организованном, поддерживаемом
и масштабируемом коде.
*/

/*
Намерение

Целью шаблона Data Mapper является перемещение данных между объектно-ориентированными объектами домена и базой данных,
сохраняя их независимость друг от друга. Он обеспечивает четкое разделение задач, что делает кодовую базу более
модульной и простой в обслуживании.
*/

/*
Также известен как

Объектно-реляционный картограф (ORM)
*/

/*
Подробное объяснение

Шаблон Data Mapper способствует четкому разделению между объектами данных и процессами, которые их сохраняют.
Без него бизнес-логика и код доступа к базе данных часто запутываются, что приводит к жестким и сложным в обслуживании
системам. Data Mapper решает эту проблему, создавая слой, который отвечает за передачу данных между бизнес-объектами
и базовым хранилищем данных.
*/

/*
Основные характеристики

- Разделение : разделяет уровень данных и уровень бизнес-логики.
- Поддерживаемость : улучшает поддерживаемость кода.
- Тестируемость : облегчает тестирование за счет изоляции взаимодействий с базой данных.
- Возможность повторного использования : поощряет повторное использование логики сопоставления данных в различных приложениях и проектах.
*/

/*
Примеры диаграмм классов

image_1.png
*/

/*
Примеры диаграмм последовательности§

image_2.png
*/

/*
Преимущества

- Разделение задач : бизнес-логика изолирована от кода доступа к данным.
- Поддерживаемость : код легче управлять и обновлять.
- Повторное использование : логику сопоставления данных можно повторно использовать в разных проектах.
- Масштабируемость : упрощает независимое масштабирование частей приложения.
*/

/*
Компромиссы

- Сложность : добавляет некоторую сложность за счет дополнительного уровня абстракции.
- Накладные расходы на производительность : потенциальные дополнительные затраты на производительность из-за преобразования объектов.
*/

/*
Когда использовать

- Когда вам необходимо четко разделить задачи между вашей бизнес-логикой и логикой доступа к данным.
- В сложных приложениях, где взаимодействие с базой данных многочисленно, а удобство обслуживания имеет решающее значение.
- Когда логика сопоставления данных повторно используется в нескольких частях проекта или в разных проектах.
*/

/*
Примеры использования

- Корпоративные приложения со сложной логикой предметной области, отделенной от логики сохранения.
- Системы, взаимодействующие с несколькими базами данных или источниками данных.
- Приложения, требующие обширного модульного тестирования и имитации взаимодействия с базой данных.
*/

/*
Когда не следует использовать и анти-шаблоны

- Для простых приложений, где прямой доступ к базе данных в рамках бизнес-логики не затрудняет поддержку.
- Когда дополнительная сложность не оправдывает выгоды, например, в небольших, тесно связанных системах.
*/

/*
Связанные шаблоны проектирования

- Active Record : в отличие от Data Mapper, Active Record сохраняет операции базы данных в пределах объектов домена.
- Репозиторий : часто используется совместно с Data Mapper для инкапсуляции логики, необходимой для доступа к источникам данных.
*/
