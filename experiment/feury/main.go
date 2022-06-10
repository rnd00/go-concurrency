package feury

import (
	"fmt"
	"time"
)

// you should never close a channel from the receiving end
// if you send a data to a closed channel
// runtime will panic and crash

func main() {
	out1 := make(chan string)
	out2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second / 2)
			out1 <- "order processed"
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second)
			out2 <- "refund processed"
		}
	}()

	for {
		// fmt.Println(<-out1)
		// fmt.Println(<-out2)
		select {
		case msg := <-out1:
			fmt.Println(msg)
		case msg := <-out2:
			fmt.Println(msg)
		}
	}
}

// ----

func concChannel() {
	out1 := make(chan string)

	go process("order", out1)
	for v := range out1 {
		fmt.Println(v)
	}
}

func process(item string, out chan string) {
	defer close(out)
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second / 2)
		out <- fmt.Sprintf("Processed %d %s", i, item)
	}
}
