from statistics import median

with open('input.txt') as f:
    positions = [int(n) for n in f.read().split(',')]

med = int(median(positions))
total = sum(abs(med - d) for d in positions)
print(total)
