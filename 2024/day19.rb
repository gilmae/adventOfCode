data = File.readlines("inputs/day19.input").map(&:chomp)

towels = data[0].split(", ")
patterns = data[2..]

def find_builds towels, pattern, found = {}
  return 1 if pattern == ""
  return found[pattern] if found.has_key? pattern

  count = 0
  towels.each {|t|
    count += find_builds(towels, pattern[t.length..], found) if pattern.start_with? t
  }

  found[pattern] = count
  count
end

buildable = 0
total_builds = 0

patterns.each{|p|
  builds = find_builds(towels, p, {})
  buildable +=1 if builds > 0
  total_builds+=builds
}

pp buildable
pp total_builds