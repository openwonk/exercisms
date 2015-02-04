package pythagorean

import "math"

type Triplet [3]int

func Range(min, max int) (result []Triplet) {
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			aSq, bSq := a*a, b*b
			c := math.Sqrt(float64(aSq) + float64(bSq))
			if math.Mod(c, 1) == 0 && int(c) <= max {
				result = append(result, Triplet{a, b, int(c)})
			}
		}
	}
	return
}

func Sum(p int) (result []Triplet) {
	for a := 1; a <= p; a++ {
		for b := a; b <= p; b++ {
			aSq, bSq := a*a, b*b
			c := math.Sqrt(float64(aSq) + float64(bSq))
			if math.Mod(c, 1) == 0 && (a+b+int(c)) == p {
				result = append(result, Triplet{a, b, int(c)})
			}
		}
	}
	return
}
