computers = {}
matches = []
File.readlines("inputs/day23.input").map{|l|l.chomp.split("-")}.each {|a,b|
  computers[a] = (computers[a] || []) + [b]
  computers[b] = (computers[b] || []) + [a]
}

computers.clone.filter{|k,_| k[0] == "t"}.each{|k,v|
  v.each{|v1| 
    computers[v1].each{|v2|
      matches << [k,v1,v2] if computers[v2].include? k
    }
  }
}
pp matches.map{|m|m.sort}.uniq.length

def get_cliques graph
  # reddit suggests Bron–Kerbosch algorithm with pivoting
  #  https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm#With_pivoting"
    
  cliques = []
  work = [[[], graph.keys, []]]

  loop {
    job = work.pop
    break if job.nil?
    r, p, x = job

    if (p.nil? || p.empty?) && (x.nil? || x.empty?)
      cliques << r
      next
    end

    u = nil
    (p | x).each{|maybe_u|
      u = maybe_u if u.nil? || graph[maybe_u].length > graph[u].length
    }

    (p - graph[u]).each {|v|
      work+=[[r | [v], p & graph[v], x & graph[v]]]
      x << v
    }

  }
  return cliques
end

pp get_cliques(computers).sort{|a,b| a.length <=> b.length}.last.sort.join(",")