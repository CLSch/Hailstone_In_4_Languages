#! usr/bin/env python3

def hailstone(n):
    yield n
    while n != 1:
        if n % 2 == 0:
            n //= 2
        else:
            n = n * 3 + 1
        yield n

def hailstone_rec(n):
    yield n
    if n != 1:
        if n % 2 == 0:
            yield from hailstone(n // 2)
        else:
            yield from hailstone(n * 3 + 1)

for n in hailstone(3): 
    print(n)

for n in hailstone_rec(3): 
    print(n)
