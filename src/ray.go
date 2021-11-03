package main

//Ray...from mathe- you know what a ray is! Requires origin (vec3 (honestly, bothering between point3 and vec3 seemed "point"less (get it? XD) in golang)) and direction (vec3)
type ray struct {
	origin, direction vec3
}

//location of ray at a point in time. Returns a vec3 with coordinates of ray at that point in time
func (v ray) at(t float64) vec3 {
	return v.origin.vadd(v.origin.cmul(t))
}

//Return colour of pixel if the ray is shot out
func ray_colour(r ray, tracker obj_tracker) vec3 {
	if tracker.ray_tracker_intersection(r) {
		return vec3{coords: [3]float64{0, 0, 1}}
	} else {
		return vec3{coords: [3]float64{0, 0, 0}}
	}
}
