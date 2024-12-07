data = File.readlines("inputs/day06.input").map(&:chomp)

start = [0,0]
FACINGS = [[0,-1], [1,0], [0,1], [-1,0]]

BOARD = {}

data.each_with_index { |line, row|
  line.chars.each_with_index { |ch, col|
    next if ch == "."
    if ch == "^"
      start = [col,row] 
    else
      BOARD[[col, row]] = ch 
    end
  }
}

HEIGHT = data.length
WIDTH = data[0].length

def on_board? pos
  xs = (0..WIDTH-1)
  ys = (0..HEIGHT-1)
  xs.include?(pos[0]) && ys.include?(pos[1])
end

def next_move board, pos, face
  x,y = pos
  (0..3).each {|t|
    dx,dy = FACINGS[(face+t)%4]  
    test = [x+dx,y+dy]
    return [test,(face+t)%4] if !board.has_key? test
  }
  return [nil,nil]
end

def travel board, start, facing
  work = [[start, facing]]
  visited = {}
  loop {
    job = work.pop
    return [visited, true] if visited.include?(job)
    visited[job] = true
    pos, facing = job

    x, y = pos
    dx,dy = FACINGS[facing]
    new_pos = [x+dx, y+dy]

    return [visited, false] if !on_board? new_pos

    if board.has_key? new_pos
      new_facing = (facing+1)%4
      work << [pos,new_facing]
    else
      work << [new_pos, facing]
    end
  }
  return [visited, false]
end

visited, _ = travel BOARD, start, 0
pp visited.map{|k,_| k[0]}.uniq.length

placements = {}
visited.each{|k,_| 
  pos, facing = k
  x,y = pos
  dx,dy = FACINGS[facing]
  extra_rock = [x+dx, y+dy]
  dboard = {}.merge(BOARD) 
  dboard = dboard.merge({extra_rock=>"O"}) if on_board?(extra_rock)
  
  _,is_loop = travel dboard, start, 0
  placements[extra_rock] = true if is_loop
  
}
pp placements.length
