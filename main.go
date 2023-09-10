package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/kbinani/screenshot"
)

type Display struct {
	name   string
	index  int
	bounds image.Rectangle
}

func getDisplayNames(displays []Display) []string {
	var displayNames []string

	for i := 0; i < len(displays); i++ {
		displayNames = append(displayNames, displays[i].name)
	}

	return displayNames
}

func takeScreenshot(displays []Display, selectedScreen string) {
	for i := 0; i < len(displays); i++ {
		if displays[i].name != selectedScreen {
			continue
		}

		img, err := screenshot.CaptureRect(displays[i].bounds)

		if err != nil {
			log.Fatal("Unable to get bounds of requested display", err)
		}

		fileName := "screenshot_" + strconv.FormatInt(time.Now().Unix(), 10) + ".png"

		file, err := os.Create(fileName)

		if err != nil {
			log.Fatal("Unable to create file", err)
		}

		defer file.Close()

		png.Encode(file, img)

	}
}

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Quick ScreenShot")

	myWindow.Resize(fyne.NewSize(300, 200))

	var displays []Display

	for i := 0; i < screenshot.NumActiveDisplays(); i++ {
		displays = append(displays, Display{"screen " + strconv.Itoa(i), i, screenshot.GetDisplayBounds(i)})
	}

	displaySelector := widget.NewSelect(getDisplayNames(displays), func(value string) {
		fmt.Println("Screen " + value + " selected")
	})

	screenshotButton := widget.NewButton("Take screenshot", func() {
		takeScreenshot(displays, displaySelector.Selected)
	})

	displayOptionsListContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), displaySelector, layout.NewSpacer())

	takeScreenShotButtonContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), screenshotButton, layout.NewSpacer())

	myWindow.SetContent(container.New(layout.NewVBoxLayout(), layout.NewSpacer(), displayOptionsListContainer, layout.NewSpacer(), takeScreenShotButtonContainer, layout.NewSpacer()))

	myWindow.ShowAndRun()
}
