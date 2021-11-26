package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func calculateDistance(reindeer Reindeer, elapsedTime int) int {

	cycleTime := reindeer.ActiveTime + reindeer.RestTime
	numberOfCycles := elapsedTime / cycleTime
	remainingTime := elapsedTime % cycleTime
	distance := reindeer.Speed * reindeer.ActiveTime * numberOfCycles

	if remainingTime > reindeer.ActiveTime {
		distance += reindeer.Speed * reindeer.ActiveTime
	} else {
		distance += reindeer.Speed * remainingTime
	}

	return distance

}

var inputFile = flag.String("inputFile", "inputs/day14.input", "Relative file path to use as input.")
var raceTime = flag.Int("raceTime", 2503, "How long the race should run")

type Reindeer struct {
	Name       string
	Speed      int
	ActiveTime int
	RestTime   int
}

func (r *Reindeer) isRunning(elapsedTime int) bool {
	cycleTime := r.ActiveTime + r.RestTime
	timeInCycle := elapsedTime % cycleTime

	return timeInCycle != 0 && timeInCycle <= r.ActiveTime
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")

	regex := regexp.MustCompile("(\\w+) can fly (\\d+) km/s for (\\d+) seconds, but then must rest for (\\d+) seconds")

	reindeer := make([]Reindeer, len(lines))
	for i := range lines {
		submatches := regex.FindStringSubmatch(lines[i])
		speed, _ := strconv.Atoi(submatches[2])
		active, _ := strconv.Atoi(submatches[3])
		rest, _ := strconv.Atoi(submatches[4])

		reindeer[i] = Reindeer{submatches[1], speed, active, rest}
	}
	fmt.Println(oldStyle(reindeer, *raceTime))
	fmt.Println(newStyle(reindeer, *raceTime))
}

func oldStyle(reindeer []Reindeer, elapsedTime int) int {
	longest := 0

	for idx := range reindeer {
		r := reindeer[idx]
		distance := calculateDistance(r, elapsedTime)
		if distance > longest {
			longest = distance
		}
	}

	return longest
}

func newStyle(reindeer []Reindeer, totalTime int) int {
	scores := make(map[string]int)
	accDistance := make(map[string]int)

	for i := range reindeer {
		scores[reindeer[i].Name] = 0
		accDistance[reindeer[i].Name] = 0
	}

	for i := 1; i <= totalTime; i++ {
		longest := 0
		for rx := range reindeer {

			r := reindeer[rx]

			if r.isRunning(i) {
				accDistance[r.Name] += r.Speed
			}

			if accDistance[r.Name] > longest {
				longest = accDistance[r.Name]
			}
		}

		for k, v := range accDistance {
			if v == longest {
				scores[k] += 1
			}
		}
	}

	fmt.Println(scores)
	fmt.Println(accDistance)

	highScore := 0
	for _, v := range scores {
		if v > highScore {
			highScore = v
		}
	}

	return highScore

}
