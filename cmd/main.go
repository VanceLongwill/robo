package main

import (
	"context"
	"log"
	"math"
	"robo"
	"time"
)

// MaxSum is the maximum allowed value for the sum of all digits
var MaxSum int64 = 23

// getSumOfDigits calculates the sum of all the individual digits in a base 10 number i.e. getSumOfDigits(123) => 6
func getSumOfDigits(n int64) int64 {
	var sum int64
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// abs gets the absolute value of an int64
func abs(n int64) int64 {
	return int64(math.Abs(float64(n)))
}

// isTravellable determines whether a node is safe or not
func isTravellable(p robo.Node) bool {
	// NB since we're using the absolute value here, we could make the exploration more efficient by only finding 1st quadrant results (i.e. x >= 0, y >= 0)
	// 	and multiplying the resulting area by 4 (as the same area would be reflected in all 4 quadrants)
	return getSumOfDigits(abs(p.X))+getSumOfDigits(abs(p.Y)) <= MaxSum
}

// findSafeNodes finds accessible nodes adjacent to a base node
func findSafeNodes(n robo.Node) []robo.Adjacent {
	var adjs []robo.Adjacent
	for d := range robo.Directions {
		node := robo.Navigate(n, d)
		if isTravellable(node) {
			adjs = append(adjs, robo.Adjacent{Node: node, Direction: d})
		}
	}
	return adjs
}

func main() {
	log.Println("Starting robot...")
	// start a context we could use to cancel the search
	ctx := context.Background()

	// create a ticker which will be used to poll the state of the robot
	t := time.NewTicker(time.Second / 10)

	done := make(chan bool)

	r := robo.NewDefaultRobot()
	// our custom scanner which discovers which Nodes are accessible
	scanner := robo.ScannerFunc(findSafeNodes)

	// initialise a DFS search to find all accessible nodes
	d := robo.NewDFS(robo.WithScanner(scanner), robo.WithRobot(r))

	var start time.Time
	// use a separate go routing for doing the search work to allow checking of progress
	go func() {
		start = time.Now()
		// run the search to completion, or until an error occurrs or the context expires
		if err := d.Run(ctx); err != nil {
			log.Fatalln(err)
		}
		done <- true
	}()

	for {
		select {
		case <-ctx.Done():
			log.Printf("Paused")
			t.Stop()
			// @TODO: take advantage of the context mechanism to allow the user to pause and restart the search without recalculation
			return
		case <-done:
			elapsed := time.Since(start)
			log.Printf("Finished exploring area in %d milliseconds. Total area: %d, Total distance travelled: %d", elapsed.Milliseconds(), d.Visited().Len(), r.DistanceTravelled())
			t.Stop()
			return
		case <-t.C:
			log.Printf("Current position: %s, Total distance travelled: %d", r.Position(), r.DistanceTravelled())
		}
	}
}
