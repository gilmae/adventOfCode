data = File.readlines("inputs/day10.input").map(&:chomp)

board = {}

# bit positions

#       8       4       2       1
#       N       E       S       W

NORTH = 8
EAST = 4
SOUTH = 2
WEST = 1

NEIGHBOURS = [
  [[-1, 0], SOUTH],
  [[1, 0], NORTH],
  [[0, 1], WEST],
  [[0, -1], EAST],
]

OPPOSITES = { NORTH => SOUTH, SOUTH => NORTH, EAST => WEST, WEST => EAST }

def get_exits(char)
  case char
  when "|"
    return NORTH | SOUTH
  when "-"
    return EAST | WEST
  when "L"
    return NORTH | EAST
  when "J"
    return NORTH | WEST
  when "7"
    return SOUTH | WEST
  when "F"
    return SOUTH | EAST
  when "S"
    return NORTH | SOUTH | EAST | WEST
  end
  return 0
end

def has_exit?(room, direction)
  return room != 0 && room & direction == direction
end

def get_moves(pos, board)
  room = board[pos]
  return nil if room.nil?
  moves = []
  NEIGHBOURS.each { |n|
    neighbour_pos = [pos[0] + n[0][0], pos[1] + n[0][1]]
    neighbour = board[neighbour_pos]

    moves << neighbour_pos if has_exit?(room, OPPOSITES[n[1]]) && has_exit?(neighbour, n[1])
  }
  moves
end

starting_pos = []
data.each_with_index { |line, row|
  line.chars.each_with_index { |ch, col|
    board[[row, col]] = get_exits ch
    starting_pos = [row, col] if ch == "S"
  }
}

visited = {}
path = []
rooms_to_check = [[starting_pos, 0]]
max = -1
while !rooms_to_check.empty?
  room, hops = rooms_to_check.shift
  next if visited[room]
  visited[room] = hops
  path << room
  max = hops if hops > max
  moves = get_moves room, board
  moves.each { |m|
    rooms_to_check << [m, hops + 1]
  }
end

pp max

## PART B

visited = {}
path = []
rooms_to_check = [[starting_pos, 0]]
max = -1
while !rooms_to_check.empty?
  room, hops = rooms_to_check.shift
  next if visited[room]
  visited[room] = hops
  path << room
  max = hops if hops > max
  moves = get_moves room, board
  moves.each { |m|
    rooms_to_check.unshift([m, hops + 1])
  }
end

# shoelace formula
area = (path.each_with_index.map { |p, idx|
  idx2 = (idx + 1) % path.length
  p2 = path[idx2]

  (p2[1] + p[1]) * (p2[0] - p[0])
}.sum / 2).abs()

pp area - (path.length / 2) + 1
