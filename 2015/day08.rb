require "./helper.rb"
require "optparse"

options = {}
OptionParser.new do |opt|
  opt.on("--input INPUT") { |o| options[:input] = o }
end.parse!

include Helper
input = options[:input] || "#{__FILE__}".gsub(/\.rb/, ".input")

def parseLine(line)
  content = line[1..-2].chars # slice off the leading and trailing "
  index = 0
  stringLength = 0
  while index < content.length
    case content[index]
    when "\\"
      case content[index + 1]
      when "\\"
        index += 1
      when "\""
        index += 1
      when "x"
        index += 3
      else
      end
    end
    stringLength += 1
    index += 1
  end
  return stringLength
end

def escapeLine(line)
  newLine = ["\""]
  line.chars.each { |c|
    case c
    when "\""
      newLine << "\\"
    when "\\"
      newLine << "\\"
    end
    newLine << c
  }
  newLine << "\""
  return newLine.join("")
end

data = get_data(input).map { |l| l.chomp }

codeLength = data.map { |l| l.length }.reduce(:+)
stringLength = data.map { |l| parseLine l }.reduce(:+)

escapedLength = data.map { |l| escapeLine(l).length }.reduce(:+)
p codeLength - stringLength
p escapedLength - codeLength
