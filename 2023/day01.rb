data = File.readlines("inputs/day01.input")

# part a
calibration = 0
data.each { |l|
  digits = l.scan(/\d/)
  calibration += digits[0].to_i * 10 + digits.last.to_i
}
puts calibration

NAMES = { "one" => 1, "two" => 2, "three" => 3, "four" => 4, "five" => 5, "six" => 6, "seven" => 7, "eight" => 8, "nine" => 9, "zero" => 0 }

def get_val(x)
  return x.to_i if x =~ /\d/
  NAMES[x]
end

def find_digits(input, digits)
  pattern = /(\d|one|two|three|four|five|six|seven|eight|nine|zero)/
  pos = input =~ pattern
  return if pos.nil?

  digits << pattern.match(input)[0]

  find_digits input[pos + 1..], digits
end

#part b
calibration = 0
data.each { |l|
  digits = []
  find_digits l, digits
  calibration += get_val(digits[0]) * 10 + get_val(digits.last)
}
puts calibration
