require "pqueue"
PIXELS = {'#'=>'#', true=>'O'}
def print_board(board)
  minx, miny, maxx, maxy = 0,0,W,H
  (miny..maxy).each { |y|
    puts (minx..maxx).map { |x|
      px = board[[x, y]]
      (px != nil)?px: " "
    }.join("")
  }
end

def manhattan src, dest
    (src[0] - dest[0]).abs() + (src[1] - dest[1]).abs()
end

data = File.readlines("inputs/day18.input").map { |line| line.chomp.scan(/-?\d+/).map &:to_i }

board = {}
data[0..1023].each {|x,y|
  board[[x,y]] = "#"
}

DIRECTIONS = [[0,-1], [1,0], [0,1], [-1,0]]
start = [0,0]
#dest = [6,6]
#H = 7
#W = 7
H = 71
W = 71
dest = [70,70]

def calculate_path(board, start, dest)
  work = PQueue.new([[start,0, []]]) { |a, b| a[1] > b[1] }
  min_steps = 1e9
  min_path = {}
  visited = {}

  loop {  
    job = work.shift
    break if job.nil?
  
    pos, cost, seen = job
    next if cost >= min_steps
    next if visited.has_key? pos
  
    if pos==dest
      if cost < min_steps
        min_steps = cost 
        min_path = seen
      end
    end
  
    next if board[pos] == "#"

    visited[pos] = cost

    DIRECTIONS.each {|d|
      dx,dy = d
      x,y = pos
      next_pos = [x+dx, y+dy]
      next unless (x+dx).between?(0,W-1) && (y+dy).between?(0,H-1)
      work << [next_pos,cost+1, seen+[pos]]  
    }
  }
  [min_steps, min_path]
end
min_steps, min_path = calculate_path(board, start, dest)
pp min_steps


data[1024..].each {|x,y|
  board[[x,y]] = "#"
  place_in_path = min_path.index [x,y]
  next if place_in_path.nil?
  ms, dpath = calculate_path board, min_path[place_in_path-1], dest
  if ms == 1e9
    pp [x,y]
    exit
  end
  min_path = min_path[0..place_in_path-1] + dpath
}