input = ARGV[0] || "inputs/day19.input"
data = File.readlines(input).map(&:chomp)

ORE = 0
CLAY = 1
OBSIDIAN = 2
GEODE = 3

ORE_FOR_ORE_BOT = 1
ORE_FOR_CLAY_BOT = 2
ORE_FOR_OBSIDIAN_BOT = 3
CLAY_FOR_OBSIDIAN_BOT = 4
ORE_FOR_GEODE_BOT = 5
OBSIDIAN_FOR_GEODE_BOT = 6

State = Struct.new(:time,
                   :ore, :clay, :obsidian, :geode,
                   :oreBot, :clayBot, :obsidianBot, :geodeBot)

blueprints = data.map { |d|
  d.scan(/\d+/).map(&:to_i)
}

def work_blueprint(blueprint, time)
  work = [State.new(time, 0, 0, 0, 0, 1, 0, 0, 0)]
  most = 0
  visited = {}
  count = 0
  while (work.size > 0)
    state = work.shift
    if (visited.include? state)
      next
    end

    visited[state] = true

    if state.time == 0
      if most < state.geode
        most = state.geode
        next
      end
    else
      next_state = State.new(state.time - 1,
                             state.ore + state.oreBot,
                             state.clay + state.clayBot,
                             state.obsidian + state.obsidianBot,
                             state.geode + state.geodeBot,
                             state.oreBot,
                             state.clayBot,
                             state.obsidianBot,
                             state.geodeBot)

      # In general, don't bother building bots to mine more of a resource than we can use in a single turn.
      # If we can build a Geode bot, do that and nothing else
      # If we can build a Obsidian bot, build one rather than building a clay or ore bot or building nothing
      # And if we have lots of resources, build something and ignore the noop timeline
      buildGeodeBot = (state.ore >= blueprint[ORE_FOR_GEODE_BOT] &&
                       state.obsidian >= blueprint[OBSIDIAN_FOR_GEODE_BOT])
      buildObsidianBot = !buildGeodeBot &&
                         (state.obsidianBot <= blueprint[OBSIDIAN_FOR_GEODE_BOT]) &&
                         (state.ore >= blueprint[ORE_FOR_OBSIDIAN_BOT] &&
                          state.clay >= blueprint[CLAY_FOR_OBSIDIAN_BOT])

      buildClayBot = (state.ore >= blueprint[ORE_FOR_CLAY_BOT]) &&
                     !buildGeodeBot && #!buildObsidianBot && # Preferencing obsidian over clay worked for Part 1, gave wrong answer for part 2
                     (state.clayBot <= blueprint[CLAY_FOR_OBSIDIAN_BOT])

      buildOreBot = (state.ore >= blueprint[ORE_FOR_ORE_BOT]) &&
                    (state.oreBot <= [blueprint[ORE_FOR_CLAY_BOT], blueprint[ORE_FOR_OBSIDIAN_BOT], blueprint[ORE_FOR_GEODE_BOT]].max) &&
                    !buildGeodeBot && !buildObsidianBot

      noop = !buildGeodeBot && !buildObsidianBot &&
             (state.ore < 2 * [blueprint[ORE_FOR_CLAY_BOT], blueprint[ORE_FOR_OBSIDIAN_BOT], blueprint[ORE_FOR_GEODE_BOT]].max) &&
             (state.clay < 2 * blueprint[CLAY_FOR_OBSIDIAN_BOT])

      if noop
        work << next_state
      end

      work << State.new(next_state.time,
                        next_state.ore - blueprint[ORE_FOR_GEODE_BOT],
                        next_state.clay,
                        next_state.obsidian - blueprint[OBSIDIAN_FOR_GEODE_BOT],
                        next_state.geode,
                        next_state.oreBot,
                        next_state.clayBot,
                        next_state.obsidianBot,
                        next_state.geodeBot + 1) if buildGeodeBot

      work << State.new(next_state.time,
                        next_state.ore - blueprint[ORE_FOR_OBSIDIAN_BOT],
                        next_state.clay - blueprint[CLAY_FOR_OBSIDIAN_BOT],
                        next_state.obsidian,
                        next_state.geode,
                        next_state.oreBot,
                        next_state.clayBot,
                        next_state.obsidianBot + 1,
                        next_state.geodeBot) if buildObsidianBot

      work << State.new(next_state.time,
                        next_state.ore - blueprint[ORE_FOR_CLAY_BOT],
                        next_state.clay,
                        next_state.obsidian,
                        next_state.geode,
                        next_state.oreBot,
                        next_state.clayBot + 1,
                        next_state.obsidianBot,
                        next_state.geodeBot) if buildClayBot

      work << State.new(next_state.time,
                        next_state.ore - blueprint[ORE_FOR_ORE_BOT],
                        next_state.clay,
                        next_state.obsidian,
                        next_state.geode,
                        next_state.oreBot + 1,
                        next_state.clayBot,
                        next_state.obsidianBot,
                        next_state.geodeBot) if buildOreBot
    end
  end
  most
end

pp blueprints.map { |blueprint|
  work_blueprint(blueprint, 24) * blueprint[0]
}.sum

pp blueprints[0..2].map { |blueprint|
  work_blueprint(blueprint, 32)
}.reduce(1, &:*)
