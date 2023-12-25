package Quaternion

import (
	"math"
)

type Quaternion struct {
	W float64
	X float64
	Y float64
	Z float64
}

type RotMat [3][3]float64

func (quat *Quaternion) Set(w, x, y, z float64) Quaternion {
	q := Quaternion{W: w, X: x, Y: y, Z: z}
	return q
}

func FromEuler(y, p, r float64) Quaternion { // YPR ZYX in radians
	// q = q4 + iq1 + jq2 + kq3
	// q1 = cos(yaw/2) cos(pitch/2) sin(roll/2) - sin(yaw/2) sin(pitch/2) cos(roll/2)
	// q2 = cos(yaw/2) sin(pitch/2) cos(roll/2) + sin(yaw/2) cos(pitch/2) sin(roll/2)
	// q3 = sin(yaw/2) cos(pitch/2) cos(roll/2) - cos(yaw/2) sin(pitch/2) sin(roll/2)
	// q4 = cos(yaw/2) cos(pitch/2) cos(roll/2) + sin(yaw/2) sin(pitch/2) sin(roll/2)

	qx := math.Cos(y*0.5)*math.Cos(p*0.5)*math.Sin(r*0.5) - math.Sin(y*0.5)*math.Sin(p*0.5)*math.Cos(r*0.5)
	qy := math.Cos(y*0.5)*math.Sin(p*0.5)*math.Cos(r*0.5) + math.Sin(y*0.5)*math.Cos(p*0.5)*math.Sin(r*0.5)
	qz := math.Sin(y*0.5)*math.Cos(p*0.5)*math.Cos(r*0.5) - math.Cos(y*0.5)*math.Sin(p*0.5)*math.Sin(r*0.5)
	qw := math.Cos(y*0.5)*math.Cos(p*0.5)*math.Cos(r*0.5) + math.Sin(y*0.5)*math.Sin(p*0.5)*math.Sin(r*0.5)
	return Quaternion{W: qw, X: qx, Y: qy, Z: qz}
}

func (quat *Quaternion) ToEuler() (float64, float64, float64) { //With ZYX order
	// q = q4 + iq1 + jq2 + kq3
	// tan(yaw) = 2(q1q2+q4q3) / (q4^2 + q1^2 - q2^2 - q3^2)
	// sin(pitch) = -2(q1q3-q4q2)
	// tan(roll)  =  2(q4q1+q2q3) / (q4^2 - q1^2 - q2^2 + q3^2)
	q := *quat
	yaw := math.Atan(2 * (q.X*q.Y + q.W*q.Z) / (q.W*q.W + q.X*q.X - q.Y*q.Y - q.Z*q.Z))
	pitch := math.Asin(-2 * (q.X*q.Z - q.W*q.Y))
	roll := math.Atan(2 * (q.W*q.X + q.Y*q.Z) / (q.W*q.W - q.X*q.X - q.Y*q.Y + q.Z*q.Z))
	return yaw, pitch, roll
}

func (quat *Quaternion) ToRotMat() RotMat {
	q := *quat
	R := RotMat{}
	R[0][0] = 1 - 2*q.Y*q.Y - 2*q.Z*q.Z
	R[0][1] = 2*q.X*q.Y - 2*q.W*q.Z
	R[0][2] = 2*q.X*q.Z + 2*q.W*q.Y

	R[1][0] = 2*q.X*q.Y + 2*q.W*q.Z
	R[1][1] = 1 - 2*q.X*q.X - 2*q.Z*q.Z
	R[1][2] = 2*q.Y*q.Z - 2*q.W*q.X

	R[2][0] = 2*q.X*q.Z - 2*q.W*q.Y
	R[2][1] = 2*q.Y*q.Z + 2*q.W*q.X
	R[2][2] = 1 - 2*q.X*q.X - 2*q.Y*q.Y

	return R
}

func Sgn(x float64) float64 {
	switch {
	case x < 0:
		return -1
	case x > 0:
		return +1
	}
	return 0
}

func RotMatToQuat(RotMat) Quaternion {
	R := RotMat{}
	q := Quaternion{}
	q.W = math.Sqrt((R[0][0] + R[1][1] + R[2][2] + 1.0)) / 2.0
	q.X = math.Sqrt((R[0][0] - R[1][1] - R[2][2] + 1.0)) / 2.0
	q.Y = math.Sqrt((-R[0][0] + R[1][1] - R[2][2] + 1.0)) / 2.0
	q.Z = math.Sqrt((-R[0][0] - R[1][1] + R[2][2] + 1.0)) / 2.0

	if q.W < 0.0 {
		q.W = 0.0
	}
	if q.X < 0.0 {
		q.X = 0.0
	}
	if q.Y < 0.0 {
		q.Y = 0.0
	}
	if q.Z < 0.0 {
		q.Z = 0.0
	}

	if q.W >= q.X && q.W >= q.Y && q.W >= q.Z {
		q.W *= 1.0
		q.X *= Sgn(R[2][1] - R[1][2])
		q.Y *= Sgn(R[0][2] - R[2][0])
		q.Z *= Sgn(R[1][0] - R[0][1])
	}

	return q
}
