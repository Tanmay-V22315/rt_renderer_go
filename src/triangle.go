package main

type triangle struct {
	vertex1, vertex2, vertex3 vec3
}

//Return Normals of triangle
func (tri triangle) N() vec3 {
	var edge1 vec3 = tri.vertex1.vadd(tri.vertex2.inv())
	var edge2 vec3 = tri.vertex2.vadd(tri.vertex3.inv())
	return cross(edge1, edge2)
}
