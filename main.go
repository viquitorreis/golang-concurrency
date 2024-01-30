package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var count uint32 = 0

func main() {
	now := time.Now()

	plchan := make(chan string, 100)
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go hello(plchan, wg)
		go counter(plchan, wg)
		wg.Wait()
	}
	close(plchan)

	for pl := range plchan {
		fmt.Println(pl)
	}

	fmt.Println(time.Since(now))
}

func hello(payload chan string, wg *sync.WaitGroup) {
	count++
	pl := fmt.Sprintf("payload: Hello world! contagem: %d", count)
	payload <- pl

	defer wg.Done()
}

func counter(payload chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddUint32(&count, uint32(1))
	pl := fmt.Sprintf("Estou apenas contando. contagem: %d", count)
	payload <- pl

}
