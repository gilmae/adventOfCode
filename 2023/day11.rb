data = File.readlines("inputs/day11.input").map(&:chomp)

height = data.length
width = data[0].length

empty_rows = height.times.map { |d| d }
empty_cols = width.times.map { |d| d }

def manhattan(p1, p2)
  return (p1[0] - p2[0]).abs() + (p1[1] - p2[1]).abs()
end

galaxies = []

data.each_with_index { |line, row|
  line.chars.each_with_index { |c, col|
    if c == "#"
      galaxies << [row, col]
      empty_rows.delete(row)
      empty_cols.delete(col)
    end
  }
}

def get_growth_amount(p1, p2, empty_rows, empty_cols)
  rows = [p1[0], p2[0]].sort
  cols = [p1[1], p2[1]].sort

  empty_cols.filter { |c| c > cols[0] && c < cols[1] }.length + empty_rows.filter { |c| c > rows[0] && c < rows[1] }.length
end

#pp get_growth_amount [2, 0], [4, 6], empty_rows, empty_cols

#pp manhattan([5, 1], [9, 4]) + get_growth_amount([5, 1], [9, 4], empty_rows, empty_cols)
distances = []
galaxies.each_with_index { |g1, index|
  galaxies[index + 1..].each { |g2|
    distances << (manhattan(g1, g2) + get_growth_amount(g1, g2, empty_rows, empty_cols))
  }
}

pp distances.sum

distances = []
galaxies.each_with_index { |g1, index|
  galaxies[index + 1..].each { |g2|
    distances << (manhattan(g1, g2) + 999999 * get_growth_amount(g1, g2, empty_rows, empty_cols))
  }
}

pp distances.sum
