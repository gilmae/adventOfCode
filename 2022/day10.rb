data = File.readlines("inputs/day10.input").map(&:chomp).map { |l| l.split(" ") }

x_values = {}

def find_x_value_at(values, cycle)
  possibility = nil
  values.each { |k, v|
    break if k > cycle
    possibility = v
  }
  possibility
end

x = 1
cycle = 1

# cpu
data.each { |l|
  cmd = l[0]
  x_values[cycle] = x
  case cmd
  when "noop"
    cycle += 1
  when "addx"
    cycle += 2
    x += l[1].to_i
  end
}

x_values[cycle] = x

pp [20, 60, 100, 140, 180, 220].map { |cycle|
  find_x_value_at(x_values, cycle) * cycle
}.sum

screen = []
(1..240).each { |c|
  pixel = (c % 40) - 1
  sprite_middle = find_x_value_at(x_values, c)
  pixel_value = "."
  pixel_value = "#" if pixel >= sprite_middle - 1 && pixel <= sprite_middle + 1
  screen[c - 1] = pixel_value
}

screen.each_slice(40).to_a.each { |line|
  pp line.join("")
}
