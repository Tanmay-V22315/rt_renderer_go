package main

//Only triangles are supported (will probably change later). If using blender or any dcc, don't forget to triangulate the model. More fields will be added later on.
type model struct {
	name string
	tris []triangle
}

func (m model) add(t triangle) {
	m.tris[len(m.tris)+1] = t
}

//Models ray intersection check
func (m model) ray_model_intersection(r ray) bool {
	for i := range m.tris {
		if ray_tri_intersection(m.tris[i], r) {
			return true
		}
	}
	return false
}
