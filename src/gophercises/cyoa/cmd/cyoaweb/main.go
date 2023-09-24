package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/m/src/gophercises/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "CYOAアプリケーションの起動ポート")

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

	h := cyoa.NewHandler(story)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}