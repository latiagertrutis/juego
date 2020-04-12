// ///////////////////////////////////////////////////////////////////
// Filename: animation.go
// Description:
// Author: Mateo Rodriguez Ripolles (teorodrip@posteo.net)
// Maintainer:
// Created: Sun Apr 12 17:40:55 2020 (+0200)
// ///////////////////////////////////////////////////////////////////

package main

import (
	"time"
)

type Animation struct {
	SpriteID        uint
	Status          uint
	LastTick        time.Time
	AnimationFrames []uint
}

func (a *Animation) Init(spriteID uint) {

}
