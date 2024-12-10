data = File.readlines("inputs/day08.input").map(&:chomp)

index = {}
H = data.length
W = data[0].length

data.each_with_index { |line, y|
  line.chars.each_with_index { |ch, x|
    next if ch == "."
    index[ch] ||= []
    index[ch] << [x,y]
  }
}
antinodes = []
resonant_antinodes = []

def on_board? pos
  xs = (0..W-1)
  ys = (0..H-1)
  xs.include?(pos[0]) && ys.include?(pos[1])
end

def get_antinodes p1,p2, with_resonance
  return [] if p1.nil? || p2.nil?
  x1, y1 = p1
  x2, y2 = p2

  dx = x2-x1
  dy = y2-y1
  
  if with_resonance
    n = [p1,p2]
    idx = 1
    loop {
      dp1 = [x1-dx*idx,y1-dy*idx]
      dp2 = [x2+dx*idx,y2+dy*idx]
      n += [dp1, dp2]
      idx+=1
      break if !on_board?(dp1) && !on_board?(dp2)
    }
    n.delete_if{|a| !on_board? a}
  else
    [[x1-dx,y1-dy], [x2+dx,y2+dy]].delete_if{|a| !on_board? a}
  end 
end

index.each {|_,v|
  v.each_with_index {|node, idx|
    (idx+1..v.length-1).each {|node2|
      antinodes += get_antinodes node, v[node2], false
      resonant_antinodes += get_antinodes node, v[node2], true
    }
  }
}

puts antinodes.uniq.length
puts resonant_antinodes.uniq.length