package main

import (
	"encoding/csv"
	"image/color"

	// deep "github.com/patrikeh/go-deep"
	// "github.com/patrikeh/go-deep/training"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type dataPoint struct {
	Date                string
	RentedBikes         int
	Hour                int
	Temperature         float64
	Humidity            int
	WindSpeed           float64
	Visibility          int
	DewPointTemperature float64
	SolarRadiation      float64
	Rainfall            float64
	Snowfall            float64
	Season              string
	Holiday             string
	FunctioningDay      string
}

type dataSet []dataPoint

func main() {
	data := loadData()
	plotData(data)
}

func loadData() dataSet {
	csvFile, err := os.Open("SeoulBikeData.csv")
	if err != nil {
		log.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Println(err)
	}

	var data dataSet
	for _, line := range csvLines {
		if line[13] == "No" {
			continue
		}
		rentedBikes, _ := strconv.Atoi(line[1])
		hour, _ := strconv.Atoi(line[2])
		temperature, _ := strconv.ParseFloat(line[3], 64)
		humidity, _ := strconv.Atoi(line[4])
		windSpeed, _ := strconv.ParseFloat(line[5], 64)
		visibility, _ := strconv.Atoi(line[6])
		dewPointTemperature, _ := strconv.ParseFloat(line[7], 64)
		solarRadiation, _ := strconv.ParseFloat(line[8], 64)
		rainfall, _ := strconv.ParseFloat(line[9], 64)
		snowfall, _ := strconv.ParseFloat(line[10], 64)
		dataPoint := dataPoint{
			Date:                line[0],
			RentedBikes:         rentedBikes,
			Hour:                hour,
			Temperature:         temperature,
			Humidity:            humidity,
			WindSpeed:           windSpeed,
			Visibility:          visibility,
			DewPointTemperature: dewPointTemperature,
			SolarRadiation:      solarRadiation,
			Rainfall:            rainfall,
			Snowfall:            snowfall,
			Season:              line[11],
			Holiday:             line[12],
			FunctioningDay:      line[13],
		}
		// fmt.Println(dataPoint.Date, dataPoint.RentedBikes, dataPoint.Hour, dataPoint.Temperature, dataPoint.Humidity, dataPoint.WindSpeed,
		// 	dataPoint.Visibility, dataPoint.DewPointTemperature, dataPoint.SolarRadiation, dataPoint.Rainfall, dataPoint.Snowfall,
		// 	dataPoint.Season, dataPoint.Holiday, dataPoint.FunctioningDay)
		data = append(data, dataPoint)
	}
	return data
}

func plotData(data dataSet) {
	p := plot.New()
	p.X.Label.Text = "Rented Bikes"
	p.Y.Label.Text = "Wind Speed"
	p.Add(plotter.NewGrid())

	var n = len(data)
	log.Println(n)
	scatterData := make(plotter.XYs, n)
	for i := range scatterData {
		scatterData[i].X = float64(data[i].RentedBikes)
		scatterData[i].Y = float64(data[i].WindSpeed)
	}

	s, err := plotter.NewScatter(scatterData)
	if err != nil {
		log.Println(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}

	p.Add(s)
	p.Legend.Add("scatter", s)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "plot.png"); err != nil {
		log.Println(err)
	}
}
