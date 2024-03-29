package main

type InMemoryPlayerScore struct {
	store map[string]int
}

// コンストラクタ
func NewInMemoryPlayerStore() *InMemoryPlayerScore {
	return &InMemoryPlayerScore{map[string]int{}}
}

func (i *InMemoryPlayerScore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerScore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerScore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}