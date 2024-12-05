data = File.readlines("inputs/day05.input").map(&:chomp)

predecessors = {}
idx = 0

while data[idx] != ""
  p,a = data[idx].split("|")
  predecessors[p] ||= []
  predecessors[p] << a
  idx+=1
end 

partA = 0
partB = 0

data[idx+1..].each {|u|
  pages = u.split(",")
  test = pages.sort{|a,b| (predecessors.has_key?(a) && predecessors[a].include?(b))?-1:1}.join(",")
  
  if test != u
    partB += pages[pages.length/2].to_i
  else
    partA += pages[pages.length/2].to_i
  end
}

pp partA, partB