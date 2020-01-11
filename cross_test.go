package cross

import (
	"testing"
)

func TestCrosspointCount(t *testing.T) {
	BasicTestCrosspointCount(t, []int{2, 4, 1, 2, 5, 3}, []int{3, 2, 3, 1}, 6)
	BasicTestCrosspointCount(t, []int{1, 1}, []int{1, 1}, 1)
	BasicTestCrosspointCount(t, []int{1, 1}, []int{100}, 2)
	BasicTestCrosspointCount(t, []int{2, 2, 2}, []int{10, 10, 10}, 6)
}
func BasicTestCrosspointCount(t *testing.T, xAxis, yAxis []int, expected int) {
	if output := CrosspointCount(xAxis, yAxis); output != expected {
		t.Errorf(`
			Test Fail
			Input: 
				xAxis = %v
				yAxis = %v
			Expected: %v
			But got: %v`,
			xAxis, yAxis, expected, output,
		)
	}
}
func TestCrosspointCountString(t *testing.T) {
	BasicTestCrosspointCountString(t, "2 4 1 2 5 3-3 2 3 1", 6, false)
	BasicTestCrosspointCountString(t, "2 4 1 2 5 3-a 2 3 1", 6, true)
	BasicTestCrosspointCountString(t, "2 4 1 2 5 3 3 2 3 1", 6, true)
	BasicTestCrosspointCountString(t, "2 4 1 2 5 3/3 2 3 1", 6, true)
}
func BasicTestCrosspointCountString(t *testing.T, input string, expected int, expectingErr bool) {
	output, err := CrosspointCountString(input)
	if expectingErr != (err != nil) {
		t.Errorf(`
			Test Fail
			Expecting error = %v
			But got error: %v`,
			expectingErr, err,
		)
	}
	if output != expected && !expectingErr {
		t.Errorf(`
			Test Fail
			Input: %s
			Expected: %v
			But got: %v`,
			input, expected, output,
		)
	}
}
