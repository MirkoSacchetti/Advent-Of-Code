def parse_file(filename):
    data = []
    with open(filename, "r") as file:
        for line in file:
            data.append(line.strip())
    return data

def part_one(data):
    gamma_rate = ["0"] * len(data[0])
    epsilon_rate = ["0"] * len(data[0])
    for index in range(len(data[0])):
        if most_common(data, index) == "1":
            # Gamma rate is made of most common bit in each position.
            gamma_rate[index] = "1"
        else:
            # Epsilon rate is made of least common bit in each position.
            epsilon_rate[index] = "1"
            
    return int("".join(gamma_rate), 2) * int("".join(epsilon_rate), 2) 

def part_two(data):
    oxygen_rate = ""
    scrub_rate = ""
    data_filtered_oxygen = data.copy()
    data_filtered_scrub = data.copy()
    
    for index in range(len(data[0])):
        # Process oxygen generator rating
        if len(data_filtered_oxygen) > 1:
            val_o = most_common(data_filtered_oxygen, index)
            data_filtered_oxygen = filter_values(data_filtered_oxygen, index, val_o)
            if len(data_filtered_oxygen) == 1:
                oxygen_rate = data_filtered_oxygen[0]
        # Process CO2 scrubber rating
        if len(data_filtered_scrub) > 1:
            val_s = less_common(data_filtered_scrub, index)
            data_filtered_scrub = filter_values(data_filtered_scrub, index, val_s)
            if len(data_filtered_scrub) == 1:
                scrub_rate = data_filtered_scrub[0]
        # Exit if we have all
        if oxygen_rate and scrub_rate:
            break
            
    return int(oxygen_rate, 2) * int(scrub_rate, 2)

def filter_values(data, index, val):
    return [r for r in data if r[index] == val]

def most_common(data, index):
    ones_count = sum(1 for r in data if r[index] == "1")
    if ones_count >= len(data) / 2:
        return "1"
    return "0"

def less_common(data, index):
    ones_count = sum(1 for r in data if r[index] == "1")
    if ones_count >= len(data) / 2:
        return "0"
    return "1"


if __name__ == "__main__":
    data = parse_file("input.txt")
    print("Part one:", part_one(data))
    print("Part two:", part_two(data))
