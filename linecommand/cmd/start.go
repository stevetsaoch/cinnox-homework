/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [absolute path of your linebot project]",
	Short: "Start linebot server.",
	Long:  `Start linebot server.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		argument := args[0]
		commands := []string{"docker_start", "server_start"}

		for _, command := range commands {
			e := exec.Command("make", command)
			e.Dir = argument
			var out bytes.Buffer
			e.Stdout = &out
			err := e.Run()
			if err != nil {
				log.Fatal(err)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
