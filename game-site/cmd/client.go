package cmd

import (
	"os"

	"github.com/kertox662/game-site/pkg/games/tictactoe"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Runs a console client for game-site",
	Run: func(cmd *cobra.Command, args []string) {
		v := tictactoe.NewClient(os.Stdin, os.Stdout)
		v.RunClient()
	},
}
