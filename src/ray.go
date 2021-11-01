package main

//Ray...from mathe- you know what a ray is! Requires origin (vec3 (honestly, bothering between point3 and vec3 seemed pointless in golang)) and direction (vec3)
type ray struct {
	origin    vec3
	direction vec3
}

//location of ray at a point in time. Returns a vec3 with coordinates of ray at that point in time
func (v ray) at(t float64) vec3 {
	return v.origin.vadd(v.origin.cmul(t))
}
