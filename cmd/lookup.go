package cmd

import (
	"io"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
)

var lookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "Look up an ActivityPub object",
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{}

		req, err := http.NewRequest("GET", args[0], nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Add("Accept", "application/activity+json")

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		colorfulJSON := pretty.Color(pretty.PrettyOptions(body, &pretty.Options{
			Width:  80,
			Prefix: "",
			Indent: "\t",
		}), pretty.TerminalStyle)

		color.New(color.FgHiGreen).Println(string(colorfulJSON))
	},
}

func init() {
	rootCmd.AddCommand(lookupCmd)
}
