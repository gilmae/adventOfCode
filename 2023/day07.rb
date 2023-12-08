data = File.readlines("inputs/day07.input").map(&:chomp).map { |l| l.split(" ") }
card_weights = "AKQJT98765432".reverse

def categorise_hand(hand)
  return 6 if hand[0][1] == 5
  return 5 if hand[0][1] == 4

  return 4 if hand[0][1] == 3 && hand[1][1] == 2
  return 3 if hand[0][1] == 3

  return 2 if hand[0][1] == 2 && hand[1][1] == 2
  return 1 if hand[0][1] == 2
  return 0
end

winnings = 0
ranked_hands = data.map { |hand, bid|
  hand_type = Hash.new(0)

  hand.chars.each { |c| hand_type[c] += 1 }

  [hand_type.map { |k, v| [k, v] }.sort { |a, b|
    [b[1], card_weights.index(b[0])] <=> [a[1], card_weights.index(a[0])]
  }, bid, hand]
}.sort { |a, b|
  ([categorise_hand(b[0])] + b[2].chars.map { |c| card_weights.index(c) }) <=> ([categorise_hand(a[0])] + a[2].chars.map { |c| card_weights.index(c) })
}.reverse.each_with_index { |hand, index|
  winnings += hand[1].to_i * (index + 1)
}

pp winnings

# part B
# J is now Joker, which is lowest ranked for count backs, but is wild for categorising
card_weights = "AKQT98765432J".reverse
winnings_part2 = 0
ranked_hands = data.map { |hand, bid|
  hand_type = Hash.new(0)

  hand.chars.each { |c| hand_type[c] += 1 }

  num_jokers = hand_type["J"]
  hand_type.delete "J"
  if hand_type.empty?
    hand_type = { "J" => 0 }
  end

  v = [hand_type.map { |k, v| [k, v] }.sort { |a, b|
    [b[1], card_weights.index(b[0])] <=> [a[1], card_weights.index(a[0])]
  }, bid, hand]
  v[0][0][1] += num_jokers

  v
}.sort { |a, b|
  ([categorise_hand(b[0])] + b[2].chars.map { |c| card_weights.index(c) }) <=> ([categorise_hand(a[0])] + a[2].chars.map { |c| card_weights.index(c) })
}.reverse.each_with_index { |hand, index|
  winnings_part2 += hand[1].to_i * (index + 1)
}
pp winnings_part2
