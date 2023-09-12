package main

import (
	"fmt"
	"sync"
)

// A broker is responsible for holding topics that cache messages, as well as a waitGroup
type broker struct {
	topics    map[string]topic
	waitGroup sync.WaitGroup
}

// A topic is a virtual queue that caches messages sent by a publisher to be consumed by a subscriber
type topic struct {
	channel chan string
}

// initTopic initialises a topic on the broker of a given name 't'
func (b *broker) initTopic(t string) {
	b.topics[t] = topic{channel: make(chan string)}
	go func() {
		b.waitGroup.Wait()
		close(b.topics[t].channel)
	}()
}

// publish publishes a message string 'm' to a topic channel
func (b *broker) publish(t string, m string) {
	defer b.waitGroup.Done()
	// checks that the topic exists
	_, ok := b.topics[t]
	if !ok {
		panic(fmt.Errorf("The topic does not exist!"))
	}
	b.topics[t].channel <- m
}

// subscribe iterates over the topic's channel and reads the messages on it
func (b *broker) subscribe(t string, h func(m string)) {
	for i := range b.topics[t].channel {
		h(i)
	}
}

func main() {
	b := broker{topics: make(map[string]topic), waitGroup: sync.WaitGroup{}}
	b.initTopic("foo")
	b.initTopic("bar")

	messageOne := "first test"
	messageTwo := "second test"

	b.waitGroup.Add(1)
	go b.publish("foo", messageOne)
	b.waitGroup.Add(1)
	go b.publish("bar", messageTwo)
	b.waitGroup.Add(1)
	go b.publish("foo", messageTwo)
	b.waitGroup.Add(1)
	go b.publish("bar", messageOne)
	// This will panic
	// b.waitGroup.Add(1)
	// go b.publish("foobar", messageOne)

	go b.subscribe("foo", func(m string) { fmt.Println("foo message:", m) })
	go b.subscribe("bar", func(m string) { fmt.Println("bar message:", m) })

	b.waitGroup.Wait()
	fmt.Println("finished.")
}
