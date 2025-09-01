package goconcepts

import (
	"context"
	"fmt"
	"time"
)

func DoWork() {
	timeout := 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	wkChan := make(chan bool, 1)

	go func() {
		fmt.Println("## DOING WORK .....")
		time.Sleep(2 * time.Second)
		fmt.Println("Things are working well. WORK done")
		wkChan <- true
	}()
	select {
	case <-ctx.Done():
		fmt.Println("timeout occured")
	case <-wkChan:
		fmt.Println("job completed")

	}

}
