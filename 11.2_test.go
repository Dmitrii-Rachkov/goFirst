package main

import "testing"

// Функция называется с 'Test'
// В тесте ниже если 1 + 2 != 3 то
func TestPlus(t *testing.T) {
	result := plus(1, 2)
	if result != 3 {
		// t.Logf("Expected 3, actual %v", result) // выводим в Log
		// t.Fail()                                // помечаем что тест не прошёл
		// t.FailNow() // помечаем тест не прошёл и выходим из программы
		// .Logf() и .Fail() можно заменить на .Errorf()
		t.Errorf("Expected 3, actual %v", result)
	}

	// Если один тест провален и использован метод .Fail(), то будет выполнен следующий тест
	// Если тест провален и использован метод .FailNow(), то следующие тесты не выполнятся
	result = plus(2, 2)
	if result != 4 {
		t.Logf("Expected 4, actual %v", result) // выводим в Log
		t.Fail()                                // помечаем что тест не прошёл
	}
}

/*
Сабтесты
Позволяют разделить наши тесты на какие-то логичесике отдельные части
и возможно вынести какую-то общую инициализацию.
Можем запускать отдельные сабтесты или сразу все в одном.

Сабтесты удобны для интеграционных тестов, когда нам например нужно подключиться к базе данных один раз,
а потом использовать эти данные в каждом из сабтестов.
*/

func TestPlusSub(t *testing.T) {
	// Сабтест для положительных чисел
	t.Run("Positive numbers", func(t *testing.T) {
		result := plus(1, 2)
		if result != 3 {
			t.Errorf("Expected 3, actual %v", result)
		}

		result = plus(2, 2)
		if result != 4 {
			t.Errorf("Expected 4, actual %v", result)
		}
	})

	// Сабтест для отрицательных чисел
	t.Run("Negative numbers", func(t *testing.T) {
		result := plus(-1, -2)
		if result != -3 {
			t.Errorf("Expected -3, actual %v", result)
		}

		result = plus(-2, -2)
		if result != -4 {
			t.Errorf("Expected -4, actual %v", result)
		}
	})
}

// Пример с структурами данных
func TestPlusStruct(t *testing.T) {
	var tests = []struct {
		a, b, result int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{5, 5, 10},
	}

	for _, tt := range tests {
		result := plus(tt.a, tt.b)
		if result != tt.result {
			t.Errorf("Expected %v, actual %v", tt.result, result)
		}
	}
}

// запуск всех тестов в проекте - go test ./...
// запуск тестов с подробностями - go test -v ./...
// запуск всех benchmark в проекте - go test -bench=. ./.
