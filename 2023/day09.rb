data = File.readlines("inputs/day09.input").map(&:chomp)

class Sequence
  attr_accessor :sequences

  def initialize(seq)
    @sequences = [seq]
    self.parse
  end

  def parse
    while true
      next_seq = []
      @sequences.last.each_cons(2) { |a, b|
        next_seq << b - a
      }

      @sequences << next_seq

      break if next_seq.all? { |d| d == 0 }
    end
  end

  def get_next
    idx = @sequences.length - 1
    while idx > 0
      @sequences[idx - 1] << (@sequences[idx - 1].last + @sequences[idx].last)
      idx -= 1
    end
    @sequences[0].last
  end

  def get_previous
    idx = @sequences.length - 1
    while idx > 0
      @sequences[idx - 1].prepend(@sequences[idx - 1].first - @sequences[idx].first)
      idx -= 1
    end
    @sequences[0].first
  end
end

sequences = data.map { |line|
  s = Sequence.new(line.scan(/-?\d+/).map(&:to_i))
}

pp sequences.map { |s|
  s.get_next
}.sum

pp sequences.map { |s|
  s.get_previous
}.sum
