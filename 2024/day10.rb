data = File.readlines("inputs/day10.input").map(&:chomp)

trailheads = {}
distinct_trailheads = {}
work = []
map = {}
data.each_with_index { |line, y|
  line.chars.each_with_index { |ch, x|
    map[[x,y]] = ch=="."?-99:ch.to_i
    if ch == "0"
      trailheads[[x,y]] = []
      distinct_trailheads[[x,y]] = []
      work << [[x,y], [x,y], {}]
    end
  }
}

H = data.length
W = data[0].length

loop do
  job = work.shift
  #pp job
  break if job.nil?

  th, cur, seen = job
  height = map[cur]

  if map[cur] == 9
    trailheads[th] << cur
    distinct_trailheads[th] << seen
    next
  end

  [[1,0], [0,1], [-1,0], [0,-1]].each {|d|
    x,y = cur
    dx,dy = d
    next_pos = [x+dx, y+dy]
    next if seen.has_key? next_pos
    next if map[next_pos] != height+1
    if (0..W-1).include?(next_pos[0])  && (0..H-1).include?(next_pos[1])
      work << [th, next_pos, {next_pos=>true}.merge(seen)]
    end
  }
end

pp trailheads.map{|k,v| v.uniq.length}.sum
pp distinct_trailheads.map{|k,v| v.uniq.length}.sum