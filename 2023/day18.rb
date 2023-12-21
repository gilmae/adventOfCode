data = File.readlines("inputs/day18.input").map(&:chomp)
board = {}
actions = data.map { |line|
  line.split(" ")
}

board = { [0, 0] => "#000000" }
path = [[0, 0]]
pos = [0, 0]
path_length = 1

path2 = [[0, 0]]
pos2 = [0, 0]
path_length2 = 1

actions.each { |a|
  case a[0]
  when "U"
    delta = [0, -1]
  when "D"
    delta = [0, 1]
  when "L"
    delta = [-1, 0]
  when "R"
    delta = [1, 0]
  end
  #   a[1].to_i.times { |_|
  #     pos = [pos[0] + delta[0], pos[1] + delta[1]]
  #     board[pos] = a[2]
  #     path << pos
  #   }

  distance = a[1].to_i
  pos = [pos[0] + distance * delta[0], pos[1] + distance * delta[1]]
  path << pos
  path_length += distance

  case a[2][7]
  when "3" # up
    delta = [0, -1]
  when "2" # left
    delta = [-1, 0]
  when "1" # down
    delta = [0, 1]
  when "0" # right
    delta = [1, 0]
  end
  distance = a[2][2..6].to_i(16)
  path_length2 += distance

  pos2 = [pos2[0] + distance * delta[0], pos2[1] + distance * delta[1]]
  path2 << pos2
}

area = (path.each_with_index.map { |p, idx|
  idx2 = (idx + 1) % path.length
  p2 = path[idx2]

  (p[1] + p2[1]) * (p[0] - p2[0])
}.sum / 2).abs()

pp area + (path_length / 2) + 1

area2 = (path2.each_with_index.map { |p, idx|
  idx2 = (idx + 1) % path2.length
  p2 = path2[idx2]

  (p[1] + p2[1]) * (p[0] - p2[0])
}.sum / 2).abs()

pp area2 + (path_length2 / 2) + 1
