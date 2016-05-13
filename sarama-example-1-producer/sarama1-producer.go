// https://github.com/tcnksm-sample/sarama/blob/master/async-producer/main.go

package main

import (
    "fmt"
    "log"
    "os"
    "os/signal"
    "time"

    "strconv"

    "github.com/Shopify/sarama"  // provides Kafka protocol
)

func main() {
    config := sarama.NewConfig()  // Setup config

    // level of acknowledgement reliability needed from the broker
    config.Producer.RequiredAcks = sarama.WaitForAll

    config.Producer.Retry.Max = 5  // total number of tries to retry sending a message, default 3

    //brokers := []string{"localhost:2181"}
    brokers := []string{"localhost:9092"}
    producer, err := sarama.NewAsyncProducer(brokers, config)
    if err != nil {
        // Should not reach here
        panic(err)
    }

    defer func() {
        if err := producer.Close(); err != nil {
            // Should not reach here
            panic(err)
        }
    }()

    signals := make(chan os.Signal, 1)
    signal.Notify(signals, os.Interrupt)

    var enqueued, errors int
    doneCh := make(chan struct{})
    go func() {
        for {
            time.Sleep(500 * time.Millisecond)
            strTime := strconv.Itoa(int(time.Now().Unix()))
            msg := &sarama.ProducerMessage{
                Topic: "test",
                Key: sarama.StringEncoder(strTime),
                Value: sarama.StringEncoder("more stuff"),
            }
            select {
            case producer.Input() <- msg:
                enqueued++
                fmt.Println("Produce message")
            case err := <-producer.Errors():
                errors++
                fmt.Println("Failed to produce message:", err)
            case <-signals:
                doneCh <- struct{}{}
            }
        }
    }()

    <-doneCh
    log.Printf("Enqueued: %d; errors: %d\n", enqueued, errors)
}
