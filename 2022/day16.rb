input = ARGV[0] || "inputs/day16.input"
data = File.readlines(input).map(&:chomp)

def deep_clone(o)
  Marshal.load(Marshal.dump(o))
end

class Room
  attr_accessor :valve_flow, :name, :exits

  def initialize(name, flow, exits)
    @name = name
    @valve_flow = flow
    @exits = exits
  end
end

class State
  attr_accessor :open, :time
end

rooms = {}
interesting_valves = ["AA"]
valves_off = []
data.each { |d|
  name, flow, exits = d.scan(/Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? ([\w,\s]+)/)[0]
  valves_off << name
  rooms[name] = Room.new name, flow.to_i, exits.split(",").map(&:strip)
  interesting_valves << name if flow.to_i > 0
}

def get_shortest_path(rooms, start, dest)
  work = [[start, []]]

  visited = {}
  quickest = nil
  while !work.empty?
    current, steps = work.shift

    next if visited[current]

    quickest = steps if current == dest && (quickest == nil || steps.length < quickest.length)

    visited[current] = true
    if rooms[current] == nil
      pp current
      exit
    end
    rooms[current].exits.each { |target|
      work.append [target, steps + [current]]
    }
  end

  return quickest[1..-1]
end

paths = {}
interesting_valves.each { |s|
  paths[s] = {}
  interesting_valves.each { |d|
    next if s == d
    paths[s][d] = get_shortest_path rooms, s, d
  }
}

def where_we_are_going_we_do_not_need_roads(a, b)
  return [a, b] if b == nil

  return [a, b] if a[0] > b[0]
  return [b, a]
end

def possible_flow_left(opened_valves, valves, time_left)
  valves.each { |name, valve|
    opened_valves.include?(name) ? 0 : valve.valve_flow
  }.sum * time_left
end

def release_pressure(a, b, paths, valves, opened_valves, run_time)
  now, future = where_we_are_going_we_do_not_need_roads a, b
  now, current = now

  pressures_relieved = []
  paths[current].each { |next_room, path|
    next if opened_valves.include? next_room
    steps = path.length + 2  # go to valve and open it
    next if steps > now
    next_now = now - steps
    pressures_relieved << release_pressure([next_now, next_room],
                                           future,
                                           paths, valves,
                                           {}.replace(opened_valves).merge({ next_room => next_now }),
                                           run_time)
  }
  return pressures_relieved.max if !pressures_relieved.empty?
  # no moves left
  return opened_valves.map { |valve, opened_at|
           valves[valve].valve_flow * opened_at
         }.sum
end

pp release_pressure([30, "AA"], nil, paths, rooms, {}, 30)
pp release_pressure([26, "AA"], [26, "AA"], paths, rooms, {}, 26)
