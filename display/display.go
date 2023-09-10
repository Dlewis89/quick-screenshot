package display

import (
	"image"
	"strconv"
	"time"

	"github.com/kbinani/screenshot"
)

type Display struct {
	name   string
	index  int
	bounds image.Rectangle
}

func (d *Display) GetName() string {
	return d.name
}

func (d *Display) GetBounds() image.Rectangle {
	return d.bounds
}

func GetDisplays() []Display {
	var displays []Display

	for i := 0; i < screenshot.NumActiveDisplays(); i++ {
		displays = append(displays, Display{"screen " + strconv.Itoa(i), i, screenshot.GetDisplayBounds(i)})
	}

	return displays
}

func CreateDisplayName() string {
	return "screenshot_" + strconv.FormatInt(time.Now().Unix(), 10) + ".png"
}

func GetDisplayNames(displays []Display) []string {
	var displayNames []string

	for i := 0; i < len(displays); i++ {
		displayNames = append(displayNames, displays[i].name)
	}

	return displayNames
}
