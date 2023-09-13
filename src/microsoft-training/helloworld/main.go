// // main: すべての 実行可能プログラムは main パッケージに含まれる必要がある。
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // import (
// // 	"errors"
// // 	"fmt"
// // )

// // import (
// // 	"reflect"
// // 	"fmt"
// // )

// // import: 別のパッケージ内のコードからプログラムにアクセス。
// // fmt: 標準ライブラリパッケージ
// // import "fmt"

// // main(): package main 全体で main()関数は1つだけ。
// // func main() {
// // 	fmt.Println("helo")
// // }

// /*
//   変数の定義方法
// */
// // 1
// // var firstName, lastName string
// // var age int
// // 2
// // var (
// // 	firstName, lastName string
// // 	age int
// // )
// // 3
// // var (
// // 	firstName = "ichiro" // 型が推定される
// // 	lastName = "tanaka"
// // 	age = 20
// // )
// // func main() {
// // firstName, lastName := "ichiro", "tanaka"
// // age := 20 // int
// // age = age /3 // = 6 になる（int = 整数 のため小数が切り捨てられた）
// // fmt.Println(firstName, lastName, age)
// // }

// /*
//   定数の定義方法
// */
// // const (
// // 	StatusOk = 0 // 変数のように := は使わない
// // 	StatusError = 1
// // )

// /*
//   型
// */
// // 浮動小数点数(float32, float64)
// // var shousuu1 float32 = 2.11
// // shousuu2 := 2.11

// // 真偽値(bool)
// // 他の言語のように暗黙的に0, 1に変換しない。明示的に行う必要あり。

// // 文字列(string)
// // ダブルクォーテーション (") で囲む。シングルクォーテーション (')で囲むのは、rune。

// // 型変換・キャスト
// // import "strconv"

// // func main() {
// //     i, _ := strconv.Atoi("-42") // Atoi: 文字列 → 数値へパース。Atoi は返り値が2つあり、2つ目は使わないため、 _ で、これから使わないことを明示。
// //     s := strconv.Itoa(-42)
// //     println(i, s)
// // }

// // func main() {
// // 	num1 := 12
// // 	num2 := 1.23
// // 	// println(reflect.TypeOf(num1))
// // 	// println(reflect.TypeOf(num2))
// // 	fmt.Println(reflect.TypeOf(num1))
// // 	fmt.Println(reflect.TypeOf(num2))
// // }

// // 標準入力
// // import (
// // 	"os"
// // 	"strconv"
// // )

// // func main() {
// // 	sum := sum(os.Args[1], os.Args[2])
// // 	println("合計:", sum)
// // }

// // // func 関数名(引数名 型) 返り値の型 { ... }
// // func sum(number1 string, number2 string) int {
// // 	int1, _ := strconv.Atoi(number1)
// // 	int2, _ := strconv.Atoi(number2)
// // 	return int1 + int2
// // }

// // // func 関数名(引数名 型) (返り値名 返り値の型) { ... }
// // func sum(number1 string, number2 string) (result int) {
// // 	int1, _ := strconv.Atoi(number1)
// // 	int2, _ := strconv.Atoi(number2)
// // 	result = int1 + int2
// // 	return result
// // }

// // ポインター
// // func main() {
// // 	first := "ジョン"
// // 	updateName(&first) // ポインター(メモリアドレス)を渡す
// // 	println(first) // 田中 と出力される
// // }

// // func updateName(name *string) { // 文字列へのポインター を示す
// // 	*name = "田中" // ポインタをupdate
// // }

// // 自作（calculator）パッケージを使う
// // import (
// // 	"github.com/myuser/calculator"
// // 	"rsc.io/quote"
// // )

// // func main() {
// // 	total := calculator.Sum(3, 5)
// // 	println(total)
// // 	println("version: ", calculator.Version)
// // 	println((quote.Hello()))
// // }

// // import "fmt"

// // func givemeanumber() int {
// // 	return -1
// // }

// // func main() {
// // 	// 変数numを定義して、それをif内で使う
// // 	if num := givemeanumber(); num < 0 {
// // 		fmt.Println(num, "is negative")
// // 	} else if num < 10 {
// // 		fmt.Println(num, "has only one digit")
// // 	} else {
// // 		fmt.Println(num, "has multiple digits")
// // 	}

// // 	// if文で定義した変数numは、if外で使えない
// // 	// fmt.Println(num) // エラー undifed: num
// // }

// // import (
// // 	"fmt"
// // 	"math/rand"
// // 	"time"
// // )

// // func main() {
// // 	sec := time.Now().Unix() // UNIXタイム
// // 	rand.Seed(sec) // seed値としてUNIXタイムを使い、実行の度(1秒毎?)にランダムな値を生成
// // 	i := rand.Int31n(10) // 0~10 の乱数
// // 	fmt.Println(i)

