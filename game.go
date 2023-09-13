package actionrpg2d

import (
	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/hajimehoshi/ebiten/v2"
	resources "github.com/luigivanacore/actionrpg2d/Resources"
)

const (
	WindowWidth  = 800
	WindowHeight = 640
)

type Game struct {
	warrior_sprite *ebiten_extended.Sprite
}

const (
	warrior_idle_down = iota
)

func (g *Game) Initialize() {
	ebiten_extended.ResourceManager().LoadImage(resources.Spr_warrior_idle_down)
	g.warrior_sprite = ebiten_extended.NewSprite(ebiten_extended.ResourceManager().GetTexture(warrior_idle_down), true)
	g.warrior_sprite.SetPosition(WindowWidth/2, WindowHeight/2)
}

func (g *Game) Run() {

}

func (g *Game) Update() error{
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	g.warrior_sprite.Draw(screen,op)
}

func (g *Game) Layout(int, int) (screenwidth int, screenheight int) {
	return WindowWidth, WindowHeight
}
