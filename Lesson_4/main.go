package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//Написать программу, которая при получении в канал сигнала
//SIGTERM останавливается не позднее, чем за одну секунду (установить таймаут).

func main() {
	//timer()
		ctx := context.Background()
	ctx,cancel := context.WithCancel(ctx)
	signalOS := make(chan os.Signal, 1)
	signal.Notify(signalOS, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {

			select {
			case <-signalOS:
					cancel()

			default:
				fmt.Println("working")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	select {
	case <-ctx.Done():
		fmt.Println("context")
	}
}


func timer() {
	signalOS := make(chan os.Signal, 1)
	signal.Notify(signalOS, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan time.Time, 1)
	go func() {
		for {
			select {
			case <-signalOS:
				done <- time.Now()
			default:
				fmt.Println("working")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	end := <-done
	fmt.Printf("interrupt time: %f", time.Since(end).Seconds())
}
