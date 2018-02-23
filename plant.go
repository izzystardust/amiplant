package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type gender int

const (
	undef gender = iota
	male
	female
)

type person struct {
	weight    float64
	height    float64
	wentVegan time.Duration
	gender    gender
}

func main() {
	bits["fat"].percentMass = fatfn
	w, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("wat weight", err)
		panic(err)
	}
	h, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("wat height", err)
		panic(err)
	}
	v, err := time.Parse("2006-01-02", os.Args[3])
	if err != nil {
		panic(err)
	}
	p := person{
		weight:    w,
		height:    h,
		wentVegan: time.Since(v),
		gender:    male,
	}
	if os.Args[4][0] == 'f' {
		p.gender = female
	}
	sum := 0.0
	for b, c := range bits {
		fmt.Println("Calculating ", b)
		sum += c.percentPlantMass(p)
	}

	fmt.Println("You are", sum*100, "percent plant")
}
