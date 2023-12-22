data = File.readlines("inputs/day21.input").map(&:chomp)
HEIGHT = data.length - 1
WIDTH = data[0].length - 1
board = {}
start = []

GOAL = 64

data.each_with_index { |line, row|
  line.chars.each_with_index { |ch, col|
    board[[col, row]] = ch
    start = [col, row] if ch == "S"
  }
}

def get_neighbours(cur, board)
  neighbours = []
  [[1, 0], [0, 1], [-1, 0], [0, -1]].each { |dx, dy|
    poss = [cur[0] + dx, cur[1] + dy]
    if poss[0].between?(0, WIDTH) && poss[1].between?(0, HEIGHT) && board[poss] != "#"
      neighbours << poss
    end
  }
  neighbours
end

SEEN = {}
queue = [[start, 0]]
while !queue.empty?
  cur, steps = queue.shift
  next if SEEN.has_key? cur

  SEEN[cur] = steps
  get_neighbours(cur, board).each { |n|
    queue << [n, steps + 1]
  }
end

pp SEEN.filter { |k, v| v <= GOAL && v % 2 == 0 }.keys.uniq.length

# Part B is math. We've worked out shortest path to all spots in the garden.
# The map has a straight run from starting pos on both axis, so effectively
# the furtherest we can travel is 26501365 steps in all cardinal directions
# We can more or less fill in a diamond shape by entirely filling in plots
# as we expand out but eventually we run out of steps to fill in the corners
# of the furtherest plots out. See below for having 458 steps, enough to expand
# three additional plots

#    E
#   EOE
#  EOEOE
# EOEOEOE
#  EOEOE
#   EOE
#    E

# We know how many positions in the garden can be reached by odd numbers of steps
# We know when we move between plots that that what was odd in the first is even in
# the second. So we can work out how many plots 26501365 will take us (202300)
# We can also see a pattern in the "polarity" of plots
#   Plot Expansion       Num Odd     Num Even
#       0                   1           0
#       1                   1           4
#       2                   9           4
#       3                   9           16
#
# If expansions n is odd, then there are (n+1) ** 2 even plots and n**2 odd
# if n is even, there are (n+1)**2 odd plots and n**2 even
# 202300 is even, so there are 202301**2 even plots and 202300**2 odd plots
# The number of steps we take is odd, so position ploarity is odd. in odd plots and
# even in even plots. If it were an even number it would be the other way around
# So we know we have occupied some or all of all of those plots, so multiply the
# number of even plots by even positions in the plot and odd plots of odd positions
# Then remove the corners because it's a diamond, not a square.
# Work out all the corner spots in the original plot. You can reach the edge in 65 steps
# so the corners are the ones you cannot reach in 65 or fewer steps.
# If you have even expansion, there are (n+1) * 4 corners in odd plots that
# are unreachable, and n n*4 corners in even plots that have been reached.
# So get the odd positions in the *4* corners of the original plot and
# subtract n+1 of them. Get the even positions in the *4* corners of the original and
# subtract n of them
# Flip polarity for odd expansion.

expansion = (26501365 - (WIDTH / 2)) / (WIDTH + 1)

even_positions_in_corners = SEEN.filter { |k, v| v % 2 == 0 && v > 65 }.keys.uniq.length
odd_positions_in_corners = SEEN.filter { |k, v| v % 2 == 1 && v > 65 }.keys.uniq.length

even_positions = SEEN.filter { |k, v| v % 2 == 0 }.keys.uniq.length
odd_positions = SEEN.filter { |k, v| v % 2 == 1 }.keys.uniq.length

even_plots = expansion ** 2
odd_plots = (expansion + 1) ** 2

pp odd_plots * odd_positions + even_plots * even_positions - (expansion + 1) * odd_positions_in_corners + expansion * even_positions_in_corners - expansion
