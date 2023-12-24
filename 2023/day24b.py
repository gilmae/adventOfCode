import z3

infile = open("inputs/day24.input")
stones = []
vectors = []

for line in infile:
    line = line.strip()
    p, v = line.split(" @ ")
    stones.append(tuple(map(float, p.split(", "))))
    vectors.append(tuple(map(float, v.split(", "))))

x = z3.Real('x')
y = z3.Real('y')
z = z3.Real('z')
vx = z3.Real('vx')
vy = z3.Real('vy')
vz = z3.Real('vz')

s = z3.Solver()

for i in range(len(stones)):
    x_i, y_i, z_i = stones[i]
    vx_i, vy_i, vz_i = vectors[i]
    ti = z3.Real(f"t{i}")
    s.add(x_i + vx_i * ti == x + vx * ti)
    s.add(y_i + vy_i * ti == y + vy * ti)
    s.add(z_i + vz_i * ti == z + vz * ti)

print(s.check())
m = s.model()
print(m[x], m[y], m[z])
