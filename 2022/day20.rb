input = ARGV[0] || "inputs/day20.input"
data = File.readlines(input).map(&:chomp).map(&:to_i)

def get_grove_number(data, iterations, multiplier)
  mixed = data.each_with_index.map { |n, i| [n * multiplier, i] }
  iterations.times {
    mixed.length.times { |counter|
      idx = mixed.find_index { |d| d[1] == counter }
      val = mixed.delete_at(idx)
      mixed.insert((idx + val[0]) % mixed.length, val)
    }
  }
  zero = mixed.find_index { |d| d[0] == 0 }
  return mixed[(zero + 1000) % data.length][0] + mixed[(zero + 2000) % data.length][0] + mixed[(zero + 3000) % data.length][0]
end

# part A

pp get_grove_number data, 1, 1
pp get_grove_number data, 10, 811589153
