package main

import (
	// "bytes"
	"fmt"
	"io"
	"os"
)

// func Greet(name string) {
// stdoutに出力されるため、テスト（期待した文字列が出力されているか）が困難。
// fmt.Printf("Hello, %s", name)

// → fmt.Printfは、Fprintf(os.Stdout, 略)という実装になっている。
// → Fprintfは、Fprintf(w io.Writer, 略)なので、
//   Greetには（fmt.Printfでなく、）Fprintfを使うようにして、
//   引数としてio.Writerインターフェースを満たすものを渡せば（依存性の注入 DI）すれば、
//   テストしやすい関数になる。
// → テストではbytes.Bufferを渡す
// }

// ↓

// bytes.Bufferが注入される想定
// func Greet(writer *bytes.Buffer, name string) {
// 	fmt.Fprintf(writer, "Hello, %s", name)
// }

// func main() {
// 	Greet(os.Stdout, "World")
// 	// → エラー
// 	//   Greet(writer *bytes.Buffer, name string) では os.Stdout を注入できない
// }

// ↓

// ↑に書いたとおり、os.Stdoutも使える、より汎用的なものにするため、io.Writerにする
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "World")
	// → エラー
	//   Greet(writer *bytes.Buffer, name string) では os.Stdout を注入できない
}