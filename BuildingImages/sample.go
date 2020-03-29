package main

import (
	"math/rand"
	"os"
//	"net/http"
//	"log"
	svg "github.com/ajstarks/svgo"
)

func rn(n int) int { return rand.Intn(n) }

func main() {
	canvas := svg.New(os.Stdout)
	data := []struct {
		Month string
		Usage int
	}{
		{"Jan", 112},
		{"Feb", 145},
		{"Mar", 139},
		{"Apr", 127},
		{"May", 119},
		{"Jun", 117},
	}
	width := len(data)*60 + 10
	height := 150
	threshold := 160
	max := 0
	for _, item := range data {
		if item.Usage > max {
			max = item.Usage
		}
	}
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:white")
	for i, val := range data {
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
