package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
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

		var prettyJSON bytes.Buffer
		error := json.Indent(&prettyJSON, body, "", "\t")
		if error != nil {
			log.Fatal("JSON parse error: ", error)
		}

		fmt.Println(prettyJSON.String())
	},
}

func init() {
	rootCmd.AddCommand(lookupCmd)
}
