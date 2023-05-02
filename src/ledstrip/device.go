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
}

type device struct {
	engine        wsEngine
	isInitialized bool
}

func (dv *device) waveAnimation() error {
	for led := 0; led < len(dv.engine.Leds(0)); led++ {
		dv.engine.Leds(0)[led] = uint32(0x0000ff)
		if err := dv.engine.Render(); err != nil {
			return err
		}
		time.Sleep(50 * time.Millisecond)
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
