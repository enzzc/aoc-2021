total_lines = 0
pos_count = [0] * 12

with open('input.txt', 'r') as f:
    for line in f:
        total_lines += 1
        for pos, bit in enumerate(line):
            if bit == '1':
                pos_count[pos] += 1

gamma_rate = ''
epsilon_rate = ''
for ones in pos_count:
    if ones > total_lines / 2:
        gamma_rate += '1'
        epsilon_rate += '0'
    else:
        gamma_rate += '0'
        epsilon_rate += '1'

gamma_rate = int(gamma_rate, 2)
epsilon_rate = int(epsilon_rate, 2)

print(gamma_rate * epsilon_rate)
