require './helper.rb'
require 'optparse'

options = {}
OptionParser.new do |opt|
  opt.on('--input INPUT') { |o| options[:input] = o }
end.parse!

include Helper
input = options[:input] || "#{__FILE__}".gsub(/\.rb/, ".input")

data = get_data(input).join("")
floor = 0

first_time_in_basemennt = nil
data.chars.each_with_index{|f, i|
    floor += f == "(" ? 1:-1
    first_time_in_basemennt ||= i if floor == -1
}

puts floor, first_time_in_basemennt+1




