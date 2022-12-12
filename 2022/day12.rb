input = ARGV[0] || "inputs/day12.input"
data = File.readlines(input).map(&:chomp)

board = {}

start = []
dest = []

STARTTILE = "S".bytes[0]
ENDTILE = "E".bytes[0]
data.each_with_index { |line, row|
  line.bytes.each_with_index { |c, col|
    start = [col, row] if c == STARTTILE
    dest = [col, row] if c == ENDTILE
    board[[col, row]] = c
  }
}

def get_neightbours(board, point)
  x, y = point
  points = []
  (-1..1).each { |col|
    (-1..1).each { |row|
      next if col == 0 && row == 0
      points << [x + col, y + row]
    }
  }
  points
end

def can_travel(board, cur, dest)
  curX, curY = cur
  destX, destY = dest

  return false if curX != destX && curY != destY

  return board[dest] && (board[cur] = STARTTILE || board[dest] - board[cur] == 1)
end

def AStarElevations(board, start, dest)
  work = { start => true }

  gScore = { start => 0 }
  fScore = { start => 1 }

  path = {}

  while work.length > 0
    current = nil
    currentScore = (2 ** (0.size * 8 - 2) - 1)

    work.each { |key, _|
      score = fScore[key]

      if score < currentScore
        current = key
        currentScore = score
      end
    }
    work.delete(current)

    if current == dest
      pp "found it"
      return path.length
    else
      get_neightbours(board, current).reject { |p| !can_travel(board, current, p) }.each { |n|
        tentative_score = gScore[current] + 1
        previous_score = gScore[n]

        if !previous_score || tentative_score < previous_score
          path[n] = current
          gScore[n] = tentative_score
          fScore[n] = tentative_score + 2
          work[n] = true if !work[n]
        end
      }
    end
  end
end

pp AStarElevations(board, start, dest)
