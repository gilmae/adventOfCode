require_relative "utils"
input = ARGV[0] || "inputs/day24.input"
data = File.readlines(input).map(&:chomp)

board = {}
WEATHER_CONDITIONS = []

State24 = Struct.new(:steps, :pos)

MINX = 0
MAXX = 1
MINY = 2
MAXY = 3
BOUNDS = [1, data[0].length - 2, 1, data.length - 2] # There are no storms that can move into the start or end position

# Read Board
data.each_with_index { |line, y|
  line.chars.each_with_index { |ch, x|
    next if ch == "#"
    point = [x, y]
    case ch
    when "."
      board[point] = nil
    when "<"
      board[point] = [[-1, 0]]
    when ">"
      board[point] = [[1, 0]]
    when "^"
      board[point] = [[-0, -1]]
    when "v"
      board[point] = [[0, 1]]
    end
  }
}

def update_board(board)
  newboard = {}
  board.each { |k, v|
    if v == nil
      newboard[k] = nil unless newboard.include? k
      next
    end
    px, py = k
    v.each { |dx, dy|
      n = [px + dx, py + dy]

      # Wrap
      n[0] = BOUNDS[MAXX] if n[0] < BOUNDS[MINX]
      n[0] = BOUNDS[MINX] if n[0] > BOUNDS[MAXX]

      n[1] = BOUNDS[MAXY] if n[1] < BOUNDS[MINY]
      n[1] = BOUNDS[MINY] if n[1] > BOUNDS[MAXY]

      newboard[n] = [] if newboard[n] == nil
      newboard[n] << [dx, dy]
      newboard[k] = nil unless newboard.include? k
    }
  }
  newboard
end

# Trust me, the number of unique weather conditions is the lowest common multiplier
# of the board dimensions. I've checked
WEATHER_CYCLES = lcm BOUNDS[MAXX], BOUNDS[MAXY]
WEATHER_CYCLES.times {
  WEATHER_CONDITIONS << board
  board = update_board board
}

def get_next_positions(cur, board)
  x, y = cur
  if cur == [0, 1]
    positions = [[1, 1]]
  else
    positions = []
    positions << [x, y + 1]
    positions << [x, y - 1]
    positions << [x + 1, y]
    positions << [x - 1, y]
  end
  positions << [x, y]
  positions.delete_if { |p| !board.include?(p) || board[p] != nil }
end

def travel(start, dest, start_time)
  work = []
  work << State24.new(0, start)
  quickest = 1e15
  visited = {}

  while !work.empty?
    cur_state = work.shift
    cur = cur_state.pos

    steps = cur_state.steps
    next if steps >= quickest

    next if visited.include?([cur, steps])
    visited[[cur, steps]] = true

    if cur == dest
      if steps < quickest
        quickest = steps
      end
      next
    end
    next_board = WEATHER_CONDITIONS[(start_time + steps + 1) % WEATHER_CYCLES]
    positions = get_next_positions(cur, next_board)
    positions.each { |n|
      work << State24.new(steps + 1, n)
    }
  end
  quickest
end

start = [data[0].index("."), 0]
dest = [data.last.index("."), data.length - 1]
there = travel start, dest, 0
pp there

whoops_forgot_something = travel dest, start, there
and_back_again = travel start, dest, there + whoops_forgot_something

pp and_back_again + whoops_forgot_something + there
