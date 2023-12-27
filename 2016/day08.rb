data = File.readlines("inputs/day08.input").map(&:chomp)

# pixels will be accessed as pixels[x][y], i.e. x is columns, y is rows
pixels = {}
HEIGHT = 6
WIDTH = 50

def display(pixels)
  puts

  (0..HEIGHT - 1).each { |row|
    line = []
    (0..WIDTH - 1).each { |col|
      if pixels.has_key?([row, col]) && pixels[[row, col]] == true
        line << "#"
      else
        line << "."
      end
    }
    puts line.join("")
  }
  puts
end

def rect(pixels, x, y)
  x.times { |col|
    y.times { |row|
      pixels[[row, col]] = true
    }
  }
  pixels
end

def rotate_x(pixels, col, amount)
  height = pixels.map { |k, _| k[0] }.max + 1
  newpixels = {}
  pixels.each { |k, v|
    if k[1] == col
      new_coord = [k[0] + amount, k[1]]

      new_coord[0] = new_coord[0] % (HEIGHT)

      newpixels[new_coord] = v
    else
      newpixels[k] = v
    end
  }
  newpixels
end

def rotate_y(pixels, row, amount)
  newpixels = {}
  pixels.each { |k, v|
    if k[0] == row
      new_coord = [k[0], k[1] + amount]

      new_coord[1] = new_coord[1] % (WIDTH)

      newpixels[new_coord] = v
    else
      newpixels[k] = v
    end
  }
  newpixels
end

data.each { |line|
  m = /rect (\d+)+x(\d+)/.match(line)
  if !m.nil?
    pixels = rect pixels, m[1].to_i, m[2].to_i
  else
    m = /rotate (\w+) \w=(\d+) by (\d+)/.match(line)
    next if m.nil?
    case m[1]
    when "row"
      pixels = rotate_y pixels, m[2].to_i, m[3].to_i
    when "column"
      pixels = rotate_x pixels, m[2].to_i, m[3].to_i
    else
      puts "ERROR: #{line}"
    end
  end
}
display pixels

pp pixels.length
