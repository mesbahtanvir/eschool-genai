package models

type Quiz struct {
	Question string `json:"question"`
	OptionA  string `json:"option_a"`
	OptionB  string `json:"option_b"`
	OptionC  string `json:"option_c"`
	OptionD  string `json:"option_d"`
	Answer   string `json:"answer"`
}

type Content struct {
	Knowledge string `json:"knowledge"`
	SubModule Module `json:"sub_module"`
	Quiz      Quiz   `json:"Quiz"`
}
type Module struct {
	Title       string    `json:"title"`
	Explanation string    `json:"description"`
	Content     []Content `json:"modules"`
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
