/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var pushmessageCmd = &cobra.Command{
	Use:   "pushmessage",
	Short: "Push message to users.",
	Long:  `Push message to users who had sent message to your line official account.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		argument := args[0]
		values := map[string]string{"message": argument}
		json_data, err := json.Marshal(values)

		resp, err := http.Post("http://127.0.0.1:8000/push", "application/json",
			bytes.NewBuffer(json_data))

		if err != nil {
			log.Fatal(err)
		}
		var j interface{}
		err = json.NewDecoder(resp.Body).Decode(&j)
		fmt.Println(j)
	},
}

func init() {
	rootCmd.AddCommand(pushmessageCmd)

}
