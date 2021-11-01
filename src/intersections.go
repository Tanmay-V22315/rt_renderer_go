package main

//Using the Möller–Trumbore intersection algorithm, check whether a solution for intersection of ray and triangle exists. If not found or solution is "incorrect", return false. Else return true. Returns bool (source: https://en.wikipedia.org/wiki/M%C3%B6ller%E2%80%93Trumbore_intersection_algorithm)
func ray_tri_intersection(tri triangle, r ray) bool {
	const EPSILON = 0.000001
	var edge1 vec3 = tri.vertex1.vadd(tri.vertex2.inv())
	var edge2 vec3 = tri.vertex2.vadd(tri.vertex3.inv())
	var h vec3 = cross(r.direction, edge2)
	var a float64 = dot(h, edge1)
	if a > -EPSILON && a < EPSILON {
		return false //Ray parallel to triangle test
	}
	var f float64 = 1.0 / a
	var s vec3 = r.origin.vadd(tri.vertex1.inv())
	var u float64 = f * dot(s, h)
	if u < 0.0 || u > 1.0 {
		return false
	}
	q := cross(s, edge1)
	v := f * dot(r.direction, q)
	if v < 0.0 || u+v > 1.0 {
		return false
	}

	var t float64 = f * dot(edge2, q)
	if t > EPSILON {
		return true
	} else {
		return false
	}
}