// // 	switch i {
// // 	case 0:
// // 		fmt.Println("zero")
// // 	case 1:
// // 		fmt.Println("one")
// // 	case 2:
// // 		fmt.Println("two")
// // 	default: // デフォルト
// // 		fmt.Println("default output")
// // 	}

// // 	fmt.Println("ok")
// // }

// // import (
// // 	"fmt"
// // 	"time"
// // )

// // func main() {
// // 	switch time.Now().Weekday().String() {
// // 	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
// // 		fmt.Println("平日")
// // 	default:
// // 		fmt.Println("休日")
// // 	}

// // 	fmt.Println(time.Now().Weekday().String())
// // }

// // import "fmt"

// // func main() {

// // 	switch num := 15; {
// // 	case num < 50:
// // 		fmt.Printf("%d は50以下\n", num)
// // 		fallthrough
// // 	case num > 100:
// // 		fmt.Println("100以上")
// // 		fallthrough // caseを検証しない
// // 	case num < 200:
// // 		fmt.Println("200以下")
// // 	}
// // 	// -> 15 は50以下
// // 	//    100以上  (fallthroughによって、case文(num > 100)が無視される)
// // 	//    200以下
// // }

// // import "fmt"

// // func main() {
// // 	sum := 0
// // 	for i := 1; i < 100; i++ {
// // 		sum += i
// // 	}
// // 	fmt.Println("sumは", sum)
// // }

// // import (
// // 	"fmt"
// // 	"math/rand"
// // 	"time"
// // )

// // func main() {
// // 	var num int64 // 空の変数
// // 	rand.Seed(time.Now().Unix()) // 乱数 生成器

// // 	// numがたまたま 5 にならない限りループ (= while)
// // 	for num != 5 {
// // 		num = rand.Int63n(15) // 0~15の乱数
// // 		fmt.Println(num)
// // 	}
// // }

// // import (
// // 	"fmt"
// // )

// // func main() {
// // 	var num int32
// // 	sec := time.Now().Unix()
// // 	rand.Seed(sec)

// // 	for {
// // 		fmt.Print("ループ中")

// // 		if num = rand.Int31n(10); num == 5 { // 変数を定義して、それをif内で使う
// // 			fmt.Println("終了")
// // 			break // ループから抜ける
// // 		}

// // 		// break実行時には、この行は実行されない
// // 		fmt.Println(num)
// // 	}
// // }

// // func main() {
// // 	sum := 0
// // 	for num := 1; num <= 10; num++ {
// // 		// numが3の倍数なら、sumに足さない
// // 		if num%3 == 0 {
// // 			continue // continue以下の処理は行わず、次のループに移る
// // 		}
// // 		fmt.Println(num)
// // 		sum += num
// // 	}
// // 	fmt.Println(sum)
// // }

// // import "fmt"

// // func main() {
// // 	for i := 1; i <= 3; i++ {
// // 		defer fmt.Println("defer", -i) // [1, 2, 3]がスタック的に保管され、3, 2, 1の順で遅延実行される（取り出される）
// // 		fmt.Println("regular", i) // ここで"regular 3"が終わると、deferが遅延実行される

// // 		/* 出力
// // 		regular 1
// // 		regular 2
// // 		regular 3
// // 		defer -3
// // 		defer -2
// // 		defer -1
// // 		*/
// // 	}
// // }

// // import (
// // 	"fmt"
// // 	"io"
// // 	"io/ioutil"
// // 	"os"
// // )

// // func main() {
// // 	f, err := os.Create("notes.txt") // ファイル新規作成
// // 	if err != nil {
// // 		return
// // 	}
// // 	defer f.Close() // 閉じる。忘れないようにdeferで遅延実行を予約しておく（以下に大量のコードがある場合など）

// // 	// if文で変数を定義するパターン
// // 	if _, err = io.WriteString(f, "これが書き込まれる"); err != nil {
// // 		return
// // 	}

// // 	out, _ := ioutil.ReadFile("notes.txt") // ファイル読み込み
// // 	fmt.Println(string(out)) // 出力

// // 	f.Sync() // メモリ上のファイル内容をディスクに書き出す（内容を同期）
// // }

// // import "fmt"

// // func main() {
// // 	g(0) // 関数g()を実行
// // 	fmt.Println("finish") // panicにより、この行は実行されない
// // }

// // func g(i int) {
// // 	if i > 3 {
// // 		fmt.Println("パニック前") // 2. panic前の処理は、ふつうに2番目に実行される
// // 		panic("パニック") // 4. deferの後に実行される
// // 	}

// // 	defer fmt.Println("defer:", i) // 3. panic("パニック")より先に実行される

// // 	fmt.Println("print:", i) // 1. まずはこの行が実行される
// // 	g(i + 1) // 再帰(panicが無ければ、無限に繰り返す)
// // }

// // import "fmt"

