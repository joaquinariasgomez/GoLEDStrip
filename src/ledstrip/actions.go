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
	dev.ledDisposition = Full
	dev.currColor = uint32(0xff0000) // Red
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
	case ChangeMode:
		StartModeAction(a)
	case SetColor:
		SetColorAction(a)
	case SetBrightness:
		SetBrightnessAction(a)
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
	command := a.Command
	switch command.Instruction {
	case OfficeLights:
		device.officeLightsMode()
		// TODO: llamar con args: device.officeLightsMode(command.args)
	case StaticColor:
		device.staticColorMode(command.Args)
	}
}

func SetColorAction(a Action) {
	device := GetDeviceInstance()
	if !device.isInitialized {
		return
	}

	command := a.Command
	color := string(command.Instruction)
	device.setColor(color)
}

func SetBrightnessAction(a Action) {
	device := GetDeviceInstance()
	if !device.isInitialized {
		return
	}

	command := a.Command
	if command.Instruction == Decrease {
		device.decreaseBrightness(command.Args)
	} else if command.Instruction == Increase {
		device.increaseBrightness(command.Args)
	}
}

func SetStopState() {
	device := GetDeviceInstance()
	device.state = "stop"
}
