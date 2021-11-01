package main

import (
	"fmt"
	"strconv"

	"github.com/schollz/progressbar/v3"
)

func ray_colour(r ray) vec3 {
	var test_triangle triangle = triangle{vec3{coords: [3]float64{1, 0, 3}}, vec3{coords: [3]float64{0, -1, 3}}, vec3{coords: [3]float64{-2, 0, 4}}}
	if ray_tri_intersection(test_triangle, r) {
		return vec3{coords: [3]float64{test_triangle.N().x() + 1, test_triangle.N().y() + 1, test_triangle.N().z() + 1}}.cmul(0.5)
		// return vec3{coords: [3]float64{1, 0, 0}}
	} else {
		return vec3{coords: [3]float64{0, 0, 0}}
	}
}

func main() {
	var aspect_ratio float64 = 16.0 / 9.0
	var image_width int = 3840
	var image_height = int(float64(image_width) / aspect_ratio)

	var viewport_h float64 = 2.0
	var viewport_w float64 = aspect_ratio * viewport_h
	focal_length := 1.0

	var origin vec3 = vec3{coords: [3]float64{0, 0, 0}}
	var horizontal vec3 = vec3{coords: [3]float64{viewport_w, 0, 0}}
	var vertical vec3 = vec3{coords: [3]float64{0, viewport_h, 0}}

	var lower_left vec3 = origin.vadd(horizontal.cmul(0.5).inv()).vadd(vertical.cmul(0.5).inv()).vadd(vec3{coords: [3]float64{0, 0, focal_length}})

	fmt.Println("P3\n" + strconv.FormatInt(int64(image_width), 10) + " " + strconv.FormatInt(int64(image_height), 10) + "\n255\n")

	bar := progressbar.Default(int64(image_height - 1))

	for j := image_height - 1; j >= 0; j-- {
		bar.Add(1)
		for i := 0; i <= image_width-1; i++ {
			var pixel_colour vec3
			u := float64(i) / float64(image_width-1)
			v := float64(j) / float64(image_height-1)
			var r ray = ray{origin: origin, direction: lower_left.vadd(horizontal.cmul(u)).vadd(vertical.cmul(v)).vadd(origin.inv())}
			pixel_colour = ray_colour(r)
			fmt.Println(int(pixel_colour.x()*255.99), int(pixel_colour.y()*255.99), int(pixel_colour.z()*255.99))
		}
	}
}
