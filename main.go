// ///////////////////////////////////////////////////////////////////
// Filename: main.go
// Description:
// Author: Mateo Rodriguez Ripolles (teorodrip@posteo.net)
// Maintainer:
// Created: Sat Apr 11 16:46:58 2020 (+0200)
// ///////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// All Possible game states
const (
	StatMenu = iota
	StatPlay
)

var (
	// Variable holding current game state
	GameStat int
)

func run() {
	GameStat = StatMenu

	err := GlobalDB.Init("./resources/data.db")
	if err != nil {
		panic(err)
	}

	cfg := pixelgl.WindowConfig{
		Title:  "Cosa",
		Bounds: GetMainMonitorResolution(),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	sh := Spritesheet{}
	err = sh.Init(1)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Firebrick)

	i := 0
	fram := 0
	second := time.Tick(time.Second)
	// last := time.Now()

	T_Spr := sh.GetSprite(0)

	T_Spr.UpdateMatrix(pixel.IM.Moved(pixel.V(100, 100)))
	// T_Spr.UpdateMatrix(pixel.IM.Scaled(win.Bounds().Center(), 9))
	// mat1 := pixel.IM.Moved(win.Bounds().Center().Sub(pixel.V(70, -150)))
	// mat1 = mat1.Scaled(win.Bounds().Center(), 3)
	// mat2 := pixel.IM.Moved(win.Bounds().Center().Sub(pixel.V(30, -150)))
	// mat2 = mat2.Scaled(win.Bounds().Center(), 3)
	// mat3 := pixel.IM.Moved(win.Bounds().Center().Sub(pixel.V(-10, -150)))
	// mat3 = mat3.Scaled(win.Bounds().Center(), 3)
	// mat4 := pixel.IM.Moved(win.Bounds().Center().Sub(pixel.V(-50, -150)))
	// mat4 = mat4.Scaled(win.Bounds().Center(), 3)
	// last := time.Now()
	for !win.Closed() {
		// dt := time.Since(last).Seconds()

		sh.Batch.Clear()
		T_Spr.DrawSprite(sh.Batch, i)
		// sh.SetMatrix(1, mat2)
		// sh.WriteSprite(1, (i+8)%16)
		// sh.SetMatrix(0, mat3)
		// sh.WriteSprite(0, (i+4)%16)
		// sh.SetMatrix(1, mat4)
		// sh.WriteSprite(1, (i+10)%16)
		i = (i + 1) % 16

		win.Clear(colornames.Firebrick)
		sh.Batch.Draw(win)
		win.Update()

		fram++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, fram))
			fram = 0
		default:
		}
		// time.Sleep(time.Millisecond * 70)
	}
}

func main() {
	pixelgl.Run(run)
}
