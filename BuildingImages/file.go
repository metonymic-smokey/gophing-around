package main

import (
	"encoding/csv"
	"flag"
	//"fmt"
	svg "github.com/ajstarks/svgo"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func rn(n int) int { return rand.Intn(n) }

func main() {
	canvas := svg.New(os.Stdout)
	var file string
	flag.StringVar(&file, "input", "input.csv", "input file")
	flag.Parse()

	csvfile, err := os.Open(file)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	r := csv.NewReader(csvfile)

	type mon struct {
		Month string
		Usage int
	}

	const n = 6
	dat := [n]mon{}
	i := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		dat[i].Month = record[0]
		x, _ := strconv.Atoi(record[1])
		dat[i].Usage = x
		i++

	}

	width := len(dat)*60 + 10
	height := 150
	threshold := 160
	max := 100

	for _, item := range dat {
		if item.Usage > max {
			max = item.Usage
		}
	}

	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:white")
	for i, val := range dat {
		percent := val.Usage * (height - 50) / max
		canvas.Rect(i*60+10, (height-50)-percent, 50, percent, "fill:rgb(77,200,232)")
		canvas.Text(i*60+35, height-24, val.Month, "font-size:14pt;fill:rgb(150, 150, 150);text-anchor:middle")
	}
	threshPercent := threshold * (height - 50) / max
	canvas.Line(0, height-threshPercent, width, height-threshPercent, "stroke: rgb(255,100,100); opacity: 0.8; stroke-width: 2px")
	canvas.Rect(0, 0, width, height-threshPercent, "fill:rgb(255, 100, 100); opacity: 0.1")
	canvas.Line(0, height-50, width, height-50, "stroke: rgb(150, 150, 150); stroke-width:2")

	canvas.End()
}
