input = ARGV[0] || "inputs/day23.input"
data = File.readlines(input).map(&:chomp)

NORTH = 0
EAST = 3
SOUTH = 1
WEST = 2

DIRECTIONS = [NORTH, SOUTH, WEST, EAST]

NEIGHBOURS = {
  NORTH => [[-1, -1], [0, -1], [1, -1]],
  SOUTH => [[-1, 1], [0, 1], [1, 1]],
  WEST => [[-1, -1], [-1, 0], [-1, 1]],
  EAST => [[1, -1], [1, 0], [1, 1]],
}

board = {}
data.each_with_index do |row, y|
  row.chars.each_with_index do |point, x|
    board[[x, y]] = true if point == "#"
  end
end

def get_bounds(board)
  xs = []
  ys = []
  board.each { |k, _|
    x, y = k
    xs << x
    ys << y
  }
  return [xs.min, ys.min, xs.max, ys.max]
end

def needs_to_move?(elf, board)
  ex, ey = elf
  (-1..1).each { |y|
    (-1..1).each { |x|
      next if x == 0 && y == 0
      return true if board[[ex + x, ey + y]]
    }
  }
  return false
end

def get_move_to(p, direction, board)
  px, py = p

  #puts "Get move to #{direction}"
  dx, dy = NEIGHBOURS[direction][1]

  [px + dx, py + dy]
end

def no_neighbours_to?(point, direction, board)
  px, py = point
  NEIGHBOURS[direction].each do |n|
    dx, dy = n
    return false if board[[px + dx, py + dy]]
  end

  return true
end

def get_proposed_move(elf, dir, board)
  return nil if !needs_to_move?(elf, board)
  (0..3).each do |dd|
    proposed_direction = (dir + dd) % DIRECTIONS.length
    #puts "#{elf} tries to move towards #{proposed_direction}"
    if no_neighbours_to? elf, proposed_direction, board
      move = get_move_to(elf, proposed_direction, board)
      #puts "\t\t and will move to #{move}"
      return move
    end
    #puts "\t\tbut found neighbours"
  end
  return nil
end

def tick(board, direction_index)
  proposals = Hash.new([])

  board.each do |k, _|
    move = get_proposed_move k, direction_index, board
    next if move == nil
    proposals[move] += [k]
  end
  return nil if proposals.length == 0
  proposals.each do |k, v|
    if v.length == 1
      board.delete(v[0])
      board[k] = true
    end
  end

  board
end

def run_til_no_moves(board)
  tick = 1
  direction_index = 0
  while true
    new_board = tick(board, direction_index)
    if new_board == nil
      return [board, tick]
    end
    board = new_board
    direction_index = (direction_index + 1) % DIRECTIONS.length

    tick += 1
  end
  return [board, tick]
end

def count_empty_ground(board) (board)
  minx, miny, maxx, maxy = get_bounds(board)
  (maxx + 1 - minx) * (maxy + 1 - miny) - board.length end

def run_for(x, board)
  direction_index = 0
  x.times do
    new_board = tick(board, direction_index)
    return board if new_board == nil
    board = new_board
    direction_index = (direction_index + 1) % DIRECTIONS.length
  end
  return board
end

part_a_board = board.clone
part_a_board = run_for 10, part_a_board
pp count_empty_ground part_a_board

part_b_board = board.clone
result = run_til_no_moves part_b_board
pp result[1]
