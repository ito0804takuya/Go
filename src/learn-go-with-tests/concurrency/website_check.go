package concurrency

type WebsiteChecker func(string) bool
type result struct {
	// フィールド名をつけるのが難しいときやつけなくていいとき、型名の定義だけで、型名をそのまま使える
	string
	bool
}

// WebsiteChecker関数をDIする
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		// ゴルーチンを使うときは、関数呼び出しの形にする必要があるので、無名関数をしばしば使う
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}
	
	return results
}
