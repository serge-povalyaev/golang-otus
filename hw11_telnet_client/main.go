package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const timeoutDefault = 10

func main() {
	timeout := flag.Duration("timeout", timeoutDefault*time.Second, "connection timeout")
	flag.Parse()
	argsCount := len(flag.Args())
	if argsCount < 2 || argsCount > 3 {
		log.Fatalf("Неверный вариант запуска программы")
	}

	host, port := flag.Arg(argsCount-2), flag.Arg(argsCount-1)
	address := net.JoinHostPort(host, port)
	client := NewTelnetClient(address, *timeout, os.Stdin, os.Stdout)

	err := client.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("...Connected to", address)

	defer client.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func() {
		err = client.Receive()
		if err != nil {
			log.Fatalln(err)
			return
		}
		fmt.Println("...Connection was closed by peer")

		cancel()
	}()

	go func() {
		err = client.Send()
		if err != nil {
			log.Fatalln(err)
			return
		}

		fmt.Println("...EOF")

		cancel()
	}()

	<-ctx.Done()
}
