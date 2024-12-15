data = File.readlines("inputs/day15.input").map(&:chomp)

def sum_gps(board)
  board.each.map{|k,v| 
  if ['[', 'O'].include? v 
    x,y = k
    x+y*100
  else
    0
  end
}.sum
end

def double_board board
  dboard = {}
  board.each{|k,v|
    x,y = k
    x = x*2
    cx = x+1
    if v == "#"
      dboard[[x,y]] = v
      dboard[[cx,y]] = v
    elsif v == "O"
      dboard[[x,y]] = "["
      dboard[[cx,y]] = "]"
    end
  }
  dboard
end

def get_moves board, pos, d
  companions = {'['=>1, ']'=>-1}
  x,y = pos
  dx,dy = d
  dpos = [x+dx,y+dy]
  return [[pos,dpos]] if board[dpos].nil?
  return nil if board[dpos] == '#'
  n = get_moves board, dpos, d
  return if n.nil? 
  return ([[pos,dpos]] + n) if board[dpos] == 'O'

  # we're into [] land now, comrades
  # if we're moving horizontally, it just works itself out, don't worry about it
  return [[pos,dpos]] + n if dy == 0 
  companion_is_at_x_delta = companions[board[dpos]]
  return nil if companion_is_at_x_delta.nil?
  cpos = [dpos[0]+companion_is_at_x_delta, dpos[1]]
  cn = get_moves board, cpos, d
  
  return nil if cn.nil? # other side of the box is blocked
  return [[pos,dpos]] + n + cn
end

def process_command board, cur, m
  moves = get_moves board, cur, DIRECTIONS[m]
  return [board,cur] if moves.nil?
  dboard = {}.merge board
  moves[1..].each{|s,f| 
    dboard.delete s
    dboard.delete f
  }
  moves[1..].each{|s,f|
    dboard[f]=board[s]
  }
  [dboard,moves[0][1]]
end

board = {}
cur = nil
y = 0

loop {
  line = data[y]
  break if line == ""
  line.chars.each_with_index { |ch, x|
    cur = [x,y] if ch == "@"
    board[[x,y]] = ch if ["#", "O"].include? ch
  }
  y+=1
}
W = data[0].length
H = y-1
wider_board = double_board board
wider_cur = [cur[0]*2, cur[1]]
DIRECTIONS = {'<'=>[-1,0], '^' => [0,-1], '>'=>[1,0], 'v'=>[0,1]}

data[y+1..].join("").chars.each {|m|
  board, cur = process_command board, cur, m
}
pp sum_gps board
data[y+1..].join("").chars.each {|m|
   wider_board, wider_cur = process_command wider_board, wider_cur, m
}
pp sum_gps wider_board