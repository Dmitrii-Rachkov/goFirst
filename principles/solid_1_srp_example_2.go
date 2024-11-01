package principles

/*
1. Single Responsibility Principle (SRP)
Каждый тип должен иметь одну ответственность. Интерфейсы могут помочь разделить функциональность.
*/

type Logger interface {
	Log(message string)
}

type FileLogger struct{}

func (f FileLogger) Log(message string) {
	// Логирование в файл
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	// Логирование в консоль
}

type UserService struct {
	logger Logger
}

func (u UserService) CreateUser(name string) {
	// Создание пользователя
	u.logger.Log("User created: " + name)
}
