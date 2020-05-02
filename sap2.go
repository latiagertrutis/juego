// ///////////////////////////////////////////////////////////////////
// Filename: sap2.go
// Description: sap with boxes holding endpoints
// Author: Mateo Rodriguez Ripolles (teorodrip@posteo.net)
// Maintainer:
// Created: Sun May  3 01:18:37 2020 (+0200)
// ///////////////////////////////////////////////////////////////////

package main

// import "github.com/faiface/pixel"

type EndPoint struct {
	// Position of the endpoint inside Endpoints array
	Data  uint16
	Value float64
}

func (e EndPoint) IsMin() bool {
	// It is a minimum if the position in Points array is 0 or 1
	return (e.Data & 0x3) < 2
}

func (e EndPoint) Box() uint16 {
	return e.Data >> 1
}

type Point struct {
	X EndPoint
	Y EndPoint
}

type SapBox struct {
	Max Point
	Min Point
}

type SweepPrune struct {
	EndPointsX []*EndPoint
	EndPointsY []*EndPoint
}

// Perform insertion sort that is very efficient when only
// few elements are unsorted
func SortEndPoints(endpoints []EndPoint) {
	var i int
	var swapper EndPoint

	for j, endpoint := range endpoints {
		val := endpoint.Value
		i = j - 1
		for i >= 0 && endpoints[i].Value > val {
			swapper = endpoints[i]
			if endpoint.IsMin() && !swapper.IsMin() {
				// TODO check overlap and write to pair array
			}
		}
	}
}
