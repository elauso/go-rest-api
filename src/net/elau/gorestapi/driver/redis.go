package driver

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var RedisClient *redis.Client

func InitRedis(url string) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "",
		DB:       0,
	})
}

type callback func(string, string)

type Subscriber struct {
	pubsub   *redis.PubSub
	channel  string
	callback callback
}

func NewSubscriber(ch string, fn callback) *Subscriber {

	s := &Subscriber{
		pubsub:   RedisClient.Subscribe(ctx, ch),
		channel:  ch,
		callback: fn,
	}
	go s.listen()
	return s
}

func (s *Subscriber) listen() {
	for {
		msgi, err := s.pubsub.Receive(ctx)
		if err != nil {
			log.Printf("Failed to receive message channel, %v", err)
			break
		}

		switch msg := msgi.(type) {
		case *redis.Subscription:
			log.Println("Subscribed to ", msg.Channel)
		case *redis.Message:
			go s.callback(msg.Channel, msg.Payload)
		default:
			panic("Unreached channel listen.")
		}
	}
	s.pubsub.Close()
}
