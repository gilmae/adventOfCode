def get_movement(gap)
  if gap == 0
    return 0
  elsif gap < 0
    return -1
  else
    return 1
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

def print_board(board)
  minx, miny, maxx, maxy = get_bounds(board)
  (miny..maxy).each { |y|
    puts "#{y}|" + (minx..maxx).map { |x|
      px = board[[x, y]]
      px != nil ? px : "."
    }.join("")
  }
end

input = ARGV[0] || "inputs/day14.input"
data = File.readlines(input).map(&:chomp)

def get_board(data)
  board = {}
  lowest_ys = {}
  data.each { |line|
    points = line.split(" -> ").map { |p| p.split(",").map(&:to_i) }
    idx = 0
    last_start = points.length - 1
    while idx < last_start
      x, y = points[idx]
      tx, ty = points[idx + 1]

      while x != tx || y != ty
        board[[x, y]] = "#"
        lowest_ys[x] = y if lowest_ys[x] == nil || lowest_ys[x] < y
        x += get_movement(tx - x)
        y += get_movement(ty - y)
      end
      board[[x, y]] = "#"
      idx += 1
    end
  }

  [board, lowest_ys]
end

def drop_sand(x, y, board, lowest_ys, floor)
  movements = [[0, 1], [-1, 1], [1, 1]]

  blocked = false
  while !blocked
    moved = false
    movements.each { |m|
      moved = false
      dx, dy = m
      if board[[x + dx, y + dy]] == nil
        x += dx
        y += dy
        moved = true
        break
      end
    }
    if lowest_ys != nil
      lowest_y = lowest_ys[x]

      # Into the white!
      return nil if lowest_y == nil || lowest_y < (y)
    end

    blocked = true if floor != nil && y + 1 == floor

    blocked = true if !moved
  end
  board[[x, y]] = "o"
  [x, y]
end

idx = 0

board, lowest_ys = get_board data

min_x, _, max_x, max_y = get_bounds(board)
(min_x..max_x).each { |x|
  lowest_ys[x] = 0 if lowest_ys[x] == nil
}

p = drop_sand 500, 0, board, lowest_ys, nil
while p != nil
  idx += 1
  p = drop_sand 500, 0, board, lowest_ys, nil
end
pp idx

board, lowest_ys = get_board data
min_x, _, max_x, max_y = get_bounds(board)

floor = max_y + 2

idx = 1
p = drop_sand 500, 0, board, nil, floor
while p != [500, 0]
  idx += 1
  p = drop_sand 500, 0, board, nil, floor
end
pp idx
