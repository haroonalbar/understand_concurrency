package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	// FistEg()
	// TestFunc()
	// BlockingGo()
	// NonBlockingGo()
  LastWork()
}

func LastWork() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

  // almost uses 100% of a core to run this single goroutine
	go worker(jobs, results)
  // if we add more workers more threads will run concurrently 
  // using more cpu cores
  // but in this example the results might not be in order
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	// adding n to jobs
	for i := 0; i < 100; i++ {
		jobs <- i
	}

	// getting from results
	for i := 0; i < 100; i++ {
		fmt.Println(<-results)
	}

}

// only recive from job and only send to results
// workers goes through the jobs and adds the nth fibonacci to results
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

// return nth fibonacci number
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func NonBlockingGo() {
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
		// Here which ever is ready will be executed first so the channels wouldn't block each other.
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)

		}
	}
}

func BlockingGo() {
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

	/// here this will execute one by one so 2s will block the 500ms one which is ready
	// cause that it moves only to next only if the first one is done.
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

func FistEg() {

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
