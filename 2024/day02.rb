data = File.readlines("inputs/day02.input")

def within_margin? x
  x.abs >=1 && x.abs <= 3
end

def is_safe v
  if v[0].abs < 1 || v[0].abs > 3
    return false
  end
  dir = 0 <=> v[0]
  v[1..].each_with_index {|x,idx|
    safe_check = ((0 <=> x) == dir) && within_margin?(x)
    if !safe_check
      return false
    end
  }
  return true
end

def get_vectors levels
  memo = []

  levels[1..].each_with_index {|x,idx|
    memo<< levels[idx] - x
  }
  memo
end

parta = 0
partb = 0
data.each{ |line|
  levels = line.split(" ").map(&:to_i)
  v = get_vectors(levels)
  safe = is_safe(v)
  
  if safe
    parta+=1
    partb+=1
    next
  end

  levels.each_with_index {|l, idx|
    test = levels.clone
    test.delete_at(idx)
    v = get_vectors(test)
    if is_safe(v)
      partb+=1
      break  
    end
  }

}

pp parta
pp partb
