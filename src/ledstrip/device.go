package ledstrip

import (
	"goledserver/src/constants"
	"sync"
)

var deviceLock = &sync.Mutex{}

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
	OnlyBack ledDispEnum = "onlyback"
	Full     ledDispEnum = "full"
)

const (
	Static modeEnum = "static"
	Pulse  modeEnum = "pulse"
	Wave   modeEnum = "wave"
)

type device struct {
	engine         wsEngine
	isInitialized  bool
	state          string
	ledDisp        ledDispEnum
	mode           modeEnum
	currBrightness int
}

func (dv *device) decreaseBrightness() {
	// This will decrease brightness by 50, for example
	decreaseAmount := 50
	if dv.currBrightness-decreaseAmount <= 0 {
		dv.currBrightness = 0
	} else {
		dv.currBrightness -= decreaseAmount
	}
	dv.engine.SetBrightness(0, dv.currBrightness)
}

func (dv *device) increaseBrightness() {
	// This will increase brightness by 50, for example
	increaseAmount := 50
	if dv.currBrightness+increaseAmount >= constants.MAX_BRIGHTNESS {
		dv.currBrightness = constants.MAX_BRIGHTNESS
	} else {
		dv.currBrightness += increaseAmount
	}
	dv.engine.SetBrightness(0, dv.currBrightness)
}

var deviceInstance *device

func GetDeviceInstance() *device {
	if deviceInstance == nil {
		deviceLock.Lock()
		defer deviceLock.Unlock()

		if deviceInstance == nil {
			deviceInstance = &device{}
		}
	}

	return deviceInstance
}
