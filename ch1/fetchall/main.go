package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	goRoutineChannel := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, goRoutineChannel)
	}

	for range os.Args[1:] {
		fmt.Println(<-goRoutineChannel) // receive from channel goRoutineChannel
	}

	fmt.Printf("%.2fs elapsed \n", time.Since(start).Seconds())

}

func fetch(url string, channel chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		channel <- fmt.Sprint(err) // send error to channel
		return
	}
	// we want to discard the body
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		channel <- fmt.Sprintf("while reading %s , %v ", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	channel <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
