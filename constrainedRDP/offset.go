package main

import (
	"github.com/TopoSimplify/offset"
	"github.com/intdxdt/geom"
)

var offsetDictionary = map[string]func(geom.Coords) (int, float64){
	"dp":  offset.MaxOffset,
	"sed": offset.MaxSEDOffset,
}
