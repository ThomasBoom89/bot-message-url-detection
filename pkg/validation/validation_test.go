package validation

import "testing"

func TestValidator_IsParseable(t *testing.T) {
	cases := []struct {
		url      string
		expected bool
	}{
		{"duckduckgo.com/?q=github&t=h_&ia=web", true},
		{"duckduckgo.com/?q=github&t=h_&ia=web#element", true},
		{"https://duckduckgo.com/q=github&t=h_&ia=web", false},
		{"htps://duckduckgo.com//q=github&t=h_&ia=web", false},
		{"https://com//q=github&t=h_&ia=web", false},
		{"", false},
	}

	for _, c := range cases {
		detector := new(Validator)
		result, err := detector.IsParseable(c.url)
		if err != nil {
			t.Log("should not throw err: ", err)
		}
		if result != c.expected {
			t.Log("result should be equal: ", result, c.expected)
		}
	}
}

func TestValidator_IsReachable(t *testing.T) {
	cases := []struct {
		url      string
		expected bool
	}{
		{"duckduckgo.com/?q=github&t=h_&ia=web", true},
		{"duckduckgo.com/?q=github&t=h_&ia=web#element", true},
		{"https://duckduckgo.com/q=github&t=h_&ia=web", true},
		{"htps://duckduckgo.com//q=github&t=h_&ia=web", false},
		{"https://com//q=github&t=h_&ia=web", false},
		{"", false},
	}

	for _, c := range cases {
		detector := new(Validator)
		result, err := detector.IsParseable(c.url)
		if err != nil {
			t.Log("should not throw err: ", err)
		}
		if result != c.expected {
			t.Log("result should be equal: ", result, c.expected)
		}
	}
}
