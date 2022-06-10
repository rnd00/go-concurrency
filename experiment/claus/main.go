package main

import (
	"fmt"
	"time"
)

func main() {
	// init
	c := make(chan int)
	q := make(chan int)

	go func() {
		defer close(c)
		for i := 1; true; i++ {
			if _, _, ok := TryReceive(q); ok {
				break
			}
			time.Sleep(time.Second)
			c <- i
		}
	}()

	go func() {
		defer close(q)
		time.Sleep(time.Second * 5)
		q <- 1
	}()

	for {
		data, more, ok := TryReceiveWithTimeout(c, time.Second/4)
		if !more {
			break
		}
		fmt.Println("c:", data, more, ok)
	}
}

// ----

func testTryReceive() {
	// init
	c := make(chan int)
	q := make(chan int)

	go func() {
		defer close(c)
		for i := 0; true; i++ {
			if _, _, ok := TryReceive(q); ok {
				break
			}
			time.Sleep(time.Second)
			c <- i
		}
	}()

	go func() {
		defer close(q)
		time.Sleep(time.Second * 10)
		q <- 1
	}()

	for {
		time.Sleep(time.Second / 4)
		fmt.Println(TryReceive(c))
	}
}

func TryReceive(c <-chan int) (data int, more, ok bool) {
	select {
	case data, more = <-c:
		return data, more, true
	default: // processed when c is blocking
		return 0, true, false
	}
}

func TryReceiveWithTimeout(c <-chan int, duration time.Duration) (data int, more, ok bool) {
	select {
	case data, more = <-c:
		return data, more, true
	case <-time.After(duration):
		return 0, true, false
	}
}

// Fanout
// One source, multiple output
func Fanout(In <-chan int, OutA, OutB chan int) {
	for data := range In { // receive until closed
		select { // send to first non blocking channel
		case OutA <- data:
		case OutB <- data:
		}
	}
}

// Funnel
// multiple source, one output
func Funnel(InA, InB <-chan int, Out chan int) {
	var data int
	var more bool
	for {
		// receive from first non-blocking
		select {
		case data, more = <-InA:
		case data, more = <-InB:
		}
		if !more {
			return
		}
		select {
		case Out <- data:
		}
	}
}

// Turnout
// n source, m output
func Turnout(Quit <-chan int, InA, InB, OutA, OutB chan int) {
	var data int
	var more bool

	for {
		// receive from first non-blocking
		select {
		case data, more = <-InA:
		case data, more = <-InB:
		case <-Quit:
			// remember: close generates a message
			// actually this is an anti pattern, but you can argue that quit acts as a delegate
			close(InA)
			close(InB)

			// fanout is here to flush the remaining data
			Fanout(InA, OutA, OutB)
			Fanout(InB, OutA, OutB)
			return
		}
		if !more {
			// ...?
			return
		}
		// send to first non-blocking
		select {
		case OutA <- data:
		case OutB <- data:
		}
	}
}
