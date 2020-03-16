package main

import (
	"bytes"
	"encoding/json"
	"github.com/intdxdt/fileutil"
	"github.com/naoina/toml"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//Opts
type Cfg struct {
	Input                  string  `json:"input"`
	Output                 string  `json:"output"`
	Constraints            string  `json:"constraints"`
	SimplificationType     string  `json:"simplificationtype"`
	Threshold              float64 `json:"threshold"`
	MinDist                float64 `json:"mindist"`
	RelaxDist              float64 `json:"relaxdist"`
	IsFeatureClass         bool    `json:"isfeatureclass"`
	PlanarSelf             bool    `json:"planarself"`
	NonPlanarSelf          bool    `json:"nonplanarself"`
	AvoidNewSelfIntersects bool    `json:"avoidself"`
	GeomRelation           bool    `json:"geomrelate"`
	DistRelation           bool    `json:"distrelate"`
	SideRelation           bool    `json:"siderelate"`
}

func (opt Cfg) String() string {
	var cfgbytes, err = json.Marshal(opt)
	if err != nil {
		panic(err)
	}
	return string(cfgbytes)
}

func readConfig() Cfg {
	f, err := os.Open(ConfigPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}
	buf = bytes.Replace(buf, []byte(`\`), []byte(`/`), -1)
	var config = Cfg{}
	if err := toml.Unmarshal(buf, &config); err != nil {
		log.Fatalln(err)
	}

	if !fileutil.IsFile(config.Input) {
		log.Println("input file not found")
		usageHelp()
		os.Exit(1)
	}

	if strings.TrimSpace(config.Constraints) != "" && !fileutil.IsFile(config.Constraints) {
		log.Println("constraints file not found")
		usageHelp()
		os.Exit(1)
	}
	return config
}
