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

var initCmd = &cobra.Command{
	Use:   "init [absolute path of your linebot project]",
	Short: "Initialize linebot server.",
	Long:  `Initialize linebot server.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		commands := []string{"docker_init", "docker_start", "server_start"}
		argument := args[0]
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
	rootCmd.AddCommand(initCmd)
}
