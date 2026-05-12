func canCompleteCircuit(gas []int, cost []int) int {
    n := len(gas)
	start, end := n - 1, n - 1

	gasBalance, stationsProcessed := 0, 0

	for stationsProcessed < n {
		gasBalance += gas[end] - cost[end]
		stationsProcessed++

		end = (end + 1) % n
		for gasBalance < 0 && stationsProcessed < n {
			start--
			gasBalance += gas[start] - cost[start]
			stationsProcessed++
		}
	}
	if gasBalance < 0 {
		return -1
	}
	return start
}