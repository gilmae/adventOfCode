data = File.readlines("inputs/day01.input")

lists = data.map {|l| 
  /(\d+)\s\s\s(\d+)/.match(l)[1..2].map(&:to_i)
}.transpose.map(&:sort)

occurences = lists[1].each_with_object(Hash.new(0)) do |x, memo|
  memo[x] += 1 
end

pp lists[0].map.with_index {|n, idx|
  [(n- lists[1][idx]).abs, n * (occurences[n] || 0)]
}.transpose.map(&:sum)