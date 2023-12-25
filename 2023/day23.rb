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
    queue = [[n, 0]]
    seen = {}
    while !queue.empty?
      cur, steps = queue.shift

      next unless seen[cur].nil?
      seen[cur] = true
      if nodes.include?(cur) && steps > 0
        distances[[n, cur]] = steps
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
#distancesB.delete_if { |k, _| k[0] == start || k[1] == start }

nodes.each { |n|
  queue = [[n, 0]]
  seen = {}
  while !queue.empty?
    cur, steps = queue.shift

    next unless seen[cur].nil?
    seen[cur] = true
    if nodes.include?(cur) && steps > 0
      distancesA[[n, cur]] = steps
      next
    end

    get_neighbours(cur, board, true).each { |n|
      queue << [n, steps + 1]
    }
  end
}

def get_neighbour_nodes(cur, distances)
  distances.filter { |k, _| k[0] == cur }.map { |k, v| [k[1], v] }
end

#pp distances

def solve(start, goal, distances, partA)
  queue = [[start, 0, []]]
  seen = []
  max_distance = -1
  while !queue.empty?
    cur, steps, path = queue.shift

    next if path.include? cur

    if cur == goal
      max_distance = steps if steps > max_distance
      next
    end

    get_neighbour_nodes(cur, distances).each { |nn|
      queue.unshift([nn[0], steps + nn[1], path + [cur]])
    }
  end
  max_distance
end

pp solve(start, goal, distancesA, true)
#pp solve(start, goal, distancesB, false)

pp nodes.map { |n| [n, distancesB.filter { |k, _| k[1] == n }.length] }
