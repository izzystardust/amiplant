package main

import (
	"fmt"
	"math"
	"time"
)

type component struct {
	turnover    time.Duration
	percentMass func(height, weight float64, g gender) float64
}

func id(p float64) func(float64, float64, gender) float64 {
	return func(a, b float64, c gender) float64 {
		return p
	}
}

// returns percent of total body mass contributed by plants in this component
func (c component) percentPlantMass(p person) float64 {
	percentReplaced := p.wentVegan.Hours() / c.turnover.Hours()
	fmt.Println("d:", p.wentVegan)
	fmt.Println("t:", c.turnover)
	fmt.Println("Percent replaced:", percentReplaced)
	return math.Min(percentReplaced, 1.0) * c.percentMass(p.height, p.weight, p.gender)

}

func fatfn(h, w float64, g gender) float64 {
	remaining := 1.0
	for t, v := range bits {
		if t != "fat" {
			remaining -= v.percentMass(h, w, g)
		}
	}
	return remaining
}

// skin from formed to shed is about a month
// https://www.ncbi.nlm.nih.gov/books/NBK26865/
var bits = map[string]*component{
	"skin": {
		turnover:    24 * 30 * time.Hour,
		percentMass: id(.16),
	},
	"skeleton": {
		turnover:    10 * 365 * 24 * time.Hour, // no source lol
		percentMass: id(.2),
	},
	"lean": {
		turnover: 5 * 365 * 24 * time.Hour, // lol sources
		percentMass: func(h, w float64, g gender) float64 {
			// https://www.ncbi.nlm.nih.gov/pmc/articles/PMC473290/pdf/jclinpath00363-0085.pdf
			var a, b, c float64
			if g == male {
				a = 0.32810
				b = 0.33929
				c = 29.5336
			} else {
				a = 0.29569
				b = 0.41813
				c = 43.2933
			}
			return ((a*w + b*h - c) / w) - .36
		},
	},
	"fat": {
		turnover:    10 * 365 * 24 * time.Hour,
		percentMass: nil,
	},
}
