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
		z := int(sqrt + 1)
		zz := z * z
		fmt.Println(length_of_plain, z, zz)

		z = (zz - length_of_plain)

		for i := 0; i < z; i++ {
			plain += " "
		}

		y += 1
		x += 1
	}

	fmt.Println("pre-table-loop: ", plain+"|")
	for i := 0; i <= y; i++ { // vertical
		fmt.Println("v", "h", "e", "-")

		for j := 0; j <= x; j++ { // horizontal
			e := j + (y * j) + i
			a := string(plain[e])
			result += a
			fmt.Println(i, j, e, a)
		}
		if i != y {
			result += " "
		}
		fmt.Println("result(", i, "): ", result)

	}
	r = regexp.MustCompile("[ ]+")
	result = r.ReplaceAllString(result, " ")
	// result = strings.Replace(result, "  ", " ", -1)
	result = strings.Trim(result, " ")
	return result // plain
}
