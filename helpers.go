package main

import (
	"image/color"
	"math/rand"
)

func RandomAgent() Agent {
	var c color.RGBA
	if rand.Intn(2) != 0 {
		c = blue
	} else {
		c = red
	}
	return Agent{
		X:   float32(rand.Intn(gWidth)),
		Y:   float32(rand.Intn(gHeight)),
		R:   5,
		Col: c,
	}
}
