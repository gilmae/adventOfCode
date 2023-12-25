data = File.readlines("inputs/day23.input").map(&:chomp)

def get_neighbours(p, board, partA)
  candidates = []
  directions = [[0, 1], [0, -1], [1, 0], [-1, 0]]
  if partA
    case board[p]
    when "<"
      directions = [[-1, 0]]
    when ">"
      directions = [[1, 0]]
    when "^"
      directions = [[0, -1]]
    when "v"
      directions = [[0, 1]]
    end
  end

  directions.each { |dx, dy|
    px, py = p
    candidate = [px + dx, py + dy]

    next if candidate[0] < 0 || candidate[0] > WIDTH || candidate[1] < 0 || candidate[1] > HEIGHT
    next if board[candidate] == "#"
    candidates << candidate
  }
  candidates
end

board = {}
HEIGHT = data.length - 1
WIDTH = data[0].length - 1

data.each_with_index { |line, row|
  line.chars.each_with_index { |ch, col|
    board[[col, row]] = ch if ch != "."
  }
}
start = [1, 0]
goal = [data[0].length - 2, data.length - 1]
GOAL = goal

nodes = [start, goal]

HEIGHT.times { |y|
  WIDTH.times { |x|
    next if board.has_key? [x, y]
    nodes << [x, y] if get_neighbours([x, y], board, false).length >= 3
  }
}

def get_distances(nodes, board, partA)
  distances = {}
  nodes.each { |n|
    distances[n] = []
    queue = [[n, 0]]
    seen = {}
    while !queue.empty?
      cur, steps = queue.shift

      next unless seen[cur].nil?
      seen[cur] = true
      if nodes.include?(cur) && steps > 0
        distances[n] << [cur, steps]
        next
      end

      get_neighbours(cur, board, partA).each { |n|
        queue << [n, steps + 1]
      }
    end
  }
  distances
end

distancesA = get_distances nodes, board, true
distancesB = get_distances nodes, board, false

def get_neighbour_nodes(cur, distances)
  distances.filter { |k, _| k[0] == cur }.map { |k, v| [k[1], v] }
end

def solve(start, goal, distances)
  queue = [[start, 0, []]]
  seen = []
  max_distance = -1
  while !queue.empty?
    cur, steps, path = queue.shift
    pp path.length
    if cur == goal
      max_distance = steps if steps > max_distance
      next
    end
    next if path.include? cur

    distances[cur].each { |neighbour, dist|
      queue.unshift[neighbour, steps + dist, path + [cur]]
    }
  end
  max_distance
end

SEEN = {}
@answer = -1

def dfs(start, steps, distances)
  if SEEN[start]
    return 0
  end

  SEEN[start] = true
  if start == GOAL
    @answer = [steps, @answer].max
  end
  distances[start].each { |neighbour, dist|
    dfs(neighbour, steps + dist, distances)
  }
  SEEN[start] = false
end

dfs(start, 0, distancesA)
pp @answer
@answer = 0
dfs(start, 0, distancesB)
pp @answer
