package main

import (
	"fmt"
	"strconv"

	"github.com/schollz/progressbar/v3"
)

func main() {
	var aspect_ratio float64 = 16.0 / 9.0
	var image_width int = 600
	var image_height = int(float64(image_width) / aspect_ratio)
	fmt.Println("P3\n" + strconv.FormatInt(int64(image_width), 10) + " " + strconv.FormatInt(int64(image_height), 10) + "\n255\n")

	bar := progressbar.Default(int64(image_height - 1))

	for j := image_height - 1; j >= 0; j-- {
		bar.Add(1)
		for i := 0; i <= image_width-1; i++ {
			var pixel_colour vec3
			u := float64(i) / float64(image_width-1)
			v := float64(j) / float64(image_height-1)
			pixel_colour = ray_colour(c.get_ray(u, v), folder1)
			fmt.Println(int(pixel_colour.x()*255.99), int(pixel_colour.y()*255.99), int(pixel_colour.z()*255.99))
		}
	}
}
