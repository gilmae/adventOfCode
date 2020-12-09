require './helper.rb'
require 'optparse'

options = {}
OptionParser.new do |opt|
  opt.on('--input INPUT') { |o| options[:input] = o }
end.parse!

include Helper
input = options[:input] || "#{__FILE__}".gsub(/\.rb/, ".input")

data = get_data(input).join("")

class Santa
    attr_accessor :map, :position
    def initialize
        @map = Hash.new(0)
    end

    def visit c
        @map[c.to_a] +=1
        @position = c
    end
end

class Coords
    attr_accessor :x, :y

    def initialize x,y
        @x = x
        @y = y
    end

    def to_a
        [@x,@y]
    end
end

def get_coords direction, coords
    case direction
    when '<'
        return Coords.new(coords.x-1, coords.y)
    when '>'
        return Coords.new(coords.x+1, coords.y)
    when '^'
        return Coords.new(coords.x, coords.y+1)
    when 'v'
        return Coords.new(coords.x, coords.y-1)
    end
end


def track_santas paths, santas
    c = Coords.new(0,0)
    santas.each{|s| s.visit c}

    index = 0

    paths.chars.each{|d|
        c = get_coords d, santas[index].position
        santas[index].visit c
        index = (index + 1) % santas.length
    }
end

santas = [Santa.new]
track_santas data, santas

p santas.reduce([]){|m,v| m + v.map.keys}.uniq.length


santas = [Santa.new, Santa.new]
track_santas data, santas

p santas.reduce([]){|m,v| m + v.map.keys}.uniq.length