package handler

import (
	"errors"
	"fmt"
	"os"
)

type (
	Figure struct {
		DataFileName  string  `json:"data_file_name"`
		Title         string  `json:"title"`
		TitleSize     int     `json:"title_size"`
		XLabel        string  `json:"x_label"`
		YLabel        string  `json:"y_label"`
		ImageSize     float64 `json:"image_size"`
		Font          string  `json:"font"`
		KeyValue      string  `json:"key_value"`
		KeyPosition   string  `json:"key_position"`
		PointsOrLine  string  `json:"points_or_line"`
		PointsSize    int     `json:"points_size"`
		PointType     string  `json:"point_type"`
		LineWidth     string  `json:"line_width"`
		PlotLineColor string  `json:"plot_line_color"`
		XTics         bool    `json:"xtics"`
		YTics         bool    `json:"ytics"`
		Grid          bool    `json:"grid"`
		XRangeStart   string  `json:"x_range_start"`
		XRangeEnd     string  `json:"x_range_end"`
		YRangeStart   string  `json:"y_range_start"`
		YRangeEnd     string  `json:"y_range_end"`
		PlotScript    string  `json:"plot_script"`
	}
)

func makeScript(fileName string, figure *Figure) (string, string, error) {
	outputPath := "/app/images/" + fileName + "_plot.png"
	if figure.DataFileName == "" {
		return "", "", errors.New("no data")
	} else {
		// Check if figure.DataFileName exists in /app/data
		dataFilePath := "/app/data/" + figure.DataFileName
		if _, err := os.Stat(dataFilePath); os.IsNotExist(err) {
			return "", "", errors.New("data file does not exist")
		}
	}
	if figure.ImageSize == 0 {
		figure.ImageSize = 1
	} else if figure.ImageSize > 7 {
		figure.ImageSize = 7
	}
	if figure.TitleSize == 0 {
		figure.TitleSize = 10
	}
	if figure.KeyPosition == "" {
		figure.KeyPosition = "top left"
	}
	if figure.PlotLineColor == "" {
		figure.PlotLineColor = "#000000"
	}
	if figure.PointsOrLine == "" {
		figure.PointsOrLine = "points"
	}
	if figure.PointsSize == 0 {
		figure.PointsSize = 1
	}
	if figure.PointType == "" {
		figure.PointType = "7"
	}
	if figure.LineWidth == "" {
		figure.LineWidth = "1"
	}
	xTics := ""
	yTics := ""
	if figure.XTics {
		xTics = "#"
	}
	if figure.YTics {
		yTics = "#"
	}
	grid := ""
	if !figure.Grid {
		grid = "#"
	}
	isPoint := `pt ` + figure.PointType + ` pointsize ` + fmt.Sprint(figure.PointsSize) + `*image_size linewidth ` + fmt.Sprint(figure.LineWidth) + `*image_size lc rgb "` + figure.PlotLineColor + `"`
	isLine := `linecolor "` + figure.PlotLineColor + `" linewidth ` + figure.LineWidth + `*image_size`
	if figure.PointsOrLine == "points" {
		isLine = ""
	}
	if figure.PointsOrLine == "lines" {
		isPoint = ""
	}
	if figure.PointsOrLine == "linespoints" {
		isPoint = `pt ` + figure.PointType + ` pointsize ` + fmt.Sprint(figure.PointsSize) + `*image_size lc rgb "` + figure.PlotLineColor + `"`
		isLine = `linewidth ` + figure.LineWidth + `*image_size`
	}
	script := `
	set encoding utf8

	image_size = ` + fmt.Sprint(figure.ImageSize) + `
	set terminal pngcairo enhanced font "arial,".(12*image_size) size 640*image_size,480*image_size # PNG形式で出力
	set output '` + outputPath + `'

	set title font ",".(` + fmt.Sprint(figure.TitleSize) + `*image_size)
	set border linewidth 1*image_size
	
	set title '` + figure.Title + `'
	set xlabel "` + figure.XLabel + `"
	set ylabel "` + figure.YLabel + `"

	` + xTics + `unset xtics
	` + yTics + `unset ytics
	` + grid + `set grid

	set xrange [` + figure.XRangeStart + `:` + figure.XRangeEnd + `]
	set yrange [` + figure.YRangeStart + `:` + figure.YRangeEnd + `]
	
	set style data linespoints
	set key ` + figure.KeyPosition + `
	
	plot "../data/` + figure.DataFileName + `" using 1:2 title "` + figure.KeyValue + `" with ` + figure.PointsOrLine + ` ` + isPoint + ` ` + isLine
	return script, outputPath, nil
}
