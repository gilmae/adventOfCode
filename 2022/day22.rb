input = ARGV[0] || "inputs/day22.input"
data = File.readlines(input).map(&:chomp)

RIGHT = 0
DOWN = 1
LEFT = 2
UP = 3

MOVES = [[1, 0], [0, 1], [-1, 0], [0, -1]] # RIGHT, DOWN, LEFT, UP
BOARD = {}
COMMANDS = data.last.scan(/(\d+|[LR])/).flatten

# Read board
data[0..-3].each_with_index { |line, y|
  line.chars.each_with_index { |point, x|
    next if point == " "
    BOARD[[x, y]] = point == "."
  }
}

def row_bounds(row)
  cols = BOARD.find_all { |k, v| k[1] == row }.map { |k, _| k[0] }
  return [cols.min, cols.max]
end

def column_bounds(col)
  rows = BOARD.find_all { |k, v| k[0] == col }.map { |k, _| k[1] }
  return [rows.min, rows.max]
end

def get_bounds
  rows = BOARD.map { |k, _| k[0] }.uniq.sort
  cols = BOARD.map { |k, _| k[1] }.uniq.sort

  return [rows.min, cols.min, rows.max, cols.max]
end

def change_facing(cur, change)
  cur -= 1 if change == "L"
  cur += 1 if change == "R"

  cur %= 4
  cur
end

def wrap(n, facing)
  if facing == LEFT || facing == RIGHT
    bounds = row_bounds n[1]
    if facing == RIGHT
      n = [bounds[0], n[1]]
    else
      n = [bounds[1], n[1]]
    end
  else
    bounds = column_bounds n[0]
    if facing == UP
      n = [n[0], bounds[1]]
    else
      n = [n[0], bounds[0]]
    end
  end
  return n
end

def move(facing, start, length)
  delta = MOVES[facing]
  pos = start.clone

  length.times {
    # Get proposed next position
    n = [pos[0] + delta[0], pos[1] + delta[1]]

    # If position out of bounds, warp
    if BOARD[n] == nil
      n = wrap(n, facing)
    end

    # If position is blocked, return current
    return pos unless BOARD[n]
    # Else update current to repeat
    pos = n
  }
  pos
end

bounds = get_bounds
pos = [row_bounds(0)[0], 0]
facing = RIGHT
COMMANDS.each { |cmd|
  if cmd == "L" || cmd == "R"
    facing = change_facing facing, cmd
  else
    pos = move facing, pos, cmd.to_i
  end
}
# pp wrap [12, 5], RIGHT
pp (pos[0] + 1) * 4 + (1000 * (pos[1] + 1)) + facing
