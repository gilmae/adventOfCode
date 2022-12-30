def lcm(a, b)
  x, y = a, b
  x, y = b, a if a < b

  rem = x % y
  while (rem != 0)
    x = y
    y = rem
    rem = x % y
  end
  a * b / y
end
