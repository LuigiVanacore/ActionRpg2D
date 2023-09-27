package resources

import (
	"github.com/LuigiVanacore/ebiten_extended"
)

const (
	Warrior_idle_down = iota
	Warrior_idle_left
	Warrior_idle_right
	Warrior_idle_up
	Warrior_walk_down
	Warrior_walk_up
	Warrior_walk_left
	Warrior_walk_right
	Sword
	Aim
)



func LoadData() {
	ebiten_extended.ResourceManager().LoadImage(Spr_warrior_idle_down)
	ebiten_extended.ResourceManager().LoadImage(Spr_warrior_idle_left)
	ebiten_extended.ResourceManager().LoadImage(Spr_warrior_idle_right)
	ebiten_extended.ResourceManager().LoadImage(Spr_warrior_idle_up)
	ebiten_extended.ResourceManager().LoadImage(Spr_warrior_walk_down)
	ebiten_extended.ResourceManager().LoadImage(Spr_warrior_walk_up)
	ebiten_extended.ResourceManager().LoadImage(Spr_warrior_walk_left)
	ebiten_extended.ResourceManager().LoadImage(Spr_warrior_walk_right)
	ebiten_extended.ResourceManager().LoadImage(Spr_sword)
	ebiten_extended.ResourceManager().LoadImage(Spr_aim)
}
 