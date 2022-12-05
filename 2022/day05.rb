data = File.readlines("inputs/day05.input")
line = 0
rex = /(\[[A-Z]\]\s?|\s{3}\s?)/
md = data[line].scan(rex)
stacks9k = []

while md.length >0
    md.each_with_index { | item, idx | 
    if item[0].strip != ""
        stacks9k[idx+1] = (stacks9k[idx+1] || "") + item[0][1]
    end
    }
    line +=1
    md = data[line].scan(rex)
end
stacks9k1 = stacks9k.clone
line +=1
rex = /move (\d+) from (\d+) to (\d+)/
while line < data.length 
    count,src,dest =  data[line].scan(rex)[0].map(&:to_i)

    pile9k = stacks9k[src].slice(0,count).reverse
    pile9k1 = stacks9k1[src].slice(0,count)

    stacks9k[src] = stacks9k[src][count..-1]
    stacks9k1[src] = stacks9k1[src][count..-1]
    
    stacks9k[dest] = pile9k + stacks9k[dest]
    stacks9k1[dest] = pile9k1 + stacks9k1[dest]
    
    line +=1
end

stacks9k = stacks9k[1..-1]
stacks9k1 = stacks9k1[1..-1]

puts(stacks9k.map{|i|i[0]}.join(""))
puts(stacks9k1.map{|i|i[0]}.join(""))