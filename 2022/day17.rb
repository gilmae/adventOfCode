input = ARGV[0] || "inputs/day17.input"
data = File.readlines(input).map(&:chomp)[0].chars

board = {}

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

def print_board(board)
  minx, miny, maxx, maxy = get_bounds(board)
  maxy += 3
  (0..maxy).each { |y|
    puts (0..6).map { |x|
      px = board[[x, maxy - y]]
      px != nil ? px : "."
    }.join("")
  }
  puts "-------"
end

LINE = 0
CROSS = 1
ELL = 2
EYE = 3
ROCK = 4

SHAPES = {
  ELL => [[0, 0], [1, 0], [2, 0], [2, 1], [2, 2]],
  LINE => [[0, 0], [1, 0], [2, 0], [3, 0]],
  CROSS => [[1, 0], [1, 1], [1, 2], [0, 1], [2, 1]],
  EYE => [[0, 0], [0, 1], [0, 2], [0, 3]],
  ROCK => [[0, 0], [0, 1], [1, 0], [1, 1]],
}

def get_height_at(target_cycle, cycle_now, cycle_before, height_before, height_now)
  time_left = target_cycle - cycle_now - 1
  time_between_patterns = cycle_now - cycle_before
  height_added = (height_now + 1) - height_before
  cycles_left, remaining = time_left.divmod(time_between_patterns)

  if remaining == 0
    return (height_now + 1) + height_added * cycles_left + 1
  end
  return nil
end

def will_collide(rock, x, y, board)
  shape = SHAPES[rock]

  shape.each { |s|
    dx, dy = s
    dx = dx + x
    dy = dy + y

    return true if dx < 0 || dx > 6
    return true if dy < 0

    return true if board.include?([dx, dy])
  }
  return false
end

def add_to_board(rock, x, y, board, sigil)
  shape = SHAPES[rock]
  shape.each { |s|
    dx, dy = s
    board[[x + dx, y + dy]] = sigil
  }
end

patterns = {}
rock = 0
jet_index = 0

1e12.to_i.times { |cycle|
  _, _, _, height = get_bounds(board)
  height = -1 if height == nil
  height += 1 # Because the board index starts at 0, le sigh
  pp height if cycle == 2022
  key = [rock, jet_index]
  if patterns.include?(key)
    pattern_cycle, pattern_height = patterns[key]

    projected_height = get_height_at 1e12, cycle, pattern_cycle, pattern_height, height
    if projected_height
      puts projected_height
      exit
    end
  else
    patterns[key] = [cycle, height + 1]
  end
  x = 2
  y = height + 3

  while true
    wind_direction = data[jet_index]
    jet_index = (jet_index + 1) % data.length

    dx = wind_direction == ">" ? 1 : -1

    x += dx if !will_collide(rock, x + dx, y, board)
    if !will_collide(rock, x, y - 1, board)
      y -= 1
    else
      break
    end
  end
  add_to_board rock, x, y, board, "#"
  rock = (rock + 1) % SHAPES.length
}
