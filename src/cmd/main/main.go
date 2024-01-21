package main

import (
	"AkwardSilents/log"
	"AkwardSilents/pkg/interfaces/socket"
	"fmt"
	"sync"
	"AkwardSilents/pkg/service/dbfunctions"
)

func main() {

	log.Info("Server started")
	fmt.Println("Server started")

	var wg sync.WaitGroup
	wg.Add(2)
	err := db.InitDB()
	if err != nil {
		panic(err)
	}
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