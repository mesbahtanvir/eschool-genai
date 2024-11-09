package models

type Module struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Quiz        string   `json:"quiz"`
	SubModules  []Module `json:"sub_modules"`
}

type CourseBlueprint struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Modules     []Module `json:"modules"`
}

type Course struct {
	CourseID        string          `json:"course_id"`
	CourseBlueprint CourseBlueprint `json:"course_blueprint"`
}
