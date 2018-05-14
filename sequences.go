package main

import (
	"time"

	"github.com/TeamNorCal/animation"
)

// Define test sequences

func CreateTest1() *animation.Sequence {
	var s1 = animation.Step{UniverseID: 0, Effect: animation.NewInterpolateSolidHexRGB(0xff0000, 0x00ff00, 3*time.Second)}
	var s2 = animation.Step{UniverseID: 0, Effect: animation.NewInterpolateSolidHexRGB(0x00ff00, 0x0000ff, 3*time.Second)}
	var s3 = animation.Step{UniverseID: 0, Effect: animation.NewInterpolateSolidHexRGB(0x0000ff, 0xff0000, 3*time.Second)}
	s1.ThenDoImmediately("s2")
	s2.ThenDoImmediately("s3")
	s3.ThenDoImmediately("s1")

	var t1 = animation.Step{UniverseID: 1, Effect: animation.NewInterpolateSolidHexRGB(0xf4ab22, 0x22f445, 1*time.Second)}
	var t2 = animation.Step{UniverseID: 1, Effect: animation.NewInterpolateSolidHexRGB(0x22f445, 0xd122f4, 1*time.Second)}
	var t3 = animation.Step{UniverseID: 1, Effect: animation.NewInterpolateSolidHexRGB(0xd122f4, 0xf4ab22, 1*time.Second)}
	t1.ThenDoImmediately("t2")
	t2.ThenDoImmediately("t3")
	t3.ThenDoImmediately("t1")

	// Test1 is a simple linear interpolation between successive colors
	seq := animation.NewSequence()
	seq.AddInitialStep("s1", &s1)
	seq.AddStep("s2", &s2)
	seq.AddStep("s3", &s3)
	seq.AddInitialStep("t1", &t1)
	seq.AddStep("t2", &t2)
	seq.AddStep("t3", &t3)

	return seq
}
