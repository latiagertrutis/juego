package main

import (
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
)

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

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Cosa",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)

	pic, err := loadPicture("./resources/cosa.png")
	if err != nil {
		panic(err)
	}

	var frames []pixel.Rect
	for x := pic.Bounds().Min.X; x < pic.Bounds().Max.X; x += 1000 {
		frames = append(frames, pixel.R(x, 0, x+1000, 1000))
	}

	win.Clear(colornames.Firebrick)

	mat := pixel.IM
	i := 0
	for !win.Closed() {
		win.Clear(colornames.Firebrick)

		sprite := pixel.NewSprite(pic, frames[i])
		mat := mat.Moved(win.Bounds().Center())
		sprite.Draw(win, mat)
		i = (i + 1) % len(frames)

		time.Sleep(time.Millisecond * 25)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
