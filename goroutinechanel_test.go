package latgo

import (
	"fmt"
	"testing"
	"time"
)

func MakeChanel(t *testing.T) {
	chanel := make(chan string)
	defer close(chanel)

	go func() {
		time.Sleep(2 * time.Second)
		chanel <- "suppa dheepssupreme"
		fmt.Println("mengrim data chanel")
	}()

	data := <-chanel
	fmt.Println(data)
	time.Sleep(2* time.Second)

}
