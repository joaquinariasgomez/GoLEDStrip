package ledstrip

import (
	"goledserver/src/constants"
	"goledserver/src/utils"
	"time"
)

// Cada animación, indiferentemente de si es una animación o una posición
// estática, debe acabar con la función staticFinalPartWaitToStop(), pera esperar
// a la siguiete animación pero que siga actualizándose con cosas como cambios
// en el brillo de los LEDs, etc.

const (
	WhiteColor     = uint32(0xffffff)
	WarmWhiteColor = uint32(0xfdf4dc)
	RedColor       = uint32(0xff0000)
	BlackColor     = uint32(0x000000)
)

/*=================== FUNCIONES PRIVADAS ===================*/

func (dv *device) staticFinalPartWaitToStop() {
	// Static final part of animation
	for {
		if dv.state == "stop" {
			break
		}

		// Update color
		for led := 0; led < len(dv.engine.Leds(0)); led++ {
			dv.engine.Leds(0)[led] = dv.currColor
		}

		// Render
		if err := dv.engine.Render(); err != nil {
			panic(err)
		}
		time.Sleep(time.Millisecond)
	}
}

func (dv *device) shutdownLights() {
	for bright := dv.currBrightness; bright >= 0; bright-- {
		dv.engine.SetBrightness(0, bright)
		if err := dv.engine.Render(); err != nil {
			panic(err)
		}
		time.Sleep(time.Millisecond)
	}
	for led := 0; led < len(dv.engine.Leds(0)); led++ {
		dv.engine.Leds(0)[led] = BlackColor
	}
	// Render
	if err := dv.engine.Render(); err != nil {
		panic(err)
	}
}

/*=================== ANIMACIONES ===================*/

func (dv *device) startupAnimation() {
	dv.state = "running"

	for led := 0; led < len(dv.engine.Leds(0)); led++ {
		if dv.state == "stop" {
			break
		}

		dv.engine.Leds(0)[led] = WhiteColor // All the startup animation will be in White color
		if err := dv.engine.Render(); err != nil {
			panic(err)
		}
		time.Sleep(time.Millisecond)
	}
	dv.staticFinalPartWaitToStop()
}

func (dv *device) officeLightsMode() {
	dv.state = "running"

	dv.currColor = WarmWhiteColor
	for led := 0; led < len(dv.engine.Leds(0)); led++ {
		dv.engine.Leds(0)[led] = dv.currColor
	}

	// Fade in part
	for bright := 50; bright <= constants.MAX_BRIGHTNESS; bright += 2 {
		dv.engine.SetBrightness(0, bright)
		if err := dv.engine.Render(); err != nil {
			panic(err)
		}
		time.Sleep(time.Millisecond / 10)
	}
	// Fade out
	for bright := constants.MAX_BRIGHTNESS - 1; bright >= dv.currBrightness; bright-- {
		dv.engine.SetBrightness(0, bright)
		if err := dv.engine.Render(); err != nil {
			panic(err)
		}
		time.Sleep(time.Millisecond)
	}

	dv.staticFinalPartWaitToStop()
}

func (dv *device) staticColorMode(args []string) {
	dv.state = "running"

	if len(args) > 0 {
		dv.currColor = utils.StringToUint32(args[0])
	}

	dv.staticFinalPartWaitToStop()
}

func (dv *device) rainbowBallsMode() {
	dv.state = "running"

	for it := 0; ; it++ {
		for led := 0; led < len(dv.engine.Leds(0)); led++ {
			dv.engine.Leds(0)[led] = utils.GetRainbowBallColor((led*256/constants.MAX_LEDS + it) & 255)
		}
		if dv.state == "stop" {
			break
		}

		if err := dv.engine.Render(); err != nil {
			panic(err)
		}
		time.Sleep(20 * time.Millisecond)
	}

	dv.staticFinalPartWaitToStop()
}

func (dv *device) rainbowContinuousMode() {
	dv.state = "running"

	for it := 0; ; it++ {
		for led := 0; led < len(dv.engine.Leds(0)); led++ {
			dv.engine.Leds(0)[led] = utils.GetRainbowBallColor((it + led) & 255)
		}
		if dv.state == "stop" {
			break
		}

		if err := dv.engine.Render(); err != nil {
			panic(err)
		}
		time.Sleep(20 * time.Millisecond)
	}

	dv.staticFinalPartWaitToStop()
}

func (dv *device) breathingAnimation() error {
	startLed := 0
	endLed := len(dv.engine.Leds(0)) - 1
	switch dv.ledDisposition {
	case OnlyBack:
		startLed = 80
		endLed -= 80
	}

	minBright := 100
	maxBright := 255
	color := uint32(0xffffff)
	// black := uint32(0x0)
	// otherColor := uint32(0xffff00)
	for led := startLed; led <= endLed; led++ {
		dv.engine.Leds(0)[led] = color
	}

	for {
		for bright := minBright; bright <= maxBright; bright++ {
			dv.engine.SetBrightness(0, bright)
			if err := dv.engine.Render(); err != nil {
				return err
			}
			time.Sleep(time.Millisecond)
		}
		for bright := maxBright; bright >= minBright; bright-- {
			dv.engine.SetBrightness(0, bright)
			if err := dv.engine.Render(); err != nil {
				return err
			}
			time.Sleep(time.Millisecond)
		}
	}

	return nil
}

func (dv *device) waveAnimation() error {
	startLed := 0
	endLed := len(dv.engine.Leds(0)) - 1
	switch dv.ledDisposition {
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
