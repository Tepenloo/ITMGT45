package main

import (
	"fmt"
	"strings"
)

func relationshipStatus(fromMember string, toMember string, socialGraph map[string]map[string]string) string {
	fromFollowsTo := strings.Contains(socialGraph[fromMember]["following"], toMember)
	toFollowsFrom := strings.Contains(socialGraph[toMember]["following"], fromMember)

	if fromFollowsTo && toFollowsFrom {
		return "friends"
	} else if fromFollowsTo {
		return "follower"
	} else if toFollowsFrom {
		return "followed by"
	}
	return "no relationship"
}

func ticTacToe(board [][]string) string {
	size := len(board)

	for i := 0; i < size; i++ {
		if allEqual(board[i]) {
			return board[i][0]
		}
		column := []string{}
		for j := 0; j < size; j++ {
			column = append(column, board[j][i])
		}
		if allEqual(column) {
			return column[0]
		}
	}

	diag1, diag2 := []string{}, []string{}
	for i := 0; i < size; i++ {
		diag1 = append(diag1, board[i][i])
		diag2 = append(diag2, board[i][size-i-1])
	}
	if allEqual(diag1) {
		return diag1[0]
	}
	if allEqual(diag2) {
		return diag2[0]
	}

	return "NO WINNER"
}

func allEqual(arr []string) bool {
	if arr[0] == "" {
		return false
	}
	for _, v := range arr {
		if v != arr[0] {
			return false
		}
	}
	return true
}

func eta(firstStop string, secondStop string, routeMap map[string]map[string]int) int {
	time := 0
	current := firstStop

	for {
		for next, travelTime := range routeMap[current] {
			time += travelTime
			current = next
			if current == secondStop {
				return time
			}
		}
	}
}

func main() {
	socialGraph := map[string]map[string]string{
		"@joaquin": {"following": "@chums,@jobenilagan"},
		"@chums":   {"following": "@bongolpoc,@miketan,@rudyang,@joeilagan"},
	}
	fmt.Println(relationshipStatus("@joaquin", "@chums", socialGraph))

	board := [][]string{
		{"X", "X", "O"},
		{"O", "X", "O"},
		{"", "O", "X"},
	}
	fmt.Println(ticTacToe(board))

	routeMap := map[string]map[string]int{
		"upd":  {"admu": 10},
		"admu": {"dlsu": 35},
		"dlsu": {"upd": 55},
	}
	fmt.Println(eta("upd", "dlsu", routeMap))
}
