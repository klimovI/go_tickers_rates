package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

const pairFlag = "pair"

var rateCmd = &cobra.Command{
	Use:   "rate",
	Short: "Get price of given pair",
	Run: func(cmd *cobra.Command, args []string) {
		pair, _ := cmd.Flags().GetString(pairFlag)
		price := fetchPairPrice(pair)
		fmt.Println(price)
	},
}

func init() {
	rootCmd.AddCommand(rateCmd)

	rateCmd.PersistentFlags().String(pairFlag, "", "Pair price to get")
	if err := rateCmd.MarkPersistentFlagRequired(pairFlag); err != nil {
		log.Fatalln(err)
	}
}

const apiRatesUrl = "http://localhost:3001/api/v1/rates"

func fetchPairPrice(pair string) float64 {
	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, apiRatesUrl, nil)

	query := request.URL.Query()
	query.Add("pairs", pair)
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var data map[string]float64

	if err = json.Unmarshal(body, &data); err != nil {
		log.Fatalln(err)
	}

	return data[pair]
}
