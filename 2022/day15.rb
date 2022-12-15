input = ARGV[0] || "inputs/day15.input"
row = (ARGV[1] || 2000000).to_i
data = File.readlines(input).map(&:chomp)

board = {}
no_beacon = {}
rex = /.*=+(-?\d+).*=+(-?\d+).*=+(-?\d+).*=+(-?\d+)/
data.each_with_index { |line, idx|
  sensor_x, sensor_y, beacon_x, beacon_y = line.scan(rex)[0].map(&:to_i)

  distance = (beacon_x - sensor_x).abs + (beacon_y - sensor_y).abs

  if (row - sensor_y).abs < distance
    exclude_distance = distance - (row - sensor_y).abs
    ((-1 * exclude_distance)..exclude_distance).each { |x|
      no_beacon[x + sensor_x] = true if no_beacon[x + sensor_x] == nil
    }
    no_beacon[beacon_x] = false if beacon_y == row
  end
}

pp no_beacon.map { |_, v| v ? 1 : 0 }.sum
