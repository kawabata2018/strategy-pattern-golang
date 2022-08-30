package main

import (
	"math/rand"
)

type WinningStrategy struct {
	random   *rand.Rand
	won      bool
	prevHand Hand
}

func NewWinningStrategy(seed int64) *WinningStrategy {
	return &WinningStrategy{
		random: rand.New(rand.NewSource(seed)),
	}
}

func (s *WinningStrategy) NextHand() Hand {
	if !s.won {
		s.prevHand = Hand(s.random.Intn(3))
	}
	return s.prevHand
}

func (s *WinningStrategy) Study(win bool) {
	s.won = win
}
