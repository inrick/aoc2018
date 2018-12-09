package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
)

type Event int

const (
	BeginShift Event = iota
	Wake
	Sleep
)

type Entry struct {
	t  time.Time
	ev Event
	id int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var log []Entry
	for sc.Scan() {
		line := sc.Text()
		ts, msg := line[1:17], line[19:]

		var e Entry
		var err error
		e.t, err = time.Parse("2006-01-02 15:04", ts)
		if err != nil {
			panic(err)
		}
		switch msg {
		case "wakes up":
			e.ev = Wake
		case "falls asleep":
			e.ev = Sleep
		default:
			e.ev = BeginShift
			fmt.Sscanf(msg, "Guard #%d begins shift", &e.id)
		}
		log = append(log, e)
	}
	sort.Slice(log, func(i, j int) bool { return log[i].t.Before(log[j].t) })

	// map from guard id to shift schedule
	scheds := make(map[int][]int)
	var guard int        // current guard
	var shift []int      // current shift schedule
	var asleep time.Time // when current guard fell asleep
	for _, e := range log {
		switch e.ev {
		case BeginShift:
			guard = e.id
			// init schedule array if first time seeing guard
			if shift = scheds[guard]; len(shift) == 0 {
				shift = make([]int, 60)
				scheds[guard] = shift
			}
		case Wake:
			for i := asleep.Minute(); i < e.t.Minute(); i++ {
				shift[i]++
			}
		case Sleep:
			asleep = e.t
		default:
			panic(nil)
		}
	}

	// a
	var sleepiestGuard, sleepiestMin, maxSlept int
	for guard, shift := range scheds {
		sleepiest, slept := 0, 0
		for minute, times := range shift {
			slept += times
			if times > shift[sleepiest] {
				sleepiest = minute
			}
		}
		if slept > maxSlept {
			sleepiestGuard = guard
			maxSlept = slept
			sleepiestMin = sleepiest
		}
	}
	fmt.Printf("a) %d\n", sleepiestGuard*sleepiestMin)

	// b
	var maxTimes int
	for guard, shift := range scheds {
		sleepiest := 0
		for minute := range shift {
			if shift[minute] > shift[sleepiest] {
				sleepiest = minute
			}
		}
		if shift[sleepiest] > maxTimes {
			sleepiestGuard = guard
			maxTimes = shift[sleepiest]
			sleepiestMin = sleepiest
		}
	}
	fmt.Printf("b) %d\n", sleepiestGuard*sleepiestMin)
}
