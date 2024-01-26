package menu

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	ints := make(chan int, 3)
	ints <- 1
	ints <- 2
	ints <- 3
	close(ints)

	for i := 0; i < 10; i++ {
		select {
		case a := <-ints:
			fmt.Println(a)
		}
	}

}
