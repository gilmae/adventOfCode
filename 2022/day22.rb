input = ARGV[0] || "inputs/day22.input"
data = File.readlines(input).map(&:chomp)

SIDELENGTH = 50

RIGHT = 0
DOWN = 1
LEFT = 2
UP = 3

MOVES = [[1, 0], [0, 1], [-1, 0], [0, -1]] # RIGHT, DOWN, LEFT, UP
BOARD = {}
COMMANDS = data.last.scan(/(\d+|[LR])/).flatten

FACES = {
  1 => [(50..99), (0..49)],
  2 => [(50..99), (50..99)],
  3 => [(0..49), (100..149)],
  4 => [(100..149), (0..49)],
  5 => [(0..49), (150..199)],
  6 => [(50..99), (100..149)],
}

PORTALS = {
  #[face, direction] => [newface, [position transforms], new facing]
  [1, UP] => [5, ["swap"], RIGHT], # Tested
  [1, LEFT] => [3, ["flipy"], RIGHT], # Tested
  [2, LEFT] => [3, ["swap"], DOWN], # Tested
  [2, RIGHT] => [4, ["swap"], UP], # Tested
  [3, UP] => [2, ["swap"], RIGHT], # Tested
  [3, LEFT] => [1, ["flipy"], RIGHT], # Tested
  [4, DOWN] => [2, ["swap"], LEFT], # Tested
  [4, RIGHT] => [6, ["flipy"], LEFT], # Tested
  [4, UP] => [5, ["flipy"], UP],
  [5, LEFT] => [1, ["swap"], DOWN], # Tested
  [5, RIGHT] => [6, ["swap"], UP], # Tested
  [5, DOWN] => [4, ["flipy"], DOWN],
  [6, DOWN] => [5, ["swap"], LEFT], # Tested
  [6, RIGHT] => [4, ["flipy"], LEFT], # Tested
}

def get_face(point)
  px, py = point
  FACES.each do |face, ranges|
    return face if ranges[0].include?(px) && ranges[1].include?(py)
  end
end

def get_relative_in_face(point, face = nil)
  face = get_face point if face == nil
  px, py = point
  x_range, y_range = FACES[face]
  return [px - (x_range.min), py - (y_range.min)]
end

def get_absolute_in_face(point, face, facing)
  x_range, y_range = FACES[face]
  px, py = point
  dx = x_range.min
  dy = y_range.min

  return [px + dx, py + dy]
end

def swap(p)
  px, py = p
  [py, px]
end

def flipy(p)
  px, py = p
  [px, SIDELENGTH - 1 - py]
end

def flipx(p)
  px, py = p
  [SIDELENGTH - 1 - px, py]
end

def warp(pos, heading)
  face = get_face pos
  transition = PORTALS[[face, heading]]
  relative_position = get_relative_in_face pos, face
  transition[1].each do |t|
    relative_position = send(t, relative_position)
  end

  pos = get_absolute_in_face relative_position, transition[0], transition[2]
  return [pos, transition[2]]
end

###########################################################################
# Tests

def test(start, heading, expected_pos, expected_facing)
  puts "Test #{start} heading #{heading}"
  actual_pos, actual_facing = warp start, heading
  puts "Expected #{expected_pos}, got #{actual_pos}"
  puts "Expected #{expected_facing}, got #{actual_facing}"
  puts
end

# # Face 1, UP
# test [55, 0], UP, [0, 155], RIGHT

# # Face 5, LEFT
# test [0, 155], LEFT, [55, 0], DOWN

# # Face 3, UP
# test [5, 100], UP, [50, 55], RIGHT

# # Face 2, LEFT
# test [50, 55], LEFT, [5, 100], DOWN

# # Face 3, LEFT
# test [0, 105], LEFT, [50, 44], RIGHT

# # Face 1, LEFT
# test [50, 44], LEFT, [0, 105], RIGHT

# # Face 5, RIGHT
# test [49, 155], RIGHT, [55, 149], UP

# # Face 6, DOWN
# test [55, 149], DOWN, [49, 155], LEFT

# # Face 2, RIGHT
# test [99, 55], RIGHT, [105, 49], UP

# # FACE 4, DOWN
# test [105, 49], DOWN, [99, 55], LEFT

# # Face 4, RIGHT
# test [149, 5], RIGHT, [99, 144], LEFT

# # Face 6, RIGHT
# test [99, 144], RIGHT, [149, 5], LEFT

# # Face 5, DOWN
# test [5, 199], DOWN, [105, 0], DOWN

# # Face 4, UP
# test [105, 0], UP, [5, 199], UP

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

def move_with_wrap(facing, start, length)
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

def move_with_warp(facing, start, length)
  pos = start.clone

  length.times {
    delta = MOVES[facing]
    # Get proposed next position
    n = [pos[0] + delta[0], pos[1] + delta[1]]
    nfacing = facing
    # If position out of bounds, warp
    if BOARD[n] == nil
      warped_pos, warped_facing = warp(pos, facing)
      n = [warped_pos[0], warped_pos[1]]
      nfacing = warped_facing
    end

    # If position is blocked, return current
    return [pos, facing] unless BOARD[n]
    # Else update current to repeat
    pos = n
    facing = nfacing
  }
  [pos, facing]
end

pos = [50, 0]
facing = RIGHT
COMMANDS.each { |cmd|
  if cmd == "L" || cmd == "R"
    facing = change_facing facing, cmd
  else
    pos = move_with_wrap facing, pos, cmd.to_i
  end
}
pp (pos[0] + 1) * 4 + (1000 * (pos[1] + 1)) + facing

pos = [50, 0]
facing = RIGHT
COMMANDS.each { |cmd|
  if cmd == "L" || cmd == "R"
    facing = change_facing facing, cmd
  else
    pos, facing = move_with_warp facing, pos, cmd.to_i
  end
}

pp (pos[0] + 1) * 4 + 1000 * (pos[1] + 1) + facing
