data = File.readlines("inputs/day22.input").map(&:to_i)

def prng secret
  secret = ((secret << 6) ^ secret) % 16777216
  secret = ((secret>>5) ^ secret) % 167777216
  ((secret << 11) ^ secret) % 16777216
end

def slice_sequence tallys, sequence, deltas
  (deltas.length-3).times {|idx|
   tallys[deltas[idx..idx+3]] = sequence[idx+4] if !tallys.has_key? deltas[idx..idx+3]
  }
end

tallys = Hash.new(0)
pp data.map{|secret|
  sequence = [secret]
  local_tally = {}
  2000.times {|_|
    sequence << prng(sequence.last)
  }

  ones = sequence.map{|s|s%10}
  deltas = ones.zip(ones[1..])
  slice_sequence( local_tally, ones, (deltas.first deltas.size-1).map{|a,b| b-a})
  local_tally.each {|k,v|
    tallys[k] += v
  }
  sequence.last
}.sum

pp tallys.max_by{|k,v| v}