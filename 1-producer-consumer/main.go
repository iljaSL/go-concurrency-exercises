//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(tweet chan<- *Tweet, stream Stream) {
	for {
		t, err := stream.Next()
		if err == ErrEOF {
			close(tweet)
			break
		}
		tweet <- t
	}
}

func consumer(tweet <-chan *Tweet, done chan<- bool) {
	for t := range tweet {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
	done <- true
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	tweet := make(chan *Tweet)
	done := make(chan bool)

	// Producer
	go producer(tweet, stream)

	// Consumer
	go consumer(tweet, done)

	<-done

	fmt.Printf("Process took %s\n", time.Since(start))
}
