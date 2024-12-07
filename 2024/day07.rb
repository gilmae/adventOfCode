data = File.readlines("inputs/day07.input").map(&:chomp)

def test total, nums, operators
  f = nums.shift
  work = [[f, nums, []]]
  paths = []  
  loop {
    job = work.shift
    break if job.nil?

    t, n, ops = job
    if t==total && n.empty?
      return ops
    elsif t > total
      next
    end
    num, *rest = n

    next if num.nil?

    operators.each {|op|
    
    case op
      when "+"
        work.unshift([t+num, rest, ops+[op]])
      when "*"
        work.unshift([t*num, rest, ops+[op]])
      when "||"
        work.unshift([(t.to_s+num.to_s).to_i, rest, ops+[op]])
      end
    }
  
  }

  return paths
end

partA = 0
partB = 0
data.each {|line|
  parts = line.split(": ")
  solutions = test parts[0].to_i, parts[1].split(" ").map(&:to_i), ['+', '*']
  partA+= parts[0].to_i if !solutions.empty?

  solutions = test parts[0].to_i, parts[1].split(" ").map(&:to_i), ['+', '*', "||"]
  partB+= parts[0].to_i if !solutions.empty?
}

pp partA
pp partB