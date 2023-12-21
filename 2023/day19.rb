data = File.readlines("inputs/day19.input").map(&:chomp)

buckets = { "A" => [], "R" => [] }
rules = {}
rule_order = []

idx = 0
line = data[idx]

while line != ""
  m = /(\w+){([^}]+)}/.match line

  rule_key = m[1]
  rule_def = m[2].split(",").map { |tester|
    parts = tester.split(":")
    if parts.length == 1
      [nil, parts[0]]
    else
      [[parts[0][0], parts[0][1], parts[0][2..].to_i], parts[1]]
    end
  }
  rules[rule_key] = rule_def
  rule_order << rule_key
  idx += 1
  line = data[idx]
end

parts = data[idx + 1..].map { |p|
  Hash[*(p[1..p.length - 2].split(",").map { |a|
         ap = a.split("=")
         ap[1] = ap[1].to_i
         ap
       }).flatten(1)]
}

def test_rule(rule, part)
  rule.each { |t|
    tester = t[0]
    return t[1] if tester.nil?

    if part[tester[0]].send(tester[1], tester[2])
      #pp tester
      return t[1]
    end
  }
  return nil
end

part = { "x" => 787, "m" => 2665, "a" => 1222, "s" => 2876 }

# part A
parts.each { |part|
  dests = []
  dest = "in"
  while !["R", "A"].include? dest
    dest = test_rule rules[dest], part
    dests << dest
    if dest.nil?
      puts "Error found!"
      break
    end
  end
  buckets[dest] << part
}

pp buckets["A"].map { |p| p.values.sum }.sum

def get_new_split(op, n, lo, hi)
  case op
  when "<"
    hi = [hi, n - 1].min
  when "<="
    hi = [hi, n].min
  when ">"
    lo = [lo, n + 1].max
  when ">="
    lo = [lo, n].max
  end
  return [lo, hi]
end

def split_part_by_rules(attribute, op, n, x, m, a, s)
  case attribute
  when "x"
    x = get_new_split op, n, x[0], x[1]
  when "m"
    m = get_new_split op, n, m[0], m[1]
  when "a"
    a = get_new_split op, n, a[0], a[1]
  when "s"
    s = get_new_split op, n, s[0], s[1]
  end
  return [x, m, a, s]
end

queue = [["in", [1, 4000], [1, 4000], [1, 4000], [1, 4000], []]]
sum = 0
while !queue.empty?
  job = queue.pop
  state, x, m, a, s, path = job

  next if x[0] > x[1] || m[0] > m[1] || a[0] > a[1] || s[0] > s[1]

  if state == "A"
    sum += (x[1] - x[0] + 1) * (m[1] - m[0] + 1) * (a[1] - a[0] + 1) * (s[1] - s[0] + 1)
    next
  end

  next if state == "R"

  # otherwise create new splits based on rules
  tests = rules[state]

  new_states = []
  tests.each { |t|
    if !t[0].nil?
      attribute, op, n = t[0]
      # Get the ranges of parts that match this rule and add them to the queue
      nx, nm, na, ns = split_part_by_rules attribute, op, n, x, m, a, s
      queue.append([t[1], nx, nm, na, ns, path + [state]])

      # The inverse range failed the rule and can be sent to the enxt tes
      x, m, a, s = split_part_by_rules attribute, (op == ">" ? "<=" : ">="), n, x, m, a, s
    else
      # No test, so every part goes straight through to the next state
      # if there were subsquent tests, they are unreachable so just break
      queue.append([t[1], x, m, a, s, path + [state]])
      break
    end
  }
  queue += new_states
end
pp sum
