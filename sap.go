// ///////////////////////////////////////////////////////////////////
// Filename: sap.go
// Description: full axes sweep and prune algorithm
// Author: Mateo Rodriguez Ripolles (teorodrip@posteo.net)
// Maintainer:
// Created: Thu Apr 16 20:30:05 2020 (+0200)
// ///////////////////////////////////////////////////////////////////

package main

// import "github.com/faiface/pixel"

type EndPoint struct {
	Data  uint16
	Value float64
}

func (e EndPoint) isMin() bool {
	return (e.Data & 0x1) == 0x1
}

type SapAxis struct {
	X, Y *EndPoint
}

type Box struct {
	Min SapAxis
	Max SapAxis
}

type SweepPrune struct {
	EndPointsX []EndPoint
	EndPointsY []EndPoint
}

func SortEndPoints(endpoints []EndPoint) {
	var i int
	var swapper EndPoint

	for j, endpoint := range endpoints {
		val := endpoint.Value
		i = j - 1
		for i >= 0 && endpoints[i].Value > val {
			swapper = endpoints[i]
			if endpoint.isMin() && !swapper.isMin() {
				// TODO check overlap and write to pair array
			}
		}
	}
}