// // func main() {
// // 	defer func() { // 5. panicが発生したので、deferの後で実行される
// // 		if r := recover(); r != nil { // panicが起きたときに、catchするよう予約
// // 			fmt.Println("リカバーした", r) // r = "パニック"
// // 		}
// // 	}()

// // 	g(0)
// // 	fmt.Println("finish!!") // panicにより、この行は実行されない
// // }

// // func g(i int) {
// // 	if i > 3 {
// // 		fmt.Println("パニック前") // 2. panic前の処理は、ふつうに2番目に実行される
// // 		panic("パニック") // 4. panic発生
// // 	}

// // 	defer fmt.Println("defer:", i) // 3. recover() より先に実行される

// // 	fmt.Println("print:", i) // 1. まずはこの行が実行される
// // 	g(i + 1) // 再帰(panicが無ければ、無限に繰り返す)
// // }

// // func main() {
// // 	fizzBuzz(15)
// // 	fizzBuzz(10)
// // 	fizzBuzz(9)
// // 	fizzBuzz(2)
// // }

// // func fizzBuzz(i int) {
// // 	switch {
// // 	case i%15 == 0:
// // 		fmt.Println("FizzBuzz")
// // 	case i%5 == 0:
// // 		fmt.Println("Buzz")
// // 	case i%3 == 0:
// // 		fmt.Println("Fizz")
// // 	default:
// // 		fmt.Println(i)
// // 	}
// // }

// // func main() {
// // 	guessSquare(25)
// // }

// // func guessSquare(i float64) {
// // 	// 計算途中値（初期値は1）
// // 	sqroot := 1.00
// // 	// 計算結果
// // 	guess := 0.00

// // 	// 計算するのは10回まで
// // 	for count := 1; count <= 10; count++ {
// // 		// 計算する（下記ニュートン法）
// // 		// sqroot n+1 = sqroot n − (sqroot n * sqroot n − x) / (2 * sqroot n)
// // 		guess = sqroot - (sqroot*sqroot-i)/(2*sqroot)

// // 		if sqroot == guess {
// // 			// 計算前後で結果が同じであれば、それが平方根
// // 			fmt.Println("平方根は:", guess)
// // 			break // ループ10回以下でも終了
// // 		} else {
// // 			// ループ
// // 			fmt.Println("計算途中:", guess)
// // 			sqroot = guess
// // 		}
// // 	}
// // }

// // func main() {
// // 	val := 0

// // 	// 繰り返し整数の入力を求めます。 ループの終了条件は、ユーザーが負の数を入力した場合です。
// // 	for val >= 0 {
// // 		fmt.Print("数字を入力 : ")
// // 		fmt.Scanf("%d", &val)
// // 		if val == 0 {
// // 			// 数が 0 の場合は、0 is neither negative nor positive と出力します。 数の要求を続けます。
// // 			fmt.Println(val, "is neither negative nor positive")
// // 		} else if val < 0 {
// // 			// ユーザーが負の数を入力したら、プログラムをクラッシュさせます。 その後、スタック トレース エラーを出力します
// // 			panic("負の数なので終了")
// // 		} else {
// // 			// 数が正の値の場合は、You entered: X と出力します (X は入力された数)。 数の要求を続けます。
// // 			fmt.Println("入力したのは :", val)
// // 		}
// // 	}
// // }

// // func main() {
// // 	var a [3]int
// // 	a[1] = 10

// // 	fmt.Println(a[0]) // 0 : intの場合、初期値が0なので 0 と出力される
// // 	fmt.Println(a[1]) // 10

// // 	fmt.Println(len(a)) // 3 : len()は要素数を取得する

// // 	fmt.Println(a[len(a)-1]) // 0
// // }

// // func main() {
// // 	// 5個の要素を全て埋める必要はない。4, 5個目には、初期値の空文字""がセットされている。
// // 	cities := [5]string{"東京", "大阪", "神戸"}

// // 	fmt.Println(cities) // [東京 大阪 神戸  ]
// // }

// // func main() {
// // 	cities := [...]string{"東京", "大阪", "神戸"} // 要素数を省略

// // 	fmt.Println(cities) // [東京 大阪 神戸]  上記と違って、空文字が含まれていない

// // 	fmt.Println(len(cities)) // 3
// // }

// // func main() {
// // 	var twoD [3][5]int

// // 	for i := 0; i < 3; i++ {
// // 		for j := 0; j < 5; j++ {
// // 			twoD[i][j] = (i + 1) * (j + 1)
// // 		}
// // 		fmt.Println("Row", i, ":", twoD[i])
// // 	}
// // 	fmt.Println(twoD)
// // }

// // func main() {
// // 	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
// // 	fmt.Println("length:", len(months))
// // 	fmt.Println("capacity:", cap(months))

// // 	quarter1 := months[0:3]
// // 	quarter2 := months[3:6]
// // 	quarter3 := months[6:9]
// // 	quarter4 := months[9:12]

