def parse_file(filename):
    data = []
    with open(filename, "r") as file:
        for line in file:
            values = line.strip().split()
            row = [int(x) for x in values] 
            data.append(row[0])
    return data

def part_one(data):
    counter = 0
    prev_depth = 0 
    for line in data:
        depth = line
        if prev_depth > 0 and depth > prev_depth:
            counter += 1
        prev_depth = depth
    return counter

def part_two(data):
    counter = 0
    prev_depth = 0
    for i in range(len(data)-2):
        depth = data[i]+data[i+1]+data[i+2]
        if prev_depth > 0 and depth > prev_depth:
            counter += 1
        prev_depth = depth
    return counter

if __name__ == "__main__":
    data = parse_file("input.txt")
    print( "Part one:", part_one(data))
    print( "Part two:", part_two(data))
