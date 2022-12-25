input = ARGV[0] || "inputs/day25.input"
data = File.readlines(input).map(&:chomp)

DIGITS = { "0" => 0, "1" => 1, "2" => 2, "-" => -1, "=" => -2 }
SNAFUS = { 0 => "0", 1 => "1", 2 => "2", -1 => "-", -2 => "=" }

def get_digital_from_snafu(num)
  val = 0
  num.chars.reverse.each_with_index do |ch, idx|
    val += DIGITS[ch] * 5 ** idx
  end
  val
end

def get_snafu_from_digital(d)
  snafu = []
  carry = 0

  while d > 0
    rem = d % 5
    d = d / 5
    if rem > 2
      d += 1
      rem = rem - 5
    end
    snafu << SNAFUS[rem]
  end

  if carry > 0
    snafu << SNAFUS[carry]
  end
  snafu.reverse.join("")
end

pp get_snafu_from_digital(data.map { |n| get_digital_from_snafu n }.sum)
