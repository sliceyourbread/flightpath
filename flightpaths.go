package main

import (
	"fmt"
	"os"
)

type travelData struct {
	index   int
	airport string
	cost    []int
}

var (
	combinations [][]int
	locations    []travelData
)

func init() {
	locations = []travelData{
		{0, "Castle Black", []int{0, 15, 80, 90}},
		{1, "Winterfell", []int{0, 0, 40, 50}},
		{2, "Riverrun", []int{0, 0, 0, 70}},
		{3, "King's Landing", []int{0, 0, 0, 0}},
	}

	for j := 0; j < len(locations); j++ {
		for i := 0; i < len(locations); i++ {
			combinations = append(combinations, []int{j, i})
		}
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	args := os.Args

	// if we receive more than two loctions return error
	// however an improvement for this piece could be adding more control over the locations
	if len(args) > 3 {
		return fmt.Errorf("to many parameters, please input two locations")
	}

	var departure, destination travelData
	for _, l := range locations {
		if l.airport == args[1] {
			departure = l
		}

		if l.airport == args[2] {
			destination = l
		}
	}

	if departure.index > destination.index {
		return fmt.Errorf("invalid location inputted, program only flys from north to south")
	}

	if departure.airport == "" || destination.airport == "" {
		return fmt.Errorf("invalid location inputted, please double check your spelling")
	}

	paths := flightPaths(departure.index, destination.index)
	costs := costCalculator(paths)
	stops := stopCalculator(paths, departure.airport)

	for i := range stops {
		fmt.Printf("%v: %v\n", stops[i], costs[i])
	}

	return nil
}

// flightPaths returns a slice  of all potential flight routes
func flightPaths(s, e int) [][]int {
	var stops [][]int
	var validPaths [][]int

	for i := 0; i < len(combinations); i++ {

		// can't go north to south
		if combinations[i][0] >= combinations[i][1] {
			continue
		}

		// direct flight
		if combinations[i][0] == s && combinations[i][1] == e {
			validPaths = append(validPaths, combinations[i])
			continue
		}

		// add all valid flights from the first stop to map
		if combinations[i][0] == s {
			stops = append(stops, combinations[i])
			continue
		}

		for j := range stops {
			// valid flight path - departure matches most recent arrival
			if stops[j][len(stops[j])-1] == combinations[i][0] {
				var tmp []int
				tmp = append(tmp, stops[j]...)
				tmp = append(tmp, combinations[i]...)
				stops = append(stops, tmp)
			}
		}

	}

	for j := 0; j < len(stops); j++ {
		if stops[j][0] == s && stops[j][len(stops[j])-1] == e {
			validPaths = append(validPaths, stops[j])
		}
	}

	return validPaths
}

func costCalculator(paths [][]int) []int {
	costs := make([]int, len(paths))

	for i := 0; i < len(paths); i++ {
		// array is stored as [1223] where 1 -> 2 is one flight
		// therefore we set j=j+2 so we skip to the next departure location
		for j := 0; j < len(paths[i]); j = j + 2 {
			if j == 0 {
				costs[i] = locations[paths[i][j]].cost[paths[i][j+1]]
			} else {
				costs[i] += locations[paths[i][j]].cost[paths[i][j+1]]

			}
		}
	}

	return costs
}

func stopCalculator(paths [][]int, dep string) []string {
	stops := make([]string, len(paths))

	for i := 0; i < len(paths); i++ {

		// array is stored as [1223] where 1 -> 2 is one flight
		// therefore we set j=j+2 so we skip to the next departure location
		for j := 0; j < len(paths[i]); j = j + 2 {

			// users will always depart from the first location provided
			if j == 0 {
				stops[i] = dep
			}

			// adding the destination stop
			stops[i] = stops[i] + " -> " + locations[paths[i][j+1]].airport
		}
	}
	return stops
}