// // 	fmt.Println(quarter1, len(quarter1), cap(quarter1)) // [January February March] 3 12
// // 	fmt.Println(quarter2, len(quarter2), cap(quarter2)) // [April May June] 3 9
// // 	fmt.Println(quarter3, len(quarter3), cap(quarter3)) // [July August September] 3 6
// // 	fmt.Println(quarter4, len(quarter4), cap(quarter4)) // [October November December] 3 3
// // 	// lenは、そのスライスが持つ要素数。
// // 	// capは、基になる配列を基準に、どこからスライスしたかによって変わる。

// // 	quarter2Extend := quarter2[:4] // スライスを拡張（要素数を3→4に）
// // 	fmt.Println(quarter2, len(quarter2), cap(quarter2)) // [April May June] 3 9
// // 	fmt.Println(quarter2Extend, len(quarter2Extend), cap(quarter2Extend)) // [April May June July] 4 9
// // }

// // func main() {
// // 	var numbers []int
// // 	for i := 0; i < 10; i++ {
// // 		numbers = append(numbers, i) // 要素を追加
// // 		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
// // 	}
// // }

// // func main() {
// // 	letters := []string{"A", "B", "C", "D", "E"} // スライス
// // 	remove := 2 // 削除する要素の位置

// // 	if remove < len(letters) {

// // 		fmt.Println("Before", letters, "Remove", letters[remove]) // Before [A B C D E] Remove C

// // 		fmt.Println(letters[remove]) // C
// // 		fmt.Println(letters[:remove]) // [A B] = Cの手前まで
// // 		fmt.Println(letters[remove+1:]) // [D E] = C以降
// // 		letters = append(letters[:remove], letters[remove+1:]...) // (切り出した)[A B]に、(切り出した)[D E]を追加している

// // 		fmt.Println("After", letters) // After [A B D E]

// // 	}
// // }

// // func main() {
// // 	letters := []string{"A", "B", "C", "D", "E"} // スライス
// // 	fmt.Println("Before", letters)

// // 	slice1 := letters[0:2] // A, B
// // 	slice2 := letters[1:4] // B, C, D

// // 	slice1[1] = "X"

// // 	fmt.Println("slice1", slice1)
// // 	fmt.Println("slice2", slice2)

// // 	fmt.Println("After", letters)
// // }

// // func main() {
// // 	letters := []string{"A", "B", "C", "D", "E"}
// // 	fmt.Println("Before", letters)

// // 	slice1 := letters[0:2] // A, B

// // 	slice2 := make([]string, 3) // cap=3と指定した 空のスライスが生成される
// // 	copy(slice2, letters[1:4]) // slice2にletters[1:4]をコピーする -> B, C, D

// // 	slice1[1] = "X"

// // 	fmt.Println("slice1", slice1) // slice1 [A X]
// // 	fmt.Println("slice2", slice2) // slice2 [B C D] : 要素が変更されていない!

// // 	fmt.Println("After", letters) // After [A X C D E] : slice1が基にしている配列は当然変更される
// // }

// // func main() {
// // 	studentAge := map[string]int{
// // 		"john": 32,
// // 		"bob": 31,
// // 	}
// // 	fmt.Println(studentAge) // map[bob:31 john:32]
// // 	fmt.Println(studentAge["john"]) // 32
// // }

// // func main() {
// // 	// 空のマップを作成
// // 	studentAge := make(map[string]int)
// // 	fmt.Println(studentAge) // map[]

// // 	// マップに追加
// // 	studentAge["john"] = 32
// // 	studentAge["bob"] = 31
// //   fmt.Println(studentAge) // map[bob:31 john:32]

// // 	// マップ内の項目にアクセス
// // 	fmt.Println("john's age is", studentAge["john"]) // john's age is 32

// // 	// マップ内にないキーでアクセスすると、初期値が返る（今回はintなので 0）
// // 	fmt.Println("hogehoge's age is", studentAge["hogehoge"]) // john's age is 0

// // 	// マップ内に存在するかを確認したい場合
// // 	age, exist := studentAge["hogehoge"] // 2つ目の戻り値でboolが返る
// // 	if exist {
// //     fmt.Println("hogehoge's age is", age)
// // 	} else {
// // 		fmt.Println("not exist")
// // 	}

// // 	// 項目の削除
// // 	delete(studentAge, "john") // delete(マップ, 削除する項目のキー)
// // 	fmt.Println(studentAge) // map[bob:31]

// // 	// 存在しない項目を削除
// // 	delete(studentAge, "hogehoge") // エラー（パニック）にならない
// // }

// // func main() {
// // 	studentAge := map[string]int{
// // 		"john": 32,
// // 		"bob": 31,
// // 	}

// // 	for name, age := range studentAge {
// // 		fmt.Printf("%s\t%d\n", name, age)
// // 	}

// // 	// for name := range studentAge {
// // 	// 	fmt.Printf("%s\t", name)
// // 	// }

// // 	// for _, age := range studentAge {
// // 	// 	fmt.Printf("%d\t", age)
// // 	// }
// // }

