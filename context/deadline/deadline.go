package goconcepts

import (
	"context"
	"fmt"
	"time"
)

func DoWorkWithDeadline() {
	deadline := time.Now().Add(10 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	wkChan := make(chan bool)

	go func() {
		fmt.Println("## DOING WORK .....")
		time.Sleep(20 * time.Second)
		fmt.Println("Things are working well. WORK done")
		close(wkChan)
	}()
	select {
	case <-ctx.Done():
		fmt.Println("timeout occured: ", ctx.Err())
	case <-wkChan:
		fmt.Println("job completed")

	}

}
