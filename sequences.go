package main

import (
	"image/color"
	"time"

	"github.com/TeamNorCal/animation"
)

// Define test sequences

var s1 = animation.Step{UniverseID: 0, StepID: 1, Effect: animation.NewInterpolateSolid(
	color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0xff, 0x00, 0xff}, 3*time.Second)}
var s2 = animation.Step{UniverseID: 0, StepID: 2, Effect: animation.NewInterpolateSolid(
	color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}, 3*time.Second), OnCompletionOf: 1}
var s3 = animation.Step{UniverseID: 0, StepID: 3, Effect: animation.NewInterpolateSolid(
	color.RGBA{0x00, 0x00, 0xff, 0xff}, color.RGBA{0xff, 0x00, 0x00, 0xff}, 3*time.Second), OnCompletionOf: 2}

// Test1 is a simple linear interpolation between successive colors
var Test1 = animation.Sequence{[]*animation.Step{&s1, &s2, &s3}, true}
