
stones = File.readlines("inputs/day11.input").map(&:chomp)[0].split(" ").map{|i| [i.to_i,1]}.to_h

SEEN = {}

def next_stones(stone)
  return [1] if stone == 0
  stone_str = stone.to_s
  if stone_str.length%2==0
   
    return stone_str.chars.each_slice(stone_str.length/2).map{|p| p.join("").to_i}

  end
  return [stone*2024]
end

def blink stones
  next_stones = {}
  stones.each {|k, v|
    dstones = SEEN[k]
    dstones = next_stones(k) if dstones.nil?
    SEEN[k] = dstones
    dstones.each {|s|
      next_stones[s] = (next_stones[s]||0) + v
    }
  } 
  next_stones
end

75.times {|i|
  stones = blink stones
  (pp stones.values.sum) if i == 24
}
pp stones.values.sum