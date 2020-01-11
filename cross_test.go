package cross

import (
	"testing"
)

func TestCrosspointCount(t *testing.T) {
	BasicTestCrosspointCount(t, []int{2, 4, 1, 2, 5, 3}, []int{3, 2, 3, 1}, 6)
}
func BasicTestCrosspointCount(t *testing.T, xAxis, yAxis []int, expected int) {
	if output := CrosspointCount(xAxis, yAxis); output != expected {
		t.Errorf(
			`Test Fail
			Input: 
				xAxis = %v
				yAxis = %v
			Expected: %v
			But got: %v`,
			xAxis, yAxis, expected, output,
		)
	}
}
