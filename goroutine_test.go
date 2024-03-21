package latgo

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello world!!!")
}

func TestGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("hatooo lloooo")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display ", number)
}

func TestNumber(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
}
