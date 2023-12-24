require "z3"

def get_m(a1, a2)
  (a2[1] - a1[1]) / (a2[0] - a1[0])
end

def get_b(a, m)
  m * -a[0] + a[1]
end

def get_intersection(b1, b2, m1, m2)
  #   x = (b2 - b1)/(m1 - m2)
  # y = (m1 b2 - m2 b1)/(m1 - m2)
  x = (b2 - b1) / (m1 - m2)
  y = (m1 * b2 - m2 * b1) / (m1 - m2)

  [x, y]
end

def in_bounds(p, bounds)
  px, py = p
  min, max = bounds

  return px >= min && px <= max && py >= min && py <= max
end

def intersection_in_past?(p, i)
  start, _, delta, _, _ = p
  sx, _, _ = start
  dx, _, _ = delta
  ix, _, _ = i

  (ix - sx) / dx < 0
end

data = File.readlines("inputs/day24.input").map(&:chomp)
stones = data.each_with_index.map { |line, idx|
  x, y, z, dx, dy, dz = line.scan(/(-?\d+)/).map { |i| i[0].to_i }
  p1 = [x.to_f, y.to_f, z.to_f]
  p2 = [x + dx.to_f, y + dy.to_f, z + dz.to_f]
  m = get_m p1, p2
  b = get_b p2, m
  [p1, p2, [dx, dy, dz], m, b]
}
#pp stones
collisions = 0
stones.each_with_index { |s1, idx|
  stones[idx = 1..].each { |s2|
    x, y = get_intersection s1[4], s2[4], s1[3], s2[3]

    if x == Float::INFINITY # || y == Float::INFINITY
    elsif intersection_in_past?(s1, [x, y])
    elsif intersection_in_past?(s2, [x, y])
    elsif !in_bounds [x, y], [200000000000000, 400000000000000]
    else
      collisions += 1
    end
  }
}

pp collisions / 2 # Counts A colliding with B and B colliding with A, so halve

## Why is this not working?!?
# x = Z3.Real("x")
# y = Z3.Real("y")
# z = Z3.Real("z")

# dx = Z3.Real("dx")
# dy = Z3.Real("dy")
# dz = Z3.Real("dz")

# a = (0..stones.length - 1).map { |i| Z3.Int("a#{i}") }

# r = Z3::Solver.new
# (0..stones.length - 1).each { |i|
#   Z3.Add(x + a[i] * dx - stones[i][0][0] - a[i] * stones[i][1][0] = 0)
#   Z3.Add(y + a[i] * dy - stones[i][0][1] - a[i] * stones[i][1][1] = 0)
#   Z3.Add(z + a[i] * dz - stones[i][0][2] - a[i] * stones[i][1][2] = 0)
# }

# pp r.check
# M = r.model
