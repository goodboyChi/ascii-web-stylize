package asciiArt

import (
	"strings"
	"web/asciiArt/reader"
)

func Ascii(text string, banner string) string {

    bannFile := reader.ReadingFile(banner)

    var result strings.Builder

        for j := 0; j < 8; j++ {    
            for _, char := range text{
                if char < 32 || char > 126{
                    continue
                }
                start := int(char-32)*9 + 1 
                result.WriteString(string(bannFile[start+j]))
            } 
            result.WriteString("\n")
        }

    return result.String()
}
