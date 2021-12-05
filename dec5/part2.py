from collections import defaultdict

WIDTH = 1000
diagram = defaultdict(int)  # 1000×1000 array would be a bit large...


def path(p1, p2):
    (x1, y1), (x2, y2) = p1, p2

    yield x1, y1
    while True:
        if x1 == x2 and y1 == y2:
            break

        if x1 < x2:
            x1 += 1
        elif x1 > x2:
            x1 -= 1

        if y1 < y2:
            y1 += 1
        elif y1 > y2:
            y1 -= 1

        yield x1, y1


with open('input.txt', 'r') as f:
    for line in f:
        (x1, y1), (x2, y2) = [
            [int(n) for n in part.strip().split(',')]
            for part in line.split('->')
        ]
        for x, y in path((x1, y1), (x2, y2)):
            diagram[WIDTH*x+y] += 1

total = sum(1 for x in diagram.values() if x > 1)
print(total)
