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

board[start] = 96
board[dest] = 123

def get_neighbours(board, point)
  x, y = point
  points = []
  [[0, -1], [0, 1], [-1, 0], [1, 0]].each { |d|
    dx, dy = d
    points << [x + dx, y + dy]
  }
  points
end

def can_travel(board, cur, dest)
  return board[dest] && (board[dest] - board[cur] <= 1)
end

def move_from(board, starting_points, dest)
  work = starting_points.map { |s| [s, 0] }

  visited = {}

  while !work.empty?
    # this was pop before, so I think it mean I was getting longest path.
    # I....I nearly cried when I worked it out
    current, steps = work.shift
    next if visited[current]

    return steps if current == dest

    visited[current] = true

    get_neighbours(board, current).each { |target|
      work.append [target, steps + 1] if can_travel board, current, target
    }
  end

  return nil
end

pp move_from(board, [start], dest)

starting_points = []
board.each { |k, v|
  starting_points << k if v == 97
}

pp move_from(board, starting_points, dest)
