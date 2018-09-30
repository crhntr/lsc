package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	totalFlag bool
)

func init() {
	flag.BoolVar(&totalFlag, "t", false, "print total line count")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nUsage of %s:\n", os.Args[0])
		fmt.Printf("Example:\n\t%[1]s -t *_test.go\n\t%[1]s helloWorld.txt\n", os.Args[0])
		fmt.Println("Flags")
		flag.PrintDefaults()
	}
}
func main() {
	flag.Parse()
	var total int

	wd := "./"
	if len(flag.Args()) > 0 {
		wd = flag.Arg(0)
	}

	err := filepath.Walk(wd, func(fileName string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return nil
		}
		if strings.HasSuffix(fileName, ".git") {
			return filepath.SkipDir
		}
		if info.IsDir() {
			return nil
		}
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
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if totalFlag {
		fmt.Println(total)
	}
}
