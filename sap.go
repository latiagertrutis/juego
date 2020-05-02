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
	// First 2 bits are the EndPoint position inside the
	// SapBox Points array, the rest are the Box position
	// inside the global box array
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

type SapBox struct {
	// Array containing the 2 needed points to define a rectangle.
	// The points are in form of indexes of the EnPoint slices
	// in the sweep and prune structure, in this form:
	// [X_min, Y_min, X_max, Y_max]
	Points [4]uint
}

type SweepPrune struct {
	EndPointsX []EndPoint
	EndPointsY []EndPoint
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
