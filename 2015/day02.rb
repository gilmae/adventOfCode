require './helper.rb'
require 'optparse'

options = {}
OptionParser.new do |opt|
  opt.on('--input INPUT') { |o| options[:input] = o }
end.parse!

include Helper
input = options[:input] || "#{__FILE__}".gsub(/\.rb/, ".input")

data = get_data(input)

calcs =  data.map{|box|
  dimensions = box.split("x").map{|i|i.to_i}.sort
  [
    dimensions[0]*dimensions[1]*2 + 2*dimensions[1]*dimensions[2]+ 2*dimensions[0]*dimensions[2] + dimensions[0]*dimensions[1],
    dimensions[0]*2+dimensions[1]*2 + dimensions[0]*dimensions[1]*dimensions[2]
  ]
}

puts calcs.reduce(0){|m,v|m+v[0]}
puts calcs.reduce(0){|m,v|m+v[1]}
