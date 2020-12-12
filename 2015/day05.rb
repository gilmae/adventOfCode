require './helper.rb'
require 'optparse'

options = {}
OptionParser.new do |opt|
  opt.on('--input INPUT') { |o| options[:input] = o }
end.parse!

include Helper
input = options[:input] || "#{__FILE__}".gsub(/\.rb/, ".input")

data = get_data(input)

def is_nice? str
    return !(/((.)\2+)/ =~ str).nil? && str.scan(/[aeiou]/).length >= 3 && str.scan(/ab|cd|pq|xy/).length == 0
end

def is_new_nice? str
    return !(/((.).\2+)/ =~ str).nil? && !(/((..).*\2+)/ =~ str).nil?
end

puts data.map {|line|
    (is_nice? line)?1:0
}.reduce(&:+)
puts data.map {|line|
    (is_new_nice? line)?1:0
}.reduce(&:+)