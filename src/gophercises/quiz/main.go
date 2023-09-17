package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// コマンドに-csvオプション、-limitオプションを設定
	// flag.型名("フラグ名", デフォルト値, "説明文")
	csvFilename := flag.String("csv", "problems.csv", "渡すcsvファイルは'question,answer'の形式で記述されている必要があります。")
	timeLimit := flag.Int("limit", 30, "1問あたりの回答時間(秒)")
	flag.Parse() // 実際にフラグ変数を使えるようになる

	// csvファイルを開く
	file, err := os.Open(*csvFilename) // csvFilenameは*stringなので、ファイル名（文字列）には*csvFilenameでアクセス
	if err != nil {
		exit(fmt.Sprintf("このファイルを開くのに失敗しました: %s\n", *csvFilename))
	}

	// CSVを読み込む
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("CSVファイルの読み込みに失敗しました")
	}
	problems := parseLines(lines)

	// 回答時間の制限タイマー : すべての問題を、この時間以内に解き終わらないと全問正解できない
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// 正解数
	correct := 0

problemLoop: // このループにラベルを付ける記法
	for i, p := range problems {
		// 回答者に問題を提示
		fmt.Printf("問題 #%d: %s\n", i+1, p.q)

		answerCh := make(chan string)
		go func() {
			var answer string
			// 回答（入力）を受け取る
			fmt.Scanf("%s", &answer)
			answerCh <- answer
		}()

		// チャネルのselect文（いずれかの値を受信したらXXXする）
		select {
		case <-timer.C: // 設定した時間に到達すると、timer.Cチャネルに現在時刻が送信される
			// 制限時間が到達した場合、終了
			break problemLoop // break だけだと（for文でなく）select文からしか抜けられない
		case answer := <-answerCh:
			// 正誤判定
			if answer == p.a {
				correct++
			}
		}
	}
	fmt.Printf("正解数は%d問中、%d問でした", len(problems), correct)
}

// csv形式（q,a）からproblem型{q,a}に変換
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]), // csv内のanswerにスペースが含まれていたら困るので除去
		}
	}

	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// 実行時コマンド（quizディレクトリで）
// （コード修正後はquizディレクトリで go build . した上で）
// ./quiz -csv="hoge.csv" -limit=60
