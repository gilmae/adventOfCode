data = File.readlines("inputs/day20.input").map(&:chomp)
FLIPFLOP = "%"
CONJUNCTION = "&"

modules = {}
# pulse [target, type, source]
# module [type, state, targets[], memory]
inputs = {}
data.each { |line|
  parts = line.split(" -> ")
  targets = parts[1].split(",").map(&:strip)
  _module = [nil, false, targets, {}]
  name = parts[0]

  case parts[0][0]
  when FLIPFLOP
    name = parts[0][1..]
    _module[0] = FLIPFLOP
  when CONJUNCTION
    name = parts[0][1..]
    _module[0] = CONJUNCTION
  end

  modules[name] = _module

  targets.each { |t|
    if inputs[t].nil?
      inputs[t] = [name]
    else
      inputs[t] << name
    end
  }
}

inputs.each { |k, v|
  m = modules[k]
  next if m.nil?

  v.each { |i| m[3][i] = false }
}
watchlist = inputs[inputs["rx"][0]]
pulses = { true => 0, false => 0 }

last_time_seen = {}
times_seen = {}
cycles = []
time = 1

def gcd(a, b)
  return a if b == 0
  gcd(b, a % b)
end

# A method that returns the least common multiple (LCM) of two numbers using the GCD method above
def lcm(a, b)
  (a * b) / gcd(a, b)
end

def lcm_of_array(array)
  lcm = array.reduce(1) { |lcm, n| lcm(lcm, n) }
  return lcm
end

while true
  queue = [["broadcaster", false]]

  while !queue.empty?
    target, pulse, from = queue.shift

    if !pulse
      times_target_seen = times_seen[target] || 0
      if watchlist.include?(target) && last_time_seen.has_key?(target) && times_target_seen == 2
        cycles << time - last_time_seen[target]
      end

      times_seen[target] = (times_seen[target] || 0) + 1
      last_time_seen[target] = time
    end

    if cycles.length == watchlist.length
      pp lcm_of_array(cycles)
      exit
    end

    pulses[pulse] += 1

    _module = modules[target]
    next if _module.nil?

    new_pulse = nil
    case _module[0]
    when FLIPFLOP
      if !pulse
        _module[1] = !_module[1]
        new_pulse = _module[1]
        _module[2].each { |t|
          #puts "#{target} -#{_module[1] ? "high" : "low"}-> #{t}"
          queue << [t, new_pulse, target]
        }
      end
    when CONJUNCTION
      _module[3][from] = pulse
      new_pulse = !inputs[target].map { |i| _module[3][i] }.all? { |i| i }
      _module[2].each { |t|
        #puts "#{target} -#{send_high_pulse ? "high" : "low"}-> #{t}"
        queue << [t, new_pulse, target]
      }
    else
      new_pulse = pulse
      _module[2].each { |t|
        #puts "#{target} -#{pulse ? "high" : "low"}-> #{t}"
        queue << [t, new_pulse, target]
      }
    end

    modules[target] = _module
  end
  pp pulses.values.reduce(&:*) if time == 1000
  time += 1

  #break if time == 1010
end
