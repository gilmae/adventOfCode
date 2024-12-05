data = File.readlines("inputs/day03.input")

def scan str
  muls = str.scan(/mul\((\d{1,3}),(\d{1,3})\)/)

  muls.map {|mul|
    mul[0].to_i * mul[1].to_i
  }.sum
end

s = data.join("")
pp scan s

partb = 0
while !s.nil? && s.size > 0 
 slice_end = s.index "don't()"
 sub_str = ""
  if !slice_end.nil?
   sub_str = s[0..slice_end+6]
   s = s[slice_end+6..]
  else
    sub_str = s
    s = ""
  end
 
  partb += scan(sub_str)
  slice_end = s.index "do()"
 
  s = (slice_end.nil?) ? "" : s[slice_end+3..]
end

pp partb
