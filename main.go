package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type trade struct {
	tradeId      int
	price        float64
	qty          float64
	quoteQty     float64
	time         int
	isBuyerMaker bool
	isbestMatch  bool
}

func parseCSVFile(path string) (trades []trade, _ error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitValues := strings.Split(line, ",")

		tradeIdStr := splitValues[0]
		tradeId, err := strconv.Atoi(tradeIdStr)
		if err != nil {
			return nil, err
		}

		pricestr := splitValues[1]
		price, err := strconv.ParseFloat(pricestr, 64)
		if err != nil {
			return nil, err
		}

		qtyStr := splitValues[2]
		qty, err := strconv.ParseFloat(qtyStr, 64)
		if err != nil {
			return nil, err
		}

		quoteQtyStr := splitValues[3]
		quoteQty, err := strconv.ParseFloat(quoteQtyStr, 64)
		if err != nil {
			return nil, err
		}

		timesStr := splitValues[4]
		times, err := strconv.Atoi(timesStr)
		if err != nil {
			return nil, err
		}

		isBuyerMakerStr := splitValues[5]
		isBuyerMaker, err := strconv.ParseBool(isBuyerMakerStr)
		if err != nil {
			return nil, err
		}

		isbestMatchStr := splitValues[6]
		isbestMatch, err := strconv.ParseBool(isbestMatchStr)
		if err != nil {
			return nil, err
		}

		trades = append(trades, trade{
			tradeId:      tradeId,
			price:        price,
			qty:          qty,
			quoteQty:     quoteQty,
			time:         times,
			isBuyerMaker: isBuyerMaker,
			isbestMatch:  isbestMatch,
		})
	}
	return trades, nil
}
