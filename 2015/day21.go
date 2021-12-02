package main

import "fmt"

type mob struct {
	hp, damage, ac int
}

type item struct {
	cost, buff int
}

type ring struct {
	cost, buff,
	acbuff int
}

// I think this is what is the cheapest way to get to kill rate is quicker than boss kill rate
// DPS is damage - (other) ac, minimum of 1
// kill rate is HP / (other) DPS
// If my kill rate is slower than boss's kill rate, I win

func main() {

	boss := mob{100, 8, 2}
	weapons := []item{
		item{8, 4},
		item{10, 5},
		item{25, 6},
		item{40, 7},
		item{74, 8},
	}
	armour := []item{
		item{0, 0},
		item{13, 1},
		item{31, 2},
		item{53, 3},
		item{75, 4},
		item{102, 5},
	}

	rings := []ring{
		{0, 0, 0},
		{0, 0, 0},
		{25, 1, 0},
		{50, 2, 0},
		{100, 3, 0},
		{20, 0, 1},
		{40, 0, 2},
		{80, 0, 3},
	}

	cheapest := 999999999
	expensivist := 0
	for _, w := range weapons {
		for _, a := range armour {
			for i, r1 := range rings {
				for j, r2 := range rings {
					if i == j {
						continue // Can't wear the same ring twice
					}
					myDps := getDps(w.buff+r1.buff+r2.buff, boss.ac)

					bossDps := getDps(8, a.buff+r1.acbuff+r2.acbuff)

					myKillRate := boss.hp / myDps
					bossKillRate := 100 / bossDps
					cost := w.cost + a.cost + r1.cost + r2.cost
					if myKillRate <= bossKillRate {

						if cost < cheapest {
							cheapest = cost
							fmt.Println(cost, w, a, r1, r2, myDps, bossDps, myKillRate, bossKillRate)
						}
					} else {
						if cost > expensivist {
							expensivist = cost
							fmt.Println(cost, w, a, r1, r2, myDps, bossDps, myKillRate, bossKillRate)
						}
					}

				}
			}
		}
	}

	fmt.Println(cheapest, expensivist)
}

func getDps(damage, ac int) int {
	dps := damage - ac
	if dps < 1 {
		return 1
	}
	return dps
}
