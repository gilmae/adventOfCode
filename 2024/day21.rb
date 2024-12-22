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

def x_first_then_y start, dest, dx, dy, x_distance, y_distance, num_d_pads, layer
  memoise_get_shortest(D_PAD, num_d_pads, layer-1, "A", dx) + 
  x_distance + 
  memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, dy) + 
  y_distance + 
  memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, "A")
end

def y_first_then_x start, dest, dx, dy, x_distance, y_distance, num_d_pads, layer
  memoise_get_shortest(D_PAD, num_d_pads, layer-1, "A", dy) + 
  y_distance + 
  memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, dx) + 
  x_distance + 
  memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, "A")
end

def get_shortest pad, num_d_pads, layer, start, dest
  return 1 if layer == 0
  start = index_of(pad, start) if start.is_a? String 
  dest = index_of(pad, dest) if dest.is_a? String 
  
  dy = ['v',nil,'^'][(start[1]<=>dest[1])+1]
  dx = ['>',nil,'<'][(start[0]<=>dest[0])+1]
  x_distance = ((dest[0] - start[0]).abs - 1)
  y_distance = ((dest[1] - start[1]).abs - 1)

  return 1 if dx.nil? && dy.nil?
  
  return memoise_get_shortest(D_PAD, num_d_pads, layer-1, "A", dx) + x_distance + memoise_get_shortest(D_PAD, num_d_pads, layer-1, dx, "A") if dy.nil?
  return memoise_get_shortest(D_PAD, num_d_pads, layer-1, "A", dy) + y_distance + memoise_get_shortest(D_PAD, num_d_pads, layer-1, dy, "A") if dx.nil?
  
  if layer < num_d_pads
    return x_first_then_y(start,dest,dx,dy,x_distance, y_distance, num_d_pads, layer) if start[0] == 0
    return y_first_then_x(start,dest,dx,dy,x_distance, y_distance, num_d_pads, layer) if dest[0] == 0
    return [
      x_first_then_y(start,dest,dx,dy,x_distance, y_distance, num_d_pads, layer),
      y_first_then_x(start,dest,dx,dy,x_distance, y_distance, num_d_pads, layer)
    ].min
  else
    return x_first_then_y(start,dest,dx,dy,x_distance, y_distance, num_d_pads, layer) if start[0] == 0 && dest[1] == 3
    return y_first_then_x(start,dest,dx,dy,x_distance, y_distance, num_d_pads, layer) if dest[0] == 0 && start[1] == 3
    return [
      x_first_then_y(start,dest,dx,dy,x_distance, y_distance, num_d_pads, layer),
      y_first_then_x(start,dest,dx,dy,x_distance, y_distance, num_d_pads, layer)
    ].min
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
