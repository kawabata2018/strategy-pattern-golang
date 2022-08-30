package main

type Hand int

// じゃんけんの手を表す3つの定数
const (
	Rock     Hand = iota // 0
	Scissors             // 1
	Paper                // 2
)

// h1がh2より強いときtrue
func (h1 Hand) IsStrongerThan(h2 Hand) bool {
	return h1.fight(h2) == 1
}

// h1がh2より弱いときtrue
func (h1 Hand) IsWeakerThan(h2 Hand) bool {
	return h1.fight(h2) == -1
}

// 引き分けは0, h1の勝ちなら1, h2の勝ちなら-1
func (h1 Hand) fight(h2 Hand) int {
	if h1 == h2 {
		return 0
	} else if (h1+1)%3 == h2 {
		return 1
	} else {
		return -1
	}
}

// じゃんけんの「手」の文字列表現
func (h Hand) String() string {
	return h.Names()[h]
}

// じゃんけんの「手」の文字列表現を返すためのマップ
func (h Hand) Names() []string {
	return []string{
		"グー",
		"チョキ",
		"パー",
	}
}
