package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 20

func main() {

	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(N)) * time.Second)
		return x * 2
	}
	in1 := make(chan int, N)
	in2 := make(chan int, N)
	out := make(chan int, N)

	start := time.Now()
	merge2Channels(fn, in1, in2, out, N+1)
	for i := 0; i < N+1; i++ {
		in1 <- i
		in2 <- i
	}

	orderFail := false
	EvenFail := false
	for i, prev := 0, 0; i < N; i++ {
		c := <-out
		if c%2 != 0 {
			EvenFail = true
		}
		if prev >= c && i != 0 {
			orderFail = true
		}
		prev = c
		fmt.Println(c)
	}
	if orderFail {
		fmt.Println("порядок нарушен")
	}
	if EvenFail {
		fmt.Println("Есть не четные")
	}
	duration := time.Since(start)
	if duration.Seconds() > N {
		fmt.Println("Время превышено")
	}
	fmt.Println("Время выполнения: ", duration)
}

func V1merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		wgAll := sync.WaitGroup{}
		results := make([]int, n)
		for i := 0; i < n; i++ {
			x1, x2 := <-in1, <-in2
			wgAll.Add(1)
			go func(n1, n2, ind int) {
				defer wgAll.Done()
				wg := sync.WaitGroup{}
				wg.Add(1)
				go func() {
					defer wg.Done()
					n1 = fn(n1)
				}()
				wg.Add(1)
				go func() {
					defer wg.Done()
					n2 = fn(n2)
				}()
				wg.Wait()
				results[ind] = n1 + n2
			}(x1, x2, i)
		}
		wgAll.Wait()
		for _, result := range results {
			out <- result
		}
	}()
}
func workUp(mu *sync.Mutex, wg *sync.WaitGroup, fn func(int) int, n int, res *int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		r := fn(n)
		mu.Lock()
		*res += r
		mu.Unlock()
	}()
}
func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		wg := sync.WaitGroup{}
		mu := sync.Mutex{}
		results := make([]int, n)
		for i := 0; i < n; i++ {
			workUp(&mu, &wg, fn, <-in1, &results[i])
			workUp(&mu, &wg, fn, <-in2, &results[i])
		}
		wg.Wait()
		for _, result := range results {
			out <- result
		}
	}()
}
