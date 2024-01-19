package utils

import "testing"

type PasswordTest struct {
	current string
	valid   bool
}

func TestCheckStrongPassword(t *testing.T) {
	passwords := []PasswordTest{
		{"Test2@a", true},
		{"tEst2@a", true},
		{"test@test", false},
		{"test", false},
		{"test2", false},
	}

	for _, testCase := range passwords {

		result := CheckStrongPassword(testCase.current)

		if result != nil {
			if testCase.valid != false {
				t.Errorf("Password: %s should be invalid.", testCase.current)
			}
		} else {
			if testCase.valid != true {
				t.Errorf("Password: %s should be valid.", testCase.current)
			}
		}

	}

}
