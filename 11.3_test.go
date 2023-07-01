package main

import "testing"

// Бенчмарки также располагаются в файлах '_test'
// Название функции называется на 'Benchmark' и аргумент '*testing.B'
// Выполняется цикл от 0 до N
func BenchmarkIncrementMutex(b *testing.B) {
	var a uint64 = 0
	for i := 0; i < b.N; i++ {
		IncrementMutex(&a)
	}
}

func BenchmarkIncrementAtomic(b *testing.B) {
	var a uint64 = 0
	for i := 0; i < b.N; i++ {
		IncrementAtomic(&a)
	}
}
