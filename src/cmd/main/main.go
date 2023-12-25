package main

import (
	"AkwardSilents/pkg/interfaces/socket"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	done := make(chan struct{})

	go func() {
		defer wg.Done()
		err := socket.SignupEndpoint()
		if err != nil {
			fmt.Println("Error starting SignupEndpoint:", err)
		}
		close(done)
	}()

	go func() {
		defer wg.Done()
		err := socket.ChatEndpoint()
		if err != nil {
			fmt.Println("Error starting ChatEndpoint:", err)
		}
		close(done)
	}()

	wg.Wait()

	// Wait for both servers to finish
	<-done
	<-done

	fmt.Println("Server started")
}
