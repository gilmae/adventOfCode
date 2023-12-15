data = File.readlines("inputs/day14.input").map(&:chomp)

board = {}
height = data.length
width = data[0].length

def print(board, height, width)
  height.times { |row|
    line = []
    width.times { |col|
      rock = board[[row, col]]
      line << "." if rock.nil?
      line << rock if !rock.nil?
    }
    pp line.join("")
  }
end

def rotate(board, height, width)
  new_board = {}
  board.each { |k, v|
    dr = k[1]
    dc = height - 1 - k[0]
    new_board[[dr, dc]] = v
  }

  new_board
end

def tilt(board, height, width)
  new_board = {}

  height.times { |row|
    width.times { |col|
      rock = board[[row, col]]
      if rock == "#"
        new_board[[row, col]] = rock
        next
      end
      cur = [row, col]
      while cur[0] > 0 && new_board[[cur[0] - 1, cur[1]]].nil?
        cur[0] -= 1
      end
      new_board[cur] = rock
    }
  }
  new_board
end

def get_load(board, height, width)
  total = 0
  height.times { |row|
    width.times { |col|
      rock = board[[row, col]]
      total += height - row if rock == "O"
    }
  }
  total
end

data.each_with_index { |line, row|
  line.chars.each_with_index { |ch, col|
    board[[row, col]] = ch if ch != "."
  }
}

parta_board = tilt board, height, width
pp get_load parta_board, height, width

t = 0
SEEN = {}
while t < 1e9
  4.times { |_|
    board = tilt(board, height, width)
    board = rotate(board, height, width)
  }

  if SEEN.has_key? board
    cycle_length = t - SEEN[board]
    amt = ((1e9 - t) / cycle_length).to_i
    t += amt * cycle_length
  end
  SEEN[board] = t
  t += 1
end

pp get_load board, height, width
