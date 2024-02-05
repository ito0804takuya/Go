package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhur.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http;//google.com",
		"http;//blog.hoge.com",
		"waat://furhur.geds",
	}

	want := map[string]bool{
		"http;//google.com": true,
		"http;//blog.hoge.com": true,
		"waat://furhur.geds": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}