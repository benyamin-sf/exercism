package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var totalBirds int

    for _, birdsCount := range birdsPerDay {
        totalBirds += birdsCount
    }
    return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var totalBirds int
	weekStart := (week - 1) * 7
	weekEnd := weekStart + 7

	if weekEnd > len(birdsPerDay) {
		weekEnd = len(birdsPerDay)
	}

	for _, v := range birdsPerDay[weekStart:weekEnd] {
		totalBirds += v
	}
	return totalBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for idx := range birdsPerDay {
        if idx%2 == 0 {
            birdsPerDay[idx]++
        }
    }
    return birdsPerDay
}
