package cross

func CrosspointCount(xAxis, yAxis []int) int {
	crosspointCount := 0
	for yLinePos, yLineLength := range yAxis {
		limit := yLineLength
		if len(xAxis) < yLineLength {
			limit = len(xAxis)
		}
		for i := 0; i < limit; i++ {
			if xAxis[i] < yLinePos+1 {
				continue
			}
			crosspointCount++
		}
	}
	return crosspointCount
}
