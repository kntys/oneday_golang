package main

import "fmt"

func scoreAvg(scores []int) int {
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	return sum / len(scores)
}

func sbSwr(a int, b int) int {
	return (a - b) * (a - b)
}

func avgVar(scores []int, avg int) int {
	sum := 0
	for i := 0; i < len(scores); i++ {
		if avg == 0 {
			sum += scores[i]
		} else {
			sum += sbSwr(scores[i], avg)
		}
	}
	return sum / len(scores)
}

func scoreMaxMin(scores []int) (max int, min int) {
	max = scores[0]
	min = scores[0]

	for i := 1; i < len(scores); i++ {
		if max < scores[i] {
			max = scores[i]
		}
		if min > scores[i] {
			min = scores[i]
		}
	}
	return
}

func main() {
	mathScores := []int{40, 89, 77, 68, 72, 39, 58, 87, 93, 48, 65, 74, 60}
	fmt.Printf("%d名による数学試験の結果:\n", len(mathScores))
	fmt.Printf("平均点 %d点\n", scoreAvg(mathScores))
	fmt.Printf(" 分散　%d（点・点)\n", avgVar(mathScores, avgVar(mathScores, 0)))
	max, min := scoreMaxMin(mathScores)
	fmt.Printf("最高点　%d \n", max)
	fmt.Printf("最低点　%d \n", min)
}
