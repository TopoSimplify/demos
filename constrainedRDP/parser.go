package main

import (
	"fmt"
)

func parseLinearArray(array interface{}) [][]float64 {
	var ln [][]float64
	var row = array.([]interface{})
	for _, pt := range row {
		ln = append(ln, parsePointArray(pt))
	}
	return ln
}

func parsePointArray(pt interface{}) []float64 {
	var p []float64
	for _, o := range pt.([]interface{}) {
		p = append(p, parseFloat(o))
	}
	return p
}

func parseFloat(o interface{}) float64 {
	var v float64
	switch o.(type) {
	case float64:
		v = o.(float64)
	case float32:
		v = float64(o.(float32))
	case int64:
		v = float64(o.(int64))
	case int32:
		v = float64(o.(int32))
	default:
		panic(fmt.Sprintf("incompatible type: %T , for : %v", o, o))
	}
	return v
}
