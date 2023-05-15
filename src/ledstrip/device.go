package ledstrip

import (
	"goledserver/src/constants"
	"strconv"
	"strings"
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

const (
	OnlyBack ledDispEnum = "onlyBack"
	Full     ledDispEnum = "full"
)

type device struct {
	engine         wsEngine
	isInitialized  bool
	state          string
	ledDisposition ledDispEnum
	currColor      uint32
	currBrightness int
}

func (dv *device) setColor(c string) {
	// Remove 0x preffix
	cleanC := strings.Replace(c, "0x", "", -1)
	ui32c, err := strconv.ParseUint(cleanC, 16, 32)
	if err != nil {
		panic(err)
	}
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
