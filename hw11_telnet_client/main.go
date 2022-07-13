package main

import (
	"context"
	"flag"
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

	if err := client.Connect(); err != nil {
		log.Fatalln(err)
	}

	defer client.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func() {
		err := client.Receive()
		if err != nil {
			log.Fatalln(err)
			return
		}

		cancel()
	}()

	go func() {
		err := client.Send()
		if err != nil {
			log.Fatalln(err)
			return
		}

		cancel()
	}()

	<-ctx.Done()
}
