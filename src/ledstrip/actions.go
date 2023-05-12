package ledstrip

import (
	. "goledserver/src/constants"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

/* The responsibility of this class will be to decrypt the action and to talk with the device */

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
	dev.mode = Static
	dev.currBrightness = MAX_BRIGHTNESS

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

func StartAction(a Action) {
	switch a.Type {
	case Startup:
		StartDevice()
		device := GetDeviceInstance()
		device.startupAnimation()
	case SetMode:
		StartModeAction(a)
	case SetBrightness:
		StartBrightnessAction(a)
	}
}

func StartModeAction(a Action) {
	device := GetDeviceInstance()
	if !device.isInitialized {
		return
	}

	// Aqui iria la desencriptación del comando que se quiere ejecutar
	// Un comando es un JSON, con una serie de parámetros y árbol dentro
	// Por ahora, lets keep things simple y supongamos que es un string con una orden simple
	if a.Command == "office-lights" {
		device.staticOfficeLights()
	}
}

func StartBrightnessAction(a Action) {
	device := GetDeviceInstance()
	if !device.isInitialized {
		return
	}

	if a.Command == "decrease" {
		device.decreaseBrightness()
	} else if a.Command == "increase" {
		device.increaseBrightness()
	}
}

func WipeAction() {
	device := GetDeviceInstance()
	if !device.isInitialized {
		return
	}

	device.waveAnimation()
}

func SetStopState() {
	device := GetDeviceInstance()
	device.state = "stop"
}

func PulsateAction() {
	device := GetDeviceInstance()
	if !device.isInitialized {
		return
	}

	go device.breathingAnimation()
}

func ExampleWipe() {
	// TODO: mover StartDevice() al startup sequence, donde solo sería llamado una vez
	// Esta función settea la configuración del singleton device, luego no hace falta llamarla por cada método
	StartDevice()
	WipeAction()
}

func ExamplePulsate() {
	StartDevice()
	PulsateAction()
}
