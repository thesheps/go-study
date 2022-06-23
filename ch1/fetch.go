package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// For the purposes of this example I've decided to just swallow any errors
func main() {
	for _, url := range os.Args[1:] {
		resp, _ := http.Get(url)

		b, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		fmt.Printf("%s", b)
	}
}
