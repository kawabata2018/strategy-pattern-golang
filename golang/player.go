package main

import "fmt"

type Player struct {
	name      string
	strategy  Strategy
	wincount  int
	losecount int
	gamecount int
}

// 名前と戦略を授けてプレイヤーを作る
func NewPlayer(name string, strategy Strategy) *Player {
	return &Player{
		name:     name,
		strategy: strategy,
	}
}

// 戦略におうかがいを立てて次の手を決める
func (p *Player) NextHand() Hand {
	return p.strategy.NextHand()
}

// 勝った
func (p *Player) Win() {
	p.strategy.Study(true)
	p.wincount++
	p.gamecount++
}

// 負けた
func (p *Player) Lose() {
	p.strategy.Study(false)
	p.losecount++
	p.gamecount++
}

// 引き分け
func (p *Player) Even() {
	p.gamecount++
}

func (p *Player) String() string {
	return fmt.Sprintf("[%v:%v games, %v win, %v lose]", p.name, p.gamecount, p.wincount, p.losecount)
}
