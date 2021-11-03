package main

var test_triangle triangle = triangle{
	vertex1: vec3{[3]float64{-0.5, 0.4, 2}},
	vertex2: vec3{[3]float64{0.4, -0.7, 2}},
	vertex3: vec3{[3]float64{1, 0, 2}},
}

var model1 model = model{"test_triangle", []triangle{test_triangle}}

var folder1 obj_tracker = obj_tracker{[]model{model1}}
