package main

import (
	"flag"
	"fmt"
	"github.com/TopoSimplify/opts"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/math"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
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

	var err error
	var keyIndx []string
	var polyCoords []geom.Coords

	if isWKTFile(cfg.Input) {
		polyCoords, err = readWKTInput(cfg.Input)
		if err != io.EOF {
			log.Println(fmt.Sprintf("Failed to read file: %v\nerror:%v\n", cfg.Input, err))
			os.Exit(1)
		}
		//for i := range polyCoords {
		//	keyIndx = append(keyIndx, KeyIndx{i, i})
		//}
	} else if isTomlFile(cfg.Input) {
		err = readTomlInput(cfg.Input, func(data map[string]geom.Coords) {
			var i = 0
			for key, coords := range data {
				keyIndx = append(keyIndx, key)
				polyCoords = append(polyCoords, coords)
				i++
			}
		})
		if err != nil {
			log.Println(fmt.Sprintf("Failed to read file: %v\nerror:%v\n", cfg.Input, err))
			os.Exit(1)
		}
	} else {
		panic("unknown file type, expects wkt/txt or toml")
	}

	// config output
	cfg.Output = strings.TrimSpace(cfg.Output)
	if cfg.Output == "" {
		cfg.Output = Output
	}

	// read constraints
	cfg.Constraints = strings.TrimSpace(cfg.Constraints)

	if cfg.Constraints != "" && isWKTFile(cfg.Constraints) {
		constraints, err = readConstraints(cfg.Constraints)
		if err != io.EOF {
			log.Println(fmt.Sprintf("Failed to read file: %v\nerror:%v\n", cfg.Constraints, err))
			os.Exit(1)
		}
	} else if cfg.Constraints != "" && isTomlFile(cfg.Constraints) {
		err = readTomlConstraints(cfg.Constraints, func(data ConstToml) {
			constraints = data.Geometries()
		})
		if err != nil {
			log.Println(fmt.Sprintf("Failed to read file: %v\nerror:%v\n", cfg.Constraints, err))
			os.Exit(1)
		}
	}

	// simplify
	log.Println("starting simplification ")
	var t0 = time.Now()
	if cfg.IsFeatureClass {
		simpleCoords = simplifyFeatureClass(polyCoords, &options, constraints, offsetFn)
	} else {
		simpleCoords = simplifyInstances(polyCoords, &options, constraints, offsetFn)
	}
	var t1 = time.Now()
	log.Println("done simplification ")
	log.Println(fmt.Sprintf("elapsed time: %v seconds", math.Round(t1.Sub(t0).Seconds(), 6)))

	var saved bool
	//Save output
	if isWKTFile(cfg.Input) {
		switch cfg.SimplificationType {
		case "dp":
			err = writeCoords(cfg.Output, simpleCoords, geom.WriteWKT)
		case "sed":
			err = writeCoords(cfg.Output, simpleCoords, geom.WriteWKT3D)
		}
		if err != nil {
			panic(err)
		}
		saved = true
	} else if isTomlFile(cfg.Input) {
		var coordinates [][][]float64
		var outputDict = make(map[string][][]float64, len(simpleCoords))
		var fn = func(dim int) {
			for i, simple := range simpleCoords {
				var ln [][]float64
				for _, idx := range simple.Idxs {
					ln = append(ln, simple.Pnts[idx][:dim])
				}
				coordinates = append(coordinates, ln)
				outputDict[keyIndx[i]] = ln
			}
		}

		switch cfg.SimplificationType {
		case "dp":
			fn(2)
		case "sed":
			fn(3)
		}

		if err = writeTomlCoords(cfg.Output, outputDict); err != nil {
			panic(err)
		}
		saved = true
	} else {
		panic("unknown file type, expects wkt/txt or toml")
	}
	if saved {
		log.Println("simplification save to file :", cfg.Output)
	}
}

func optsFromCfg(cfg Cfg) opts.Opts {
	return opts.Opts{
		Threshold:              cfg.Threshold,
		MinDist:                cfg.MinDist,
		RelaxDist:              cfg.RelaxDist,
		PlanarSelf:             cfg.PlanarSelf,
		NonPlanarSelf:          cfg.NonPlanarSelf,
		AvoidNewSelfIntersects: cfg.AvoidNewSelfIntersects,
		GeomRelation:           cfg.GeomRelation,
		DistRelation:           cfg.DistRelation,
		DirRelation:            cfg.SideRelation,
	}
}

func isTomlFile(fname string) bool {
	return strings.ToLower(filepath.Ext(fname)) == ".toml"
}

func isWKTFile(fname string) bool {
	var ext = strings.ToLower(filepath.Ext(fname))
	return ext == ".wkt" || ext == ".txt"
}
