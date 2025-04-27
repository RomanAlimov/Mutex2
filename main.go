package main

import (
	"fmt"
	"sync"
	"time"
)

func error() {
	start := time.Now()
	var cou int

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Nanosecond) //Данный код работает неисправно из за того что он долгий
		cou++
	}

	fmt.Println(cou)
	fmt.Println(time.Now().Sub(start).Seconds())

	// 	result =
	// 	1000
	// 	0.5152547 Отклик, Это много, сейчас я это исправлю в MutexNo
}

func MutexNo() {
	start := time.Now()
	var cou int
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1) // выполнений итерраций

		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond) // сон на наносекунду
			cou++                       // счетчик прибавления

		}()
	}
	wg.Wait() // завершение и ожидание.

	fmt.Println(cou)                             // напечатал счетчик = 981
	fmt.Println(time.Now().Sub(start).Seconds()) // время выполнения программы = 0.0026028
}

// result =
// 981 выставлено не 1000 из за DataRace
// 0.0026028 отклик самый низкий но есть проблема выше

func MutexOn() { // тот же самый код только с добавлением Mutex
	start := time.Now()
	var cou int
	var wg sync.WaitGroup
	var mx sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1) // выполнений итерраций

		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond) // сон на наносекунду

			mx.Lock()   // блокировка изменений горутин
			cou++       // счетчик прибавления
			mx.Unlock() // разблокировка изменений

		}()
	}
	wg.Wait() // завершение и ожидание.

	fmt.Println(cou)                             // напечатал счетчик = 981
	fmt.Println(time.Now().Sub(start).Seconds()) // время выполнения программы = 0.0026028
}

func main() {
	MutexOn()
}

// result =
// 1000
// 0.0025617 теперь все работает как надо, но чуть больше стала задержка кода.
