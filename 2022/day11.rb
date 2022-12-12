data = File.readlines("inputs/day11.input").map(&:chomp)

class Integer
  def prime_factors
    num = self
    (2...num).each do |fact|
      if (num % fact == 0)
        other_fact = num / fact
        return [*fact.prime_factors, *other_fact.prime_factors]
      end
    end
    return self
  end
end

class Monkey
  attr_accessor :items, :trueDest, :falseDest, :divisor, :mutator, :inspections

  def parse_monkey(lines)
    @items = lines[1].split(":")[1].split(",").map(&:to_i)
    @divisor = lines[3][21..-1].to_i
    @trueDest = lines[4][29..-1].to_i
    @falseDest = lines[5][30..-1].to_i
    @mutator = lines[2].split("=")[1].strip.split(" ").map(&:strip)
    @inspections = 0
  end

  def play(item, modulus, worry_adjuster)
    @inspections += 1

    n = item
    operand = 0
    case @mutator[2]
    when "old"
      operand = item
    else
      operand = @mutator[2].to_i
    end
    v = (n.send(@mutator[1], operand) / worry_adjuster).to_i
    v %= modulus if modulus
    v
  end

  def throw_to(item)
    if item % @divisor == 0
      @trueDest
    else
      @falseDest
    end
  end
end

def run(monkeys, rounds, modulus, worry_adjuster)
  rounds.times { |round|
    monkeys.each { |monkey|
      monkey.items.length.times { |_|
        item = monkey.items.pop
        n = monkey.play item, modulus, worry_adjuster
        who_to = monkey.throw_to n
        monkeys[who_to].items << n
      }
    }
  }
  monkeys.map { |m| m.inspections }
end

monkeys = []
data.each_slice(7).each_with_index { |lines, idx|
  monkey = Monkey.new
  monkey.parse_monkey lines
  monkeys << monkey
}

inspections = run monkeys, 20, nil, 3
pp inspections.sort.reverse[0..1].reduce(&:*)

monkeys = []
data.each_slice(7).each_with_index { |lines, idx|
  monkey = Monkey.new
  monkey.parse_monkey lines
  monkeys << monkey
}

modulus = monkeys.map { |m| m.divisor }.reduce(&:*)

inspections = run monkeys, 10000, modulus, 1
pp inspections.sort.reverse[0..1].reduce(&:*)
