// Birkey projects play area
package main

import (
	"fmt"
	"log"
	"log/syslog"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/Shopify/sarama"
)

var slog *log.Logger

func main() {
	slog = newLogger("[goplay] ")
	http.HandleFunc("/", handler)
	slog.Println("Starting http server at localhost:8080")
	slog.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	path := strings.Split(r.URL.Path, "/")
	plen := len(path)
	_, message, topic := path[plen-plen], path[plen-1], path[plen-2]
	slog.Println("topic:", topic, "message:", message)
}

//get a logger with the given prefix that logs to syslog or stdout if
//could not log it to syslog.
func newLogger(prefix string) *log.Logger {
	logger, err := syslog.NewLogger(syslog.LOG_INFO|syslog.LOG_KERN, log.LstdFlags)
	if err != nil {
		fmt.Println("Could not init syslog due to", err, "Will log to stdout instead")
		logger = log.New(os.Stdout, prefix, log.LstdFlags)
	}
	logger.SetPrefix(prefix)
	return logger
}

// publish a message to kafka topic
func toKafka(topic, message string) {
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var enqueued, errors int
	for {
		select {
		case producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(message)}:
			enqueued++
		case err := <-producer.Errors():
			log.Println("Failed to produce message", err)
			errors++
		case <-signals:
			break
		}
	}

	log.Printf("Enqueued: %d; errors: %d\n", enqueued, errors)
}

// func Query(conns []Conn, query string) Result {
//     ch := make(chan Result, len(conns))  // buffered
//     for _, conn := range conns {
//         go func(c Conn) {
//             ch <- c.DoQuery(query):
//         }(conn)
//     }
//     return <-ch
// }
