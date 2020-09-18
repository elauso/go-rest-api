package pubsub

import (
	"fmt"
	"log"
	"net"
	"os"
	"reflect"
	"strings"
	"time"

	"gopkg.in/redis.v2"
)

type PubSub struct {
	client *redis.Client
}

var Service *PubSub

type Subscriber struct {
	pubsub   *redis.PubSub
	channel  string
	callback processFunc
}

func init() {
	host := os.Getenv("REDIS_DATASOURCE")
	if len(strings.TrimSpace(host)) == 0 {
		panic(fmt.Errorf("Failed to initialize redis cause env host is empty!"))
	}
	client := redis.NewTCPClient(&redis.Options{
		Addr:     host,
		Password: "",
		DB:       0,
		PoolSize: 10,
	})
	Service = &PubSub{client}
}

type processFunc func(string, string)

func NewSubscriber(channel string, fn processFunc) (*Subscriber, error) {

	s := Subscriber{
		pubsub:   Service.client.PubSub(),
		channel:  channel,
		callback: fn,
	}

	if err := s.subscribe(); err != nil {
		return nil, err
	}

	go s.listen()
	return &s, nil
}

func (s *Subscriber) subscribe() error {
	var err error

	err = s.pubsub.Subscribe(s.channel)
	if err != nil {
		log.Println("Error subscribing to channel.")
		return err
	}
	return nil
}

func (s *Subscriber) listen() error {
	var channel string
	var payload string

	for {
		msg, err := s.pubsub.ReceiveTimeout(time.Second)
		if err != nil {
			if reflect.TypeOf(err) == reflect.TypeOf(&net.OpError{}) && reflect.TypeOf(err.(*net.OpError).Err).String() == "*net.timeoutError" {
				// Timeout, ignore
				continue
			}
			log.Print("Error in ReceiveTimeout()", err)
		}

		channel = ""
		payload = ""

		switch m := msg.(type) {
		case *redis.Subscription:
			log.Printf("Subscription Message: %v to channel '%v'. %v total subscriptions.", m.Kind, m.Channel, m.Count)
			continue
		case *redis.Message:
			channel = m.Channel
			payload = m.Payload
		case *redis.PMessage:
			channel = m.Channel
			payload = m.Payload
		}

		go s.callback(channel, payload)
	}
}
