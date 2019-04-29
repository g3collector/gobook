package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]int)
	files := os.Args[1:]

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		s := bufio.NewScanner(f)
		for s.Scan() {
			if s.Text() == "" {
				break
			}
			counts[s.Text()]++
			n := counts[s.Text()]
			if n >= 2 && filenames[f.Name()] < 2 {
				filenames[f.Name()]++
			}
		}
	}

	for fileStuff, filesWithMoreThanTwo := range filenames {
		if filesWithMoreThanTwo > 0 {
			fmt.Println(fileStuff, filesWithMoreThanTwo)
		}
	}
}
