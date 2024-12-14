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

def count_corners board, point, region
  fences = board[point][1]
  num_corners = 0
  # internal corners
  [[0,1],[1,2],[2,3],[3,0]].each {|side1,side2|
    num_corners+=1 if fences.include?(side1) && fences.include?(side2)
  }

  #external corners
  x,y = point
  npoint = [x,y-1]
  spoint = [x,y+1]
  epoint = [x+1,y]
  wpoint = [x-1,y]
  
  #NE
  num_corners+=1 if has_side?(board, region, npoint, 1) && has_side?(board, region, epoint, 0)
  #NW
  num_corners+=1 if has_side?(board, region, npoint, 3) && has_side?(board, region, wpoint, 0)

  #SE
  num_corners+=1 if has_side?(board, region, spoint, 1) && has_side?(board, region, epoint, 2)
  #SW
  num_corners+=1 if has_side?(board, region, spoint, 3) && has_side?(board, region, wpoint, 2)

  num_corners
end

def has_side? board, region, point, fence
  return false if board[point].nil?
  return false if !region.include? point
  return board[point][1].include? fence
end

board.each {|k,v|
  next if SEEN.has_key? k
  ch,_ = v
  region = fill k, board
  regions << region.uniq
}

pp regions.map {|r|
  r.length * r.map{|p|
    board[p][1].length
  }.sum
}.sum

pp regions.map {|r|
  r.length *  r.map{|p|
      count_corners board, p, r
  }.sum 
}.sum
