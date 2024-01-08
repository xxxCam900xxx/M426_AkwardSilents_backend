package main

import (
	"AkwardSilents/pkg/interfaces/socket"
	"fmt"
	"sync"
)

func main() {
	fmt.Println("start Server")	
	var wg sync.WaitGroup
	wg.Add(2)

	done := make(chan struct{})

	go func() {
		defer wg.Done()
		err := socket.SignupEndpoint()
		if err != nil {
			fmt.Println("Error starting SignupEndpoint:", err)
		}
		fmt.Println("start Server1")	

		close(done)
	}()

	go func() {
		fmt.Println("start Server1")	

		defer wg.Done()
		fmt.Println("start Server1")	

		err := socket.ChatEndpoint()

		fmt.Println("start Server1")	
		if err != nil {
			fmt.Println("Error starting ChatEndpoint:", err)
		}
		fmt.Println("start Server2")	

		close(done)
	}()

	wg.Wait()

	// Wait for both servers to finish
	<-done
	<-done

	fmt.Println("Server started")
}
