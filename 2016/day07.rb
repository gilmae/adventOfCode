data = File.readlines("inputs/day07.input").map(&:chomp)

def check_chunk_for_abba(chunk)
  (0..chunk.length - 4).each { |idx|
    return true if chunk[idx] == chunk[idx + 3] && chunk[idx + 1] == chunk[idx + 2] && chunk[idx] != chunk[idx + 1]
  }
  return false
end

def is_bracketed(chunk)
  return chunk[0] == "[" && chunk.end_with?("]")
end

pp data.map { |line|
  chunks = line.scan(/(\[?\w+\]?)/).map { |i| i[0] }

  chunks.map { |chunk|
    if check_chunk_for_abba(chunk)
      !is_bracketed(chunk)
    else
      nil
    end
  }.delete_if { |i| i.nil? }
}.delete_if { |i| i.empty? || i.any? { |j| j == false } }.length

def check_for_aba(chunk)
  abas = []
  (0..chunk.length - 3).each { |idx|
    if chunk[idx] == chunk[idx + 2] && chunk[idx] != chunk[idx + 1]
      abas << chunk[idx..idx + 1].chars
    end
  }
  return abas
end

def check_for_bab(chunk, aba)
  (0..chunk.length - 3).each { |idx|
    return true if chunk[idx] == aba[1] && chunk[idx + 2] == aba[1] && chunk[idx + 1] == aba[0]
  }
  return false
end

def supports_ssl(line)
  chunks = line.scan(/(\[?\w+\]?)/).map { |i| i[0] }
  outside_squares = chunks.filter { |c| !is_bracketed(c) }
  inside_squares = chunks.filter { |c| is_bracketed(c) }
  abas = []
  outside_squares.each { |chunk|
    abas += check_for_aba chunk
  }
  return false if abas.length == 0

  abas.each { |aba|
    next if aba.nil? || aba.empty?
    inside_squares.each { |chunk|
      return true if check_for_bab(chunk, aba)
    }
  }
  return false
end

pp data.map { |line|
  supports_ssl line
}.delete_if { |i| !i }.length
