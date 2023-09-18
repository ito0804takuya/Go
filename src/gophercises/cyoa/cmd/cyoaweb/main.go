package main

import (
	"flag"
	"fmt"
	"os"

	"example.com/m/src/gophercises/cyoa"
)

func main() {
	// コマンドにfileオプションを設定
	filename := flag.String("file", "gopher.json", "CYOAストーリーに使うJSONファイル")
	flag.Parse()
	fmt.Printf("%s.\n", *filename)

	// jsonファイルを開く
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, nil := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}