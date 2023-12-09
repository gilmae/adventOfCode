data = File.readlines("inputs/day08.input").map(&:chomp)

moves = data[0].chars

lefts = {}
rights = {}
starts = [] # For part B

next_moves = { "L" => {}, "R" => {} }
data[2..].each { |line|
  start, left, right = line.scan(/[0-9A-Z]{3}/)
  next_moves["L"][start] = left
  next_moves["R"][start] = right
  starts << start if start[2] == "A"
}

# Part A
pos = "AAA"
move = 0
while pos != "ZZZ"
  #if next_moves[moves[move % moves.length]][pos] == "L"
  pos = next_moves[moves[move % moves.length]][pos]
  #
  move += 1
end

pp move

#PART B
pos = starts
t = 0

# For each start, find how long till it gets to the first Z
# Then do LCM on all the times
times = []
starts.each { |p|
  t = 0
  while true
    break if p[2] == "Z"
    p = next_moves[moves[t % moves.length]][p]
    t += 1
  end
  times << t
}

def gcd(a, b)
  return a if b == 0
  gcd(b, a % b)
end

# A method that returns the least common multiple (LCM) of two numbers using the GCD method above
def lcm(a, b)
  (a * b) / gcd(a, b)
end

def lcm_of_array(array)
  lcm = array.reduce(1) { |lcm, n| lcm(lcm, n) }
  return lcm
end

pp lcm_of_array(times)
