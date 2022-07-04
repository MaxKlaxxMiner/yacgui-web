package lineDemo

import (
	"client/jscore/canvas"
	"client/jscore/keys"
	"math"
	"math/rand"
	"time"
)

const pih = math.Pi / 180.0
const pih2 = math.Pi / 360.0

type LineDemo struct {
	Speed     float64
	EdgeCount int
	gra       float64
	timeStamp time.Time
}

func New() LineDemo {
	return LineDemo{
		Speed:     0.5,
		EdgeCount: 21,
		timeStamp: time.Now(),
	}
}

func (demo *LineDemo) TickUpdate(c *canvas.Context, k *keys.Keys) {
	c.Save()
	c.ResetTransform()
	defer c.Restore()

	c.Clear("#000")
	c.SetLineWidth(2)

	if k.Pressed("KeyD", "ArrowRight") {
		demo.Speed += 0.025
	} else {
		if k.Pressed("KeyA", "ArrowLeft") {
			demo.Speed -= 0.025
		} else {
			if demo.Speed < 0.49 || demo.Speed > 0.51 {
				if demo.Speed < 0.5 {
					demo.Speed += 0.05
				} else {
					demo.Speed -= 0.05
				}
			}
		}
	}

	if k.Pressed("NumpadAdd", "BracketRight") {
		demo.EdgeCount += 2
		k.Release("NumpadAdd", "BracketRight")
	}

	if k.Pressed("NumpadSubtract", "Slash") && demo.EdgeCount > 3 {
		demo.EdgeCount -= 2
		k.Release("NumpadSubtract", "Slash")
	}

	tim := time.Now()
	elapsedTicks := 1.0
	elapsedTicks2 := float64(tim.Sub(demo.timeStamp)) / 16666666 // 60 ticks = 1 second
	if elapsedTicks2 > 2 {
		elapsedTicks2 -= 2
		elapsedTicks += elapsedTicks2*0.1 + 0.01
		if elapsedTicks2 > 10 {
			elapsedTicks = elapsedTicks2
		}
	}
	demo.timeStamp = demo.timeStamp.Add(time.Duration(elapsedTicks * 16666666))

	demo.gra += demo.Speed * elapsedTicks

	grad := math.Mod(demo.gra, 360)
	eg := math.Trunc(demo.gra / 360)

	for demo.gra < 0 {
		demo.gra += 720
	}
	for demo.gra > 720 {
		demo.gra -= 720
	}

	color := (int(math.Sin(pih*grad)*127)+127)<<16 | 127

	sw1 := pih * grad
	sw2 := pih2 * demo.gra

	rnd := rand.NewSource(12345)

	radStep := math.Pi * 2.0 / float64(demo.EdgeCount)
	widthH := float64(c.Width / 2)
	heightH := float64(c.Height / 2)
	widthK := widthH - widthH/5
	heightK := heightH - heightH/10
	for y := 0; y < demo.EdgeCount; y++ {
		x2 := math.Sin(sw1+radStep*float64(y)+sw2)*widthK + widthH
		y2 := -math.Cos(sw1+radStep*float64(y))*heightK + heightH

		for x := 0; x < y; x++ {
			x1 := math.Sin(sw1+radStep*float64(x)+sw2)*widthK + widthH
			y1 := -math.Cos(sw1+radStep*float64(x)+(sw1*eg))*heightK + heightH

			c.LineF(x1, y1, x2, y2, color|int(rnd.Int63()&0xff00))
		}
	}
}
