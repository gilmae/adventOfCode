require './helper.rb'
require 'optparse'

options = {}
OptionParser.new do |opt|
  opt.on('--input INPUT') { |o| options[:input] = o }
  opt.on('--wire WIRE') { |o| options[:wire] = o }
  opt.on('--debug 1|0') {|o| options[:debug] = o == 1}
end.parse!

include Helper
input = options[:input] || "#{__FILE__}".gsub(/\.rb/, ".input")

data = get_data(input)

wires = Hash.new(0)
memo = Hash.new(0)



wires = Hash[data.map{|line| 
    parts = line.split("->").map(&:strip)
    [parts[1], parts[0]]
}]


def get_value wires, memos, wire, debug = true
    puts "Get value of wire #{wire}" if debug
    return nil if wire.nil?

    if /^\d+$/ =~ wire
        memos[wire] = wire.to_i
        puts "Value (wire is numeric): #{memos[wire]}" if debug
        return memos[wire]
    end

    if memos.include? wire
        puts "Value (in memos): #{memos[wire]}" if debug
        return memos[wire]
    end

    formula = wires[wire]

    if /^\d+$/ =~ formula
        memos[wire] = formula.to_i
        puts "Value (formula is numeric): #{memos[wire]}" if debug
        return memos[wire]
    end

    puts "Parsing #{formula}" if debug

    formula_parts = formula.split(" ").map(&:strip)

    formula_parts.unshift nil if formula_parts.length == 2

    operand1 = get_value(wires, memos, formula_parts[0], debug) unless formula_parts[0].nil?

    return operand1 if formula_parts.length == 1
    operand2 = get_value(wires, memos, formula_parts[2], debug) unless formula_parts[1].nil?

    case formula_parts[1]
    when "NOT"
        memos[wire] = ~operand2
    when "AND"
        memos[wire] =  operand1 & operand2
    when "OR"
        memos[wire] =  operand1 | operand2
    when "LSHIFT"
        memos[wire] =  operand1 << operand2
    when "RSHIFT"
        memos[wire] =  operand1 >> operand2
    end
    puts "Returning #{formula} == #{memos[wire]}" if debug
    return memos[wire]
end

#p wires

wire_a = get_value wires, memo, (options[:wire] || "a"), options[:debug]
puts "Part A, Wire A receives #{wire_a}"

memo = Hash.new(0)
wires["b"] = wire_a.to_s

wire_a = get_value wires, memo, (options[:wire] || "a"), options[:debug]
puts "Part B, Wire A receives #{wire_a}"
#p memo

