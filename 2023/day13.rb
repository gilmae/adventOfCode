data = File.readlines("inputs/day13.input")

blocks = []
block = []
data.each { |line|
  if line == "\n"
    blocks << block
    block = []
  else
    block << line.chomp
  end
}
blocks << block

def get_cols(lines)
  cols = []
  lines[0].length.times { |num|
    col = []
    lines.each { |line|
      col << line[num]
    }
    cols << col.join("")
  }
  cols
end

def check_is_reflection(sideA, sideB, partB)
  lengths = [sideA.length, sideB.length]
  aedis = sideA.reverse
  unflawed = partB
  lengths.min.times { |idx|
    if aedis[idx] != sideB[idx]
      # if we have already used our one flaw, not a reflection
      return false if !unflawed

      # if it is off by one, this is our one flaw
      if off_by_one(aedis[idx], sideB[idx])
        unflawed = false
      else
        # Otherwise if it is off by more than one, not a reflection
        return false
      end
    end
  }
  return !unflawed
end

def off_by_one(a, b)
  diffs = 0
  a.length.times { |idx|
    diffs += 1 if a[idx] != b[idx]
    return false if diffs > 1
  }
  true
end

def find_mirrors(lines, partB)
  lines[0..lines.length - 2].each_with_index { |line, idx|
    if check_is_reflection(lines[0..idx], lines[idx + 1..], partB)
      return idx
    end
  }
  return nil
end

pp blocks.map { |block|
  row = find_mirrors block, false
  if row.nil?
    cols = get_cols block
    (find_mirrors(cols, false) + 1)
  else
    100 * (row + 1)
  end
}.sum

pp blocks.map { |block|
  row = find_mirrors block, true
  if row.nil?
    cols = get_cols block
    (find_mirrors(cols, true) + 1)
  else
    100 * (row + 1)
  end
}.sum
