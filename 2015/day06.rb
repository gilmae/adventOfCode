require './helper.rb'
require 'optparse'

options = {}
OptionParser.new do |opt|
  opt.on('--input INPUT') { |o| options[:input] = o }
end.parse!

include Helper
input = options[:input] || "#{__FILE__}".gsub(/\.rb/, ".input")

data = get_data(input)
# lights = Hash.new()

# data.each {|l|
#     puts l
#     /(?<cmd>.+)\s(?<startx>\d{1,3}),(?<starty>\d{1,3}) through (?<endx>\d{1,3}),(?<endy>\d{1,3})/ =~ l
#     (startx..endx).each{|x|
#         (starty..endy).each {|y|
#             case cmd
#             when "turn on"
#                 lights[[x,y]] = true
#             when "turn off"
#                 lights.delete([x,y])
#             when "toggle"
#                 if lights.include? [x,y]
#                     lights.delete([x,y])
#                 else
#                     lights[[x,y]] = true
#                 end
#             end
#         }
#     }
# }

# puts lights.keys.length

lights = Hash.new(0)

data.each {|l|
    puts l
    /(?<cmd>.+)\s(?<startx>\d{1,3}),(?<starty>\d{1,3}) through (?<endx>\d{1,3}),(?<endy>\d{1,3})/ =~ l
    (startx..endx).each{|x|
        (starty..endy).each {|y|
            case cmd
            when "turn on"
                lights[[x,y]] += 1
            when "turn off"
                lights[[x,y]] -= 1
                lights[[x,y]] = 0 if lights[[x,y]] < 0
            when "toggle"
                lights[[x,y]] += 2
            end
        }
    }
}

puts lights.values.reduce(&:+)

puts lights.keys.length

