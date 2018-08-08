package main

import (
	"sync"
	"fmt"
)

var total [100000]struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup,index int) {
	defer wg.Done()

	for i := 0; i <= 100; i++ {
		//total[index].Lock()
		total[index].value += i
		//total[index].Unlock()
	}
}

func f(index int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg,index)
	go worker(&wg,index)
	wg.Wait()

	fmt.Println(total[index].value)
}




func main() {

	for i := 0; i< 100000; i++{
		f(i)
	}
}