// // type Employee struct {
// // 	ID        int
// // 	FirstName string
// // 	LastName  string
// // }

// // func main() {
// // 	employee1 := Employee{1001, "Bob", "Hoge"}
// // 	fmt.Println(employee1) // {1001 Bob Hoge}

// // 	employee2 := Employee{FirstName: "John", LastName: "Fuga"}
// // 	fmt.Println(employee2) // {0 John Fuga}
// // 	fmt.Println(employee2.ID) // 0

// // 	employee2.ID = 1002
// // 	fmt.Println(employee2) // {1002 John Fuga}
// // }

// // func main() {
// // 	employee := Employee{FirstName: "John", LastName: "Fuga"}
// // 	fmt.Println(employee) // {0 John Fuga}

// // 	// ポインタ作成
// // 	employeeCopy := &employee
// // 	employeeCopy.FirstName = "bob"
// // 	fmt.Println(employeeCopy) // &{0 bob Fuga}

// // 	// 元の構造体も変更されている
// // 	fmt.Println(employee) // {0 bob Fuga}
// // }

// // type Person struct {
// // 	ID        int
// // 	FirstName string
// // 	LastName  string
// // }

// // type Employee struct {
// // 	Person // Personを埋め込む
// // 	ManagerID int
// // }

// // type Contractor struct {
// // 	Person // Personを埋め込む
// // 	CompanyID int
// // }

// // func main() {
// // 	employee := Employee{
// // 		// 初期化の際は、Personフィールド（構造体）を明示しないといけない
// // 		Person: Person{
// // 			FirstName: "john",
// // 		},
// // 	}
// // 	fmt.Println(employee) // {{0 john } 0}

// // 	// 初期化でないので、Person経由でなくてOK
// // 	employee.LastName = "doe"
// // 	fmt.Println(employee) // {{0 john doe} 0}
// // }

// // type Person struct {
// // 	ID        int
// // 	FirstName string `json:"name"`
// // 	LastName  string `json:"lastname,omitempty"`
// // }

// // type Employee struct {
// // 	Person
// // 	ManagerID int
// // }

// // type Contractor struct {
// // 	Person
// // 	CompanyID int
// // }

// // func main() {
// // 	employees := []Employee{
// // 		Employee{
// // 			Person: Person{
// // 				LastName: "hoge", FirstName: "john",
// // 			},
// // 		},
// // 		Employee{
// // 			Person: Person{
// // 				LastName: "fuga", FirstName: "bob",
// // 			},
// // 		},
// // 	}

// // 	data, _ := json.Marshal(employees) // 構造体をjsonに変換
// // 	fmt.Printf("%s\n", data)
// // 	// [
// // 	// 	{"ID":0,"name":"john","lastname":"hoge","ManagerID":0},
// // 	// 	{"ID":0,"name":"bob","lastname":"fuga","ManagerID":0}
// // 	// ]

// // 	var decorded []Employee
// // 	json.Unmarshal(data, &decorded) // jsonを構造体に変換(戻す)
// // 	fmt.Printf("%v", decorded)
// // 	// [
// // 	// 	{{0 john hoge} 0}
// // 	// 	{{0 bob fuga} 0}
// // 	// ]
// // }

// // func main() {
// // 	// 入力を受け取る
// // 	input_int, _ := strconv.Atoi(os.Args[1])

// // 	// フィボナッチ数列を作成
// // 	slice := generateFibonacciSequence(input_int)
// // 	fmt.Println(slice)
// // }

// // func generateFibonacciSequence(input int) []int {
// // 	if input < 2 {
// // 		// 入力が2未満の場合計算できないため、panicを起こしnilスライスを返す
// // 		panic([]int{})
// // 	} else {
// // 		// 入力が2以上の場合、フィボナッチ数列を生成する

// // 		// スライスの初期値の意味
// // 		// 0 : ロジック上あったほうが便利なので設けたダミーの数値。(不要なので後で削除する)
// // 		// 1 : 存在することが確定している、最初の値
// // 		slice := []int{0, 1}

// // 		// [入力された数値]個のフィボナッチ数列を生成
// // 		for i := 0; i < input; i++ {
// // 			// 追加する数値
// // 			add_num := slice[len(slice)-2] + slice[len(slice)-1]

// // 			// スライスに要素追加
// // 			slice = append(slice, add_num)
// // 		}

// // 		// 初期値として用意していた 0 を削除
// // 		slice = append(slice[:0], slice[1:]...)

// // 		return slice
// // 	}
// // }

// // type Employee struct {
// // 	ID        int
// // 	FirstName string
// // 	LastName  string
// // 	Address   string
// // }

// // func main() {
// // 	employee, err := getInformation(1001)
// // 	if err != nil {
// // 		// なにかする
// // 	} else {
// // 		fmt.Println(employee)
// // 	}
// // }

