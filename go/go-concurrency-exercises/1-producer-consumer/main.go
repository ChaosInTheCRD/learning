package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, tweets chan *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(tweets)
			return
		}

		tweets <- tweet
	}
}

func consumer(tweets chan *Tweet) {
	for t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	tweets := make(chan *Tweet)

	// Producer
	go producer(stream, tweets)

	// Consumer
	consumer(tweets)

	fmt.Printf("Process took %s\n", time.Since(start))
}
