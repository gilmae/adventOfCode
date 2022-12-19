input = ARGV[0] || "inputs/day18.input"
data = File.readlines(input).map(&:chomp).map { |l| l.split(",") }

def get_neighbours x,y,z, board
    deltas = [[0,0,1], [0,0,-1], [0,1,0],[0,-1,0],[1,0,0],[-1,0,0]]

    neighbours = deltas.map{|d| 
        dx,dy,dz = d

        dx +=x
        dy+=y
        dz+=z

        [dx,dy,dz]
    }
end

def get_bounds(board)
    xs = []
    ys = []
    zs = []
    board.each { |k, _|
      x, y,z = k
      xs << x
      ys << y
zs << z
    }
    return [xs.min, ys.min, zs.min,xs.max, ys.max,zs.max]
end

board = {}

data.map {|d|
    x,y,z = d

    board[[x.to_i,y.to_i,z.to_i]] = true
}

pp board.map{|k,_|
    x,y,z = k
    get_neighbours(x,y,z,board).map{|n|
        board.include?(n) ? 0 : 1
    }.sum
}.sum

bounds = get_bounds board
# duh! I need to be checking from the cubes _around_ the lava cubes
3.times {|idx|
    bounds[idx] -=1
    bounds[idx+3] +=1
}
minx,miny,minz,maxx,maxy,maxz = bounds


# Flood fill the board and find the points that are lava
# or more specifically, the lava cubes that touch a "face" of an air cube
outer_surface = []
work = [[minx,miny,minz]]
visited = {}

while work.length > 0
    current = work.shift
    next if visited.include?(current)

    x,y,z = current
    get_neighbours(x,y,z,board).each_with_index{|n,idx| #we'll treat idx as a face
        nx,ny,nz = n

        next if nx<minx || nx > maxx || ny < miny || ny > maxy || nz < minz || nz > maxz

        if board.include?(n)
            outer_surface << [n,idx]
        else # n is more empty air
            work.append(n)
        end
    }
    visited[current] = true
end

pp outer_surface.length