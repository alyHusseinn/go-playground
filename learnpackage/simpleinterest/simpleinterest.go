package simpleinterest

import "fmt"

/*
 * init function added
 */
func init() {
	fmt.Println("Simple interest package initialized")
}

//Calculate calculates and returns the simple interest for a principal p, rate of interest r for time duration t years
func Calculate(p float64, r float64, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}

/*
We capitalized the function Calculate in the Simple interest package. 
This has a special meaning in Go. Any variable or function which starts 
with a capital letter are exported names in go. 
*/
