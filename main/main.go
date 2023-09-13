package main

import (
	"log"
	actionrpg2d "github.com/luigivanacore/actionrpg2d"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(actionrpg2d.WindowWidth, actionrpg2d.WindowHeight)
	ebiten.SetWindowTitle("AirWars2D")
	game := &actionrpg2d.Game{}
	//actionrpg2d.SetDebug(true)
	game.Initialize()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}