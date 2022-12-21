input = ARGV[0] || "inputs/day21.input"
data = File.readlines(input).map(&:chomp)

monkeys = {}
data.each { |d|
  label, formula = d.split(":").map(&:strip)
  monkeys[label] = formula.split(" ")
}

def unpack(monkey, monkeys)
  formula = monkeys[monkey]

  if formula.length == 1
    return formula[0]
  end

  op1 = nil
  op2 = nil

  op1 = unpack formula[0], monkeys
  op2 = unpack formula[2], monkeys

  return op1.send(formula[1], op2) if op1.is_a?(Integer) && op2.is_a?(Integer)

  op1 = [op1] if !op1.is_a?(Array)
  op2 = [op2] if !op2.is_a?(Array)
  operand = [formula[1]]

  return ["("] + op1 + operand + op2 + [")"]
end

#part a
pp eval(unpack("root", monkeys).join(" "))

#part b
monkeys["humn"] = ["x"]
monkeys["root"][1] = "="

unpacked = (unpack("root", monkeys).join (" ")).slice(2..-3) # remove the paranthesis at the begining and end

sides = unpacked.split("=").map(&:strip)
known = nil
unknown = nil
if sides[0].include? "x"
  known = eval(sides[1])
  unknown = sides[0]
else
  known = eval(sides[0])
  unknown = sides[1]
end

min = (2 ** (0.size * 8 - 2) - 1)
max = 0

while known
  mid = (max + min) / 2

  test = unknown.gsub(/x/, mid.to_s)

  guess = eval(test)
  if guess == known
    pp mid
    break
  elsif guess < known
    min = mid
  elsif guess > known
    max = mid
  end
end
