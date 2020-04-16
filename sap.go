// ///////////////////////////////////////////////////////////////////
// Filename: sap.go
// Description: full axes sweep and prune algorithm
// Author: Mateo Rodriguez Ripolles (teorodrip@posteo.net)
// Maintainer:
// Created: Thu Apr 16 20:30:05 2020 (+0200)
// ///////////////////////////////////////////////////////////////////

package main

import (
	"github.com/faiface/pixel"
)

type EndPoint struct {
	Data  uint16
	Value float64
}

type Box struct {
	Min [2]*EndPoint
	Max [2]*EndPoint
}

type SweepPrune struct {
}
