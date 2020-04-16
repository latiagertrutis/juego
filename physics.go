// ///////////////////////////////////////////////////////////////////
// Filename: physics.go
// Description:
// Author: Mateo Rodriguez Ripolles (teorodrip@posteo.net)
// Maintainer:
// Created: Thu Apr 16 16:41:52 2020 (+0200)
// ///////////////////////////////////////////////////////////////////

package main

import (
	"github.com/faiface/pixel"
)

type PhysicsObj struct {
	Spr *Sprite
	Vel pixel.Vec // Velocity
}

func CalcPhysics(obj *PhysicsObj) {
	mat := pixel.IM.Moved(obj.Vel)
	obj.Spr.UpdateMatrix(mat)
}
