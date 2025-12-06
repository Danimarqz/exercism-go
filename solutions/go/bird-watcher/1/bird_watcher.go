package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var count int
	for i:= range birdsPerDay {
		count += birdsPerDay[i]
	}
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	initDay := 7 * (week - 1)
	lastDay := 7 * week
	var count int
	for i:= initDay; i < lastDay; i++ {
		count += birdsPerDay[i]
	}
	return count
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	var fixedArray []int = birdsPerDay
	for i:= 0; i < len(birdsPerDay); i += 2 {
		fixedArray[i]++
	}
	return fixedArray
}
