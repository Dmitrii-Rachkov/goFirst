package principles

/*
5. Dependency Inversion Principle (DIP)
Зависимости должны быть на уровне абстракций, а не конкретных реализаций.
*/

type User struct {
	name string
	id   int
}

type Repository interface {
	GetUser(id int) User
}

type UserService2 struct {
	repo Repository
}

func (u UserService2) GetUser(id int) User {
	return u.repo.GetUser(id)
}

/*
В этом примере UserService2 зависит от абстракции Repository, а не от конкретной реализации.
*/
