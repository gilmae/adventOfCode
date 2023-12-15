data = File.readlines("inputs/day12.input").map(&:chomp)

class SpringsTester
  attr_accessor :line, :nums

  def initialize
    @seen = {}
  end

  def solve(line_idx, nums_idx, size)
    # build a candidate line
    # Build it by iterating over each char in the source line, keeping track of
    #   what number we are trying to build a match for
    #   and the size of the block of #s we are currently trying to match
    # If the current char is a '#', increment the block size and check the next char
    # If the current char is a '.', check to see if we were tracking a block size.
    #   if we were and it matches the value of the current number, we are so far valid. Reset the block size and move to the next char and number
    #   if we were and it doesn't match the value of the current number, we're invalid, so stop
    #   if we weren't and move to the next char
    # If the current char is a '?', do the same tests simulating it being either a '.' or a '#'
    #
    # If we hit a state we have been in before, just return what is stored for that state
    # Otherwise if we have reached the end of the surce line,
    #      if we're in a block and the current size is the same as the last num, return 1
    #      if we're not in a block and we've tested all the numbers, return 1
    #      otherwise we're invalid and return 0
    #
    seen_key = [line_idx, nums_idx, size]
    return @seen[seen_key] if @seen.has_key? seen_key

    if line_idx == line.length
      return 1 if nums_idx == @nums.length && size == 0
      return 1 if nums_idx == @nums.length - 1 and size == @nums.last
      return 0
    end

    ans = 0
    [".", "#"].each { |c|
      if c == @line[line_idx] || @line[line_idx] == "?"
        ans += solve(line_idx + 1, nums_idx + 1, 0) if c == "." && size > 0 && nums_idx < @nums.length && size == @nums[nums_idx]
        ans += solve(line_idx + 1, nums_idx, 0) if c == "." && size == 0
        ans += solve(line_idx + 1, nums_idx, size += 1) if c == "#"
      end
    }
    @seen[seen_key] = ans
    ans
  end
end

pp data.map { |line|
  parts = line.split(" ")
  t = SpringsTester.new
  t.line = parts[0]
  t.nums = parts[1].split(",").map(&:to_i)
  t.solve 0, 0, 0
}.sum

pp data.map { |line|
  parts = line.split(" ")
  t = SpringsTester.new
  t.line = ([parts[0]] * 5).join("?")
  t.nums = parts[1].split(",").map(&:to_i) * 5
  t.solve 0, 0, 0
}.sum
