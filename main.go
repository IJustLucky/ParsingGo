
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
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

		n := len(trades)
		var total float64
		sumArr := []float64{}
		for i := 0; i < n; i++ {
			totalArr := append(sumArr, price)
			total += totalArr[len(totalArr)-1]
			fmt.Println(total / float64(i))
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

func averagePrice(trades []trade, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		t := trades[i]
		total += t.price
	}
	return total / float64(n)
}

func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

/*func avgPrice(prices []float64, n int) float64 {
	var total float64
	for _, v := range prices {
		total += v
	}
	return total / float64(n)
}*/

func main() {
	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":8081", nil)
}

func httpserver(w http.ResponseWriter, _ *http.Request) {
	result := display("E:/ETHUSDTCSV/ETHUSDT-trades-2017-08.csv")
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title: "Displaying data in Golang",
		}))
	line.SetXAxis([]string{"time"}).
		AddSeries("Category A", result).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}

func display(path string) []opts.LineData {
	items := make([]opts.LineData, 0)
	trades, err := parseCSVFile(path)
	if err != nil {
		log.Fatal(err)
	}
	answer := averagePrice(trades, len(trades))
	items = append(items, opts.LineData{Value: answer})
	return items
}
