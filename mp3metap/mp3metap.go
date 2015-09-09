package mp3metap

import (
	"fmt"
	"os"
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
	fmt.Println("Title :",str[3:33])
	fmt.Println("Artist :",str[33:63])
	fmt.Println("Album :",str[63:93])
	fmt.Println("Year :",str[93:97])
	return
}
