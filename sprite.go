// ///////////////////////////////////////////////////////////////////
// Filename: sprite.go
// Description:
// Author: Mateo Rodriguez Ripolles (teorodrip@posteo.net)
// Maintainer:
// Created: Sat Apr 11 17:06:31 2020 (+0200)
// ///////////////////////////////////////////////////////////////////

package main

import (
	"github.com/faiface/pixel"
	"image"
	_ "image/png"
	"os"
)

type Range struct {
	X, Y float64
}

// The sprite iself, that means a single image, or animation
// or whatever you want, it is declared pointing to the first
// and last frame of the sprite inside the spritesheet (Frames).
// It also have a matrix to hold it's current state.
type Sprite struct {
	// The position of the images for this sprite inside the spritesheet
	Bounds pixel.Rect
	Size   Range // Size for each frame of the sprite
	Frames []*pixel.Sprite
	Mat    pixel.Matrix
}

// The spritesheet itself
type Spritesheet struct {
	Path      string
	Pic       pixel.Picture
	Sprites   []Sprite
	Triangles pixel.TrianglesData
	Batch     *pixel.Batch
}

func (s *Spritesheet) Init(spritesheetID int) (err error) {
	var (
		nSprites                                             int
		BoundsX0, BoundsY0, BoundsXf, BoundsYf, SizeX, SizeY float64
	)

	row := GlobalDB.GetSpritesheet.QueryRow(spritesheetID)
	row.Scan(&(s.Path), &nSprites)
	s.Pic, err = loadPicture(s.Path)
	if err != nil {
		return err
	}

	rows, err := GlobalDB.GetSprites.Query(spritesheetID)
	if err != nil {
		return err
	}
	defer rows.Close()

	i := 0
	s.Sprites = make([]Sprite, nSprites)
	for rows.Next() && i < nSprites {
		// Scan the rows
		err = rows.Scan(
			&BoundsX0,
			&BoundsY0,
			&BoundsXf,
			&BoundsYf,
			&SizeX,
			&SizeY,
		)
		if err != nil {
			return err
		}

		// Init the sprite structure
		s.Sprites[i] = Sprite{
			Bounds: pixel.R(BoundsX0, BoundsY0, BoundsXf, BoundsYf),
			Size:   Range{SizeX, SizeY},
			Frames: make([]*pixel.Sprite, calcFrameNum(
				BoundsX0, BoundsY0, BoundsXf, BoundsYf, SizeX, SizeY,
			)),
			Mat: pixel.IM,
		}

		// Init the frames of the sprite
		j := 0
		for x := BoundsX0; x < BoundsXf; x += SizeX {
			for y := BoundsY0; y < BoundsYf; y += SizeY {
				s.Sprites[i].Frames[j] = pixel.NewSprite(s.Pic, pixel.R(x, y, x+SizeX, y+SizeY))
				j++
			}
		}
		i++
	}

	s.Batch = pixel.NewBatch(&s.Triangles, s.Pic)

	return nil
}

// writes sprite to the spritesheet batch
func (s *Spritesheet) WriteSprite(sprite, frame int) {
	s.Sprites[sprite].Frames[frame].Draw(s.Batch, s.Sprites[sprite].Mat)
}

func (s *Spritesheet) SetMatrix(sprite int, mat pixel.Matrix) {
	s.Sprites[sprite].Mat = mat
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

func calcFrameNum(BX0, BY0, BXf, BYf, SX, SY float64) int {
	X := (BXf - BX0)
	Y := (BYf - BY0)

	return int((X * Y) / (SX * SY))
}
