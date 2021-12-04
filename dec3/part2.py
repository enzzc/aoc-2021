from copy import deepcopy

LENGTH = 12
init_bitsets = []


def get_rating(init_bitsets, co2=False):
    bitsets = deepcopy(init_bitsets)
    for pos in range(LENGTH):
        onebit_count = 0
        total_lines = 0
        if len(bitsets) <= 1:
            break
        for bitset in bitsets:
            total_lines += 1
            bit = bitset[pos]
            if bit == '1':
                onebit_count += 1
            if co2:
                bit_criteria = '0' if onebit_count >= total_lines / 2 else '1'
            else:
                bit_criteria = '1' if onebit_count >= total_lines / 2 else '0'
        bitsets = [
            num for num in bitsets
            if num[pos] == bit_criteria
        ]
    return int(''.join(bitsets[0]), 2)


with open('input.txt', 'r') as f:
    for line in f:
        acc = []
        for pos, bit in enumerate(line):
            if bit == '1' or bit == '0':
                acc.append(bit)
        init_bitsets.append(acc)

o2 = get_rating(init_bitsets)
co2 = get_rating(init_bitsets, co2=True)

print(o2 * co2)
