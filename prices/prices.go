package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.3f", taxIncludedPrice)
	}
	// result[taxRate] = taxIncludedPrices

	fmt.Println("result=>", result)
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("could not process file opening", err)
		return
	}
	scanner := bufio.NewScanner(file) //reading line by line
	var lines []string
	// false for no line
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Reading from file content failes")
		fmt.Println(err)
		file.Close()
		return
	}

	fmt.Println("Data from file read -> ", lines)

	prices := make([]float64, len(lines))

	for index, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64) // conversion
		if err != nil {
			fmt.Println("Converting price to float failure")
			fmt.Println(err)
			file.Close()
			return
		}

		prices[index] = floatPrice

	}

	job.InputPrices = prices
}
