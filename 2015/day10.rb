# Probaly works if I am prepared to sit waiting for days. NNeeds concurrenncy
# I don't really know how to do concurrency in Ruy, so switching to a go version, day10.go

require "./helper.rb"
require "optparse"

options = {}
OptionParser.new do |opt|
  opt.on("--input INPUT") { |o| options[:input] = o }
  opt.on("--iterations ITERATIONS") { |o| options[:iterations] = o.to_i }
end.parse!

include Helper
input = options[:input] || "#{__FILE__}".gsub(/\.rb/, ".input")

data = get_data(input).first

def look_say(seed)
  last = seed.chars.first
  result = []
  count = 0
  seed.chars.each { |c|
    if last != c
      result += [count.to_s, last]
      count = 1
    else
      count += 1
    end
    last = c
  }
  result += [count.to_s, last]

  return result.join("")
end

seed = data
iterations = input = options[:iterations] || 40
iterations.times { |i|
  p i
  seed = look_say seed
}
p seed
