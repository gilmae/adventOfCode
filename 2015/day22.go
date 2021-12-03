package main

type effect struct {
	damage, protection, recharge, heal, duration, cost int
}

type entity struct {
	armour, hp, mana int
	isrecharging bool
	effects map[string]int
	opponent entity
}

func main() {
	spells := map[string]effect {
		{"magicmissile", effect{damage:4, cost:53, duration:0, recharge:0, heal:0}},
		{"drain", effect{cost:73, damage:2,heal:2, duration:0}},
		{"shield", effect{cost:113, protection:7, duration:6}},
		{"poison", effect{cost:173, duration:6, damage: 3}},
		{"recharge", effect{cost: 229, duration:5, recharge:101}},
	}

	/* loop
	* Am I about to die? Yes? 
	*	Do I have a drain effect healing me? No? Cast drain
			* Yes? Resign
	* If I do nothing, will my opponent die? Do nothing.
	* Do I have a recharge effect on? 
		* No? Will casting any spell take me below 229 Mana? Cast Recharge
	* Do I have a shield? No? Cast Shield
	* Do I have a drain effect on my opponent? Cast drain

}