// // func getInformation(id int) (*Employee, error) {
// // 	employee, err := apiCallEmployee(1000)

// // 	// 修正前
// // 	// return employee, err

// // 	// 修正後
// // 	// 1. エラーが予想されていなくても、エラーがないかどうかを常に確認します。
// // 	//   そのうえで、不要な情報がエンドユーザーに公開されないようにします。
// // 	if err != nil {
// // 		return nil, err // apiCallEmployee()からエラーが返って来ている場合は、errのみ返す
// // 	}
// // 	return employee, nil
// // }

// // // 3. 可能な限り、再利用可能なエラー変数を作成します。
// // var ErrNotFound = errors.New("Employee not found")

// // func apiCallEmployee(id int) (*Employee, error) {
// // 	if id != 1001 {
// // 		return nil, ErrNotFound
// // 	}

// // 	employee := Employee{LastName: "hoge", FirstName: "john"}
// // 	return &employee, nil
// // }

// // import (
// // 	"log"
// // )

// // func main() {
// // 	log.Print("printろぐ")
// // 	// → 2021/08/09 07:23:53 printろぐ

// // 	// log.Fatal("fatalろぐ")
// // 	// fmt.Print("これは見えない")
// //   // → 2021/08/09 07:23:53 fatalろぐ
// // 	// Fatal()により、プログラムが停止するため、Fatal()以降は実行されない

// // 	// log.Panic("panicろぐ")
// // 	// fmt.Print("これは見えない")
// // 	// → panic: panicろぐ
// // 	//   goroutine 1 [running]:
// // 	//   log.Panic(0xc0000c3f58, 0x1, 0x1)
// // 	// 	  			/usr/local/go/src/log/log.go:354 +0xae
// // 	//   main.main()
// // 	// 	  			/Users/itotakuya/projects/go/src/helloworld/main.go:814 +0xa5
// // 	// Panic()以降は実行されない

// // 	log.SetPrefix("main():")
// //   log.Print("printログ")
// // 	log.Fatal("Fatalログ")
// // 	// → main():2021/08/09 07:29:49 printログ
// // 	//   main():2021/08/09 07:29:49 Fatalログ
// // 	//   exit status 1
// // }

// // func main() {
// // 	// ファイル作成
// // 	// os.OpenFile(ファイル名, フラグ(ファイルが無ければ作成する等), パーミッション)
// // 	file, err := os.OpenFile("ingo.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

// // 	if err != nil {
// // 		log.Fatal("エラー発生")
// // 	}

// // 	// 後でファイルにログを記録するが、file.Close()を忘れないようにdefer予約
// // 	defer file.Close()

// // 	// fileにログを記録することを宣言
// // 	log.SetOutput(file)
// //   // 通常どおり、log.Print()を実行することで、ファイルに記録される
// // 	log.Print("ログ")
// // }

// // type triangle struct {
// // 	size int
// // }

// // // func (変数 構造体) メソッド名() 返却型 {
// // func (t triangle) perimeter() int {
// // 	return t.size * 3
// // }

// // // func main() {
// // // 	t := triangle{3}
// // // 	fmt.Println("この三角形の外周:", t.perimeter())
// // // }

// // // *triangle で、構造体自体でなく 構造体のポインターを渡すように定義
// // // func (t *triangle) doubleSize() {
// // // 	t.size *= 2
// // // }

// // // func main() {
// // // 	t := triangle{3}
// // // 	t.doubleSize() // ポインタを参照して tが更新される

// // // 	fmt.Println("size:", t.size) // size: 6
// // // 	fmt.Println("perimeter:", t.perimeter()) // perimeter: 18
// // // }

// // type coloredTriangle struct {
// // 	triangle
// // 	color string
// // }

// // func (t coloredTriangle) perimeter() int {
// // 	return t.size * 5
// // }

// // func main() {
// // 	t := coloredTriangle{triangle{3}, "blue"}
// // 	fmt.Println("size:", t.size)

// // 	// coloredTriangleで、coloredTriangle構造体向けのメソッドを使う
// // 	fmt.Println("perimeter:", t.perimeter()) // perimeter: 15
// // 	// coloredTriangleで、triangle構造体向けのメソッドも使える
// // 	fmt.Println("perimeter:", t.triangle.perimeter()) // perimeter: 9
// // }

// // func main() {
// // 	// t := geometry.Triangle{}
// // 	// t.SetSize(3)
// // 	// fmt.Println(t.Perimeter()) // 18

// // 	t := geometry.Triangle{}
// // 	t.SetSize(3)
// // 	fmt.Println(t.doubleSize())
// // }

// // インターフェイス ------------------
// // type Shape interface {
// // 	Perimeter() float64
// // 	Area() float64
// // }

// // // 2つの構造体 ------------------
// // // 構造体にはインターフェイスとの関連は定義されていない

// // // Square構造体
// // type Square struct {
// // 	size float64
// // }

