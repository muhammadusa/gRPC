package main

import (
	"fmt"
	"sync"
)

func Task(wg *sync.WaitGroup, name string) {

	fmt.Println(name)

	wg.Done()
}

func main () {

	var wg sync.WaitGroup

	wg.Add(3)

	go Task(&wg, "Farruh")
	go Task(&wg, "Asadbek")
	go Task(&wg, "Husniddin")

	wg.Wait()
}
