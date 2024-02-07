package main

import (
	"fmt"
	"math"
	"sync"
)

func part1() {
	const count = 10
	// buffered channel not really needed here, but we at least know
	// how many values we want
	ch := make(chan int, count*2)
	// use a done channel to block
	done := make(chan bool)
	defer close(done)
	for i := 1; i <= 2; i++ {
		go func(multi int) {
			for v := 0; v < count; v++ {
				ch <- v * multi
			}
		}(i)
	}
	go func() {
		counter := 0
		// we know how many values we want, so we keep reading from
		// ch until we're done
		for counter < count*2 {
			val, ok := <-ch
			counter++
			fmt.Println(counter, ": ", val, ok)
		}
		fmt.Println("Closing after", counter)
		// close ch for completeness
		close(ch)
		// invoke done channel
		done <- true
	}()
	<-done
}

func part2() {
	const count = 10
	// use wait group to block on both goroutines doing work
	var wg sync.WaitGroup
	wg.Add(2)
	ch1 := make(chan int, count)
	ch2 := make(chan int, count)
	go func() {
		defer wg.Done()
		for v := 0; v < count; v++ {
			ch1 <- v
		}
	}()
	go func() {
		defer wg.Done()
		for v := 0; v < count; v++ {
			ch2 <- v
		}
	}()
	wg.Wait()
	for {
		// keep reading until channels are empty
		select {
		case v1 := <-ch1:
			fmt.Println("Read ", v1, "from ch1")
		case v1 := <-ch2:
			fmt.Println("Read ", v1, "from ch2")
		default:
			return
		}
	}
}

func part3() {
	type roots map[int]float64
	const numRoots = 100_00
	counter := 0
	var generateRoots = func() roots {
		m := make(roots, numRoots)
		counter++
		for i := 0; i < numRoots; i++ {
			m[i] = math.Sqrt(float64(i))
		}
		return m
	}
	val := sync.OnceValue[roots](generateRoots)
	val()
	fmt.Println(val())
	fmt.Println(counter)
}

func main() {
	// part1()
	part2()
	// part3()
}
