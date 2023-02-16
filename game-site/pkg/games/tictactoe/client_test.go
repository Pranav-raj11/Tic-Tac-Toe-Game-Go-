package tictactoe

import (
	"fmt"
	"strings"
	"testing"
	//"github.com/stretchr/testify/assert"
)

func TestConsoleClientCreate(t *testing.T) {
	tests := []struct {
		name       string
		boardSize  string
		maxPlayers int
		toConnect  int
	}{
		{
			name:       "Random Numbers",
			boardSize:  "1",
			maxPlayers: 2,
			toConnect:  3,
		}, {
			name:       "Same",
			boardSize:  "3",
			maxPlayers: 3,
			toConnect:  3,
		}, {
			name:       "Letter",
			boardSize:  "h",
			maxPlayers: 3,
			toConnect:  4,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input := fmt.Sprintf("create %s %d %d", test.boardSize, test.maxPlayers, test.toConnect)
			r := strings.NewReader(input)
			b := &strings.Builder{}
			c := NewClient(r, b)
			c.RunClient()
			fmt.Println("b resulted in: ", b.String())
		})
	}
}

func TestConsoleClientMove(t *testing.T) {
	tests := []struct {
		name   string
		gameid string
		player int
		row    int
		col    int
	}{
		{
			name:   "",
			gameid: "",
			player: 3,
			row:    3,
			col:    3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
		})
	}
}

func TestConsoleClientData(t *testing.T) {
	tests := []struct {
		name   string
		gameid string
	}{
		{
			name:   "Printed id",
			gameid: "gam-54688-78493",
		}, {
			name:   "id Index",
			gameid: "1",
		}, {
			name:   "id Index",
			gameid: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// fmt.Println("Printing test.name: ", test.name)
		})
	}
}
