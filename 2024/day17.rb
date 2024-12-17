data = File.readlines("inputs/day17.input").map(&:chomp).join("").scan(/-?\d+/).map(&:to_i)
a,b,c,*ops = data

def run_program ops, registers
  ip=0
  stdout = []
  loop {
    registers = registers.map(&:to_i)
    break if ip >= ops.length
    instruction = ops[ip]
    literal_operand = ops[ip+1]
    combo_operand = [0,1,2,3,registers[0],registers[1],registers[2]][literal_operand]
    ip+=2
    case (instruction) 
      when 0 #adv
        registers[0] = (registers[0] / 2**combo_operand).to_i
      when 1 #bxl
        registers[1] = registers[1] ^ literal_operand
      when 2 #bst
        registers[1] = combo_operand % 8
      when 3 #jnz
        ip = literal_operand if registers[0] != 0
      when 4 #bxc
        registers[1] = registers[1] ^ registers[2]
      when 5 #out
        stdout << (combo_operand%8)
      when 6 #bdv
        registers[1] = (registers[0]/2**combo_operand).to_i
      when 7 #cdv
        registers[2] = (registers[0]/2**combo_operand).to_i
    end
  }
  stdout
end

pp run_program(ops, [a,b,c]).map(&:to_s).join(",")
#pp run_program(ops, [48,b,c]).map(&:to_s).join(",")

#49
#53

work = [[0,1]]

best_a = 1e18
loop {
  job = work.shift
  break if job.nil?
  a, depth = job
  break if depth > ops.length

  8.times {|potential|
    result = run_program ops, [a*8+potential,0,0]
    next if result != ops[ops.length-depth..]
    if depth == ops.length
      if a*8+potential < best_a
        best_a = a*8+potential
      end
    else
      work << [a*8+potential, depth+1]
    end
  }
}
pp best_a