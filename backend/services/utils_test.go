package services

import (
	"backend/models"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetStructMetadata(t *testing.T) {

	expectedOutput := `Struct: CourseBlueprint
Field: Title, Type: string
Field: Description, Type: string
Field: Modules, Type: []      Struct: Module
    Field: Title, Type: string
	Field: Explanation, Type: string
	Field: Content, Type: []`

	result := GetStructMetadata(models.CourseBlueprint{})

	if diff := cmp.Diff(expectedOutput, result); diff != "" {
		t.Errorf("GetStructMetadata(user) mismatch (-expected +got):\n%s", diff)
	}
}
