package main

import (
	"fmt"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var frameNum int
var recording = false

type Game struct {
	Agents        []Agent
	Width, Height int
	Grid          *SpatialHash
	frame         int
}

func NewGame(w, h, numAgents int) *Game {
	g := &Game{Width: w, Height: h}
	g.Grid = NewSpatialHash(gWidth, gHeight, float32(neighborRadius))
	for i := 0; i < numAgents; i++ {
		g.Agents = append(g.Agents, RandomAgent())
	}
	return g
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bg)
	for _, a := range g.Agents {
		vector.StrokeCircle(screen, a.X, a.Y, a.R, 2, a.Col, true)
	}
	if recording {
		img := screen.SubImage(screen.Bounds()).(*ebiten.Image)
        path := fmt.Sprintf("frames/frame_%05d.png", frameNum)
		frameNum++
		file, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		if err := png.Encode(file, img); err != nil {
			log.Fatal(err)
		}
	}
}

func (g *Game) Update() error {
	g.frame++
	if g.frame%10 != 0 { 
		return nil
	}
	if err := g.Grid.RebuildHash(g.Agents); err != nil {
		log.Fatal(err)
	}

	for i := range g.Agents {
		a := &g.Agents[i]
		fx, fy := float32(0), float32(0)

		for _, j := range g.Grid.Neighbors(a.X, a.Y) {
			b := &g.Agents[j]
			dx := b.X - a.X
			dy := b.Y - a.Y
			d2 := dx*dx + dy*dy
			if d2 == 0 || d2 > float32(nbr2) {
				continue
			}
			d := float32(math.Sqrt(float64(d2)))
			nx := dx / d
			ny := dy / d
			w := 1 - d/float32(interactR)
			if a.Col == b.Col {
				fx += nx * float32(cohesionK) * w
				fy += ny * float32(cohesionK) * w
			} else {
				fx -= nx * float32(separationK) * w
				fy -= ny * float32(separationK) * w
			}

			a.VX = a.VX*float32(damping) + fx
			a.VY = a.VY*float32(damping) + fy
			if s2 := a.VX*a.VX + a.VY*a.VY; s2 > float32(maxSpeed*maxSpeed) {
				s := float32(math.Sqrt(float64(s2)))
				a.VX = a.VX / s * float32(maxSpeed)
				a.VY = a.VY / s * float32(maxSpeed)
			}
			a.X += a.VX
			a.Y += a.VY
			if a.X < a.R {
				a.X = a.R
				a.VX = -a.VX
			} else if a.X > float32(gWidth)-a.R {
				a.X = float32(gWidth) - a.R
				a.VX = -a.VX
			}
			if a.Y < a.R {
				a.Y = a.R
				a.VY = -a.VY
			} else if a.Y > float32(gHeight)-a.R {
				a.Y = float32(gHeight) - a.R
				a.VY = -a.VY
			}
		}
	}
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Width, g.Height
}
