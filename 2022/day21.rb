input = ARGV[0] || "inputs/day21.input"
data = File.readlines(input).map(&:chomp)

monkeys = {}
data.each { |d|
  label, formula = d.split(":").map(&:strip)
  monkeys[label] = formula.split(" ")
}

def unpack(monkey, monkeys)
  formula = monkeys[monkey]

  return formula[0].to_i if formula.length == 1

  op1 = unpack formula[0], monkeys
  op2 = unpack formula[2], monkeys

  return op1.send(formula[1], op2)
end

pp unpack "root", monkeys
