data = File.readlines("inputs/day03.input").map(&:chomp)

board = {}
not_symbols = [".", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"]

symbols = []
parts = {}
all_part_coords = {}

reading_part = false
part = ""
part_coords = nil

data.each_with_index { |line, row|
  line.chars.each_with_index { |ch, col|
    board[[row, col]] = ch

    if ("0".."9").include? ch
      # we're reading a part
      part_coords = part_coords || [row, col]
      part << ch
    elsif !part_coords.nil?
      #stop reading a part
      parts[part_coords] = part
      part.length.times { |dcol|
        all_part_coords[[part_coords[0], part_coords[1] + dcol]] = part
      }
      part = ""
      part_coords = nil
    end

    symbols << [row, col] if !not_symbols.include? ch
  }
}

def is_part?(part, part_coords, symbols)
  (-1..1).each { |drow|
    (-1..part.length).each { |dcol|
      pos = [part_coords[0] + drow, part_coords[1] + dcol]
      if symbols.include? pos
        return pos
      end
    }
  }
  return nil
end

def get_touching_parts(co, all_part_coords)
  parts = []
  (-1..1).each { |drow|
    (-1..1).each { |dcol|
      dco = [co[0] + drow, co[1] + dcol]
      if all_part_coords.include? dco
        parts << all_part_coords[dco]
      end
    }
  }
  return parts.uniq
end

sum_of_parts = 0
parts.each { |coords, p|
  sym = is_part?(p, coords, symbols)
  if !sym.nil?
    sum_of_parts += p.to_i
  end
}
pp sum_of_parts

sum_of_gears = 0
symbols.each { |sc|
  touched_parts = get_touching_parts sc, all_part_coords
  if touched_parts.length == 2
    sum_of_gears += touched_parts[0].to_i * touched_parts[1].to_i
  end
}

pp sum_of_gears
