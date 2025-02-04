package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	store sync.Map
}

func NewCache() *Cache {
	return &Cache{
		// подумайте, нужно ли сюда что-то добавлять
		// подумал - ничего не надо
	}
}

func (c *Cache) Set(key string, value int, d time.Duration) {
	// добавить значение в мапу
	c.store.Store(key, value)
	time.AfterFunc(d, func() {
		// удалить значение из мапы
		c.store.Delete(key)
	})
}

func (c *Cache) Get(key string) (any, bool) {
	// получить и возвратить значение мапы
	value, ok := c.store.Load(key)
	return value, ok
}

func main() {
	// функция main() не требует изменений

	cache := NewCache()

	cache.Set("1", 567, 100*time.Millisecond)
	cache.Set("2", 22, 200*time.Millisecond)
	cache.Set("3", 9, 300*time.Millisecond)

	// определяем локальную фукнцию
	print := func() {
		for _, key := range []string{"1", "2", "3"} {
			v, ok := cache.Get(key)
			if ok {
				fmt.Print(v, " ")
			}
		}
		fmt.Println()
	}

	print()
	time.Sleep(150 * time.Millisecond)
	print()
	time.Sleep(100 * time.Millisecond)
	print()
}
