package main

import (
	"runtime"
	"flag"
	"fmt"
	"github.com/TopoSimplify/opts"
	"bufio"
	"os"
	"github.com/intdxdt/geom"
	"bytes"
	"strings"
	"io"
	"log"
)

var ConfigPath string

var wktEmpty = []byte("empty")
var wktPolygon = []byte("polygon")
var wktLinestring = []byte("linestring")
var wktPoint = []byte("point")

func init() {
	flag.StringVar(&ConfigPath, "c", "./config.toml", "Configuration file path")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	var cfg = readConfig()
	//var config = optsFromCfg(cfg)
	var polyCoords, err = readPolyline(cfg.Input)
	if err != io.EOF {
		log.Println(fmt.Sprintf("Failed to read file: %v\nerror:%v\n", cfg.Input, err))
		os.Exit(1)
	}

	fmt.Println(polyCoords)
}

func optsFromCfg(cfg Cfg) *opts.Opts {
	return &opts.Opts{
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

func readPolyline(fname string) ([]geom.Coords, error) {
	var polylines []geom.Coords
	var file, err = os.Open(fname)
	if err != nil {
		return polylines, err
	}
	defer file.Close()

	var reader = bufio.NewReader(file)
	var transform = func(line string) {
		var gtype = geom.WKTType(line)
		if bytes.Equal(gtype, wktLinestring) {
			var obj = geom.ReadWKT(line, geom.GeoTypeLineString)
			polylines = append(polylines, obj.ToCoordinates()...)
		}
	}
	err = readlinesFromReader(reader, transform)
	return polylines, err
}

func readlinesFromReader(reader *bufio.Reader, callback func(lnStr string)) (error) {
	var err error
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)
		if len(line) > 0 {
			callback(line)
		}
	}
	return err
}
