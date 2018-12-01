package utils

import (
	"os"
	_ "image/jpeg"
	_ "golang.org/x/image/bmp"
	_ "image/png"
	"image"
	"image/draw"
	"log"
	"danser/render/texture"
)

func LoadImage(path string) (*image.NRGBA, error) {
	file, err := os.Open(path)
	log.Println("Loading texture: ", path)
	if err != nil {
		log.Println("er1")
		return nil, err
	}
	img, _, err := image.Decode(file)
	if err != nil {
		log.Println("er2")
		return nil, err
	}
	bounds := img.Bounds()
	nrgba := image.NewNRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(nrgba, nrgba.Bounds(), img, bounds.Min, draw.Src)
	return nrgba, nil
}

/*func LoadTexture(path string) (*texture.Texture, error) {
	img, err := LoadImage(path)
	if err == nil {
		tex := glhf.NewTexture(
			img.Bounds().Dx(),
			img.Bounds().Dy(),
			4,
			true,
			img.Pix,
		)

		tex.Begin()
		tex.SetWrap(glhf.CLAMP_TO_EDGE)
		tex.End()

		return tex, nil
	}
	return nil, err
}*/

func LoadTexture(path string) (*texture.TextureSingle, error) {
	img, err := LoadImage(path)
	if err == nil {
		tex := texture.LoadTextureSingle(img, 4)

		return tex, nil
	}
	return nil, err
}

func LoadTextureToAtlas(atlas *texture.TextureAtlas, path string) (*texture.TextureRegion, error) {
	img, err := LoadImage(path)
	if err == nil {
		return atlas.AddTexture(path, img.Bounds().Dx(), img.Bounds().Dy(), img.Pix), nil
	}
	log.Println(err)
	return nil, err
}

/*func LoadTextureU(path string) (*glhf.Texture, error) {
	img, err := LoadImage(path)
	if err == nil {
		tex := glhf.NewTexture(
			img.Bounds().Dx(),
			img.Bounds().Dy(),
			0,
			true,
			img.Pix,
		)

		tex.Begin()
		tex.SetWrap(glhf.CLAMP_TO_EDGE)
		tex.End()

		return tex, nil
	}
	return nil, err
}*/
