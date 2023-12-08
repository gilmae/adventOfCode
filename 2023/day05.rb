data = File.readlines("inputs/day05.input").map(&:chomp)

seeds = data[0].scan(/(\d+)+/).map(&:first).map(&:to_i)

class Function
  attr_accessor :tuples

  def initialize(tuples)
    @tuples = tuples
  end

  def solve_for_one(x)
    tuples.each { |d, s, l|
      return x + d - s if s <= x && x <= s + l
    }
    return x
  end

  def apply_range(r)
    all_new_ranges = []
    @tuples.each { |d, s, l|
      src_end = s + l
      new_range = []
      while !r.empty?
        st, ed = r.shift
        before = [st, [ed, s].min]
        inter = [[st, s].max, [src_end, ed].min]
        after = [[src_end, st].max, ed]

        new_range << before if before[1] > before[0]
        new_range << after if after[1] > after[0]

        if inter[1] > inter[0]
          all_new_ranges << [inter[0] - s + d, inter[1] - s + d]
        end
      end
      r = new_range
    }
    return all_new_ranges + r
  end
end

def get_map(data)
  map = []

  idx = 0
  while data[idx].scan(/(\d+)+/).empty?
    idx += 1
  end
  legend = data[idx].scan(/(\d+)+/).map(&:first).map(&:to_i)
  while !legend.empty? && !data[idx].nil?
    idx += 1

    map << legend
    legend = data[idx].scan(/(\d+)+/).map(&:first).map(&:to_i) unless data[idx].nil?
  end
  return map, data[idx..]
end

data = data[2..]
maps = []

while !data.nil? && !data.empty?
  map, data = get_map data
  maps << Function.new(map)
end

locations = seeds.map { |seed|
  maps.each { |map|
    seed = map.solve_for_one seed
  }
  seed
}

pp locations.min

seed_ranges = []
(0..seeds.length - 1).step(2).each { |index| seed_ranges << [seeds[index], seeds[index + 1]] }

# #pp (locations & seeds).min
P2 = []
seed_ranges.each { |st, l|
  rs = [[st, st + l]]
  maps.each { |m|
    rs = m.apply_range rs
  }

  P2 << rs.map { |r| r[0] }.min
}

pp P2.min
# locations = seed_ranges.map { |seed_range|
#   maps.each { |map|
#     seed_range = map.apply_range [seed_range[0], seed_range[0] + seed_range[1]]
#   }
#   seed_range.map { |sr| sr[0] }.min
# }

# pp locations.min
