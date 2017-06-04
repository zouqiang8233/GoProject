// This example declares a durable Exchange, and publishes a single message to
// that Exchange with a given routing key.
//
package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

var (
	uri          = flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
	exchangeName = flag.String("exchange", "test-exchange", "Durable AMQP exchange name")
	exchangeType = flag.String("exchange-type", "direct", "Exchange type - direct|fanout|topic|x-custom")
	routingKey   = flag.String("key", "test-key", "AMQP routing key")
	body         = flag.String("body", "foobar", "Body of message")
	reliable     = flag.Bool("reliable", true, "Wait for the publisher confirmation before exiting")
)

func init() {
	flag.Parse()
}

func main() {
	if err := publish(*uri, *exchangeName, *exchangeType, *routingKey, *body, *reliable); err != nil {
		log.Fatalf("%s", err)
	}
	log.Printf("published %dB OK", len(*body))
	time.Sleep(time.Second * 1000000)
}

func publish(amqpURI, exchange, exchangeType, routingKey, body string, reliable bool) error {
	log.Printf("dialing %q", amqpURI)
	connection, err := amqp.Dial(amqpURI)
	if err != nil {
		return fmt.Errorf("Dial: %s", err)
	}
	defer connection.Close()

	log.Printf("got Connection, getting Channel")
	channel, err := connection.Channel()
	if err != nil {
		return fmt.Errorf("Channel: %s", err)
	}

	log.Printf("got Channel, declaring %q Exchange (%q)", exchangeType, exchange)
	if err := channel.ExchangeDeclare(
		exchange,     // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return fmt.Errorf("Exchange Declare: %s", err)
	}

	log.Printf("enabling publishing confirms.")
	if err := channel.Confirm(false); err != nil {
		return fmt.Errorf("Channel could not be put into confirm mode: %s", err)
	}

	confirms := channel.NotifyPublish(make(chan amqp.Confirmation, 100))

	for i := 0; i < 10000000; i++ {

		for j := 0; j < 100; j++ {
			if err = channel.Publish(
				exchange,   // publish to an exchange
				routingKey, // routing to 0 or more queues
				true,       // mandatory
				false,      // immediate
				amqp.Publishing{
					Headers:         amqp.Table{},
					ContentType:     "text/plain",
					ContentEncoding: "",
					Body:            []byte(body),
					DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
					Priority:        0,              // 0-9
					// a bunch of application/implementation-specific fields
				},
			); err != nil {
				return fmt.Errorf("Exchange Publish: %s", err)
			}
		}

		confirmOne(confirms)

	}

	return nil
}

func confirmOne(confirms <-chan amqp.Confirmation) {
	//log.Printf("waiting for confirmation of one publishing")
	for i := 0; i < 100; i++ {
		if confirmed := <-confirms; confirmed.Ack {
			log.Printf("confirmed delivery with delivery tag: %d", confirmed.DeliveryTag)
		} else {
			//log.Printf("failed delivery of delivery tag: %d", confirmed.DeliveryTag)
		}
	}

}
