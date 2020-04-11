// ///////////////////////////////////////////////////////////////////
// Filename: window.go
// Description:
// Author: Mateo Rodriguez Ripolles (teorodrip@posteo.net)
// Maintainer:
// Created: Sat Apr 11 14:52:35 2020 (+0200)
// ///////////////////////////////////////////////////////////////////

package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func GetMainMonitorResolution() pixel.Rect {
	width, height := pixelgl.PrimaryMonitor().Size()
	return pixel.R(0, 0, width, height)
}
