package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

type TimeServer struct {
	Name string
	URL  string
}

type Configuration struct {
	Servers []TimeServer
}

type Time struct {
	Timezone string
	Time     string
}

func main() {
	file, _ := os.Open("conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Configuration{}

	if err := decoder.Decode(&configuration); err != nil {
		log.Fatal(err)
	}

	for {
		var times []Time = []Time{}

		for _, server := range configuration.Servers {
			times = append(times, getTime(server))
		}

		fmt.Printf("\r%s", times[0].Time)
		time.Sleep(1 * time.Second)
	}
}

func getTime(t TimeServer) Time {
	conn, err := net.Dial("tcp", t.URL)
	defer conn.Close()

	if err != nil {
		log.Fatal(err)
	}

	status, err := bufio.NewReader(conn).ReadString('\n')
	return Time{t.Name, fmt.Sprintf("\r%s", status)}
}
