package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Normalization struct {
	NormalizationAttribute []float64
	DecimalScalingConst    DecimalScalingConst
	ZScoreConst            ZScoreConst
	MinMaxConst            MinMaxConst
}

func NewNormalization() *Normalization {

	dataset := []float64{60, 20, 35, 21, 15, 55, 39}

	return &Normalization{
		NormalizationAttribute: dataset,
	}
}

type DecimalScalingConst struct {
	MaxDecimalLength float64
}

func DecimalScaling(param DecimalScalingConst, value float64) (vAccent float64) {

	vAccent = value / param.MaxDecimalLength

	return vAccent
}

func countDecimalDigits(n float64) int {
	// Convert the float to a string to handle cases like 4.0
	str := strconv.FormatFloat(n, 'f', -1, 64)

	// Split the string at the decimal point
	parts := strings.Split(str, ".")

	// Return the number of digits after the decimal point
	return len(parts[0])
}

func (n *DecimalScalingConst) CalculateMaxDecimalLength(dataset []float64) {

	var maxdecimallength float64

	c := countDecimalDigits(dataset[0])
	fmt.Println(c)

	switch c {
	case 1:
		maxdecimallength = 10
	case 2:
		maxdecimallength = 100
	default:
		maxdecimallength = 10
	}

	for _, v := range dataset {
		c := countDecimalDigits(v)

		switch c {
		case 1:
			maxdecimallength = 10
		case 2:
			maxdecimallength = 100
		default:
			maxdecimallength = 10
		}
	}

	n.MaxDecimalLength = maxdecimallength
}

func (n *Normalization) DecimalScalingNormalization() {

	n.DecimalScalingConst.CalculateMaxDecimalLength(n.NormalizationAttribute)

	fmt.Printf("max decimal digit: %#v\n", n.DecimalScalingConst.MaxDecimalLength)

	fmt.Println("decimal scaling normalization (before | after)")
	for i, value := range n.NormalizationAttribute {
		vAccent := DecimalScaling(n.DecimalScalingConst, value)
		fmt.Printf("%#v. attribute: %#v | %#v\n", i+1, value, vAccent)
	}

}

type ZScoreConst struct {
	Mean float64 // rata-rata
	STD  float64 // standard deviasi
}

func ZScore(param ZScoreConst, value float64) (vAccent float64) {
	vAccent = (value - param.Mean) / param.STD

	return vAccent
}

func (n *ZScoreConst) CalculateMean(dataset []float64) {

	var temp float64
	for _, v := range dataset {
		temp += v
	}
	n.Mean = temp / float64(len(dataset))
}

func (n *ZScoreConst) CalculateSTD(dataset []float64) {

	var temp float64

	for _, v := range dataset {
		temp += (v - n.Mean) * (v - n.Mean)
	}

	n.STD = math.Sqrt(temp / float64(len(dataset)-1))
}

func (n *Normalization) ZScoreNormalization() {

	// mean
	n.ZScoreConst.CalculateMean(n.NormalizationAttribute)
	n.ZScoreConst.CalculateSTD(n.NormalizationAttribute)

	fmt.Printf("mean: %#v\n", n.ZScoreConst.Mean)
	fmt.Printf("std: %#v\n", n.ZScoreConst.STD)

	fmt.Println("z-score normalization (before | after)")
	for i, value := range n.NormalizationAttribute {
		vAccent := ZScore(n.ZScoreConst, value)
		fmt.Printf("%#v. attribute: %#v | %#v\n", i+1, value, vAccent)
	}
}

type MinMaxConst struct {
	minA     float64
	maxA     float64
	new_minA float64
	new_maxA float64
}

func MinMax(param MinMaxConst, value float64) (vAccent float64) {

	var (
		v    float64
		minA float64
		maxA float64

		new_maxA float64
		new_minA float64
	)

	v = value
	minA = param.minA
	maxA = param.maxA

	new_minA = param.new_minA
	new_maxA = param.new_maxA

	vAccent = (((v - minA) / (maxA - minA)) * (new_maxA - new_minA)) + new_minA

	return vAccent
}

func (n *MinMaxConst) calculateMinValue(dataset []float64) {

	var minvalue float64
	minvalue = dataset[0]
	fmt.Println(minvalue)

	for _, value := range dataset {
		if value <= minvalue {
			fmt.Printf("%#v is < %#v, therefore updated.\n", value, minvalue)
			minvalue = value
		}
	}
	n.minA = minvalue
}

func (n *MinMaxConst) calculateMaxValue(dataset []float64) {

	var maxvalue float64
	maxvalue = dataset[0]

	for _, value := range dataset {
		if value >= maxvalue {
			maxvalue = value
		}
	}
	n.maxA = maxvalue
}

func (n *Normalization) MinMaxNormalization() {

	n.MinMaxConst.calculateMinValue(n.NormalizationAttribute)
	n.MinMaxConst.calculateMaxValue(n.NormalizationAttribute)

	fmt.Printf("minvalue: %#v\n", n.MinMaxConst.minA)
	fmt.Printf("maxvalue: %#v\n", n.MinMaxConst.maxA)

	fmt.Println("min max normalization (before | after)")
	for i, value := range n.NormalizationAttribute {
		vAccent := MinMax(n.MinMaxConst, value)
		fmt.Printf("%#v. attribute: %#v | %#v\n", i+1, value, vAccent)
	}
}

func RunNormalization() {

	/* init dataset */
	var dataset []float64
	// dataset = []float64{60, 20, 35, 21, 15, 55, 39}
	dataset = []float64{1.1, 8.2, 4.2, 1.5, 7.6, 2, 3.9}

	/* decimal scaling const */
	dsc := DecimalScalingConst{}
	/* z-score const */
	zsc := ZScoreConst{}

	/* min max const */
	// var (
	// 	minvalue    float64 = 15
	// 	maxvalue    float64 = 60
	// 	newMinValue float64 = 56
	// 	newMaxValue float64 = 90
	// )

	var (
		newMinValue float64 = 56
		newMaxValue float64 = 90
	)

	mmc := MinMaxConst{
		new_minA: newMinValue,
		new_maxA: newMaxValue,
	}

	normz := Normalization{
		NormalizationAttribute: dataset,

		DecimalScalingConst: dsc,
		ZScoreConst:         zsc,
		MinMaxConst:         mmc,
	}

	normz.DecimalScalingNormalization()
	// normz.ZScoreNormalization()
	// normz.MinMaxNormalization()

}
