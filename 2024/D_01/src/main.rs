use std::fs;

fn main() {
    let the_matrix = parse_input_to_matrix("input.txt");
    println!("part one:{}", part_one(&the_matrix));
    println!("part two:{}", part_two(&the_matrix));
}

fn parse_input_to_matrix(filename: &str) -> Vec<Vec<i32>> {
    let content = fs::read_to_string(filename).expect("Impossibile leggere il file");

    content
        .lines()
        .map(|line| {
            line.split_whitespace()
                .map(|num| {
                    num.parse::<i32>()
                        .expect("Impossibile convertire in numero")
                })
                .collect::<Vec<i32>>()
        })
        .collect()
}

fn part_one(m: &Vec<Vec<i32>>) -> i32 {
    let mut left_list: Vec<i32> = m.iter().map(|pair| pair[0]).collect();
    let mut right_list: Vec<i32> = m.iter().map(|pair| pair[1]).collect();
    left_list.sort();
    right_list.sort();

    let mut total_distance = 0;
    for (i, _value) in left_list.iter().enumerate() {
        total_distance += (left_list[i] - right_list[i]).abs();
    }

    total_distance
}

fn part_two(m: &Vec<Vec<i32>>) -> i32 {
    let mut left_list: Vec<i32> = m.iter().map(|pair| pair[0]).collect();
    let mut right_list: Vec<i32> = m.iter().map(|pair| pair[1]).collect();
    left_list.sort();
    right_list.sort();

    let mut similarity_score = 0;
    for value in left_list.iter() {
        let occurrences = count_occurrences(*value, &right_list);
        similarity_score += value * occurrences;
    }
    similarity_score
}

fn count_occurrences(value: i32, list: &Vec<i32>) -> i32 {
    let mut counter = 0;
    for v in list.iter() {
        if *v == value {
            counter += 1;
        }
    }
    counter
}
