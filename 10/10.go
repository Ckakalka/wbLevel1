package main

import "fmt"

// key=-20 includes [-30;-20)
// key=-0  includes [-10;0)
// key=0   includes [0;10)
// key=10  includes [10;20)
func getKeyDecade(value float64) string {
	valueInt := int(value)
	decade := (valueInt / 10) * 10
	decadeStr := fmt.Sprint(decade)
	if value < 0 {
		if -10 <= value && value < 0 {
			decadeStr = "-0"
			// is better to compare float numbers up to epsilon
		} else if float64(valueInt) == value && valueInt%10 == 0 {
			decade += 10
			decadeStr = fmt.Sprint(decade)
		}
	}
	return decadeStr
}

// Time complexity: O(n)
func main() {
	arr := []float64{-25.4, -20, -27.0, -10, -0.5, 0.5, 19.0, 15.5, 24.5, -21.0, 32.5}
	decadeToValues := make(map[string][]float64)
	for _, value := range arr {
		keyDecade := getKeyDecade(value)
		if _, ok := decadeToValues[keyDecade]; !ok {
			decadeToValues[keyDecade] = make([]float64, 0)
		}
		decadeToValues[keyDecade] = append(decadeToValues[keyDecade], value)
	}
	fmt.Println(decadeToValues)
}
