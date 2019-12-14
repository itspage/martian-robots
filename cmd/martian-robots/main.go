package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/itspage/martian-robots/pkg/cli"
)

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println("No input.")
		fmt.Println("Usage: cat samples/sample_1.txt | martian-robots")
		return
	}

	cli := cli.CLI{}

	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		if err := cli.ReadLine(string(line)); err != nil {
			log.Fatalf("Error reading line: %s, %v", line, err)
		}
	}

	output, err := cli.Output()
	if err != nil {
		log.Fatalf("Error creating output, %v", output)
	}

	for _, o := range output {
		fmt.Println(o)
	}
}
