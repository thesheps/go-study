// This example doesn't allow concurrent connections! We'll fix this in clock2.go
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	port := flag.Int("p", 1234, "The port on which to run the clock server")
	timezone := flag.String("t", "Europe/London", "The timezone to use")
	flag.Parse()

	listener, err := net.Listen("tcp", "localhost:"+strconv.Itoa(*port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}

		go handleConnection(conn, *timezone)
	}
}

func handleConnection(c net.Conn, timezone string) {
	defer c.Close()

	t, err := time.LoadLocation(timezone)

	if err != nil {
		log.Fatal("Error loading timezone", err)
	}

	for {
		_, err := io.WriteString(c, time.Now().In(t).Format("15:04:05\n"))
		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}
}
