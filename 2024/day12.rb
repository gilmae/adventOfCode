board = {}

index = {}

data = File.readlines("inputs/day12.input").map(&:chomp)

H = data.length
W = data[0].length

#N=0, E=1, S=2, W=3
data.each_with_index { |line, y|
  line.chars.each_with_index { |ch, x|
    fences = []  
    fences << 3 if x==0
    fences << 1 if x == W-1
    fences << 1 if x < W && !line[x+1].nil? && line[x+1] != ch
    fences << 3 if x >0 && line[x-1] != ch
    
    fences << 0 if y==0
    fences << 2 if y== H-1
    fences << 2 if y < H-1 && !data[y+1].nil? && data[y+1][x] != ch
    fences << 0 if y > 0 && !data[y-1].nil? && data[y-1][x] != ch
    board[[x,y]] = [ch,fences]
  }
}
SEEN = {}

regions = []

def fill start, board
  region = []
  filling_for = board[start][0]
  work = [start]

  loop do
    cur = work.shift
    break if cur.nil?

    region << cur
    SEEN[cur] = true

    [[1,0],[-1,0],[0,1],[0,-1]].each {|d|
      x, y = cur
      dx,dy = d
      xx = x+dx
      yy = y+dy
      next if board[[xx,yy]].nil? || board[[xx,yy]][0] != filling_for
      next if SEEN.has_key? [xx,yy]
      work.unshift [xx,yy]
    }
  end
  region
end

board.each {|k,v|
  next if SEEN.has_key? k
  ch,_ = v

  
  region = fill k, board
  
  regions << region.uniq
}

pp regions.map {|r|
  area = r.length
  perimeter = r.map{|p|
    board[p][1].length
  }.sum
  area*perimeter
}.sum
