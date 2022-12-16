input = ARGV[0] || "inputs/day15.input"
data = File.readlines(input).map(&:chomp)

row = input == "inputs/day15.input" ? 2000000 : 10
cave_upper_bound = input == "inputs/day15.input" ? 4000000 : 20

def get_manhattan_distance(x1, y1, x2, y2)
  (x2 - x1).abs + (y2 - y1).abs
end

def is_in_range_of_a_sensor?(x, y, sensors)
  sensors.each { |s, b|
    sx, sy = s
    _, _, d = b

    return true if d >= get_manhattan_distance(x, y, sx, sy)
  }
  return false
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

  xlabels = (minx..maxx).map { |x| x.to_s.ljust(3) }
  3.times { |i|
    xlabels.map { |x| x.chars[i] }.join("")
    puts("      #{xlabels.map { |x| x.chars[i] }.join("")}")
  }

  (miny..maxy).each { |y|
    puts "#{y.to_s.ljust(6)}" + (minx..maxx).map { |x|
      px = board[[x, y]]
      px != nil ? px : "."
    }.join("")
  }
end

def find_locations_with_no_beacon_in_line(start_x, distance, excluded_spots)
  ((-1 * distance)..distance).each { |x|
    excluded_spots[x + start_x] = true if excluded_spots[x + start_x] == nil
  }
end

def draw_diamond(start_x, start_y, r, sensor, board)
  (0..r).each { |y|
    x = r - y.abs
    board[[start_x - x, start_y + y]] = "#"
    board[[start_x + x, start_y + y]] = "#"
    board[[start_x - x, start_y - y]] = "#"
    board[[start_x + x, start_y - y]] = "#"
  }
end

def map_locations_with_no_beacon(start_x, start_y, distance, board)
  (0..distance).each { |y|
    x_axis = distance - y.abs
    x_axis.times { |x|
      board[[start_x - x, start_y + y]] = "#"
      board[[start_x + x, start_y + y]] = "#"
      board[[start_x - x, start_y - y]] = "#"
      board[[start_x + x, start_y - y]] = "#"
    }
  }
end

sensors = {}
board = {}
no_beacon = {}
rex = /.*=+(-?\d+).*=+(-?\d+).*=+(-?\d+).*=+(-?\d+)/
data.each_with_index { |line, idx|
  sensor_x, sensor_y, beacon_x, beacon_y = line.scan(rex)[0].map(&:to_i)
  distance = get_manhattan_distance sensor_x, sensor_y, beacon_x, beacon_y
  sensors[[sensor_x, sensor_y]] = [beacon_x, beacon_y, distance]

  if (row - sensor_y).abs < distance
    exclude_distance = distance - (row - sensor_y).abs

    find_locations_with_no_beacon_in_line(sensor_x, exclude_distance, no_beacon)
    no_beacon.delete(beacon_x) if beacon_y == row
  end
}

pp no_beacon.map { |_, v| v ? 1 : 0 }.sum

def scan_for_uncovered_point(sensors, cave_upper_bound)
  rotations = [[-1, 1], [1, 1], [-1, -1], [1, -1]]
  sensors.each { |k, v|
    sensor_x, sensor_y = k
    beacon_x, beacon_y, r = v
    (0..r + 1).each { |y|
      x = 1 + r - y.abs
      rotations.each { |r|
        rx, ry = r
        px = sensor_x + x * rx
        py = sensor_y + y * ry

        next if px < 0 || px > cave_upper_bound || py < 0 || py > cave_upper_bound
        return px, py if !is_in_range_of_a_sensor? px, py, sensors
      }
    }
  }
  return nil
end

p = scan_for_uncovered_point sensors, cave_upper_bound
if p != nil
  px, py = p
  pp px * 4000000 + py
end

sensors.each { |k, v|
  sensor_x, sensor_y = k
  beacon_x, beacon_y, r = v
  distance = get_manhattan_distance sensor_x, sensor_y, beacon_x, beacon_y
  # Debugging code, prints the board so I can see if I am on the right track
  # For Cthulhu's sake, don't use it on the real input!!!
  #   map_locations_with_no_beacon(sensor_x, sensor_y, distance + 1, board)
  #   board[[sensor_x, sensor_y]] = "s"
  #   board[[beacon_x, beacon_y]] = "b"
}
#print_board board