// // func (s Square) Area() float64 {
// // 	return s.size * s.size
// // }

// // func (s Square) Perimeter() float64 {
// // 	return s.size * 4
// // }

// // // Circle構造体
// // type Circle struct {
// // 	radius float64
// // }

// // func (c Circle) Area() float64 {
// // 	return math.Pi * c.radius * c.radius
// // }

// // func (c Circle) Perimeter() float64 {
// // 	return 2 * math.Pi * c.radius
// // }

// // // 関数 ------------------
// // // パラメータとしてShapeインターフェイスを求める
// // func printInformation(s Shape) {
// // 	fmt.Printf("%T\n", s)
// // 	fmt.Println("Area:", s.Area())
// // 	fmt.Println("Perimeter:", s.Perimeter())
// // 	fmt.Println()
// // }

// // // main ------------------
// // func main() {
// // 	var s Shape = Square{3}
// // 	printInformation(s)

// // 	c := Circle{6}
// // 	printInformation(c)
// // }

// // type Person struct {
// // 	Name, Country string
// // }

// // // Stringerインターフェイス = fmt.Printfで使用されるインターフェイス
// // // 元のString()メソッドをオーバーライドする(カスタム)
// // func (p Person) String() string {
// // 	return fmt.Sprintf("%v is from %v", p.Name, p.Country)
// // }

// // func main() {
// // 	rs := Person{"john", "USA"}
// // 	ab := Person{"hoge", "UK"}

// // 	// 文字列をPrintfすると、オーバーライドした通りに出力される
// // 	fmt.Printf("%s\n%s\n", rs, ab)
// // 	// john is from USA
// //   // hoge is from UK
// // 	fmt.Printf("%q\n%q\n", rs, ab)
// //   // "john is from USA"
// //   // "hoge is from UK"

// // 	// 文字列以外をPrintfするときは、関係なし
// // 	fmt.Printf("%d\n", 4) // 4
// // }

// // type GithubResponse []struct {
// // 	Fullname string `json:"full_name"`
// // }

// // // 独自のWriteインターフェイスを定義
// // type custumWriter struct{}

// // // Writeメソッドをオーバーライド
// // func (w custumWriter) Write(p []byte) (n int, err error) {
// // 	var resp GithubResponse
// // 	json.Unmarshal(p, &resp)
// // 	for _, r := range resp {
// // 		fmt.Println(r.Fullname)
// // 	}
// // 	return len(p), nil
// // }

// // func main() {
// // 	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
// // 	if err != nil {
// // 		fmt.Println("Error:", err)
// // 		os.Exit(1)
// // 	}

// // 	// io.Copy(os.Stdout, resp.Body)
// // 	// 上記では、大量の出力になる

// // 	// 1行の見やすい出力に変更する
// // 	// Cory()関数のソースを見ると、下記になっているため、Writeインターフェイスを使用する
// // 	//   func Copy(dst Writer, src Reader) (written int64, err error)
// // 	// また、Cory()関数のソースを見ると、Writeメソッドを呼び出すことがわかるため、
// // 	// Writeメソッドをオーバーライドすれば出力を変更できる
// // 	writer := custumWriter{}
// // 	io.Copy(writer, resp.Body)
// // }

// // type dollars float32

// // // dollars型を文字列で出力する場合は $25.00 みたいに$をつけて小数点2桁まで表示するように加工する
// // func (d dollars) String() string {
// // 	return fmt.Sprintf("$%.2f", d)
// // }

// // type database map[string]dollars

// // // ListenAndServe()で呼ばれるServeHTTPメソッドをカスタム
// // func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// // 	for item, price := range db {
// // 		fmt.Fprintf(w, "%s: %s\n", item, price)
// // 	}
// // }

// // func main() {
// // 	// database型をインスタンス化
// // 	db := database{"Go T-shirt": 25, "Go Jacket": 55}
// // 	fmt.Println(db)

// // 	log.Fatal(http.ListenAndServe("localhost:8000", db))
// // 	// localhost:8000にて、
// // 	// Go Jacket:$55.00
// // 	// Go T-shirt:$25.0 と表示される

// // 	// ListenAndServeの定義 :
// // 	//   package http
// // 	//   type Handler interface {
// // 	// 	  	ServeHTTP(w ResponseWriter, r *Request)
// // 	//   }
// // 	//   func ListenAndServe(address string, h Handler) error
// // }

// // type Account struct {
// // 	firstName string
// // 	lastName  string
// // }

// // func (a *Account) ChangeName(newName string) {
// // 	a.firstName = newName
// // }

// // func (a Account) String() string {
// // 	return fmt.Sprintf("%v %v", a.firstName, a.lastName)
// // }

// // type Employee struct {
// // 	Account
// // 	credit float64
// // }

// // func CreateEmployee(firstName, lastName string, credits float64) (*Employee, error) {
// // 	return &Employee{Account{firstName, lastName}, credits}, nil
// // }

