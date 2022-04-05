package main

import (
	"fmt"
	"sync"
)

func main() {
	var m = make(map[int]int)
	var l sync.WaitGroup
	for i := 0; i < 200; i++ {
		l.Add(2)
		go func() {
			m[1] = 1
			l.Done()
		}()
		go func() {
			fmt.Println(m[1])
			l.Done()
		}()
	}
	l.Wait()
}
