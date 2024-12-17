data = File.readlines("inputs/day16.input").map(&:chomp)
board = {}
dest = nil
start = nil
data.each_with_index { |line, y|
  line.chars.each_with_index { |c, x|
    board[[x, y]] = c if c == "#"
    start = [x,y] if c == "S"
    dest = [x,y] if c == "E"
  }
}
H = data.length
W = data[0].length
directions = [[0,-1],[1,0],[0,1], [-1,0]]
work = [[start, [1,0], 0, {}]]
min_cost = 1e9
path = {}
visited = {}

PIXELS = {'#'=>'#', [0,-1]=>'^',[1,0]=>'>',[0,1]=>'v', [-1,0]=>'<'}
def print_board(board)
  minx, miny, maxx, maxy = 0,0,W,H
  (miny..maxy).each { |y|
    puts (minx..maxx).map { |x|
      px = board[[x, y]]
      px != nil ? PIXELS[px] : " "
    }.join("")
  }
end

loop {
  job = work.shift
  break if job.nil?
  
  pos, facing, cost, seen = job
  previous = visited[[pos,facing]]
  next if cost >= min_cost || (!previous.nil? && cost >= previous)

  if pos==dest
    min_cost = cost if cost < min_cost
    path = seen
    next
  end
  
  next if board[pos] == "#"

  visited[[pos,facing]] = cost

  next if cost > min_cost
  seen[pos] = facing

  directions.each {|d|
    cost_modifier = 1
    dx,dy = d
    x,y = pos
    next_pos = [x+dx, y+dy]
    next if seen.has_key? next_pos
    cost_modifier+= 1000 if d != facing
    work.push [[x+dx,y+dy],d,cost+cost_modifier, {pos=>facing}.merge(seen)]  
  }
}

pp min_cost
#print_board path.merge(board)