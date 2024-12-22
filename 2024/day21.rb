data = File.readlines("inputs/day21.input").map(&:chomp)

NUM_PAD = [
  ['7', '8', '9'],
  ['4', '5', '6'],
  ['1', '2', '3'],
  [nil, '0', 'A']
]

D_PAD = [
  [nil, '^', 'A'],
  ['<', 'v', '>']
]

NUM_D_PADS = 26

def index_of keys, key
  keys.each_with_index {|row, y|
    x = row.index key
    return [x,y] unless x.nil?
  }
end

MEMO = {}
def memoise_get_shortest pad, num_d_pads, layer, start, dest
  m = MEMO[[layer,start,dest]]
  return m unless m.nil?
  m = get_shortest pad, num_d_pads, layer, start, dest
  MEMO[[layer,start,dest]] = m
  m
end

def get_shortest pad, num_d_pads, layer, start, dest
  return 1 if layer == 0
  start = index_of(pad, start) if start.is_a? String 
  dest = index_of(pad, dest) if dest.is_a? String 
  
  if layer < num_d_pads
    dy = nil
    dx = nil
    dy = 'v' if start[1] < dest[1]
    dy = '^' if start[1] > dest[1]
    dx = ">" if start[0] < dest[0]
    dx = "<" if start[0] > dest[0]

    return memoise_get_shortest(D_PAD, num_d_pads, layer-1, "A", "A") if dx.nil? && dy.nil?
    if dy.nil?
      return memoise_get_shortest(D_PAD, num_d_pads, layer-1, "A", dx) + ((dest[0] - start[0]).abs - 1) * memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, dx) + memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, "A")
    elsif dx.nil?
      return memoise_get_shortest(D_PAD, num_d_pads, layer-1, "A", dy) + ((dest[1] - start[1]).abs - 1) * memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, dy) + memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, "A")
    else
      
        return memoise_get_shortest(D_PAD, num_d_pads, layer-1, "A", dx) + 
                    ((dest[0] - start[0]).abs - 1) * memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, dx) + 
                    memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, dy) + 
                    ((dest[1] - start[1]).abs - 1) * memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, dy) + 
                    memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, "A") if start[0] == 0
      
        return memoise_get_shortest(D_PAD, num_d_pads, layer-1, "A", dy) + 
                    ((dest[1] - start[1]).abs - 1) * memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, dy) + 
                    memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, dx) + 
                    ((dest[0] - start[0]).abs - 1) * memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, dx) + 
                    memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, "A") if dest[0] == 0
      
      
        return [
          memoise_get_shortest(D_PAD, num_d_pads, layer-1, "A", dx) + ((dest[0] - start[0]).abs - 1) * memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, dx) + memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, dy) + ((dest[1] - start[1]).abs - 1) * memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, dy) + memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, "A"),
          memoise_get_shortest(D_PAD, num_d_pads, layer-1, "A", dy) + ((dest[1] - start[1]).abs - 1) * memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, dy) + memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, dx) + ((dest[0] - start[0]).abs - 1) * memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, dx) + memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, "A")
    ].min
    end
  else
    dy = nil
    dx = nil
    dy = 'v' if start[1] < dest[1]
    dy = '^' if start[1] > dest[1]
    dx = ">" if start[0] < dest[0]
    dx = "<" if start[0] > dest[0]

    return memoise_get_shortest(D_PAD, num_d_pads, layer-1, "A", "A") if dx.nil? && dy.nil? # already in place
    if dy.nil?
      return memoise_get_shortest(D_PAD, num_d_pads, layer-1, "A", dx) + ((dest[0] - start[0]).abs - 1) * memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, dx) + memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, "A")
    elsif dx.nil?
      return memoise_get_shortest(D_PAD, num_d_pads, layer-1, "A", dy) + ((dest[1] - start[1]).abs - 1) * memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, dy) + memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, "A")
    else
        return memoise_get_shortest(D_PAD, num_d_pads, layer-1, "A", dx) + 
                    ((dest[0] - start[0]).abs - 1) * memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, dx) + 
                    memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, dy) + 
                    ((dest[1] - start[1]).abs - 1) * memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, dy) + 
                    memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, "A") if start[0] == 0 && dest[1] == 3
      
        return memoise_get_shortest(D_PAD, num_d_pads, layer-1, "A", dy) + 
                    ((dest[1] - start[1]).abs - 1) * memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, dy) + 
                    memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, dx) + 
                    ((dest[0] - start[0]).abs - 1) * memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, dx) + 
                    memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, "A") if dest[0] == 0 && start[1] == 3
      
      
        return [
          memoise_get_shortest(D_PAD, num_d_pads, layer-1, "A", dx) + ((dest[0] - start[0]).abs - 1) * memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, dx) + memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, dy) + ((dest[1] - start[1]).abs - 1) * memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, dy) + memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, "A"),
          memoise_get_shortest(D_PAD, num_d_pads, layer-1, "A", dy) + ((dest[1] - start[1]).abs - 1) * memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, dy) + memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, dx) + ((dest[0] - start[0]).abs - 1) * memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, dx) + memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, "A")
    ].min
    end
  end
end
pp data.map{|line|
  (("A"+line[0..2]).chars.zip line.chars).map {|key, key2|
    memoise_get_shortest(NUM_PAD, 3, 3, key, key2)
  }.sum * line.to_i
}.sum
pp data.map{|line|
  (("A"+line[0..2]).chars.zip line.chars).map {|key, key2|
    memoise_get_shortest(NUM_PAD, 26, 26, key, key2)
  }.sum * line.to_i
}.sum
