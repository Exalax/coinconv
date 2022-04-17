package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/Exalax/coinconv/internal/coinmarketcap"
	"github.com/Exalax/coinconv/internal/plugconfig"
)

const usage = `Usage: coinconv amount from to
For example: coinconv 123.45 USD BTC`

func main() {
	flag.Parse()
	if len(flag.Args()) != 3 {
		fmt.Println(usage)
		os.Exit(1)
	}
	amount, err := strconv.ParseFloat(flag.Arg(0), 64)
	if err != nil {
		fmt.Println("Amount is not a number:", flag.Arg(0))
		os.Exit(1)
	}

	ctx := context.Background()
	config := plugconfig.New()
	client := coinmarketcap.New(config.CoinMarketCapURL, config.CoinMarketCapKey)

	ratio, err := client.Convert(ctx, amount, flag.Arg(1), flag.Arg(2))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println("OK:", ratio)
}
