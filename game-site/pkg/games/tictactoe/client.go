package tictactoe

import (
	"fmt"
	"io"
)

var (
	x, y, z, player, row, col int
	gameid                    string
)

type Client struct {
	input   io.Reader
	output  io.Writer
	gameIds []string
}

func NewClient(input io.Reader, output io.Writer) *Client {
	return &Client{
		input:   input,
		output:  output,
		gameIds: make([]string, 0),
	}
}

func (v *Client) RunClient() {
	var command string
	for {
		fmt.Fprintln(v.output, "Enter command: ")
		_, err := fmt.Fscan(v.input, &command)
		if err != nil {
			fmt.Fprintln(v.output, "Error found", err)
			break
		}
		if command == "create" || command == "Create" || command == "c" {
			x, y, z, err = v.create()

			if err != nil {
				fmt.Fprintln(v.output, "Error found", err)
				break
			}
			fmt.Fprintf(v.output, "boardSize is %d, maxPlayers is %d, toConnect is %d\n", x, y, z)
			// append gameid
		} else if command == "data" || command == "Data" || command == "d" {
			gameid, err = v.data()
			if err != nil {
				fmt.Println("Error found", err)
				break
			}
		} else if command == "move" || command == "Move" || command == "m" {
			gameid, player, row, col, err = v.move()
			if err != nil {
				fmt.Println("Error found", err)
				break
			}
		}
	}
}

func (v *Client) create() (int, int, int, error) {
	fmt.Fprintln(v.output, "Enter board size: ")

	_, err := fmt.Fscan(v.input, &x)
	if err != nil {
		return 0, 0, 0, err
	}
	fmt.Fprintln(v.output, "Enter max players: ")

	_, err = fmt.Fscan(v.input, &y)
	if err != nil {
		return 0, 0, 0, err
	}
	fmt.Fprintln(v.output, "Enter connect target: ")

	_, err = fmt.Fscan(v.input, &z)
	if err != nil {
		return 0, 0, 0, err
	}

	return x, y, z, nil
}

func (v *Client) data() (string, error) {
	fmt.Fprint(v.output, "Enter game id: ")
	_, err := fmt.Fscan(v.input, &gameid)
	if err != nil {
		return "", err
	}
	return gameid, nil
}

func (v *Client) move() (string, int, int, int, error) {
	fmt.Print("Enter game id: ")
	_, err := fmt.Fscan(v.input, &gameid)
	if err != nil {
		fmt.Println("Error found", err)
		return "", 0, 0, 0, err
	}
	fmt.Print("Enter player: ")
	_, err = fmt.Fscan(v.input, &player)
	if err != nil {
		fmt.Println("Error found", err)
		return "", 0, 0, 0, err
	}
	fmt.Print("Enter row: ")
	_, err = fmt.Fscan(v.input, &row)
	if err != nil {
		fmt.Println("Error found", err)
		return "", 0, 0, 0, err
	}
	fmt.Print("Enter column: ")
	_, err = fmt.Fscan(v.input, &col)
	if err != nil {
		fmt.Println("Error found", err)
		return "", 0, 0, 0, err
	}

	return gameid, player, row, col, err
}
