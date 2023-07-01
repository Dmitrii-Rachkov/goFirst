package main

import (
	"fmt"
	"time"
)

// Дата и время
// В go нет варианта отдельно использовать дату и время
// Дата и время храниться вместе в go

func dateTime() {
	fmt.Println("\nДата и время")
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Day(), t.Month(), t.Year())
	// Можем преобразовать месяц в число
	fmt.Println(t.Day(), int(t.Month()), t.Year())
	// Вывести дату с помощью метода Format и готового формата
	fmt.Println(t.Format(time.RFC3339))
	// Формат сами напишем
	fmt.Println(t.Format("02.01.2006 03:04:05"))
	// Прибавим 1 час
	t2 := t.Add(time.Hour * 1)
	fmt.Println(t2)
	// Вычтем один день
	t3 := t.Add(-time.Hour * 24)
	fmt.Println(t3)
	// Используем другой метод и вычтем 1 год
	t4 := t.AddDate(-1, 0, 0)
	fmt.Println(t4)
	// Сравнение дат
	fmt.Println(t.Before(t2)) // t раньше t2 true
	fmt.Println(t3.After(t4)) // t3 позже t4 true
	fmt.Println(t.Equal(t2))  // t == t2 false
	// Разница между датами
	diffTime := t3.Sub(t4)
	fmt.Println(diffTime.Hours())
}
