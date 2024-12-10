data = File.readlines("inputs/day09.input").map(&:chomp)[0]

def calculate disk
  disk.each_with_index.map {|item, idx|
    (item.nil?)?0:item * idx
  }.sum
end

disk = []
disk_map = data.chars.map(&:to_i)
free_space_map = []
data_map = []

ptr = 0
(0..disk_map.length-1).step(2).each{|idx|
  disk += [idx/2] * disk_map[idx]
  data_map << [ptr, idx/2, disk_map[idx]]
  ptr += disk_map[idx]
  break if disk_map.length <= idx+1
  disk += [nil] * disk_map[idx+1]
  free_space_map<< [ptr, disk_map[idx+1]]
  ptr += disk_map[idx+1]
}

part_b_disk = disk.clone

idx = -1
loop {
  
  break if disk.length <= idx
  if disk[idx].nil?
    item = disk.pop
    if !item.nil?
      disk[idx] = item
      idx+=1
    end
  else
    idx+=1
  end
  
}
pp calculate disk


data_map.reverse.each {|ptr, type, length|
  idx = free_space_map.index{|sptr, slength|
    break if sptr >= ptr
    length <= slength
  }
  next if idx.nil?
  space_ptr, slength = free_space_map[idx]
  mem_slice = part_b_disk.slice!(ptr, length)
  s = part_b_disk[0..space_ptr-1]
  m = part_b_disk[space_ptr+length..ptr-1]
  f = part_b_disk[ptr..part_b_disk.length]
  _ = part_b_disk.slice!(space_ptr, length)
  part_b_disk = s + mem_slice + m + ([nil]*length) 
  part_b_disk += f if !f.nil?
  free_space_map[idx] = [space_ptr+length, (slength-length)]
}


pp calculate part_b_disk