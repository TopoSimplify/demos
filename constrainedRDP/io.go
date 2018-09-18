package main

import (
	"os"
	"bufio"
	"bytes"
	"strings"
	"github.com/intdxdt/geom"
)

func readConstraints(fname string) ([]geom.Geometry, error) {
	var geoms []geom.Geometry
	var file, err = os.Open(fname)
	if err != nil {
		return geoms, err
	}
	defer file.Close()
	err = readlinesFromReader(bufio.NewReader(file), func(line string) {
		var gtype = geom.WKTType(line)
		var acceptTypes = bytes.Equal(gtype, wktPoint) ||
			bytes.Equal(gtype, wktLinestring) || bytes.Equal(gtype, wktPolygon)
		if acceptTypes {
			geoms = append(geoms, geom.ReadGeometry(line))
		}
	})
	return geoms, err
}

func readPolyline(fname string) ([]geom.Coords, error) {
	var polylines []geom.Coords
	var file, err = os.Open(fname)
	if err != nil {
		return polylines, err
	}
	defer file.Close()

	var reader = bufio.NewReader(file)
	err = readlinesFromReader(reader, func(line string) {
		var gtype = geom.WKTType(line)
		if bytes.Equal(gtype, wktLinestring) {
			var obj = geom.ReadWKT(line, geom.GeoTypeLineString)
			polylines = append(polylines, obj.ToCoordinates()...)
		}
	})
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
