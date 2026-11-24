package main

import "testing"

func TestCarFleets(t *testing.T) {
	cases := []struct {
		target          int
		position, speed []int
		want            int
	}{
		{12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}, 3},
		{10, []int{3}, []int{3}, 1},
		{100, []int{0, 2, 4}, []int{4, 2, 1}, 1}, // rear cars catch the leader
		{10, []int{0, 4, 2}, []int{2, 1, 3}, 1},
		{10, []int{6, 8}, []int{3, 2}, 2}, // leader too slow to be caught... check: car@8 t=1, car@6 t=4/3>1: 2 fleets
		{1, nil, nil, 0},
	}
	for _, c := range cases {
		if got := CarFleets(c.target, c.position, c.speed); got != c.want {
			t.Errorf("CarFleets(%d,%v,%v) = %d, want %d", c.target, c.position, c.speed, got, c.want)
		}
	}
}
