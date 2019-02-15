package main

import (
	"../toarray"
	"fmt"
)

// CreateRanking removes double values from original scores array
func createRanking(arr []int32) []int32 {

	ranking := []int32{arr[0]}
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < max {
			max = arr[i]
			ranking = append(ranking, max)
		}
	}
	return ranking
}

func climbingLeaderboard(scores []int32, alice []int32) []int32 {

  // scores slice without double values
	ranking := createRanking(scores)

	positions := []int32{}

	rank := int32(len(ranking))

	for i := 0; i < len(alice); {
		if rank == 0 {

			positions = append(positions, 1)
			i++
		} else {

			if alice[i] < ranking[rank-1] {

				positions = append(positions, rank+1)
				i++
			} else if alice[i] == ranking[rank-1] {

				positions = append(positions, rank)
				i++
			} else if rank > 0 {

				rank--
			}
		}

	}
	return positions
}

func main() {

	s8 := toarray.GetInput("testcase8scores.txt")
	a8 := toarray.GetInput("testcase8alice.txt")
	result := climbingLeaderboard(s8, a8)
	fmt.Println(result)

}
