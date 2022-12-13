require "json"
input = ARGV[0] || "inputs/day13.input"
data = File.readlines(input).map(&:chomp).delete_if { |l| l.length == 0 }.map { |d| eval d }

def correct_order?(left, right)
  while true
    if left == right
      return nil
    else
      return left < right
    end if left.is_a?(Integer) && right.is_a?(Integer)

    left = [left] if left.is_a? Integer
    right = [right] if right.is_a? Integer

    idx = 0
    while true
      if left.length <= idx && right.length <= idx
        return nil
      elsif left.length <= idx
        return true
      elsif right.length <= idx
        return false
      end

      result = correct_order? left[idx], right[idx]
      return result if result != nil
      idx += 1
    end
  end
end

sum = 0

data.each_slice(2).each_with_index { |pairs, idx|
  left, right, _ = pairs
  sum += (idx + 1) if correct_order?(left, right)
}

pp sum

first = [[2]]
second = [[6]]
sorted_data = (data + [first, second]).sort do |a, b|
  co = correct_order? a, b
  case
  when co
    -1
  else
    1
  end
end

pp (sorted_data.index(first) + 1) * (sorted_data.index(second) + 1)
