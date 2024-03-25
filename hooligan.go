package main

import (
	"image/color"
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

var tr = [6][4]int{
	{2, 3, 4, 5},
	{6, 3, 5, 1},
	{6, 4, 2, 1},
	{1, 3, 5, 6},
	{1, 2, 4, 6},
	{2, 3, 4, 5},
}

var player = struct {
	current_town	int
	dynamite_count	int
}{
	current_town: 1,
	dynamite_count: 3,
}

func int_contains(sl [4]int, thing int) bool {
	for _, value := range sl {
		if value == thing {
			return true
		}
	}
	return false
}

func load_image(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func draw_image(screen *ebiten.Image, img *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(img, op)
}

func (g *Game) Init() error {
	return nil
}

func (g *Game) Update() error {


	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//fill screen with white
	ebitenutil.DrawRect(screen, 0, 0, 320, 240, color.White)

	draw_image(screen, load_image("img/you.png"), 200, 120)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	/*
	var tr = [6][4]int{
		{2, 3, 4, 5},
		{6, 3, 5, 1},
		{6, 4, 2, 1},
		{1, 3, 5, 6},
		{1, 2, 4, 6},
		{2, 3, 4, 5},
	}
	*/

	game := &Game{}
    	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("hooligan")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
