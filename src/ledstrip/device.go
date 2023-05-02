package ledstrip

import (
	"sync"
	"time"
)

var lock = &sync.Mutex{}

type wsEngine interface {
	Init() error
	Render() error
	Wait() error
	Fini()
	Leds(channel int) []uint32
	SetBrightness(channel int, brightness int)
}

type ledDispEnum string
type modeEnum string

const (
	OnlyBack  ledDispEnum = "onlyback"
	OnlySides ledDispEnum = "onlysides"
	Full      ledDispEnum = "full"
)

const (
	Wave  modeEnum = "wave"
	Pulse modeEnum = "pulse"
)

type device struct {
	engine        wsEngine
	isInitialized bool
	ledDisp       ledDispEnum
	currentMode   modeEnum
}

func (dv *device) waveAnimation() error {
	startLed := 0
	endLed := len(dv.engine.Leds(0)) - 1
	switch dv.ledDisp {
	case OnlyBack:
		startLed = 20
		endLed -= 20
	}

	for led := startLed; led <= endLed; led++ {
		dv.engine.Leds(0)[led] = uint32(0x0000ff)
		// dv.engine.Leds(0)[led] = uint32(0xffffff)
		if err := dv.engine.Render(); err != nil {
			return err
		}
		time.Sleep(25 * time.Millisecond)
	}

	return nil
}

func (dv *device) breathingAnimation() error {
	startLed := 0
	endLed := len(dv.engine.Leds(0)) - 1
	switch dv.ledDisp {
	case OnlyBack:
		startLed = 80
		endLed -= 80
	}

	minBright := 100
	maxBright := 200
	color := uint32(0xffffff)
	for led := startLed; led <= endLed; led++ {
		dv.engine.Leds(0)[led] = color
	}

	for {
		for bright := minBright; bright <= maxBright; bright++ {
			dv.engine.SetBrightness(0, bright)
			if err := dv.engine.Render(); err != nil {
				return err
			}
			time.Sleep(time.Millisecond / 2)
		}
		for bright := maxBright; bright >= minBright; bright-- {
			dv.engine.SetBrightness(0, bright)
			if err := dv.engine.Render(); err != nil {
				return err
			}
			time.Sleep(time.Millisecond / 2)
		}
	}

	return nil
}

var deviceInstance *device

func GetDeviceInstance() *device {
	if deviceInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if deviceInstance == nil {
			deviceInstance = &device{}
		}
	}

	return deviceInstance
}
