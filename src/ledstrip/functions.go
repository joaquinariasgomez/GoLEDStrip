package ledstrip

import (
	"fmt"

	. "goledserver/src/constants"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

const (
	sleepTime = 10
)

// type wsEngine interface {
// 	Init() error
// 	Render() error
// 	Wait() error
// 	Fini()
// 	Leds(channel int) []uint32
// }

// type colorWipe struct {
// 	ws wsEngine
// }

// func (cw *colorWipe) setup() error {
// 	return cw.ws.Init()
// }

// func (cw *colorWipe) display(color uint32) error {
// 	for i := 0; i < len(cw.ws.Leds(0)); i++ {
// 		cw.ws.Leds(0)[i] = color
// 		if err := cw.ws.Render(); err != nil {
// 			return err
// 		}
// 		time.Sleep(sleepTime * time.Millisecond)
// 	}
// 	return nil
// }

func WipeAction() {
	device := GetDeviceInstance()
	if !device.isInitialized {
		return
	}

	device.waveAnimation()
}

func PulsateAction() {
	device := GetDeviceInstance()
	if !device.isInitialized {
		return
	}

	go device.breathingAnimation()
}

func PrintDeviceStatus() {
	device := GetDeviceInstance()
	fmt.Println(device)
}

func StartDevice() {
	dev := GetDeviceInstance()
	if dev.isInitialized {
		return
	}

	opt := ws2811.DefaultOptions
	opt.Channels[0].Brightness = MAX_BRIGHTNESS
	opt.Channels[0].LedCount = MAX_LEDS

	engine, err := ws2811.MakeWS2811(&opt)
	if err != nil {
		panic(err)
	}

	// Default config for device
	dev.ledDisp = Full
	dev.engine = engine
	if err = dev.engine.Init(); err == nil {
		dev.isInitialized = true
	}
}

func ShutdownDevice() {
	device := GetDeviceInstance()
	if device.isInitialized {
		device.engine.Fini()
		device.isInitialized = false
	}
}

func ExampleWipe() {
	StartDevice()
	WipeAction()
	// ShutdownDevice()

	// fmt.Println("Starting LED Strip!")
	// opt := ws2811.DefaultOptions
	// opt.Channels[0].Brightness = brightness
	// opt.Channels[0].LedCount = ledCounts

	// dev, err := ws2811.MakeWS2811(&opt)
	// checkError(err)

	// cw := &colorWipe{
	// 	ws: dev,
	// }
	// checkError(cw.setup())
	// defer dev.Fini()

	// cw.display(uint32(0x0000ff))
	// cw.display(uint32(0x00ff00))
	// cw.display(uint32(0xff0000))
	// cw.display(uint32(0x000000))
}

func ExamplePulsate() {
	StartDevice()
	PulsateAction()
}