// // func (e *Employee) AddCredits(amount float64) (float64, error) {
// // 	if amount > 0.0 {
// // 		e.credit += amount
// // 		return e.credit, nil
// // 	}
// // 	return 0.0, errors.New("入力値が不正です")
// // }

// // func (e *Employee) RemoveCredits(amount float64) (float64, error) {
// // 	if amount > 0.0 {
// // 		if amount <= e.credit {
// // 			e.credit -= amount
// // 			return e.credit, nil
// // 		}
// // 		return 0.0, errors.New("クレジットより多くの値を引くことができません")
// // 	}
// // 	return 0.0, errors.New("入力値が不正です")
// // }

// // func (e *Employee) CheckCredit() float64 {
// // 	return e.credit
// // }

// // func main() {
// // 	bruce, _ := CreateEmployee("Bruce", "Lee", 500)
// // 	fmt.Println(bruce.CheckCredits())
// // 	credits, err := bruce.AddCredits(250)
// // 	if err != nil {
// // 		fmt.Println("Error:", err)
// // 	} else {
// // 		fmt.Println("New Credits Balance = ", credits)
// // 	}

// // 	_, err = bruce.RemoveCredits(2500)
// // 	if err != nil {
// // 		fmt.Println("Can't withdraw or overdrawn!", err)
// // 	}

// // 	bruce.ChangeName("Mark")

// // 	fmt.Println(bruce)
// // }

// // func main() {
// // 	start := time.Now()

// // 	apis := []string{
// // 		"https://management.azure.com",
// // 		"https://dev.azure.com",
// // 		"https://api.github.com",
// // 		"https://outlook.office.com/",
// // 		"https://api.somewhereintheinternet.com/",
// // 		"https://graph.microsoft.com",
// // 	}

// // 	ch := make(chan string)

// // 	for _, api := range apis {
// // 		// APIごとにgoroutineを作成
// // 		go checkAPI(api, ch)
// // 	}

// // 	// サイト数と同じだけ、channelから受信し、出力
// // 	for i := 0; i < len(apis); i++ {
// // 		fmt.Print(<-ch)
// // 	}

// // 	elapsed := time.Since(start)
// // 	fmt.Printf("Done! It took %v seconds\n", elapsed.Seconds())

// // 	// 0.6秒ほどに短縮
// // }

// // func checkAPI(api string, ch chan string) {
// // 	_, err := http.Get(api)
// // 	if err != nil {
// // 		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
// // 		return
// // 	}
// // 	// Sprintfを使うのは、まだ出力はしたくないため
// // 	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
// // }

// // チャネルに送信
// // func send(ch chan<- string, message string) {
// // 	fmt.Printf("Senging: %#v\n", message)
// // 	ch <- message
// // }
// // // チャネルに受信
// // func read(ch <-chan string) {
// // 	fmt.Printf("Receving: %#v\n", <-ch)
// // }

// // func main() {
// // 	ch := make(chan string, 1)
// // 	send(ch, "hello") // Senging: "hello"
// // 	read(ch) // Receving: "hello"
// // }

// // func process(ch chan string) {
// // 	time.Sleep(3 * time.Second)
// // 	ch <- "Done processing!"
// // }

// // func replicate(ch chan string) {
// // 	time.Sleep(1 * time.Second)
// // 	ch <- "Done replicating!"
// // }

// // func main() {
// // 	// チャネルを作成
// // 	ch1 := make(chan string)
// // 	ch2 := make(chan string)

// // 	// goroutineを作成
// // 	go process(ch1)
// // 	go replicate(ch2)

// // 	for i := 0; i < 2; i++ {
// // 		// select句でイベント(チャネルからの送信)を待ち、来たらcase内を実行する
// // 		// さらに(2つ以上の)イベントが発生するまで待機したい場合、必要に応じてループを使う必要がある(for)
// // 		select {
// // 		case process := <-ch1:
// // 			fmt.Println(process)
// // 		case replicate := <-ch2:
// // 			fmt.Println(replicate)
// // 		}
// // 		fmt.Println("ここはイベント終了毎に実行される")
// // 	}
// // }

// // フィボナッチ数列を生成
// func fib(number float64) float64 {
// 	x, y := 1.0, 1.0
// 	for i := 0; i < int(number); i++ {
// 		x, y = y, x+y
// 	}

// 	// ランダムで、0~3秒待つ
// 	r := rand.Intn(3)
// 	time.Sleep(time.Duration(r) * time.Second)

// 	return x
// }

// func main() {
// 	start := time.Now()

// 	// 1~15までのフィボナッチ数列を出力
// 	for i := 1; i < 15; i++ {
// 		n := fib(float64(i))
// 		fmt.Printf("Fib(%v): %v\n", i, n)
// 	}

// 	// 実行時間を出力
// 	elapsed := time.Since(start)
// 	fmt.Printf("Done! %v seconds\n", elapsed.Seconds())
// }
