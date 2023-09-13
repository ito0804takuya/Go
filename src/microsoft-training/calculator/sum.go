// 自作したpackage
package calculator

// パッケージの中からしか呼び出せない
var logMessage = "[LOG]"

// どこからでもアクセスできる(できてしまう)
var Version = "1.0"

// パッケージの中からしか呼び出せない
func internalSum(number int) int {
	return number - 1
}

// どこからでもアクセスできる(できてしまう)
func Sum(number1, number2 int) int {
	return number1 + number2
}
