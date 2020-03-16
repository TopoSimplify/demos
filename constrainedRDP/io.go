package main

import (
	"bufio"
	"bytes"
	"github.com/intdxdt/fileutil"
	"github.com/intdxdt/geom"
	"github.com/naoina/toml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func readWKTInput(fname string) ([]geom.Coords, error) {
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

func readTomlInput(fname string, callback func(data map[string]geom.Coords)) error {
	var file, err = os.Open(fname)
	if err != nil {
		return err
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	var data = make(map[string]interface{})
	err = toml.Unmarshal(buf, &data)
	if err != nil {
		return err
	}

	var input = make(map[string]geom.Coords, len(data))
	for key, val := range data {
		input[key] = geom.CoordinatesFromArray(parseLinearArray(val))
	}
	callback(input)
	return err
}

func readTomlConstraints(fname string, callback func(data ConstToml)) error {
	var file, err = os.Open(fname)
	if err != nil {
		return err
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}
	var data = ConstToml{}
	if err = toml.Unmarshal(buf, &data); err != nil {
		return err
	}

	callback(data)
	return err
}

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

func readlinesFromReader(reader *bufio.Reader, callback func(lnStr string)) error {
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

func writeCoords(fname string, coords []geom.Coords, writer func(*geom.WKTParserObj) string) error {
	var baseDir, _ = filepath.Split(fname)
	var err = fileutil.MakeDirs(baseDir)
	if err != nil {
		return err
	}

	fid, err := os.Create(fname)
	if err != nil {
		return err
	}

	for _, o := range coords {
		_, err = fid.WriteString(writer(
			geom.NewWKTParserObj(geom.GeoTypeLineString, o),
		) + "\n")
		if err != nil {
			break
		}
	}
	return err
}

func writeTomlCoords(fname string, dict map[string][][]float64) error {
	var baseDir, _ = filepath.Split(fname)
	var err = fileutil.MakeDirs(baseDir)
	if err != nil {
		return err
	}

	fid, err := os.Create(fname)
	if err != nil {
		return err
	}
	out, err := toml.Marshal(dict)
	if err != nil {
		return err
	}
	_, err = fid.Write(out)
	return err
}

type ConstToml struct {
	Polylines map[string]interface{}
	Polygons  map[string]interface{}
	Points    map[string]interface{}
}

func (c *ConstToml) Geometries() []geom.Geometry {
	var geoms []geom.Geometry
	for _, o := range c.parsePoints() {
		geoms = append(geoms, o)
	}
	for _, o := range c.parsePolylines() {
		geoms = append(geoms, o)
	}
	for _, o := range c.parsePolygons() {
		geoms = append(geoms, o)
	}
	return geoms
}

func (c *ConstToml) parsePoints() []geom.Point {
	var pnts []geom.Point
	for _, pt := range c.Points {
		pnts = append(pnts, geom.CreatePoint(parsePointArray(pt)))
	}
	return pnts
}

func (c *ConstToml) parsePolylines() []*geom.LineString {
	var lns []*geom.LineString
	for _, ln := range c.Polylines {
		var pts = parseLinearArray(ln)
		lns = append(lns, geom.NewLineStringFromArray(pts))
	}
	return lns
}

func (c *ConstToml) parsePolygons() []*geom.Polygon {
	var polys []*geom.Polygon
	for _, lnarray := range c.Polygons {
		var coords []geom.Coords
		for _, shell := range lnarray.([]interface{}) {
			coords = append(coords, geom.CoordinatesFromArray(parseLinearArray(shell)))
		}
		polys = append(polys, geom.NewPolygon(coords...))
	}
	return polys
}
