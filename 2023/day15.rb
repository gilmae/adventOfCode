data = File.readlines("inputs/day15.input").map(&:chomp)[0]

def hash(s)
  v = 0
  s.each_byte { |c|
    v = ((v + c) * 17) % 256
  }
  v
end

pp data.split(",").map { |p| hash p }.sum

hm = {}

def print(boxes)
  boxes.keys.sort.each { |k|
    puts "Box #{k}: #{boxes[k]}\n"
  }
  puts
end

def get_focal(boxes)
  fp = []
  boxes.keys.sort.each { |k|
    boxes[k].each_with_index { |item, idx|
      fp << (k + 1) * (idx + 1) * item[1].to_i
    }
  }
  fp.sum
end

data.split(",").each { |p|
  label, cmd, focal = (/(\w+)([=-])(\d?)/.match p)[1..]

  h = hash(label)
  box = hm[h]
  box = [] if box.nil?

  case cmd
  when "="
    idx = box.find_index { |i| i[0] == label }
    if idx.nil?
      box << [label, focal]
    else
      box[idx][1] = focal
    end
  when "-"
    idx = box.find_index { |i| i[0] == label }
    box.delete_at(idx) unless idx.nil?
  end

  hm[h] = box
}

pp get_focal hm
