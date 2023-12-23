data = File.readlines("inputs/day22.input").map(&:chomp)

# [startx, starty,startz, finishx, finishy, finishz]
bricks = data.each_with_index.map { |line, idx|
  xlo, ylo, zlo, xhi, yhi, zhi = line.scan(/(\d+)/).map { |i| i[0].to_i }
  [[xlo, xhi], [ylo, yhi], [zlo, zhi]]
}

bricks.sort { |a, b| a[2][0] <=> b[2][0] }

def supported_by(brick, bricks)
  max_height = bricks.map { |b| b[2][1] }.max
  # assumes that any bricks below brick are supported

  bottom = brick[2][0] # the lowest height of the brick
  return ["ground", "ground"] if bottom == 1
  potential_support = bricks.filter { |b| (bottom - 1) == b[2][1] }

  potential_support.filter { |b|
    xintersect = (brick[0][0]..brick[0][1]).to_a & (b[0][0]..b[0][1]).to_a
    yintersect = (brick[1][0]..brick[1][1]).to_a & (b[1][0]..b[1][1]).to_a

    xintersect.length > 0 && yintersect.length > 0
  }
end

def has_support?(brick, bricks)
  supported_by(brick, bricks).length > 0
end

def settle(bricks)
  max_height = bricks.map { |b| b[2][1] }.max
  height = 1
  new_board = []

  moved_count = 0
  while height <= max_height
    bricks_to_move = bricks.filter { |b| height == b[2][0] }
    bricks_to_move.each { |brick|
      potential_dz = -1
      new_brick = brick
      while new_brick[2][0] > 1 && !has_support?(new_brick, new_board)
        x, y, z = new_brick
        z = [z[0] - 1, z[1] - 1]

        new_brick = [x, y, z]
      end
      moved_count += 1 if brick[2][0] != new_brick[2][0]

      #puts "Move #{brick} to #{new_brick}" if brick[2][0] != new_brick[2][0]

      new_board << new_brick
    }

    height += 1
  end
  [new_board, moved_count]
end

new_board, moved_count = settle bricks
unmovable = []

new_board.each { |brick|
  supports = supported_by brick, new_board
  if supports.length == 1
    #puts "Brick #{brick} supported only by #{supports}"
    (unmovable << supports[0])
  end
}

# PART A
pp (new_board - unmovable).length

puts "PART B"
pp new_board.map { |brick|
  test_bricks = new_board.filter { |b| b != brick }
  _, moved_count = settle test_bricks
  moved_count
}.sum
