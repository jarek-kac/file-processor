package reader

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(workDirectory string, c chan string, cl chan struct{} /*, wg *sync.WaitGroup*/) {
	//defer wg.Done()
	fmt.Print("File path: ", c)
	for {
		select {
		case filePath, open := <-c:
			if open {
				fullFilePath := workDirectory + filePath
				fmt.Println("File path: ", fullFilePath)

				readFile, err := os.Open(fullFilePath)
				defer readFile.Close()

				if err != nil {
					fmt.Println(err)
				}
				fileScanner := bufio.NewScanner(readFile)

				fileScanner.Split(bufio.ScanLines)

				for fileScanner.Scan() {
					fmt.Println(fileScanner.Text())
				}
				//wg.Done()
			} else {
				fmt.Println("Channel closed")
				cl <- struct{}{}
				close(cl)
				return
			}
		default:

		}
	}

}
