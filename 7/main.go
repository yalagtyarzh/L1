package main

import (
	"fmt"
	"sync"
)

// cache, хранящий в себе мапу и mutex, позволяющий блокировать доступ к объекту для прочих горутин
type cache struct {
	m     map[byte]any
	mutex sync.Mutex
}

// wg для синхронизации с main горутиной
var wg sync.WaitGroup

func main() {
	cache := cache{
		m: map[byte]any{},
	}

	// записываем данные в мапу
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go cache.store(byte('a'+i), i)
	}

	// ждемс
	wg.Wait()

	for k, v := range cache.m {
		key := string(k)
		fmt.Printf("key: %s\t\t value: %v\n", key, v)
	}
}

// store конкурентно записывает данные в мапу внутри cache
func (c *cache) store(key byte, value any) {
	// лочим мьютекс
	c.mutex.Lock()
	// записываем
	c.m[key] = value
	// анлочим
	c.mutex.Unlock()
	wg.Done()
}
