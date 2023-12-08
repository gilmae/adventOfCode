data = File.readlines("inputs/day06.input").map(&:chomp)

def solve_quadratic(a, b, c)
  x1 = (-1 * b + (b ** 2 - (4 * a * c)) ** 0.5) / 2 * a
  x2 = (-1 * b - (b ** 2 - (4 * a * c)) ** 0.5) / 2 * a

  # if the numbers are already whole numbers, that means they _equal_ the record,
  # so they have to be adjusted to make sure they would exceed it
  x1 += 1 if x1.ceil == x1
  x2 -= 1 if x2.floor == x2
  [x1.ceil, x2.floor]
end

distances = data[0].scan(/(\d+)+/).map(&:first).map(&:to_i)
times = data[1].scan(/(\d+)+/).map(&:first).map(&:to_i)

pp (0..distances.length - 1).map { |idx|
  roots = solve_quadratic(-1, distances[idx], -1 * times[idx])

  roots[1] - roots[0] + 1
}.inject(&:*)

real_distance = distances.map(&:to_s).join("").to_i
real_time = times.map(&:to_s).join("").to_i

real_ways = solve_quadratic -1, real_distance, -real_time
pp real_ways[1] - real_ways[0] + 1
