data = File.readlines("inputs/day04.input").map(&:chomp)
board = {}

data.each_with_index { |line, row|
  line.chars.each_with_index { |ch, col|
    board[[col, row]] = ch
  }
}
directions = [[1,0],[-1,0], [0,1], [0,-1], [1,1], [1,-1], [-1,1], [-1,-1]]

HEIGHT = data.length
WIDTH = data[0].length

sum = 0
new_board = {}
board.each {|k,v| 
  next if v != "X"
  sum+=directions.map {|d|

    ex = k[0]+d[0]*3
    ey = k[1]+d[1]*3
    if ex < 0 || ex >=WIDTH || ey<0 || ey >=HEIGHT
      0
    else 
      slice = []
      
      4.times {|i|
        dx = k[0]+d[0]*i
        dy = k[1]+d[1]*i
        slice << board[[dx,dy]]
      }
      slice.join("")=="XMAS" ? 1 : 0
    end
  }.sum
}
pp sum

sum = 0
board.each {|k,v|
  next if v != "A"
  next if k[0] < 1 || k[1] < 1 || k[0] > HEIGHT-2 || k[1] > HEIGHT-2
  s1p1 = [k[0]-1 , k[1]-1 ]
  s1p2 = [k[0]+1,k[1]+1]
  s2p1 = [k[0]-1,k[1]+1]
  s2p2 = [k[0]+1,k[1]-1]

  slice1 = [board[s1p1], v, board[s1p2]]
  slice2 = [board[s2p1], v, board[s2p2]]
  word1 = slice1.join("")
  word2 = slice2.join("")
  if (word1=="MAS" || word1=="SAM") && ( word2 =="MAS" || word2=="SAM")
    sum+=1
  end
}
pp sum