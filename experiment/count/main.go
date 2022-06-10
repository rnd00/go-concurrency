package count

import "fmt"

func main() {

	cache := make(map[int]int)

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results, cache)
	go worker(jobs, results, cache)
	go worker(jobs, results, cache)
	go worker(jobs, results, cache)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}

	fmt.Println(cache)
}

func worker(jobs <-chan int, results chan<- int, cache map[int]int) {
	for n := range jobs {
		fibResult := fib(n, cache[n-1], cache[n-2])
		results <- fibResult
		cache[n] = fibResult
	}
}

func fib(n, n1, n2 int) int {
	if n <= 1 {
		return n
	}

	return n1 + n2
}
