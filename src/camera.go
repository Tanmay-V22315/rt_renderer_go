package main

type camera struct {
	aspect_ratio, viewport_h, viewport_w, focal_length float64
	origin, horizontal, vertical                       vec3
}

func (c camera) lower_left() vec3 {
	return c.origin.vadd(c.vertical.cmul(0.5).inv()).vadd(c.horizontal.cmul(0.5).inv()).vadd(vec3{[3]float64{0, 0, c.focal_length}})
}

func (c camera) get_ray(u float64, v float64) ray {
	return ray{origin: c.origin, direction: c.lower_left().vadd(c.horizontal.cmul(u)).vadd(c.vertical.cmul(v)).vadd(c.origin.inv())}
}

//Default camera parameters
var c camera = camera{aspect_ratio: 16.0 / 9.0,
	viewport_h:   2.0,
	viewport_w:   (16.0 / 9.0) * 2.0,
	focal_length: 1,

	origin:     vec3{coords: [3]float64{0, 0, 0}},
	horizontal: vec3{coords: [3]float64{((16.0 / 9.0) * 2.0), 0, 0}},
	vertical:   vec3{coords: [3]float64{0.0, 2.0, 0.0}}}
