package main

type vec3 struct {
	coords [3]float64
}

//x component of vector. Returns a float64
func (v vec3) x() float64 { return v.coords[0] }

//y component of vector. Returns a float64
func (v vec3) y() float64 { return v.coords[1] }

//z component of vector. Returns a float64
func (v vec3) z() float64 { return v.coords[2] }

//inverse of a vector (reversed direction). Returns a vec3
func (v vec3) inv() vec3 {
	return vec3{coords: [3]float64{-v.x(), -v.y(), -v.z()}}
}

//add a constant to all components of vector. Returns a vec3
func (v vec3) cadd(addend float64) vec3 {
	return vec3{coords: [3]float64{v.x() + addend, v.y() + addend, v.z() + addend}}
}

//add components of one vector to components of another vector (1-1 mapped). Returns a vec3. Use inv() for subtraction
func (v1 vec3) vadd(v2 vec3) vec3 {
	return vec3{coords: [3]float64{v1.x() + v2.x(), v1.y() + v2.y(), v1.z() + v2.z()}}
}

//multipy all components of vector with a constant. Returns a vec3. Use 1/mul for division
func (v vec3) cmul(mul float64) vec3 {
	return vec3{coords: [3]float64{v.x() * mul, v.y() * mul, v.z() * mul}}
}

//1-1 multiplication of 2 vectors. Returns a vec3.
func (v1 vec3) vmul(v2 vec3) vec3 {
	return vec3{coords: [3]float64{v1.x() * v2.x(), v1.y() * v2.y(), v1.z() * v2.z()}}
}

//Length (magnitude) of vector, Returns float64
func (v vec3) length() float64 {
	return v.x()*v.x() + v.y()*v.y() + v.z()*v.z()
}

//Dot product of 2 vectors, the result of which represents how similar are two vectors in terms of direction. Returns a scalar (float64)
func dot(v1 vec3, v2 vec3) float64 {
	return v1.x()*v2.x() + v1.y()*v2.y() + v1.z()*v2.z()
}

//Cross product of 2 vectors, returns a vector that is perpendicular to both vectors. Returns a vec3
func cross(v1 vec3, v2 vec3) vec3 {
	return vec3{coords: [3]float64{v1.y()*v2.z() - v1.z()*v2.y(),
		v1.z()*v2.x() - v1.x()*v2.z(),
		v1.x()*v2.y() - v1.y()*v2.x()}}
}

//Vector that represents direction only and drops the magnitude. Returns a vec3.
func (v vec3) unit_vector() vec3 {
	return v.cmul(1 / v.length())
}
