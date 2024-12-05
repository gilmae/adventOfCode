data = File.readlines("inputs/day05.input").map(&:chomp)

predecessors = {}
antecessors = {}

idx = 0

while data[idx] != ""
  p,a = data[idx].split("|")
  predecessors[p] ||= []
  predecessors[p] << a
  antecessors[a] ||= []
  antecessors[a] << p
  idx+=1
end 

idx+=1
updates = data[idx..]

pp updates.map {|u|#
  pages = u.split(",")

  valid = true
  pages[1..].each_with_index {|p,idx|
    pre = predecessors[pages[idx]]
    valid &= !pre.nil? && pre.include?(p)
  }
  valid ? pages[pages.length/2].to_i : 0
}.sum
