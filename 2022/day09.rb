data = File.readlines("inputs/day09.input").map(&:chomp).map{|l|l.split(" ")}

class Coords 
    attr_accessor :x, :y, :last_x, :last_y

    def initialize x, y
        @x=x
        @y=y
    end

    def is_touching? c
        return (x-c.x).abs < 2 && (y-c.y).abs < 2
    end

    def to_s
        "#{x},#{y}"
    end

    def move d
        case d
        when "U"
            @y+=1
        when "D"
            @y-=1
        when "L"
            @x-=1
        when "R"
            @x+=1
        end
    end

    def move_towards c
        @x += get_movement (c.x-x)
        @y += get_movement (c.y-y)
    end

    def get_movement gap
        if gap == 0
            return 0
        elsif gap < 0
            return -1
        else
            return 1
        end
    end


end

def run_chain moves, chain
    tail_visited = {}
    
    moves.each {|line|
        line[1].to_i.times {|_|
            tail_visited[chain.last.to_s] = 1
            head = chain[0]
            head.move line[0]
            next_chain = [head]
            chain.slice(1,9).each_with_index {| knot,idx |
                friend = next_chain[idx]
                knot.move_towards(friend) if !knot.is_touching?(friend)
                next_chain << knot
            }

        }
        tail_visited[chain.last.to_s] = 1
    }
    tail_visited.length
end

pp run_chain data, [Coords.new(0,0),Coords.new(0,0)]

chain = []
10.times {|_|
    chain << Coords.new(0,0)
}

pp run_chain data, chain