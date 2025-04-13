def parse_file(filename):
    data = []
    with open(filename, "r") as file:
        for line in file:
            values = line.strip()
            data.append(values)
    return data

def part_one(data):
    gamma_rate = ["0"] * len(data[0])
    epsilon_rate = ["0"] * len(data[0])
    for index in range(len(data[0])):
        # find the most/less common and update the corrispont rate
        if most_common(data,index) == "1":
            gamma_rate[index] = "1"
        else:
            epsilon_rate[index] = "1"
    return int("".join(gamma_rate),2) *  int("".join(epsilon_rate),2) 

def part_two(data):
    oxygen_rate = [] 
    scrub_rate = [] 
    data_filtered_oxygen= data
    data_filtered_scrub= data
    completed = 0
    for index in range(len(data[0])):
        # find the most/less common for each filtered dataset 
        val_o = most_common(data_filtered_oxygen, index)
        val_s = less_common(data_filtered_scrub, index)
        # filter each dataset with the updated most/less value at the curret index
        data_filtered_oxygen = filter(data_filtered_oxygen, index, val_o)
        data_filtered_scrub = filter(data_filtered_scrub, index, val_s)
        # if the filter return only 1 save it and end the loop
        if len(data_filtered_oxygen) == 1:
            oxygen_rate=data_filtered_oxygen[0]
            completed +=1
        if len(data_filtered_scrub) == 1:
            scrub_rate=data_filtered_scrub[0]
            completed +=1
        if completed == 2:
            break
    return int("".join(oxygen_rate),2) *  int("".join(scrub_rate),2) 

def filter(data, index, val):
    filtered = []
    for r in data:
        if r[index] == val:
            filtered.append(r)
    return filtered

def most_common(data, index):
    counter = 0
    for r in data:
        if r[index] == "1":
            counter +=1
    if counter >= len(data) - counter:
        return "1"
    return  "0"

def less_common(data, index):
    counter = 0
    for r in data:
        if r[index] == "1":
            counter +=1
    if counter >= len(data) - counter:
        return "0"
    return "1"


if __name__ == "__main__":
    data= parse_file("input.txt")
    print( "Part one:", part_one(data))
    print( "Part two:", part_two(data))
