# BFS from all nodes to all other nodes, count the number of times an edge is visited
data = File.readlines("inputs/day25.input").map(&:chomp)

graph = {}

data.each { |line|
  from, to = line.split(": ")
  to = to.split(" ")
  graph[from] ||= []
  graph[from] += to

  to.each { |t|
    graph[t] ||= []
    graph[t] << from
  }
}

def count_edges(graph)
  encountered = {}
  graph.each { |k, _|
    walk_graph k, graph, encountered
  }
  encountered
end

def walk_graph(start, graph, encountered)
  seen = {}
  queue = [start]

  while !queue.empty?
    cur = queue.shift

    graph[cur].each { |neigh|
      next if seen.has_key? neigh
      queue << neigh
      seen[neigh] = true
      key = [cur, neigh].sort
      encountered[key] = (encountered[key] || 0) + 1
    }
  end
end

def count_reachable_nodes(from, graph)
  seen = {}
  queue = [from]

  while !queue.empty?
    cur = queue.shift
    next if seen.has_key? cur
    seen[cur] = true

    graph[cur].each { |neigh|
      queue << neigh
    }
  end

  seen.keys
end

def remove_edge(edge, graph)
  graph[edge[0]].delete(edge[1])
  graph[edge[1]].delete(edge[0])
  graph
end

node = nil

3.times { |_|
  edges = count_edges graph
  max = -1
  max_edge = nil
  edges.each { |edge, traversals|
    if traversals > max
      max = traversals
      max_edge = edge
    end
  }

  graph = remove_edge max_edge, graph

  node ||= max_edge[0]
}

set1 = count_reachable_nodes node, graph
set2 = graph.keys - set1

pp set1.length * set2.length
