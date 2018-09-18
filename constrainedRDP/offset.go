package main

import (
	"github.com/intdxdt/geom"
	"github.com/TopoSimplify/offset"
)

var offsetDictionary = map[string]func(geom.Coords) (int, float64){
	"dp":  offset.MaxOffset,
	"sed": offset.MaxSEDOffset,
}
