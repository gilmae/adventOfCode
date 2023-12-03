data = File.readlines("inputs/day02.input").map(&:chomp)

MAX_RED = 12
MAX_GREEN = 13
MAX_BLUE = 14

def is_possible(balls)
  return (balls["blue"] || 0) <= MAX_BLUE && (balls["green"] || 0) <= MAX_GREEN && (balls["red"] || 0) <= MAX_RED
end

def check_fewest(fewest, balls, colour)
  return if balls[colour].nil?
  fewest[colour] = balls[colour] if balls[colour] > fewest[colour]
end

score = 0
sum_of_powers = 0

data.each_with_index { |line, game|
  fewest = { "green" => -2 ^ 31, "blue" => -2 ^ 31, "red" => -2 ^ 31 }
  rounds = line[line.index(":") + 2..].split("; ")
  game_score = game + 1
  rounds.each { |round|
    retrieved = round.split(", ")
    balls = {}
    retrieved.each { |r|
      md = /(\d+)\s(\w+)/.match(r)
      balls[md[2]] = md[1].to_i
    }
    check_fewest fewest, balls, "blue"
    check_fewest fewest, balls, "green"
    check_fewest fewest, balls, "red"

    if !is_possible balls
      game_score = 0
    end
  }
  power = fewest["blue"] * fewest["green"] * fewest["red"]
  sum_of_powers += power
  score += game_score
}

pp score
pp sum_of_powers
