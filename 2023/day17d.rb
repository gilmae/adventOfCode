require "pqueue"
data = File.readlines("inputs/day17.input").map(&:chomp)
board = {}
data.each_with_index { |line, row|
  line.chars.each_with_index { |c, col|
    board[[col, row]] = c.to_i
  }
}
HEIGHT = data.length
WIDTH = data[0].length

DIRECTIONS = [[-1, 0], [0, 1], [1, 0], [0, -1]]

def solve(board, partB)
  seen = {}

  # Work items are
  # 0: cur position
  # 1: direction heading
  # 2: steps in current heading
  # 3: total heat loss incurred

  work = PQueue.new([[[0, 0], -1, -1, 0]]) { |a, b| a[3] > b[3] }

  while !work.empty?
    w = work.shift
    cur, dir, steps_in_dir, heat = w
    next if seen.has_key?([cur, dir, steps_in_dir])

    seen[[cur, dir, steps_in_dir]] = heat

    DIRECTIONS.each_with_index { |d, dir_idx|
      new_pos = [cur[0] + d[1], cur[1] + d[0]]

      new_steps_in_dir = steps_in_dir

      if dir_idx == dir
        new_steps_in_dir = new_steps_in_dir + 1
      else
        new_steps_in_dir = 1
      end

      is_valid = new_steps_in_dir <= 3
      if partB
        is_valid = new_steps_in_dir <= 10 && (dir == dir_idx || steps_in_dir >= 4 || dir == -1)
        is_valid = is_valid & (new_pos != [WIDTH - 1, HEIGHT - 1] || new_steps_in_dir >= 4)
      end

      not_reverse = ((dir_idx + 2) % 4 != dir)

      if is_valid && not_reverse && new_pos[0].between?(0, WIDTH - 1) && new_pos[1].between?(0, HEIGHT - 1)
        cost = board[new_pos]

        work << [new_pos, dir_idx, new_steps_in_dir, heat + cost]
      end
    }
  end

  ans = 1e9
  seen.keys.each { |cur, dir, steps_in_dir|
    next if cur != [WIDTH - 1, HEIGHT - 1]
    heat = seen[[cur, dir, steps_in_dir]]
    ans = [ans, heat].min
  }
  ans
end

pp solve(board, false)
pp solve(board, true)
