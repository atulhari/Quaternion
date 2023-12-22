package Quaternion

import (
	"math"
	"testing"
)

var (
	eulerQ   = [4]float64{0.7038543, -0.0705648, 0.070621, 0.703294}
	eulerOut = [3]float64{1.57, 0.2, 0.0}
	qIn      = Quaternion{W: 0.7038543, X: -0.0705648, Y: 0.070621, Z: 0.703294}
)

func TestSet(t *testing.T) {
	quat := Quaternion{}
	q := quat.Set(1.0, 1.57, 0.6, 0.0)
	if q.W != 1.0 && q.X != 1.57 && q.Y != 0.6 && q.Z != 0.0 {
		t.Fail()
	}
}

func TestFromEuler(t *testing.T) {
	quat := FromEuler(1.57, 0.2, 0.0)
	if math.Abs(quat.W-eulerQ[0]) > 1e-5 ||
		math.Abs(quat.X-eulerQ[1]) > 1e-5 ||
		math.Abs(quat.Y-eulerQ[2]) > 1e-5 ||
		math.Abs(quat.Z-eulerQ[3]) > 1e-5 {
		t.Fail()
	}

}

func TestToEuler(t *testing.T) {

	yaw, pitch, roll := qIn.ToEuler()
	if math.Abs(yaw-eulerOut[0]) > 1e-5 ||
		math.Abs(pitch-eulerOut[1]) > 1e-5 ||
		math.Abs(roll-eulerOut[2]) > 1e-5 {
		t.Fail()
	}

}
