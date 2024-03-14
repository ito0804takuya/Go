package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	// ファイルの読み込み位置を先頭に戻す
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}