package main

// Начиная с Go 1.25 можно вот так работать с WaitGroup
/*

	var wg sync.WaitGroup

	for _, item := range items {
		wg.Go(func() {
			process(item) // Просто передаем логику
		})
	}

	wg.Wait()
}
*/

// Под капотом это работает так
/*
func (wg *WaitGroup) Go(f func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		f()
	}()
}
*/
