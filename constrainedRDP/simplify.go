package main

import (
	"github.com/intdxdt/geom"
	"github.com/intdxdt/iter"
	"github.com/TopoSimplify/opts"
	"github.com/TopoSimplify/constdp"
	"github.com/TopoSimplify/offset"
)

func simplifyInstances(lns []geom.Coords, opts *opts.Opts, constraints []geom.Geometry) []*geom.LineString {
	var id = iter.NewIgen()
	var forest []*constdp.ConstDP
	var junctions = make(map[int][]int, 0)

	for _, ln := range lns {
		forest = append(forest, constdp.NewConstDP(
			id.Next(), ln, constraints, opts, offset.MaxOffset,
		))
	}
	constdp.SimplifyInstances(id, forest, junctions)

	return extractSimpleSegs(forest)
}

func simplifyFeatureClass(lns []geom.Coords, opts *opts.Opts, constraints []geom.Geometry) []*geom.LineString {
	var id = iter.NewIgen()
	var forest []*constdp.ConstDP
	for _, ln := range lns {
		forest = append(forest, constdp.NewConstDP(
			id.Next(), ln, constraints, opts, offset.MaxOffset,
		))
	}

	constdp.SimplifyFeatureClass(id, forest, opts)
	return extractSimpleSegs(forest)
}

func extractSimpleSegs(forest []*constdp.ConstDP) []*geom.LineString {
	var simpleLns []*geom.LineString
	for _, tree := range forest {
		var coords = tree.Coordinates()
		coords.Idxs = make([]int, 0, tree.SimpleSet.Size())
		for _, o := range tree.SimpleSet.Values() {
			coords.Idxs = append(coords.Idxs, o.(int))
		}
		simpleLns = append(simpleLns, geom.NewLineString(coords))
	}
	return simpleLns
}
