package stack

import "sort"

func CarFleet(target int, position []int, speed []int) int {
	pair := []carInfo{}
	stack := []float64{}
	for i := range position {
		pair = append(pair, carInfo{position: position[i], speed: speed[i]})
	}

	sort.Slice(pair, func(i, j int) bool {
		return pair[i].position < pair[j].position
	})

	for i := len(pair) - 1; i >= 0; i-- {
		stack = append(stack, float64(target-pair[i].position)/float64(pair[i].speed))
		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack)
}

type carInfo struct {
	position int
	speed    int
}
