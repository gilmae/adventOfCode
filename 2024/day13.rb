data = File.readlines("inputs/day13.input").map { |line| line.scan(/-?\d+/).map &:to_i }

def calculate_moves ax,ay,bx,by,px,py
  lowest_common_multiplier_of_ax_and_ay = ax.lcm(ay)
  ay_multiplier = lowest_common_multiplier_of_ax_and_ay / ay
  ax_multiplier = lowest_common_multiplier_of_ax_and_ay / ax

  dj = bx*ax_multiplier-by*ay_multiplier
  dp = px*ax_multiplier-py*ay_multiplier

  j =  dp/dj
  
  if dp%dj!=0
    return 0
  elsif 0 != (px-(j*bx))%ax
    return 0
  else
    i = (px-(j*bx))/ax
    return 3*i + j
  end
end

suma = 0
sumb = 0
(0..data.length-1).step(4){|idx|
  ax, ay = data[idx]
  bx,by = data[idx+1]
  px,py = data[idx+2]

  suma += calculate_moves ax,ay,bx,by,px,py
  sumb += calculate_moves ax,ay,bx,by,px+10000000000000,py+10000000000000
}

pp suma
pp sumb