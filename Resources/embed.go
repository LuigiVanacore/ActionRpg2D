package resources


import (
	_ "embed"
)

var (
	//go:embed players/warrior/spr_warrior_idle_down.png
	Spr_warrior_idle_down []byte

	//go:embed players/warrior/spr_warrior_idle_up.png
	Spr_warrior_idle_up []byte

	//go:embed players/warrior/spr_warrior_idle_left.png
	Spr_warrior_idle_left []byte

	//go:embed players/warrior/spr_warrior_idle_right.png
	Spr_warrior_idle_right []byte

	//go:embed players/warrior/spr_warrior_walk_down.png
	Spr_warrior_walk_down []byte

	//go:embed players/warrior/spr_warrior_walk_up.png
	Spr_warrior_walk_up []byte

	//go:embed players/warrior/spr_warrior_walk_left.png
	Spr_warrior_walk_left []byte

	//go:embed players/warrior/spr_warrior_walk_right.png
	Spr_warrior_walk_right []byte

	//go:embed projectiles/spr_arrow.png
	Spr_arrow []byte

	//go:embed projectiles/spr_dagger.png
	Spr_dagger []byte

	//go:embed projectiles/spr_magic_ball.png
	Spr_magicBall []byte

	//go:embed projectiles/spr_sword.png
	Spr_sword []byte

	//go:embed ui/spr_aim.png
	Spr_aim []byte
)
