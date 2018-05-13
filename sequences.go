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

var t1 = animation.Step{UniverseID: 1, StepID: 4, Effect: animation.NewInterpolateSolid(
	color.RGBA{0xf4, 0xab, 0x22, 0xff}, color.RGBA{0x22, 0xf4, 0x45, 0xff}, 1*time.Second)}
var t2 = animation.Step{UniverseID: 1, StepID: 5, Effect: animation.NewInterpolateSolid(
	color.RGBA{0x22, 0xf4, 0x45, 0xff}, color.RGBA{0xd1, 0x22, 0xf4, 0xff}, 1*time.Second), OnCompletionOf: 4}
var t3 = animation.Step{UniverseID: 1, StepID: 6, Effect: animation.NewInterpolateSolid(
	color.RGBA{0xd1, 0x22, 0xf4, 0xff}, color.RGBA{0xf4, 0xab, 0x22, 0xff}, 1*time.Second), OnCompletionOf: 5}

// Test1 is a simple linear interpolation between successive colors
var Test1 = animation.Sequence{[]*animation.Step{&s1, &s2, &s3, &t1, &t2, &t3}, true}
