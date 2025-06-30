package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	reminder := [3]string{primary, secondary, tertiary}
	var cost [3]int
	cost[0] = len(primary)
	cost[1] = len(secondary) + cost[0]
	cost[2] = len(tertiary) + cost[1]
	return reminder, cost

}
