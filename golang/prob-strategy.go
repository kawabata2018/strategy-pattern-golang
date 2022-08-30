package main

import "math/rand"

type ProbStrategy struct {
	random      *rand.Rand
	prevHand    Hand
	currentHand Hand
	history     [][]int
}

func NewProbStrategy(seed int64) *ProbStrategy {
	return &ProbStrategy{
		random:      rand.New(rand.NewSource(seed)),
		prevHand:    0,
		currentHand: 0,
		history: [][]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		},
	}
}

func (s *ProbStrategy) NextHand() Hand {
	bet := s.random.Intn(s.getSum(s.currentHand))
	var hand Hand
	if bet < s.history[s.currentHand][0] {
		hand = Rock
	} else if bet < s.history[s.currentHand][0]+s.history[s.currentHand][1] {
		hand = Scissors
	} else {
		hand = Paper
	}
	s.prevHand = s.currentHand
	s.currentHand = hand
	return hand
}

func (s *ProbStrategy) getSum(hand Hand) int {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += s.history[hand][i]
	}
	return sum
}

func (s *ProbStrategy) Study(win bool) {
	if win {
		s.history[s.prevHand][s.currentHand]++
	} else {
		s.history[s.prevHand][(s.currentHand+1)%3]++
		s.history[s.prevHand][(s.currentHand+2)%3]++
	}
}
