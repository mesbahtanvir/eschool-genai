package models

type Quiz struct {
	Question string `json:"question"`
	OptionA  string `json:"a"`
	OptionB  string `json:"b"`
	OptionC  string `json:"c"`
	OptionD  string `json:"d"`
	Answer   string `json:"ans"`
}

type Content struct {
	Knowledge string `json:"knowledge"`
	SubModule Module `json:"sub_module"`
	Quiz      Quiz   `json:"quiz"`
}
type Module struct {
	Title       string    `json:"title"`
	Explanation string    `json:"description"`
	Content     []Content `json:"content"`
}

type CourseBlueprint struct {
	Title           string   `json:"title"`
	TinyTitle       string   `json:"tiny_title"`
	Description     string   `json:"description"`
	TinyDescription string   `json:"tiny_description"`
	Modules         []Module `json:"modules"`
}

type Course struct {
	CourseID        string          `json:"course_id"`
	CourseBlueprint CourseBlueprint `json:"course_blueprint"`
}
