package main

import "fmt"

func main() {
	prices := []float64{10.0, 20.0, 30.0}
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	var result = make(map[float64][]float64)

	for _, taxRate := range taxRates {
		taxIncludedPrices := make([]float64, len(prices))
		for priceIndex, price := range prices {
			taxIncludedPrices[priceIndex] = price * (1 + taxRate)
		}
		result[taxRate] = taxIncludedPrices
	}

	fmt.Println("result=>", result)
}
