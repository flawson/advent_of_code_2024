#!/usr/bin/env python3

import re

INCLUDE = 1
EXCLUDE = 0

def parse_all_mults(input):
    regex = re.compile(r'mul\((\d+),(\d+)\)')
    matches = regex.findall(input)
    return matches

def parse_valid_mults(input):
    regex = re.compile(r"(do(n't)?\(\))|(mul\((\d+),(\d+)\))")
    matches = regex.findall(input)
    return matches

def part1(data):
    parsed_data = parse_all_mults(data)
    
    sum = 0
    for entry in parsed_data:
        product = int(entry[0])*int(entry[1])
        sum += product
    
    print(sum)

def part2(data):
    parsed_data = parse_valid_mults(data)
    
    sum = 0
    state = INCLUDE
    for entry in parsed_data:
        if entry[0] == '' and entry[2] == '':
            continue

        print(entry)
        if "don't()" in entry[0]:
            state = EXCLUDE
            print('EXCLUDE')
        if "do()" in entry[0]:
            state = INCLUDE
            print('INCLUDE')
        if state == INCLUDE:
            if 'mul' in entry[2]:
                product = int(entry[3])*int(entry[4])
                sum += product
    print(sum)

def main():
    filename = 'input.txt'
    data = ''

    with open(filename) as fh:
        data = fh.read()

    test_input = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

    part1(data)
    part2(data)


if __name__ == '__main__':
    main()
