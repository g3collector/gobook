package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, url := range args {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		outFile, err := os.Create("./test")
		// handle err
		defer outFile.Close()
		written, err := io.Copy(outFile, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		written, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%T", written)
	}
}
