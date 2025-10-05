package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	gWidth         = 1200
	gHeight        = 960
	neighborRadius = 50
	interactR      = 20
	cohesionK      = .0
	separationK    = .6
	damping        = .20
	maxSpeed       = 10
)

var (
	bg   = color.RGBA{20, 24, 33, 255}
	red  = color.RGBA{230, 80, 90, 255}
	blue = color.RGBA{80, 120, 240, 255}
	nbr2 = neighborRadius * neighborRadius
)

type Agent struct {
	X, Y   float32
	VX, VY float32
	R      float32
	Col    color.RGBA
}

func main() {
	game := NewGame(gWidth, gHeight, 800)
	ebiten.SetWindowSize(gWidth, gHeight)
	ebiten.SetWindowTitle("Your game's title")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
