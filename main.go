package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")
	myCanvas := myWindow.Canvas()

	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	rect := canvas.NewCircle(blue)
	myCanvas.SetContent(rect)

	// анонимная функция
	go func() {
		for i := 0; true; i++ {
			time.Sleep(time.Second)
			green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
			blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}

			if i % 2 == 0 {
				rect.FillColor = green
				rect.Refresh()	
			} else{
				rect.FillColor = blue
				rect.Refresh()
			}
		}
	}()

	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.ShowAndRun()
}