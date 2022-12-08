data = File.readlines("inputs/day07.input")


class Node 
    attr_accessor :parent, :contents, :size

        def initialize parent, size
            @parent = parent
            @size = size
            @contents = {}
        end
end

def parse_item ctx, line 
    pre,post = line.split(" ")

      if pre == "dir"
        newF = Node.new(ctx, 0)
        newF.contents[".."] = ctx
        ctx.contents[post] = newF
      else
        ctx.contents[post] = Node.new(ctx, pre.to_i)
      end
end

def process_command ctx, line
    parts = line.split(" ")
    if parts[1] == "ls"
        return ctx
    end
    newctx = ctx.contents[parts[2]] if ctx.contents[parts[2]] 
    newctx||ctx
end


fs = Node.new("",0)
fs.contents["/"] = Node.new(nil, 0)
ctx = fs

data.each {|line|
    ctx = process_command(ctx, line) if line[0] == "$"
    parse_item(ctx, line) unless line[0] == "$"
}

def calcSize item
    item.contents.each {|k,v| (item.size += calcSize(v)) unless k==".."}
    item.size
end


def scanForNoMoreThan memo,Node, size
    if Node.contents.length > 0
        if Node.size <=size
            memo << Node.size
        end
        Node.contents.each{|k,item| scanForNoMoreThan(memo, item,size) unless k == ".."}
    end
end

def scanForAtLeast memo,Node, size
    if Node.contents.length > 0
        if Node.size >= size
            memo << Node.size
        end
        Node.contents.each{|k,item| scanForAtLeast(memo, item, size) unless k == ".."}
    end
end

calcSize(fs.contents["/"])

root = fs.contents["/"]

memo = []
scanForNoMoreThan(memo, root, 100000)
pp memo.sum

memo = []
scanForAtLeast(memo, root, 30000000 - (70000000-root.size))
pp memo.min


