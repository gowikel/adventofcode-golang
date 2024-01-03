package day06

type Race struct {
	Time     int
	Distance int
}

func NewRace(time, distance int) Race {
	return Race{Time: time, Distance: distance}
}

func DistanceRan(holdTime, raceTimeLimit int) int {
	speed := holdTime
	raceTime := raceTimeLimit - holdTime

	if raceTime <= 0 {
		return 0
	}

	return speed * raceTime
}

func CountWaysToWin(r Race) int {
	var result int

	for i := 0; i <= r.Time; i++ {
		d := DistanceRan(i, r.Time)

		if d > r.Distance {
			result++
		}
	}

	return result
}
