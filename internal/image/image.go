package image

import (
	"image"
	"os"

	"github.com/creepitall/platformer/internal/domain"
	"github.com/faiface/pixel"
)

type heroPlayerSpritePath struct {
	name string
	path string
}

type frontLayerSpritePath struct {
	name string
	path string
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func FillHeroPlayerSprite() {
	spritePath := []heroPlayerSpritePath{
		{name: "running", path: domain.ReturnFilePath("assets/KnightRun_scale.png")},
		{name: "staying", path: domain.ReturnFilePath("assets/KnightIdle_scale.png")},
		{name: "jumping", path: domain.ReturnFilePath("assets/KnightJump_scale.png")},
	}

	for _, sprite := range spritePath {
		assets, err := loadPicture(sprite.path)
		if err != nil {
			panic(err)
		}

		switch sprite.name {
		case "running":
			{
				domain.HeroPlayerRunAssets = assets
				domain.HeroPlayerRunFrames = returnFrames(assets, 32.0, 32.0)
			}
		case "staying":
			{
				domain.HeroPlayerStayAssets = assets
				domain.HeroPlayerStayFrames = returnFrames(assets, 32.0, 32.0)
			}
		case "jumping":
			{
				domain.HeroPlayerJumpAssets = assets
				domain.HeroPlayerJumpFrames = returnFrames(assets, 32.0, 32.0)
			}
		}

	}
}

func FillFrontSpriteByScene(CurrentScene string) map[string][]*pixel.Sprite {
	m := make(map[string][]*pixel.Sprite)
	switch CurrentScene {
	case "start":
		{
			m["front"] = returnStartSceneAssets_Front()
			m["back"] = returnStartSceneAssets_Background()
		}

	}
	return m
}

func returnStartSceneAssets_Front() []*pixel.Sprite {
	//
	spritePath := []frontLayerSpritePath{
		//{name: "front_frames", path: domain.ReturnFilePath("assets/build_3.png")},
		{name: "front", path: domain.ReturnFilePath("assets/test1.png")},
	}

	sprites := make([]*pixel.Sprite, 0)
	for _, sprite := range spritePath {
		assets, err := loadPicture(sprite.path)
		if err != nil {
			panic(err)
		}

		// frames := returnFrames(assets, 32.0, 32.0)

		// for _, frameValue := range frames {
		// 	sprites = append(sprites, pixel.NewSprite(assets, frameValue))
		// }
		sprites = append(sprites, pixel.NewSprite(assets, assets.Bounds()))
	}

	return sprites
}

func returnStartSceneAssets_Background() []*pixel.Sprite {
	spritePath := []frontLayerSpritePath{
		{name: "back", path: domain.ReturnFilePath("assets/background1.png")},
		{name: "back", path: domain.ReturnFilePath("assets/background3.png")},
		{name: "back", path: domain.ReturnFilePath("assets/background4b.png")},
	}

	sprites := make([]*pixel.Sprite, 0)
	for _, sprite := range spritePath {
		assets, err := loadPicture(sprite.path)
		if err != nil {
			panic(err)
		}

		sprites = append(sprites, pixel.NewSprite(assets, assets.Bounds()))
	}

	return sprites
}

func returnFrames(assets pixel.Picture, sizeX, sizeY float64) []pixel.Rect {
	var frames []pixel.Rect
	for y := 0.0; y < assets.Bounds().Max.Y; y += sizeY {
		for x := 0.0; x < assets.Bounds().Max.X; x += sizeX {
			frames = append(frames, pixel.R(x, y, x+sizeX, y+sizeY))
		}
	}

	return frames
}
