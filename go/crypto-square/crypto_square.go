package cryptosquare

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

const TestVersion = 1

func Encode(plain string) (result string) {
	// Remove all spaces and punctuation
	r := regexp.MustCompile("[^0-9A-Za-z_]")
	// Lowercase all characters
	plain = r.ReplaceAllString(strings.ToLower(plain), "")
	// Measure length of string
	length_of_plain := len(plain)
	// Find the square root of the lenght of string
	sqrt := math.Sqrt(float64(length_of_plain))
	// Assign length of rows and columns
	x := int(sqrt) - 1 // number of horizontal columns
	y := x             // number of vertical rows

	// if not a perfect square...
	if rem := math.Mod(sqrt, 1); rem != 0 {
		// fmt.Println(length_of_plain, int(sqrt), 1-rem)
		z := int(sqrt + 1)
		z = z * z
		z = (z - length_of_plain)
		for i := 0; i < z; i++ {
			plain += " "
			y += 1

		}
	}

	fmt.Println("pre-table-loop: ", plain)
	for i := 0; i <= y; i++ { // vertical
		for j := 0; j <= x; j++ { // horizontal
			a := string(plain[j+(y*j)+i])
			result += a
			fmt.Println(":: ", a)
		}
		if i != y {
			result += " "
		}
		fmt.Println("result(", i, "): ", result)

	}
	result = strings.Trim(result, " ")
	return result // plain
}
