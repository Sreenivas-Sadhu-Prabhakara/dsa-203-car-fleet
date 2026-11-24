package main

import (
	"fmt"
	"sort"
)

// CarFleets returns how many fleets arrive at the target: cars drive toward
// it, can't pass each other, and a faster car catching a slower one slows to
// its speed, merging into one fleet. Sort by starting position (closest to
// target first); a car whose solo arrival time exceeds every fleet ahead
// starts a new fleet.
func CarFleets(target int, position, speed []int) int {
	n := len(position)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(a, b int) bool { return position[idx[a]] > position[idx[b]] })
	fleets := 0
	leadTime := 0.0 // arrival time of the fleet directly ahead
	for _, i := range idx {
		t := float64(target-position[i]) / float64(speed[i])
		if t > leadTime {
			fleets++
			leadTime = t // slower than everything ahead: a new fleet forms
		}
		// t <= leadTime: this car catches the fleet ahead and merges.
	}
	return fleets
}

func main() {
	fmt.Println(CarFleets(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))
}
