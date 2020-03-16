package main

import (
	"github.com/TopoSimplify/constdp"
	"github.com/TopoSimplify/opts"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/iter"
)

func simplifyInstances(lns []geom.Coords, opts *opts.Opts, constraints []geom.Geometry,
	offsetFn func(geom.Coords) (int, float64)) []geom.Coords {
	var id = iter.NewIgen()
	var forest []*constdp.ConstDP
	var junctions = make(map[int][]int, 0)

	for _, ln := range lns {
		forest = append(forest, constdp.NewConstDP(
			id.Next(), ln, constraints, opts, offsetFn,
		))
	}
	constdp.SimplifyInstances(id, forest, junctions)

	return extractSimpleSegs(forest)
}

func simplifyFeatureClass(lns []geom.Coords, opts *opts.Opts, constraints []geom.Geometry,
	offsetFn func(geom.Coords) (int, float64)) []geom.Coords {
	var id = iter.NewIgen()
	var forest []*constdp.ConstDP
	for _, ln := range lns {
		forest = append(forest, constdp.NewConstDP(
			id.Next(), ln, constraints, opts, offsetFn,
		))
	}

	constdp.SimplifyFeatureClass(id, forest, opts)
	return extractSimpleSegs(forest)
}

func extractSimpleSegs(forest []*constdp.ConstDP) []geom.Coords {
	var simpleCoords []geom.Coords
	for _, tree := range forest {
		var coords = tree.Coordinates().Clone()
		coords.Idxs = make([]int, 0, tree.SimpleSet.Size())
		for _, o := range tree.SimpleSet.Values() {
			coords.Idxs = append(coords.Idxs, o.(int))
		}
		simpleCoords = append(simpleCoords, coords)
	}
	return simpleCoords
}
