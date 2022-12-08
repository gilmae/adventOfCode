def scanForPart data, start, length
    pos = start
    loop {
        part = data[pos..pos+length-1].split("")
        return pos+length if part.length == part.uniq.length
        pos +=1    
    }       
end
data = File.readlines("inputs/day06.input")

data = data[0]
pos = scanForPart(data,0,4)

pp pos
pp scanForPart(data,pos,14)