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
	RentedBikes         float64
	Hour                float64
	Temperature         float64
	Humidity            float64
	WindSpeed           float64
	Visibility          float64
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
		rentedBikes, _ := strconv.ParseFloat(line[1], 64)
		hour, _ := strconv.ParseFloat(line[2], 64)
		temperature, _ := strconv.ParseFloat(line[3], 64)
		humidity, _ := strconv.ParseFloat(line[4], 64)
		windSpeed, _ := strconv.ParseFloat(line[5], 64)
		visibility, _ := strconv.ParseFloat(line[6], 64)
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

	// Sort data into better formats for plots
	rentedBikesSlice := make([]float64, 64)
	hourSlice := make([]float64, 64)
	temperatureSlice := make([]float64, 64)
	humiditySlice := make([]float64, 64)
	windSpeedSlice := make([]float64, 64)
	visibilitySlice := make([]float64, 64)
	dewPointTemperatureSlice := make([]float64, 64)
	solarRadiationSlice := make([]float64, 64)
	rainfallSlice := make([]float64, 64)
	snowfallSlice := make([]float64, 64)
	for _, dataPoint := range data {
		rentedBikesSlice = append(rentedBikesSlice, dataPoint.RentedBikes)
		hourSlice = append(hourSlice, dataPoint.Hour)
		temperatureSlice = append(temperatureSlice, dataPoint.Temperature)
		humiditySlice = append(humiditySlice, dataPoint.Humidity)
		windSpeedSlice = append(windSpeedSlice, dataPoint.WindSpeed)
		visibilitySlice = append(visibilitySlice, dataPoint.Visibility)
		dewPointTemperatureSlice = append(dewPointTemperatureSlice, dataPoint.DewPointTemperature)
		solarRadiationSlice = append(solarRadiationSlice, dataPoint.SolarRadiation)
		rainfallSlice = append(rainfallSlice, dataPoint.Rainfall)
		snowfallSlice = append(snowfallSlice, dataPoint.Snowfall)
	}
	Slices := [][]float64{rentedBikesSlice, hourSlice, temperatureSlice, humiditySlice, windSpeedSlice, visibilitySlice,
		dewPointTemperatureSlice, solarRadiationSlice, rainfallSlice, snowfallSlice}

	// For each column of data, plot against another column of data
	scatterData := make(plotter.XYs, n)
	for indexA, sliceA := range Slices {
		var n = len(sliceA)
		for indexB, sliceB := range Slices {
			if indexA != indexB {
				for _, x := range sliceA {
					for _, y := range sliceB {

					}
				}
			}
		}
	}

	//========================================
	p := plot.New()
	p.X.Label.Text = "Rented Bikes"
	p.Y.Label.Text = "Temperature"
	p.Add(plotter.NewGrid())

	n = len(data)
	scatterData = make(plotter.XYs, n)
	for i := range scatterData {
		scatterData[i].X = float64(rentedBikesSlice[i])
		scatterData[i].Y = float64(temperatureSlice[i])
	}

	s, err := plotter.NewScatter(scatterData)
	if err != nil {
		log.Println(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}

	p.Add(s)
	p.Legend.Add("scatter", s)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "plots/plot.png"); err != nil {
		log.Println(err)
	}
}
