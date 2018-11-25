package main

import (
	"fmt"
	"math"
	"sort"

	"gonum.org/v1/gonum/stat"
)

func main() {
	xs := []float64{
		32.32, 56.98, 21.52, 44.32,
		55.63, 13.75, 43.47, 43.34,
		12.34,
	}

	fmt.Printf("data: %v\n", xs)

	// calcula a média ponderada do conjunto de dados.
	// não temos pesos (ex: todos os pesos são 1)
	// então apenas passamos uma fatia nula.
	mean := stat.Mean(xs, nil)
	variance := stat.Variance(xs, nil)
	stddev := math.Sqrt(variance)

	// stat.Quantile precisa que a fatia de entrada seja classificada.
	sort.Float64s(xs)
	fmt.Printf("data: %v (sorted)\n", xs)

	// calcula a mediana do conjunto de dados.
	// aqui também, passamos uma fatia nula como pesos.
	median := stat.Quantile(0.5, stat.Empirical, xs, nil)

	fmt.Printf("média=          %v\n", mean)
	fmt.Printf("mediana=        %v\n", median)
	fmt.Printf("variância=      %v\n", variance)
	fmt.Printf("desvio padrão=  %v\n", stddev)
}
