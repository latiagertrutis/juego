// ///////////////////////////////////////////////////////////////////
// Filename: animation.go
// Description:
// Author: Mateo Rodriguez Ripolles (teorodrip@posteo.net)
// Maintainer:
// Created: Sun Apr 12 17:40:55 2020 (+0200)
// ///////////////////////////////////////////////////////////////////

package main

import (
	"github.com/faiface/pixel"
)

type AnimationFrame struct {
	FramePos uint
	Trans    pixel.Matrix
}

type Animation struct {
	SpriteID        uint
	Status          uint
	AnimationFrames []AnimationFrame
}
