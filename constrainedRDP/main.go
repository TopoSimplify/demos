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

	var polyCoords, err = readPolyline(cfg.Input)
	if err != io.EOF {
		log.Println(fmt.Sprintf("Failed to read file: %v\nerror:%v\n", cfg.Input, err))
		os.Exit(1)
	}

	var constraints []geom.Geometry
	cfg.Constraints = strings.TrimSpace(cfg.Constraints)
	if cfg.Constraints != "" {
		constraints, err = readConstraints(cfg.Constraints)
		if err != io.EOF {
			log.Println(fmt.Sprintf("Failed to read file: %v\nerror:%v\n", cfg.Input, err))
			os.Exit(1)
		}
	}

	var outputLns []*geom.LineString

	if cfg.IsFeatureClass {
		outputLns = simplifyFeatureClass(polyCoords, &options, constraints)
	} else {
		outputLns = simplifyInstances(polyCoords, &options, constraints)
	}
	for _, o := range outputLns {
		fmt.Println(o.WKT())
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
