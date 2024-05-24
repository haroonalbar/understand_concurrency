package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	// TestFunc()
  SelectWorking()
}

func SelectWorking() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		for {
			ch1 <- "Every 500ms"
			time.Sleep(500 * time.Millisecond)
		}
	}()
	go func() {
		for {
			ch2 <- "Every 2 second"
			time.Sleep(2 * time.Second)
		}
	}()

	for {
    fmt.Println(<-ch1)
    fmt.Println(<-ch2)
	}
}

func count(s string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- s
		time.Sleep(500 * time.Millisecond)
	}
	close(c)
}

func TestFunc() {
	ch := make(chan string)
	go count("bro", ch)

	for s := range ch {
		fmt.Println(s)
	}

}

func fistEg() {

	// var counter int = 1000

	errChan := make(chan error)

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		defer wg.Done()
		go func(i int) {
			err := randError(i)
			errChan <- err
		}(i)
	}
	// wg.Add(1)
	go func() {
		for err := range errChan {
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	wg.Wait()
	wg.Done()
	defer close(errChan)
}

func randError(index int) error {
	if rand.Float32() < 0.02 {
		return fmt.Errorf("ERROR on index %d", index)
	}

	return nil
}
