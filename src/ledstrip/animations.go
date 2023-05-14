package ledstrip

import "time"

// Cada animación, indiferentemente de si es una animación o una posición
// estática, debe acabar con la función staticFinalWaitToStop(), pera esperar
// a la siguiete animación pero que siga actualizándose con cosas como cambios
// en el brillo de los LEDs, etc.

const (
	WarmWhiteColor = uint32(0xfdf4dc)
)

/*=================== FUNCIONES PRIVADAS ===================*/

func (dv *device) staticFinalWaitToStop() {
	// Static final part of animation
	for {
		if dv.state == "stop" {
			break
		}
		if err := dv.engine.Render(); err != nil {
			panic(err)
		}
		time.Sleep(time.Millisecond)
	}
}

/*=================== ANIMACIONES ===================*/

func (dv *device) startupAnimation() {
	dv.state = "running"

	for led := 0; led < len(dv.engine.Leds(0)); led++ {
		if dv.state == "stop" {
			break
		}

		dv.engine.Leds(0)[led] = uint32(0xffffff)
		if err := dv.engine.Render(); err != nil {
			panic(err)
		}
		time.Sleep(time.Millisecond)
	}
	dv.staticFinalWaitToStop()
}

func (dv *device) officeLightsMode() {
	dv.state = "running"

	for led := 0; led < len(dv.engine.Leds(0)); led++ {
		dv.engine.Leds(0)[led] = WarmWhiteColor
	}
	dv.staticFinalWaitToStop()
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
