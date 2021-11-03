package main

//Acts as "folder" or "scene" or "world" from popular DCCs. Can be used for rendering multiple scenes in succession. You have to append your model to a particular obj_tracker for it to show up in the render.
type obj_tracker struct {
	models []model
}

//Add a model to the object tracker (think of it as a folder)
func (tracker obj_tracker) add(m model) obj_tracker {
	intermediate := append(tracker.models, m)
	return obj_tracker{models: intermediate}
}

//Check for intersection within models
func (tracker obj_tracker) ray_tracker_intersection(r ray) bool {
	for i := range tracker.models {
		if tracker.models[i].ray_model_intersection(r) {
			return true
		}
	}
	return false
}

//TODO remove model
