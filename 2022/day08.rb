data = File.readlines("inputs/day08.input").map(&:chomp)
board = []
visibilities = []
scenicness = []
trees_seeable = []

def get_scenicness trees, tree 
    trees.reverse.each_with_index {|t,idx|
        if t >= tree
            return idx+1 
        end
    }
    return trees.length
end

def get_visible_trees trees, tree 
    count = 0
    max_interceding_height = -1
    trees.each_with_index {|t,idx|
        if t >= max_interceding_height
            max_interceding_height = t
            count+=1
        end
    }
    return count
end

data.each_with_index {| line, y| 
    board[y] = []
    visibilities[y] = Array.new(line.length,0)
    line.chars.each_with_index { | tree, x | 
        board[y][x] = tree.to_i

    }
}

board.each_with_index{|col, y|
    col.each_with_index {|tree, x|
        if x==0 || y == 0 || y==board.length-1 || x==board[y].length-1 # on the edge
            visibilities[y][x] = 1
        else
            up = (0..y-1).map{|dy|board[dy][x]}
            down = (y+1..board.length-1).map{|dy| board[dy][x]}
            left = (0..x-1).map{|dx| board[y][dx]}
            right = (x+1..board.length-1).map{|dx| board[y][dx]}
            scenicness << [get_scenicness(up,tree),get_scenicness(down.reverse,tree),get_scenicness(left,tree),get_scenicness(right.reverse,tree)].inject(:*)
            trees_seeable << [get_visible_trees(up,tree),get_visible_trees(down.reverse,tree),get_visible_trees(left,tree),get_visible_trees(right.reverse,tree)].inject(:*)
            tree = board[y][x] 
            visibilities[y][x] = 1 if tree > up.max || tree > down.max || tree > left.max || tree > right.max
        end
    }
}

pp visibilities.map{|col| col.sum}.sum
pp scenicness.max
pp trees_seeable.max