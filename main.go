package main

import (
	"errors"
	"fmt"
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	w = 640
	h = 360
)

var USER_QUIT = errors.New("User quit")
var sprite *ebiten.Image

func main() {

	// Load sprite...

	f, _ := os.Open("ship.png")
	img, _ := png.Decode(f)
	sprite = ebiten.NewImageFromImage(img)
	f.Close()

	// Start loop...

	g := NewGame(w, h)
	ebiten.SetFullscreen(true)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetMaxTPS(240)

	err := ebiten.RunGame(g)
	if err != nil && err != USER_QUIT {
		fmt.Printf("%v\n", err)
	}
}

type Game struct {
	image *ebiten.Image
	width int
	height int
	shipx float64
	shipy float64
	cd bool
}

func NewGame(width int, height int) *Game {
	ret := new(Game)
	ret.image = ebiten.NewImage(width, height)
	ret.width = width
	ret.height = height
	ret.shipx = 64
	ret.shipy = 64
	return ret
}

func (self *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return self.width, self.height
}

func (self *Game) Draw(screen *ebiten.Image) {

	sprite_width, sprite_height := sprite.Size()
	opts := new(ebiten.DrawImageOptions)
	opts.GeoM.Translate(self.shipx - (float64(sprite_width) / 2), self.shipy - (float64(sprite_height) / 2))

	self.image.Clear()
	self.image.DrawImage(sprite, opts)
	screen.DrawImage(self.image, nil)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f   FPS: %0.2f   Vsync: %v (SPACE to toggle)   WASD to move",
		ebiten.CurrentTPS(), ebiten.CurrentFPS(), ebiten.IsVsyncEnabled()))
}

func (self *Game) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		self.shipx += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		self.shipx -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		self.shipy += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		self.shipy -= 2
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if self.cd == false {
			ebiten.SetVsyncEnabled(!ebiten.IsVsyncEnabled())
			self.cd = true
		}
	} else {
		self.cd = false
	}

	if (self.shipx < 0) {
		self.shipx = 0
	}
	if (self.shipx > float64(self.width)) {
		self.shipx = float64(self.width)
	}
	if (self.shipy < 0) {
		self.shipy = 0
	}
	if (self.shipy > float64(self.height)) {
		self.shipy = float64(self.height)
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return USER_QUIT
	} else {
		return nil
	}
}
