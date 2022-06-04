package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
)

// counter хранить в себе сам счетчик, мьютекс для конкурентного доступа к счетчику
// и вэйт группа для синхронизации с вызывающей горутиной
type counter struct {
	count int
	mutex sync.Mutex
	wg    sync.WaitGroup
}

func main() {
	// Получаем с cli количество итераций счетчика
	tries := flag.Int("tries", 5, "tries for increasing counter")
	if *tries < 1 {
		log.Fatal("invalid tries")
	}

	c := newCounter()
	// Запускаем горутины, увеличивающие счетчик
	for i := 0; i < *tries; i++ {
		c.wg.Add(1)
		go c.increase()
	}

	// ждемс выполнения всех горутин
	c.wg.Wait()

	fmt.Println(c.count)
}

// newCounter возвращает новый объект counter
func newCounter() *counter {
	return &counter{}
}

// increase конкурентно увеличивает счетчик на единицу
func (c *counter) increase() {
	c.mutex.Lock()
	c.count++
	c.mutex.Unlock()
	c.wg.Done()
}
