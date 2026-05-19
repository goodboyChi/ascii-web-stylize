package reader

import (
    "bufio"
	"fmt"
	"os"
	"io"
	"strings"
)

func ReadingFile (banner string) [] string {
// To open our Standard,txt file
	read, err := os.Open("asciiArt/banners/" + banner + ".txt")
	if err != nil {
		fmt.Println("ERROR opening file:", err)
		return nil
	}
	defer read.Close()

	file := bufio.NewReader(read)
	
	var inputFile [] string 
	// AN infinite loop that read our banner file still the end
	for {
		open, err := file.ReadString('\n')
		open1 := strings.TrimRight(open, "\n")
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
		}
		inputFile = append(inputFile, open1)
	}
	return inputFile
}
