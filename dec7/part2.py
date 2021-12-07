from statistics import mean


def distance_cost(start, end):
    start, end = min(start, end), max(start, end)
    acc = 0
    last = 0
    for _ in range(start, end):
        acc += 1 + last
        last += 1
    return acc


with open('input.txt') as f:
    positions = [int(n) for n in f.read().split(',')]

mu = int(mean(positions))
total = sum(distance_cost(mu, d) for d in positions)
print(total)
