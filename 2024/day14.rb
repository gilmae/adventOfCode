data = File.readlines("inputs/day14.input").map { |line| line.scan(/-?\d+/).map &:to_i }

H=103
W=101

part_a_board = []
data.each {|x,y,vx,vy|
  part_a_board << [((x+vx*100)%W), ((y+vy*100)%H),vx,vy]
}

def get_board_at board, t
  board.map {|x,y,vx,vy|  
    [((x+vx*t)%W), ((y+vy*t)%H), vx,vy]
  }
end

xsplit = W/2
ysplit = H/2
part_a_board = get_board_at data, 100

quadrants = [0,0,0,0]
part_a_board.map {|x,y,_,_|
  quadrants[0] += 1 if x < xsplit && y < ysplit
  quadrants[1] += 1 if x > xsplit && y < ysplit
  quadrants[2] += 1 if x > xsplit && y > ysplit
  quadrants[3] += 1 if x < xsplit && y > ysplit
}

safety = 1
quadrants.each{|q| safety*=q}
pp safety

(H*W).times {|t|
  b = get_board_at(data, t)
  xmax = b.map{|x,_,_,_,_| x}.tally.values.max
  ymax = b.map{|_,y,_,_,_| y}.tally.values.max
  
  if xmax>20 &&  ymax> 20
    puts t
    break
  end
}
