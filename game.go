package actionrpg2d

import (
	"bytes"
	"image"
	"log"
	"time"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	resources "github.com/luigivanacore/actionrpg2d/Resources"
)

const (
	WindowWidth  = 800
	WindowHeight = 640
)

type AnimationSprite struct {
	transform ebiten_extended.Transform
	spriteSheet *ebiten_extended.Sprite
	image  *ebiten.Image
	frameSize ebiten_extended.Vector2D
	currentFrame uint
	frameCount uint
	duration time.Duration
	elapsedTime float64
	isLooping bool
}

func NewAnimationSprite(spriteSheet *ebiten_extended.Sprite, frameCount uint, duration time.Duration, isLooping bool) *AnimationSprite {
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}
	runnerImage = ebiten.NewImageFromImage(img)
	textureSize := spriteSheet.GetTextureRect()
	frameSize := ebiten_extended.Vector2D{ X: textureSize.Width / float64(frameCount), Y: textureSize.Height}
	return &AnimationSprite{ spriteSheet:  spriteSheet, image: spriteSheet.GetTexture(), frameSize: frameSize, frameCount: frameCount, duration: duration, isLooping: isLooping}
}


func (a *AnimationSprite)GetTransform() *ebiten_extended.Transform {
	return &a.transform
}
func (a *AnimationSprite) SetTransform(transform ebiten_extended.Transform) {
	a.transform = transform
}

var (
	runnerImage *ebiten.Image
)
const (
	screenWidth  = 320
	screenHeight = 240

	frameOX     = 0
	frameOY     = 0
 
	frameCount  = 5

	
)
func (a *AnimationSprite) Update(dt float64) {


	// // TODO: remove this Y axis overwrite?
	// a.sprite.FrameOffset.Y = a.offsetY

	// finished := false
	// a.frameTicker += delta
	// var frame int
	// if a.frameTicker >= a.animationSpan {
	// 	finished = true
	// 	if a.repeated {
	// 		rem := math.Mod(a.frameTicker, a.animationSpan)
	// 		a.frameTicker = rem
	// 		frame = int(a.frameTicker / a.deltaPerFrame)
	// 	} else {
	// 		a.frameTicker = a.animationSpan
	// 		frame = a.numFrames - 1
	// 	}
	// } else {
	// 	frame = int(a.frameTicker / a.deltaPerFrame)
	// }

	// if a.Mode == AnimationBackward {
	// 	frame = (a.numFrames - 1) - frame
	// }

	// framesDelta := frame - a.frame
	// a.frame = frame
	// if framesDelta != 0 {
	// 	// A small optimization: don't call Emit if there are no listeners.
	// 	// This is more useful for repeated animations as they're less likely to have
	// 	// any frame event listeners.
	// 	if !a.EventFrameChanged.IsEmpty() {
	// 		a.EventFrameChanged.Emit(framesDelta)
	// 	}
	// 	a.sprite.FrameOffset.X = a.frameWidth * float64(frame)
	// }

	// return finished



	// durationInSecond := a.duration.Seconds()
	// timePerFrame := float64( durationInSecond / float64(a.frameCount))
	// a.elapsedTime +=  1.0 / 60.0
	// var rect image.Rectangle
	// if a.currentFrame == 0 {
	// 	rect = image.Rectangle{image.Point{}, image.Point{int(a.frameSize.X), int(a.frameSize.Y)}}
	// }
	
	// for a.elapsedTime >= timePerFrame && !a.IsEnded() {

	// 	rect = image.Rectangle{image.Point{int(a.currentFrame),0},image.Point{int(a.frameSize.X),int(a.frameSize.Y)}}

	// 	a.elapsedTime -= timePerFrame
	// 	if a.isLooping {
	// 		a.currentFrame = (a.currentFrame + 1) % a.frameCount
	// 		// if a.currentFrame == 0 {
	// 		// 	rect = image.Rectangle{image.Point{}, image.Point{int(a.frameSize.X), int(a.frameSize.Y)}}
	// 		// }
	// 	} else {
	// 		a.currentFrame++
	// 	}
	// }
	// a.image = a.spriteSheet.GetTexture().SubImage(rect).(*ebiten.Image)

	//a.count = (a.count / durationInSecond) % frameCount



	// durationInSecond := 12.0
	// timePerFrame :=  float64(1 / durationInSecond) //float64( float64(durationInSecond) / float64(a.frameCount))
	// a.elapsedTime +=  1.0 / 60.0
	// frameWidth:= a.spriteSheet.GetTexture().Bounds().Dx() / int(a.frameCount)
	// frameHeight := a.spriteSheet.GetTexture().Bounds().Dy()
	// for a.elapsedTime >= timePerFrame && !a.IsEnded() {
	// a.elapsedTime -= timePerFrame
	// a.currentFrame = (a.currentFrame + 1) % a.frameCount
	// sx, sy := int(frameOX+a.currentFrame*uint(frameWidth)), frameOY
	// a.image = a.spriteSheet.GetTexture().SubImage(image.Rect(sx, sy, sx+frameWidth, frameHeight)).(*ebiten.Image)
	// }


// 	frameDuration := 8
// 	a.count++
// 	i := (a.count / frameDuration) % frameCount
// 	sx, sy := frameOX+i*frameWidth, frameOY
// 	a.image = runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image)
// 	if i >= frameCount {
// 		a.count = 0
//    }


// {
// 	if (!m_animIsPlaying)
// 		return;

// 	m_elapsedTime += dt;

// 	if (m_elapsedTime > m_timeToChangeFrame)
// 	{
// 		m_currentFrame++;
// 		if (m_currentFrame > m_lastFrame)
// 		{
// 			if (m_loop)
// 				m_currentFrame = m_activeStateInfo->beg;
// 			else
// 			{
// 				m_currentFrame = m_activeStateInfo->end;
// 				m_animIsPlaying = false;
// 			}
// 		}

// 		m_activeStateInfo->atlas->SetSpriteTextureByIndex(*m_sprite, m_currentFrame);
// 		m_elapsedTime = 0.f;
// 	}


}

func (a *AnimationSprite) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	 op.GeoM.Translate(-float64(a.spriteSheet.GetTexture().Bounds().Dx())/2, -float64(a.spriteSheet.GetTexture().Bounds().Dy())/2)
	 op.GeoM.Translate(screenWidth/2, screenHeight/2)
	 
	 
	 target.DrawImage(a.image, op)
	
	 
	// target.DrawImage(a.image, op)
}

func (a *AnimationSprite) IsEnded() bool {
	return a.currentFrame >= a.frameCount && !a.isLooping
}


type Game struct {
	warrior_sprite *ebiten_extended.Sprite
}

const (
	warrior_idle_down = iota
	warrior_walk_down
)

func (g *Game) Initialize() {
	ebiten_extended.ResourceManager().LoadImage(resources.Spr_warrior_idle_down)
	ebiten_extended.ResourceManager().LoadImage(resources.Spr_warrior_walk_down)
	g.warrior_sprite = ebiten_extended.NewSprite(ebiten_extended.ResourceManager().GetTexture(warrior_walk_down), true)
	g.warrior_sprite.SetPosition(WindowWidth/2, WindowHeight/2)
	animationSprite := NewAnimationSprite(g.warrior_sprite,8, time.Duration(3  * time.Second), true)
	node := ebiten_extended.NewNode(animationSprite)
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
