data = File.readlines("inputs/day20.input").map(&:chomp)

board = {}
start = nil
data.each_with_index { |line, y|
  line.chars.each_with_index { |c, x|
    board[[x, y]] = c if c != "#"
    start = [x,y] if c == "S"
  }
}
H = data.length
W = data[0].length
steps_at_point = {start=>0}

work=[start]
loop {
  pos = work.shift
  break if pos.nil?

  [[-1,0],[1,0],[0,1],[0,-1]].map {|d|
    x,y=pos
    dx,dy = d
    next_pos = [x+dx,y+dy]
    
    next if !board.has_key? next_pos
    next if steps_at_point.has_key? next_pos
    steps_at_point[next_pos] = steps_at_point[pos]+1
    work << next_pos
  }
}

part_a_cheats = 0
part_b_cheats = 0

steps_at_point.keys.each {|p1|
  (-20..20).each {|x|
    (-20..20).each {|y|
      next if x==0 && y==0
      m = (x).abs + (y).abs
      next if m > 20
      p2 = [p1[0]+x,p1[1]+y]
      next unless steps_at_point.has_key? p2

      cheat = (steps_at_point[p2]-steps_at_point[p1])
      next if (cheat-m)<100

      part_a_cheats+=1 if m == 2
      part_b_cheats+=1 if m <= 20
    }
  }
}
pp part_a_cheats
pp part_b_cheats
