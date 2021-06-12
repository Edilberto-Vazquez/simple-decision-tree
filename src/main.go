package main

import (
	"math"
)

func p_and_n(play_tennis []bool) (p, n float64, t int) {
	t = len(play_tennis)
	for _, v := range play_tennis {
		if v {
			p++
		} else {
			n++
		}
	}
	return p, n, t
}

func entropy(p, n float64, t int) (s float64) {
	t_S := float64(t)
	p = -(p / t_S) * (math.Log2(p / t_S))
	n = -(n / t_S) * (math.Log2(n / t_S))
	s = p + n
	return s
}

func main() {

	sky := []string{"Soleado", "Soleado", "Lluvioso", "Lluvioso", "Lluvioso", "Lluvioso", "Lluvioso", "Soleado", "Soleado", "Lluvioso", "Soleado", "Lluvioso", "Lluvioso", "Lluvioso"}
	sky_values := []string{"Soleado", "Lluvioso"}
	temp := []string{"Alta", "Alta", "Alta", "Media", "Fria", "Fria", "Fria", "Media", "Fria", "Media", "Media", "Media", "Alta", "Media"}
	temp_values := []string{"Alta", "Media", "Fria"}
	hum := []string{"Alta", "Alta", "Alta", "Alta", "Normal", "Normal", "Normal", "Alta", "Normal", "Normal", "Normal", "Alta", "Normal", "Alta"}
	hum_values := []string{"Alta", "Normal"}

	wind := []string{"Debil", "Fuerte", "Debil", "Debil", "Debil", "Fuerte", "Fuerte", "Debil", "Debil", "Debil", "Fuerte", "Fuerte", "Debil", "Fuerte"}
	wind_values := []string{"Debil", "Fuerte"}

	play_tennis := []bool{false, false, true, true, true, false, true, false, true, true, true, true, true, false}

}
