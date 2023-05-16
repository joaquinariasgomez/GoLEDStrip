package utils

import (
	"github.com/go-playground/colors"
	"strconv"
	"strings"
)

func RgbToUint32(rgb *colors.RGBColor) uint32 {
	hex := rgb.ToHEX().String()
	cleanHex := hex[1:]
	ui32c, _ := strconv.ParseUint(cleanHex, 16, 32)
	return uint32(ui32c)
}

func StringToUint32(cStr string) uint32 {
	// Input for now is something like: "0xffffff"
	// TODO: support for inputs like: "rgb(255, 255, 255)"
	// Remove 0x prefix
	cleanC := strings.Replace(cStr, "0x", "", -1)
	ui32c, err := strconv.ParseUint(cleanC, 16, 32)
	if err != nil {
		panic(err)
	}
	return uint32(ui32c)
}

func GetRainbowBallColor(pos int) uint32 {
	if pos < 85 {
		rgb, err := colors.ParseRGB("rgb(" + strconv.Itoa(pos*3) + "," + strconv.Itoa(255-pos*3) + "," + strconv.Itoa(0) + ")")
		if err != nil {
			panic(err)
		}
		return RgbToUint32(rgb)
	} else {
		if pos < 170 {
			pos -= 85
			rgb, err := colors.ParseRGB("rgb(" + strconv.Itoa(255-pos*3) + "," + strconv.Itoa(0) + "," + strconv.Itoa(pos*3) + ")")
			if err != nil {
				panic(err)
			}
			return RgbToUint32(rgb)
		} else {
			pos -= 170
			rgb, err := colors.ParseRGB("rgb(" + strconv.Itoa(0) + "," + strconv.Itoa(pos*3) + "," + strconv.Itoa(255-pos*3) + ")")
			if err != nil {
				panic(err)
			}
			return RgbToUint32(rgb)
		}
	}
}
