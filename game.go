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
}
 

func (g *Game) Initialize() {

	// ebiten_extended.ResourceManager().LoadImage(resources.Spr_warrior_idle_down)
	// ebiten_extended.ResourceManager().LoadImage(resources.Spr_warrior_walk_down)
	// warrior_sprite := ebiten_extended.ResourceManager().GetTexture(warrior_walk_down)
	// animationSprite := ebiten_extended.NewAnimationSet(warrior_sprite,8, 12, true)
	// animationPlayer := ebiten_extended.NewAnimationPlayer()
	// animationPlayer.AddAnimation(animationSprite, 0)
	// animationPlayer.Start()
	resources.LoadData()
	player := NewPlayer()
	node := ebiten_extended.NewNode(player)
	ebiten_extended.GameManager().AddNode(node)
}

func (g *Game) Run() {

}

func (g *Game) Update() error{
	ebiten_extended.GameManager().Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	ebiten_extended.GameManager().Draw(screen, op)
}

func (g *Game) Layout(int, int) (screenwidth int, screenheight int) {
	return WindowWidth, WindowHeight
}
