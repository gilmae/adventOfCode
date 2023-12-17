data = File.readlines("inputs/day16.input").map(&:chomp)

board = {}
height = data.length
width = data[0].length
data.each_with_index { |line, row|
  line.chars.each_with_index { |ch, col|
    board[[col, row]] = ch if ch != "."
  }
}

def print_energised(energised, height, width)
  height.times { |row|
    line = []
    width.times { |col|
      if energised.include? [col, row]
        line << "#"
      else
        line << "."
      end
    }
    puts "#{line.join("")}\n"
  }
end

def print_mirrors(board, energised, height, width)
  height.times { |row|
    line = []
    width.times { |col|
      p = board[[col, row]]
      if !p.nil?
        line << p
      elsif energised.include? [col, row]
        line << "#"
      else
        line << "."
      end
    }
    puts "#{line.join("")}\n"
  }
end

def apply_vertical_spitter(beam)
  return [beam] if beam[1][0] == 0

  # if moving horizontally, split
  beam[1] = [0, 1]
  new_beam = [beam[0], [0, -1]]
  return [beam, new_beam]
end

def apply_horizontal_spitter(beam)
  return [beam] if beam[1][1] == 0

  # if moving vertically, split
  beam[1] = [1, 0]
  new_beam = [beam[0], [-1, 0]]
  return [beam, new_beam]
end

def apply_forward_mirror(beam)
  # mirror is a \
  pos, delta = beam

  if delta[0] == 0 # north or south
    return [[pos, [1 * delta[1], 0]]]
  else
    return [[pos, [0, 1 * delta[0]]]]
  end
end

def apply_backward_mirror(beam)
  # mirror is a /
  pos, delta = beam

  if delta[0] == 0 # North or South
    return [[pos, [-1 * delta[1], 0]]]
  else
    return [[pos, [0, -1 * delta[0]]]]
  end
end

def apply_delta(beam)
  pos, delta = beam
  new_pos = [pos[0], pos[1]]
  new_pos[0] += delta[0]
  new_pos[1] += delta[1]
  [new_pos, delta]
end

def out_of_bounds(beam, height, width)
  return false if beam[0][0].between?(0, height - 1) && beam[0][1].between?(0, width - 1)
  #puts "out of bounds at #{beam[0]}"
  return true
end

def test(beam, board, height, width)
  beams = [beam]
  energised = []
  seen = { beams[0] => true }

  while !beams.empty?
    # Move beam
    # Check if beam is on a splitter or mirror and modify beam and/or add new beams
    new_beams = []
    beams.each_with_index { |beam, idx|
      new_beam = apply_delta beam

      next if out_of_bounds(new_beam, height, width)
      next unless seen[new_beam].nil?

      seen[new_beam] = true
      energised << new_beam[0]

      case board[new_beam[0]]
      when "|"
        new_beams += apply_vertical_spitter(new_beam)
      when "-"
        new_beams += apply_horizontal_spitter(new_beam)
      when "\\"
        new_beams += apply_forward_mirror(new_beam)
      when "/"
        new_beams += apply_backward_mirror(new_beam)
      when nil
        new_beams += [new_beam]
      end
    }

    beams = new_beams
  end
  energised.uniq.length
end

pp test [[-1, 0], [1, 0]], board, height, width
beams = []
width.times { |col|
  beams << [[col, -1], [0, 1]]
  beams << [[col, height], [0, -1]]
}
height.times { |row|
  beams << [[-1, row], [1, 0]]
  beams << [[width, row], [-1, 0]]
}

max = -1
pp beams.map { |b|
  test b, board, height, width
}.max
