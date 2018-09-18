package main

import (
	"os"
	"io"
	"log"
	"fmt"
	"flag"
	"strings"
	"runtime"
	"github.com/intdxdt/geom"
	"github.com/TopoSimplify/opts"
)

var Output = "./out.txt"
var ConfigPath string
var wktPoint = []byte("point")
var wktPolygon = []byte("polygon")
var wktLinestring = []byte("linestring")

func init() {
	flag.StringVar(&ConfigPath, "c", "./config.toml", "Configuration file path")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	var cfg = readConfig()
	var options = optsFromCfg(cfg)
	var constraints []geom.Geometry
	var simpleCoords []geom.Coords

	//config simplification type
	cfg.SimplificationType = strings.ToLower(strings.TrimSpace(cfg.SimplificationType))
	var offsetFn = offsetDictionary[cfg.SimplificationType]
	if offsetFn == nil {
		log.Println(`Supported Simplification Types : "DP" or "SED", Fix config.toml file`)
		os.Exit(1)
	}

	var polyCoords, err = readPolyline(cfg.Input)
	if err != io.EOF {
		log.Println(fmt.Sprintf("Failed to read file: %v\nerror:%v\n", cfg.Input, err))
		os.Exit(1)
	}

	// config output
	cfg.Output = strings.TrimSpace(cfg.Output)
	if cfg.Output == "" {
		cfg.Output = Output
	}

	// config constraints
	cfg.Constraints = strings.TrimSpace(cfg.Constraints)
	if cfg.Constraints != "" {
		constraints, err = readConstraints(cfg.Constraints)
		if err != io.EOF {
			log.Println(fmt.Sprintf("Failed to read file: %v\nerror:%v\n", cfg.Input, err))
			os.Exit(1)
		}
	}

	// simplify
	if cfg.IsFeatureClass {
		simpleCoords = simplifyFeatureClass(polyCoords, &options, constraints, offsetFn)
	} else {
		simpleCoords = simplifyInstances(polyCoords, &options, constraints, offsetFn)
	}

	switch cfg.SimplificationType {
	case "dp":
		writeCoords(cfg.Output, simpleCoords, geom.WriteWKT)
	case "sed":
		writeCoords(cfg.Output, simpleCoords, geom.WriteWKT3D)
	}
}

func optsFromCfg(cfg Cfg) opts.Opts {
	return opts.Opts{
		Threshold:              cfg.Threshold,
		MinDist:                cfg.MinDist,
		RelaxDist:              cfg.RelaxDist,
		PlanarSelf:             cfg.PlanarSelf,
		AvoidNewSelfIntersects: cfg.AvoidNewSelfIntersects,
		GeomRelation:           cfg.GeomRelation,
		DistRelation:           cfg.DistRelation,
		DirRelation:            cfg.SideRelation,
	}
}
