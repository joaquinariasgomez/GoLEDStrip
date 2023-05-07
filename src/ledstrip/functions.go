package ledstrip

import (
	. "goledserver/src/constants"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

func WipeAction() {
	device := GetDeviceInstance()
	if !device.isInitialized {
		return
	}

	device.waveAnimation()
}

func TestAction(a Action) {
	StartDevice() // TODO: eliminar de aqui
	device := GetDeviceInstance()
	if !device.isInitialized {
		return
	}

	// Aqui iria la desencriptación del objeto "Action"
	// para enviarle a device el setup exacto

	device.testAnimation()
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
	// TODO: mover StartDevice() al startup sequence, donde solo sería llamado una vez
	// Esta función settea la configuración del singleton device, luego no hace falta llamarla por cada método
	StartDevice()
	WipeAction()
}

func ExamplePulsate() {
	StartDevice()
	PulsateAction()
}
