from collections import defaultdict

WIDTH = 1000
diagram = defaultdict(int)  # 1000×1000 array would be a bit large...

with open('input.txt', 'r') as f:
    for line in f:
        (x1, y1), (x2, y2) = [
            [int(n) for n in part.strip().split(',')]
            for part in line.split('->')
        ]

        if x1 == x2 == y1 == y2:
            diagram[WIDTH*x1+y1] += 1

        elif x1 == x2:
            for y in range(min(y1, y2), max(y1, y2)+1):
                diagram[WIDTH*x1+y] += 1

        elif y1 == y2:
            for x in range(min(x1, x2), max(x1, x2)+1):
                diagram[WIDTH*x+y1] += 1


total = sum(1 for x in diagram.values() if x > 1)
print(total)
