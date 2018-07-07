package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	totalFlag bool
)

func init() {
	flag.BoolVar(&totalFlag, "t", false, "print total line count")
}
func main() {
	flag.Parse()
	var total int
	for _, fileName := range flag.Args() {
		func(fileName string) {
			f, err := os.Open(fileName)
			if err != nil {
				fmt.Printf("Something went wrong opening: %q\n%s\n", fileName, err)
				os.Exit(1)
			}
			defer f.Close()
			r := bufio.NewReader(f)
			var count int
			for {
				_, err := r.ReadBytes('\n')
				if err != nil {
					if err == io.EOF {
						break
					} else {
						fmt.Printf("Something went wrong reading %q\n%s\n", fileName, err)
						os.Exit(1)
					}
				}
				count++
			}
			fmt.Printf("%d\t%s\n", count, fileName)
			total += count
		}(fileName)
	}
	if totalFlag {
		fmt.Println(total)
	}
}
