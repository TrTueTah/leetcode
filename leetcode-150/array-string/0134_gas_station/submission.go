package gasstation

func canCompleteCircuit(gas []int, cost []int) int {
	if sum(gas) < sum(cost) {
		return -1
	}
	start, tank := 0, 0
	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			tank = 0
			start = i + 1
		}
	}

	return start
}

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
