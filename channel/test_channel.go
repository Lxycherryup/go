package test_channel

import "fmt"

func main() {

	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for {
			if value, ok := <-ch1; ok {
				ch2 <- value * 2
			} else {
				close(ch2)
				break
			}
		}
	}()

	for i := range ch2 {
		fmt.Println("Received from ch2:", i)
	}

}
