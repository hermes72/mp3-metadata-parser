package main

import ("fmt";"os";"github.com/metadataparser/mp3metap")

func main() {
	if len(os.Args)<2 {
		fmt.Println("wrong")
	}else{
		filename:=os.Args[1]
		mp3metap.Metaparse(filename)
	}
}