package env

import (
	"fmt"
	"os"
	"testing"
)

func TestGetAndValidateF(t *testing.T) {
	key, val := "TEST", "value"
	os.Setenv(key, val)
	out := GetAndValidateF(key)
	if val != out {
		t.Errorf("Values not equal, expected: %s received: %s", val, out)
	}
}

func TestGetAndValidateF_Error(t *testing.T) {
	origLogFatalf := logFatalf

	// After this test, replace the original fatal function
	defer func() { logFatalf = origLogFatalf }()

	errors := []string{}
	logFatalf = func(format string, args ...interface{}) {
		if len(args) > 0 {
			errors = append(errors, fmt.Sprintf(format, args))
		} else {
			errors = append(errors, format)
		}
	}

	key, val := "TEST", " 	" // Whitespaces
	os.Setenv(key, val)
	GetAndValidateF(key)
	if len(errors) != 1 {
		t.Errorf("excepted one error, actual %v", len(errors))
	}
}
