use std::{fs, i32};

fn main() {
    let lines = parse_input("input.txt");
    println!("part one:{}", part_one(&lines));
    //    println!("part two:{}", part_two(&the_matrix));
}

fn parse_input(filename: &str) -> Vec<(i32, Vec<i32>)> {
    fs::read_to_string(filename)
        .expect("Impossibile leggere il file")
        .lines()
        .filter_map(|l| {
            let mut parts = l.split(':');
            let result = parts.next()?.trim().parse::<i32>().ok()?;
            let numbers = parts
                .next()?
                .trim()
                .split_whitespace()
                .filter_map(|n| n.parse::<i32>().ok())
                .collect::<Vec<i32>>();
            Some((result, numbers))
        })
        .collect()
}

fn part_one(equations: &Vec<(i32, Vec<i32>)>) -> i32 {
    let mut counter = 0;
    for e in equations {
        for ele in e.1.to_vec() {
            if e.0 == ele {
                counter += 1;
            }
        }
    }
    return counter;
}
