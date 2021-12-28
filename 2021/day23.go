package main

import (
	"container/heap"
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day02.input", "Relative file path to use as input.")
var part = flag.String("part", "a", "Which part to solve")

type move struct {
	from, fromSlot, to, toSlot int
	mob                        amphipod
	cost                       int
	burrow                     string
}

func (m move) String() string {
	return fmt.Sprintf("Move %s from %d[%d] to %d[%d] for cost %d", string(m.mob), m.from, m.fromSlot, m.to, m.toSlot, m.cost)
}

var EmptyMove move = move{from: 0, fromSlot: 0, to: 0, cost: 0}

type amphipod byte

func (a amphipod) MovementCost() int {
	switch a {
	case A:
		return 1
	case B:
		return 10
	case C:
		return 100
	case D:
		return 1000
	}
	return 0
}

func (a amphipod) String() string {
	switch a {
	case A:
		return "A"
	case B:
		return "B"
	case C:
		return "C"
	case D:
		return "D"
	}
	return "."
}

type Room struct {
	habitat amphipod
	spots   []amphipod
}

func (r Room) copy() Room {
	resp := Room{habitat: r.habitat, spots: make([]amphipod, len(r.spots))}
	for i, a := range r.spots {
		resp.spots[i] = a
	}
	return resp
}

func (r *Room) isComplete() bool {
	for _, m := range r.spots {
		if m != r.habitat {
			return false
		}
	}
	return true
}

func (r *Room) freeSpot() int {
	for i := len(r.spots) - 1; i >= 0; i-- {
		if r.spots[i] == Free {
			return i
		}
	}
	return -1
}

func (r *Room) isFree(a amphipod) bool {
	if r.habitat == Free {
		return r.freeSpot() != -1
	} else {
		for _, o := range r.spots {
			if o != a && o != Free {
				return false
			}
		}
		return true
	}
}

type Burrow [9]Room

func (b Burrow) Print() {
	fmt.Println("#############")

	// Print hallway
	fmt.Print("#")
	fmt.Printf("%s%s", b[0].spots[1], b[0].spots[0])
	for i := 1; i < 8; i++ {
		if i%2 == 0 {
			fmt.Printf("%s", b[i].spots[0])
		} else {
			fmt.Print(".")
		}
	}
	fmt.Printf("%s%s#\n", b[8].spots[0], b[8].spots[1])

	// Print Home Rooms Line 1

	fmt.Printf("###%s#%s#%s#%s###\n", b[1].spots[0], b[3].spots[0], b[5].spots[0], b[7].spots[0])
	fmt.Printf("  #%s#%s#%s#%s#  \n", b[1].spots[1], b[3].spots[1], b[5].spots[1], b[7].spots[1])
	fmt.Println("  #########  ")
}

func (b Burrow) Copy() Burrow {
	resp := Burrow{}
	for i, r := range b {
		resp[i] = r.copy()
	}

	return resp
}

func (b Burrow) hasOpenPath(mob amphipod, from, to int) bool {
	if !burrow[to].isFree(mob) {
		return false
	}
	if to > from {
		for i := from + 1; i < to; i++ {
			// If trying to go through a stash spot and it is already occpied, then path is blocked.
			// Note, this only looks at the sstash spots _between_ home rooms. The ends of the corridor are not looked at
			if b[i].habitat == Free && !b[i].isFree(mob) {
				return false
			}
		}
	} else {
		for i := from - 1; i > to; i-- {
			// If trying to go through a stash spot and it is already occpied, then path is blocked.
			// Note, this only looks at the sstash spots _between_ home rooms. The ends of the corridor are not looked at
			if b[i].habitat == Free && !b[i].isFree(mob) {
				return false
			}
		}
	}

	if to == 0 && b[to].spots[0] != Free {
		return false
	}
	if to == 8 && b[to].spots[0] != Free {
		return false
	}
	return true
}

func (b Burrow) travelCost(from, spot, to int) int {
	targetSpot := b[to].freeSpot()
	return b.travelCostToSpot(from, spot, to, targetSpot)
}
func (b Burrow) travelCostToSpot(from, spot, to, targetSpot int) int {
	cost := spot + targetSpot
	if to%2 == 1 {
		cost += 1 // The home burrows cost an extra step to get into
	}
	if from%2 == 1 {
		cost += 1 // The home burrows cost an extra step to get into
	}
	cost += int(math.Abs(float64(from) - float64(to)))
	return cost
}

func (b Burrow) applyMove(m move) Burrow {
	resp := b.Copy()
	mob := resp[m.from].spots[m.fromSlot]
	//fmt.Printf("Moving from %d->%d to %d->%d\n", m.from, m.fromSlot, m.to, toSlot)
	resp[m.to].spots[m.toSlot] = mob
	resp[m.from].spots[m.fromSlot] = Free

	return resp
}

func (b Burrow) MinCostToAllHome(homeRooms map[amphipod]int) int {
	cost := 0
	for amphipod, homeroom := range homeRooms {
		slot := 0
		for roomNum, room := range b {
			for slowNum, mob := range room.spots {
				if mob == amphipod {
					cost += mob.MovementCost() * b.travelCostToSpot(roomNum, slowNum, homeroom, slot)
					slot++
				}
			}
		}
	}

	return cost
}

func NewBurrow() Burrow {
	burrow := Burrow{}
	burrow[0] = Room{habitat: Free, spots: make([]amphipod, 2)}

	burrow[1] = Room{habitat: A, spots: make([]amphipod, 2)}
	burrow[2] = Room{habitat: Free, spots: make([]amphipod, 1)}
	burrow[3] = Room{habitat: B, spots: make([]amphipod, 2)}
	burrow[4] = Room{habitat: Free, spots: make([]amphipod, 1)}

	burrow[5] = Room{habitat: C, spots: make([]amphipod, 2)}
	burrow[6] = Room{habitat: Free, spots: make([]amphipod, 1)}
	burrow[7] = Room{habitat: D, spots: make([]amphipod, 2)}

	burrow[8] = Room{habitat: Free, spots: make([]amphipod, 2)}

	return burrow
}

func CreateBurrowFromKey(key string) Burrow {

	resp := NewBurrow()

	for _, mobKey := range strings.Split(key, "::") {
		parts := strings.Split(mobKey, ":")

		room, _ := strconv.Atoi(parts[1])
		spot, _ := strconv.Atoi(parts[2])
		bc, _ := strconv.Atoi(parts[0])
		resp[room].spots[spot] = amphipod(byte(bc))
	}

	return resp
}

func (b Burrow) key() string {
	keys := make([]string, 8)
	i := 0
	for amphipod, _ := range homeRooms {
		for roomNum, room := range b {
			for slowNum, mob := range room.spots {
				if mob == amphipod {
					keys[i] = fmt.Sprintf("%d:%d:%d", amphipod, roomNum, slowNum)
					i++
				}
			}
		}
	}
	return strings.Join(keys, "::")
}

const (
	A    amphipod = 'A'
	B    amphipod = 'B'
	C    amphipod = 'C'
	D    amphipod = 'D'
	Free amphipod = 0
)

var burrow Burrow
var homeRooms map[amphipod]int

func main() {
	flag.Parse()
	// bytes, err := ioutil.ReadFile(*inputFile)
	// if err != nil {
	// 	return
	// }

	// contents := string(bytes)
	// lines := strings.Split(contents, "\n")

	burrow = NewBurrow()
	burrow[1].spots[0] = D
	burrow[1].spots[1] = D

	burrow[3].spots[0] = B
	burrow[3].spots[1] = A

	burrow[5].spots[0] = C
	burrow[5].spots[1] = B

	burrow[7].spots[0] = C
	burrow[7].spots[1] = A

	homeRooms = map[amphipod]int{A: 1, B: 3, C: 5, D: 7}
	EmptyMove = move{burrow: burrow.key(), cost: 0}
	score := AStarBurrow(EmptyMove)
	fmt.Println(score)
}

func AllHome(b Burrow, homeRooms map[amphipod]int) bool {
	for _, room := range homeRooms {
		if !b[room].isComplete() {
			return false
		}
	}

	return true
}

func GetAllValidMoves(burrow Burrow) []move {
	allMoves := make([]move, 0)
	for i, room := range burrow {
		for spot := range room.spots {
			allMoves = append(allMoves, burrow.validDestinations(i, spot)...)
		}
	}
	return allMoves
}
func (b Burrow) validDestinations(room int, slot int) []move {
	destinations := make([]move, 0)
	mob := b[room].spots[slot]
	key := b.key()

	if mob == Free {
		return destinations // This is not a mob
	}

	energyCost := mob.MovementCost()
	if slot > 0 && burrow[room].spots[slot-1] != Free {
		return destinations // Blocked from moving by mob in front
	}
	// if burrow[room].freeSpot() != slot-1 {
	// 	return destinations
	// }
	if burrow[room].habitat == mob && burrow[room].isFree(mob) { // mob is home
		return destinations
	}
	if burrow[room].habitat == Free { // Mob is in the hall

		homeRoom := homeRooms[mob]
		if burrow[homeRoom].isFree(mob) && burrow.hasOpenPath(mob, room, homeRoom) {

			destinations = append(destinations, move{burrow: key, from: room, fromSlot: slot, to: homeRoom, toSlot: burrow[homeRoom].freeSpot(), mob: mob, cost: energyCost * burrow.travelCost(room, slot, homeRoom)})
		}
	} else {
		for i, r := range burrow {
			if r.habitat != Free && r.habitat != mob {
				// ignore other home rooms
				continue
			}
			if burrow.hasOpenPath(mob, room, i) {
				for spot := range r.spots {
					if r.spots[spot] == Free {
						destinations = append(destinations, move{burrow: key, from: room, fromSlot: slot, to: i, toSlot: spot, mob: mob, cost: energyCost * burrow.travelCostToSpot(room, slot, i, spot)})
					}
				}

			}
		}
	}
	return destinations
}

type CostMap map[move]int
type HeapQueue struct {
	Elems            *[]move
	Score, Positions CostMap
}

func (h HeapQueue) Len() int           { return len(*h.Elems) }
func (h HeapQueue) Less(i, j int) bool { return h.Score[(*h.Elems)[i]] < h.Score[(*h.Elems)[j]] }
func (h HeapQueue) Swap(i, j int) {
	h.Positions[(*h.Elems)[i]], h.Positions[(*h.Elems)[j]] = h.Positions[(*h.Elems)[j]], h.Positions[(*h.Elems)[i]]
	(*h.Elems)[i], (*h.Elems)[j] = (*h.Elems)[j], (*h.Elems)[i]
}

func (h HeapQueue) Push(x interface{}) {
	h.Positions[x.(move)] = len(*h.Elems)
	*h.Elems = append(*h.Elems, x.(move))
}

func (h HeapQueue) Pop() interface{} {
	old := *h.Elems
	n := len(old)
	x := old[n-1]
	*h.Elems = old[0 : n-1]
	delete(h.Positions, x)
	return x
}

func (h HeapQueue) Position(x move) int {
	if pos, ok := h.Positions[x]; ok {
		return pos
	}
	return -1
}

func AStarBurrow(src move) int {

	start := CreateBurrowFromKey(src.burrow)
	gScore := map[move]int{src: 0}
	fScore := map[move]int{src: start.MinCostToAllHome(homeRooms)}

	path := make(map[move]move)
	work := HeapQueue{&[]move{src}, fScore, make(CostMap)}
	heap.Init(&work)

	for len(*work.Elems) != 0 {
		current := heap.Pop(&work).(move)

		burrow = CreateBurrowFromKey(current.burrow)
		if current != EmptyMove {
			burrow = burrow.applyMove(current)
		}

		if AllHome(burrow, homeRooms) {
			moves := make([]move, 0)
			// burrow.Print()
			// // We're there, backtrace through path to calculate score
			// score := 0
			score := gScore[current]
			for current != src {
				moves = append(moves, current)
				current = path[current]
			}

			burrow := CreateBurrowFromKey(src.burrow)
			burrow.Print()
			fmt.Println()

			for i := len(moves) - 1; i >= 0; i-- {
				move := moves[i]
				burrow = burrow.applyMove(move)

				burrow.Print()
				fmt.Printf("Move %s from %d[%d] to %d[%d]\n", string(move.mob), move.from, move.fromSlot, move.to, move.toSlot)
				fmt.Println("costing: ", move.cost)
				fmt.Println()
			}
			return score
		}
		validMoves := GetAllValidMoves(burrow)
		// if len(validMoves) == 0 {
		// 	burrow.Print()
		// }
		for _, n := range validMoves {
			tentativeScore := gScore[current] + n.cost
			if previousScore, ok := gScore[n]; !ok || tentativeScore < previousScore {
				path[n] = current
				gScore[n] = tentativeScore
				fScore[n] = tentativeScore + burrow.applyMove(n).MinCostToAllHome(homeRooms)
				if pos := work.Position(n); pos == -1 {
					heap.Push(&work, n)
				} else {
					heap.Fix(&work, pos)
				}
			}
		}

	}

	return -1
}
