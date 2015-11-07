package mp3metap

import (
	"fmt"
	"os"
	"strings"
)

func Metaparse(filename string) {
	//filename:=  os.Args[1]
	file, error := os.Open(filename)
	if error != nil {
		return
	}
	info, error := file.Stat()
	if error != nil {
		return
	}
	bs := make([]byte, info.Size())
	_, error = file.Read(bs)
	if error != nil {
		return
	}
	str := string(bs)
	str=str[info.Size()-128:]
	fmt.Println("Title :",strings.Trim(str[3:33],string(0)))
	fmt.Println("Artist :",strings.Trim(str[33:63],string(0)))
	fmt.Println("Album :",strings.Trim(str[63:93],string(0)))
	fmt.Println("Year :",strings.Trim(str[93:97],string(0)))
	return
}
