package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"io/fs"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from backend!")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "API")
	})

	e.GET("/api/v1", func(c echo.Context) error {
		// コマンドの実行
		makePlotFile()
		return c.String(http.StatusOK, string("output"))
	})

	e.Logger.Fatal(e.Start(":8085"))
}

func makePlotFile() {
	currentTime := time.Now().String()
	removeChar := []string{":", " ", "+"}
	for _, char := range removeChar {
		currentTime = strings.Replace(currentTime, char, "_", -1)
	}
	script := `set terminal pngcairo enhanced font 'Arial,12'
    set output '/app/images/` + currentTime + `_plot.png'
    
    set title "Sucrose Concentration vs. Refractive Index"
    set xlabel "Concentration (wt%)"
    set ylabel "Refractive Index"
    
    set style data linespoints
    set key top left
    
    plot "-" using 1:2 title "Data" with points pt 7 lc rgb "blue"
    0.00 1.3331
    9.44 1.3469
    20.41 1.3649
    31.47 1.3845
    40.63 1.4012
    e
    `
	err := os.WriteFile("/app/scripts/"+currentTime+".plt", []byte(script), fs.FileMode(0644))
	if err != nil {
		log.Printf("ファイルの書き込み中にエラーが発生しました: %v", err)
	}
}
