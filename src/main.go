package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"reader"
	"time"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func main() {
	var (
		workDirectory = flag.String("work.ditectory", `C:\development\file-processor\src\ths\`, "work directory")
	)
	flag.Parse()

	c := make(chan string, 3)
	cl := make(chan struct{})
	//var wg sync.WaitGroup
	go reader.ReadFile(*workDirectory, c, cl)

	go func() {

		//wg.Add(1)

		var filesList = make([]string, 0)
		for {
			files, err := ioutil.ReadDir(*workDirectory)
			if err == nil {
				for _, fl := range files {
					if !contains(filesList, fl.Name()) {
						filesList = append(filesList, fl.Name())
						fmt.Print("found file: ", fl.Name())
						c <- fl.Name()
					}
				}
			} else {
				fmt.Println("Problem with reading directory", err)
				cl <- struct{}{}
			}
			time.Sleep(5 * time.Second)
		}
	}()

	<-cl

}
