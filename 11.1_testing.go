package main

import (
	sync2 "sync"
	"sync/atomic"
)

/*
файлы с тестами должны заканчиваться на '_test
В нашем случае файл с тестами это файл '11.2_test.go'
*/

func plus(a, b int) int {
	// return a + b // тест пройдёт
	return 0 // тест упадет
}

// Далее функции для тестирования бенчмарков (производительности) наших функции
var mu sync2.Mutex

func IncrementMutex(x *uint64) {
	mu.Lock()
	*x += 1
	mu.Unlock()
}

func IncrementAtomic(x *uint64) {
	atomic.AddUint64(x, 1)
}
