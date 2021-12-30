/*

No Code. Worked it out by hand

Repeatedly performs an 18 line function


Takes three inputs

Line +5  DIVIDER
Line +6  CHECK
Line +16

And reads one digit from inputs

line +1

Set w = input
Set x = (z%26) + CHECK
set z = z / DIVIDER
set x = x!=w?1:0
set z = z * ((25*x)+1)
set z = z + ((w+offset)*x)

if DIVIDER == 26
    Pop last digit from z
if last digit of z + CHECK == w then
    z = z
else
    z = z * 26 + w + OFFSET


DIVIDER is always 26 if CHECK <= 0

SO

If CHECK >0
    z = z * 26 + (w+offset) // PUSH
else
    POP last digit
    mem = z % 26
    z = z / 26
    if mem+check !+ w
        z = 26*z+(w+offset)

z is a radix 26 number that is having digits popped off and pushed on like a stack

So to get z==0 at the end, we must empty the stack, i.e. pop the last digit off z

... follow along with https://github.com/kemmel-dev/AdventOfCode2021/blob/master/day24/AoC%20Day%2024.pdf which is a better write up than I am doing
*/





