package ledstrip

import (
	"fmt"
	"goledserver/src/constants"
	"strconv"
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
	OnlyBack ledDispEnum = "onlyBack"
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
	ledDisposition ledDispEnum
	mode           modeEnum
	currColor      uint32
	currBrightness int
}

func (dv *device) setColor(c string) {
	ui32c, err := strconv.ParseUint(c, 10, 32)
	if err != nil {
		panic(err)
	}
	fmt.Println("String %v convertido a color es %v", c, uint32(ui32c))
	dv.currColor = uint32(ui32c)
}

func (dv *device) decreaseBrightness(args []string) {
	// This will decrease brightness by the ammount specified in args
	decAmount := 25
	if len(args) > 0 {
		var err error
		decAmount, err = strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
	}

	if dv.currBrightness-decAmount <= 0 {
		dv.currBrightness = 0
	} else {
		dv.currBrightness -= decAmount
	}
	dv.engine.SetBrightness(0, dv.currBrightness)
}

func (dv *device) increaseBrightness(args []string) {
	// This will increase brightness by the ammount specified in args
	incAmount := 25
	if len(args) > 0 {
		var err error
		incAmount, err = strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
	}

	if dv.currBrightness+incAmount >= constants.MAX_BRIGHTNESS {
		dv.currBrightness = constants.MAX_BRIGHTNESS
	} else {
		dv.currBrightness += incAmount
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
