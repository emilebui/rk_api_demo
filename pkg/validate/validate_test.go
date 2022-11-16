package validate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate(t *testing.T) {
	type TestStruct struct {
		Name  string `validate:"required,startswith=a"`
		Email string `validate:"email"`
	}

	testSubject := TestStruct{Name: "aok", Email: "blah@google.com"}
	err := Validate(testSubject)
	assert.Nil(t, err)

	testSubject = TestStruct{Name: "", Email: ""}
	err = Validate(testSubject)
	assert.NotNil(t, err)
}
