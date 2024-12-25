data = File.readlines("inputs/day24.input").map(&:chomp)

wires = {}
gates = []

idx = 0
loop {
  break if data[idx] == ""
  wire, value = data[idx].split(": ")
  wires[wire.strip] = value == "1"
  idx+=1
}
idx+=1
#puts "strict digraph {"
data[idx..].each {|line|
  gate = line.scan(/(.+)\s([A-Z]+)\s(.+)\s->\s(.+)/)[0]
 # puts "#{gate[0]} -> #{gate[3]} [label=#{gate[1]}]"
  #puts "#{gate[2]} -> #{gate[3]} [label=#{gate[1]}]"
  gates << gate
}
#puts "}"

def simulate wires, gates
  state = wires.clone
  loop {
    new_state = state.clone
    gates.each{|in1, op, in2, out|
      case op
      when "AND"
        new_state[out] = state[in1] && state[in2]
      when "OR"
        new_state[out] = state[in1] || state[in2]
      when "XOR"
        new_state[out] = state[in1] ^ state[in2]
      end
    }

    break if state == new_state
    state = new_state
  }
  state
end

def whats_the_number wires
  value = 0
  wires.map{|k,v| [k,v]}.sort{|a,b| a[0] <=> b[0]}.each_with_index {|kv,idx| 
    value = value + ((kv[1] ? 1 : 0)<< idx)
  }
  value
end

def find_bad_z_wire z, xy
  bits = z.to_s(2).chars.reverse.zip xy.to_s(2).chars.reverse
  bits.each_with_index {|pair, idx|
    return idx if pair[0] != pair[1]
  }
  return nil
end

def test wires, gates
  result = simulate wires, gates
  znum = whats_the_number(result.filter{|k,_| k[0] == "z"})
  xnum = whats_the_number(result.filter{|k,_| k[0] == "x"})
  ynum = whats_the_number(result.filter{|k,_| k[0] == "y"})

  return [znum==xnum+ynum, znum, xnum, ynum]
end

def does_xor_to_z wire, gates
  gates.filter{|i,op,i2,o|
    (i == wire || i2==wire) && op=="XOR" && o[0] == "z"
}.length>0
end

is_correct, znum,xnum,ynum = test wires, gates
pp znum

  # A z wire must be the result of a XOR
  # A z wire must not be XORd from the x and y wires at the same digit
  #   e.g. z01, x01, and y01 are 01 digit
  # The x and y wires at the same level must XOR to a non-z, which is then used to XOR to the z of the same level
wires_to_fix = gates.filter{|i,op,i2,o| 
  xy = ["x","y"]
  (op=="XOR" && o[0] != "z" && !xy.include?(i[0]) && !xy.include?(i2[0])) || 
  (o[0] == "z" && o != "z45" && op!="XOR") # ignore z45, its weird because it is the output

}.map{|_,_,_,o| o}

bad_z_xors = gates.filter{|i,op,i2,o| 
  # Find wires that are an AND of an X and y but XOR into a z. They should be a XOR of x & y
  # Ignore x00 and y00 because they are seed wires and are weird
  (["x","y"].include?(i[0]) && i != "x00" and i2 != "x00" && ["AND", "OR"].include?(op) && does_xor_to_z(o,gates)) 
}.map{|_,_,_,o| o}
wires_to_fix+= bad_z_xors

bad_z_xors.each {|bad|
  pi,_,_,_ = (gates.filter {|i,op,i2,o| 
    o == bad
  }[0])
  
  child_gates = gates.filter{|i,op,i2,o| 
    i == pi || i2 == pi
  }
  child_gates.each {|_,_,_,o| wires_to_fix << o}
}
pp wires_to_fix.uniq.sort.join(",")