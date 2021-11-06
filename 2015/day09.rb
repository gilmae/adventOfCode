require "./helper.rb"
require "optparse"

options = {}
OptionParser.new do |opt|
  opt.on("--input INPUT") { |o| options[:input] = o }
end.parse!

include Helper
input = options[:input] || "#{__FILE__}".gsub(/\.rb/, ".input")

data = get_data(input).map { |l| l.chomp }

intercities = {}
cities = []

parser = /(\w+)\sto\s(\w+) = (\d+)/
data.each { |l|
  m = l.match parser

  cities << m[1]
  cities << m[2]
  intercities[m[1]] ||= {}
  intercities[m[2]] ||= {}
  intercities[m[2]][m[1]] = m[3].to_i
  intercities[m[1]][m[2]] = m[3].to_i
}

cities = cities.uniq
#p cities

def find_path(cur_path, intercities, paths, length)
  #p cur_path
  #p intercities[cur_path.last]
  valid_destinations = intercities[cur_path.last].clone
  valid_destinations.delete_if { |k, v| cur_path.include? k }
  valid_destinations.keys.each { |city|
    path = cur_path + [city]
    find_path path, intercities, paths, length + valid_destinations[city]
  }
  paths[cur_path] = length
end

paths = {}
cities.each { |c|
  find_path [c], intercities, paths, 0
}

paths.delete_if { |k, v| k.length < cities.length }
p paths.values.min
p paths.values.max
