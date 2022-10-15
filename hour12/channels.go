package main

import (
	"fmt"
	"time"
)

func slowFunc(c chan string) {
	// time.Sleep(time.Second * 2)
	// c <- "slowFunc() finished"
	for msg := range c {
		fmt.Println(msg)
	}
}

func pinger(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "ping"
		<-t.C
	}
}

func channelReader(messages <-chan string) {
	msg := <-messages
	fmt.Println(msg)
}

func channelWriter(messages chan<- string) {
	messages <- "Hello World"
}

func channelReaderAndWriter(messages chan string) {
	msg := <-messages
	fmt.Println(msg)
	messages <- "Hello World"
}

func ping1(c chan string) {
	time.Sleep(time.Second * 1)
	c <- "ping on channel1"
}

func ping2(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "ping on channel2"
}

func sender(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "I'm sending a message"
		<-t.C
	}
}

func main() {
	// c := make(chan string)
	// go slowFunc(c)
	// msg := <-c
	// fmt.Println(msg)

	// messages := make(chan string, 2)
	// messages <- "hello"
	// messages <- "world"
	// close(messages)
	// fmt.Println("Pushed two messages onto Channel with no receivers")
	// time.Sleep(time.Second * 1)
	// slowFunc(messages)

	// messages := make(chan string)
	// go pinger(messages)
	// for i := 0; i < 5; i++ {
	// 	msg := <-messages
	// 	fmt.Println(msg)
	// }

	// channel1 := make(chan string)
	// channel2 := make(chan string)

	// go ping1(channel1)
	// go ping2(channel2)

	// select {
	// case msg1 := <-channel1:
	// 	fmt.Println("received", msg1)
	// case msg2 := <-channel2:
	// 	fmt.Println("received", msg2)
	// case <-time.After(500 * time.Millisecond):
	// 	fmt.Println("no messages received. giving up.")
	// }

	messages := make(chan string)
	stop := make(chan bool)
	go sender(messages)
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Time's up!")
		stop <- true
	}()
	for {
		select {
		case <-stop:
			return
		case msg := <-messages:
			fmt.Println(msg)
		}
	}
}
