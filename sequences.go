package main

import (
	"time"

	"github.com/TeamNorCal/animation"
)

// Define test sequences

// CreateTest1 creates a test sequence
func CreateTest1() *animation.Sequence {
	var s1 = animation.Step{UniverseID: 0, Effect: animation.NewInterpolateSolidHexRGB(0xff0000, 0x00ff00, 3*time.Second)}
	var s2 = animation.Step{UniverseID: 0, Effect: animation.NewInterpolateSolidHexRGB(0x00ff00, 0x0000ff, 3*time.Second)}
	var s3 = animation.Step{UniverseID: 0, Effect: animation.NewInterpolateSolidHexRGB(0x0000ff, 0xff0000, 3*time.Second)}

	var t1 = animation.Step{UniverseID: 1, Effect: animation.NewInterpolateSolidHexRGB(0xf4ab22, 0x22f445, 1*time.Second)}
	var t2 = animation.Step{UniverseID: 1, Effect: animation.NewInterpolateToHexRGB(0xd122f4, 1*time.Second)}
	var t3 = animation.Step{UniverseID: 1, Effect: animation.NewInterpolateToHexRGB(0xf4ab22, 1*time.Second)}

	u1 := animation.Step{UniverseID: 2, Effect: animation.NewInterpolateToHexRGB(0xffffff, 1500*time.Millisecond)}
	u2 := animation.Step{UniverseID: 2, Effect: animation.NewInterpolateToHexRGB(0x000000, 2*time.Second)}

	// Test1 is a simple linear interpolation between successive colors
	seq := animation.NewSequence()
	seq.AddInitialStep("s1", &s1)
	seq.AddStep("s2", &s2)
	seq.AddStep("s3", &s3)
	seq.CreateStepCycle("s1", "s2", "s3")
	seq.AddInitialStep("t1", &t1)
	seq.AddStep("t2", &t2)
	seq.AddStep("t3", &t3)
	seq.CreateStepCycle("t1", "t2", "t3")
	seq.AddInitialStep("u1", &u1)
	seq.AddStep("u2", &u2)
	seq.CreateStepCycle("u1", "u2")

	return seq
}

// CreateTest2 creates a sequence that runs animations in order across multiple universes
func CreateTest2() *animation.Sequence {
	s1 := animation.Step{UniverseID: 0, Effect: animation.NewInterpolateSolidHexRGB(0x8d209f, 0x03f164, 2*time.Second)}
	s2 := animation.Step{UniverseID: 1, Effect: animation.NewInterpolateSolidHexRGB(0x848953, 0x194847, 2*time.Second)}
	s3 := animation.Step{UniverseID: 2, Effect: animation.NewInterpolateSolidHexRGB(0xabcdef, 0x654321, 2*time.Second)}

	seq := animation.NewSequence()
	seq.AddInitialStep("s1", &s1)
	seq.AddStep("s2", &s2)
	seq.AddStep("s3", &s3)
	seq.CreateStepCycle("s1", "s2", "s3")

	return seq
}
