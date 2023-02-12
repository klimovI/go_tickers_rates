package cmd

import (
	"encoding/json"
	"errors"
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
		price, err := fetchPairPrice(pair)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

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

func fetchPairPrice(pair string) (float64, error) {
	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, apiRatesUrl, nil)

	query := request.URL.Query()
	query.Add("pairs", pair)
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)

	if err != nil {
		return 0, err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	if response.StatusCode != http.StatusOK {
		return 0, errors.New(response.Status)
	}

	var data map[string]float64

	if err = json.Unmarshal(body, &data); err != nil {
		return 0, err
	}

	return data[pair], nil
}
