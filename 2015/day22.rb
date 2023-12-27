=begin
    Every round has an attacker (alternate between player and Boss)
    In player round, player can cast a spell. Spell has a duration
    At start of every round apply damage from spell effects
    At start of every round reduce duration of spell effects
    At end of Boss round, apply Boss damage

=end

#spell [cost, damage, healing, ac, manaheal, duration]
COST = 0
DAMAGE = 1
HEALING = 2
AC = 3
MANA = 4
DURATION = 5

SPELLS = {
  "mmissile" => [53, 4, 0, 0, 0, 0],
  "drain" => [73, 2, 2, 0, 0, 0],
  "shield" => [113, 0, 0, 7, 0, 6],
  "poison" => [173, 3, 0, 0, 0, 6],
  "recharge" => [229, 0, 0, 0, 101, 5],
}

#effects = {name=>roundsleft}
effects = []

# Queue job [hp, mana, boss hp, boss damage, effects, dps, player_turn, cost, turns]
queue = [[50, 500, 71, 10, {}, [], true, 0, 1]]
#queue = [[10, 250, 13, 8, {}, [], true, 0, 1]]
#queue = [[10, 250, 14, 8, {}, [], true, 0, 1]]
spell_script = ["recharge", "shield", "drain", "poison", "mmissile"]

inc = 0

def play(starting_hp, starting_mana, starting_boss_hp, starting_boss_damage, is_hard)
  min_cost = 1e9
  min_script = nil

  queue = [[starting_hp, starting_mana, starting_boss_hp, starting_boss_damage, {}, [], true, 0, 1]]
  while !queue.empty?
    job = queue.shift

    hp, mana, boss_hp, boss_damage, effects, dps, player_turn, cost, turns = job
    next if cost > min_cost
    if is_hard and player_turn
      hp -= 1
      if hp <= 0
        #puts "We died with the boss on #{boss_hp} HP after #{turns} turns, they did #{damage} damage leaving us on #{hp - damage} hp"
        #pp dps + [[damage, boss_damage, ac, hp - damage]] if dps[0] == "recharge"
        #puts

        next
      end
    end
    ac = 0
    next_effects = []
    # Apply effects

    effects.each { |name, duration|
      effect = SPELLS[name]
      hp += effect[HEALING]
      boss_hp -= effect[DAMAGE]
      ac += effect[AC]
      ac = [7, ac].min
      mana += effect[MANA]
      duration -= 1

      next_effects << [name, duration] if duration > 0
    }

    if boss_hp <= 0
      if cost < min_cost
        min_cost = cost
        min_script = dps
        #puts "Killed the boss for #{min_cost} mana after #{turns - 1} turns with #{hp} hp remaining"
      end
      next
    end

    if player_turn
      available_spells = (SPELLS.keys - next_effects.map { |name, _| name })
      available_spells = available_spells.filter { |s| SPELLS[s][COST] <= mana }
      if available_spells.length == 0
        #puts "We died with 0 mana left to cast available spells #{available_spells} after #{turns} turns"
        next
      end
      # Queue job [hp, mana, boss hp, boss damage, effects, dps, player_turn, cost]

      available_spells.each { |name|

        #    name = spell_script[(turns / 2.0).floor]
        spell = SPELLS[name]

        next_job = [hp, mana - spell[COST], boss_hp, boss_damage, next_effects, dps + [name], false, cost + spell[COST], turns + 1]

        # If it is an instant spell, apply now
        if spell[DURATION] == 0
          if (boss_hp - spell[DAMAGE]) <= 0
            if cost + spell[COST] < min_cost
              min_cost = cost + spell[COST]
              min_script = dps

              #puts "Killed the boss for #{min_cost} mana after #{turns} turns with #{hp} remaining."
            end
            next
          end
          next_job[0] += spell[HEALING]
          next_job[2] -= spell[DAMAGE]
        else
          ne = next_effects + [[name, spell[DURATION]]]
          next_job[4] = ne
        end
        queue.unshift next_job
      }
    else
      damage = [boss_damage - ac, 1].max
      if hp - damage <= 0
        #puts "We died with the boss on #{boss_hp} HP after #{turns} turns, they did #{damage} damage leaving us on #{hp - damage} hp"
        #pp dps + [[damage, boss_damage, ac, hp - damage]] if dps[0] == "recharge"
        #puts

        next
      end
      queue.unshift [hp - damage, mana, boss_hp, boss_damage, next_effects, dps + [[damage, boss_damage, ac, hp - damage]], true, cost, turns + 1]
    end
  end
  min_cost
end

pp play(50, 500, 71, 10, false)
pp play(50, 500, 71, 10, true)
#pp min_script
