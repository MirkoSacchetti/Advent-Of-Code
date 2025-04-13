def parse_file(filename):
    data = []
    with open(filename, "r") as file:
        for line in file:
            values = line.strip().split()
            row =  {"comm": values[0], "unit": int(values[1])}
            data.append(row)
    return data

def part_one(data):
    position = 0
    depth = 0 
    for c in data:
        if  c["comm"] == "forward":
            position+= c["unit"]
        if  c["comm"] == "down":
            depth+= c["unit"]
        if  c["comm"] == "up":
            depth-= c["unit"]
    return position * depth

def part_two(data):
    position = 0
    depth = 0 
    aim = 0 
    for c in data:
        if  c["comm"] == "forward":
            position+= c["unit"]
            depth+= c["unit"] * aim
        if  c["comm"] == "down":
            aim+= c["unit"]
        if  c["comm"] == "up":
            aim-= c["unit"]
    return position * depth


if __name__ == "__main__":
    data = parse_file("input.txt")
    print( "Part one:", part_one(data))
    print( "Part two:", part_two(data))
