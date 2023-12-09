data = File.readlines("inputs/day04.input").map(&:chomp)

total = 0

cards_possessed = {}
data.length.times { |g| cards_possessed[g + 1] = 1 }

data.each_with_index { |line, idx|
  game = idx + 1
  parts = line[line.index(":")..].split("|")
  winning = parts[0].split(" ")
  numbers = parts[1].split(" ")

  matches = numbers.length - (numbers - winning).length
  total += 2 ** (matches - 1) if matches > 0

  matches.times { |m|
    winnings = game + 1 + m
    continue if winnings > data.length
    cards_possessed[winnings] += cards_possessed[game]
  }
}

pp total

pp cards_possessed.values.inject(:+)
