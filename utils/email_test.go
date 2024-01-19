package utils

import "testing"

type EmailTest struct {
	current string
	valid   bool
}

func TestCheckValidEmail(t *testing.T) {
	emails := []EmailTest{
		{"test2.a@dwa.to", true},
		{"test@test", true},
		{"test", false},
		{"test2@", false},
	}

	for _, testCase := range emails {

		result := CheckValidEmail(testCase.current)

		if result != nil {
			if testCase.valid != false {
				t.Fatal("Email validation output is not as expected")
			}
		} else {
			if testCase.valid != true {
				t.Fatal("Email validation output is not as expected")
			}
		}

	}

}
