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
    let mut safe_reports = 0;
    for report in m.iter() {
        let mut is_stable = false;
        let mut is_decreasing = false;
        let mut is_increasing = false;
        for index in 0..report.len() - 1 {
            let next_level = report[index + 1];
            match check_increment(report[index], next_level) {
                1 => is_increasing = true,
                -1 => is_decreasing = true,
                _ => is_stable = true,
            }
        }
        if !is_stable && !(is_increasing && is_decreasing) {
            safe_reports += 1;
        }
    }
    safe_reports
}

fn part_two(m: &Vec<Vec<i32>>) -> i32 {
    let mut safe_reports = 0;
    for mut report in m.iter().cloned() {
        let mut problem_dumpener = 0;
        let mut is_safe = false;
        while problem_dumpener < 2 && !is_safe {
            let mut is_decreasing = false;
            let mut is_increasing = false;
            let mut is_stable = false;
            for index in 0..report.len() - 1 {
                let next_level = report[index + 1];
                match check_increment(report[index], next_level) {
                    1 => is_increasing = true,
                    -1 => is_decreasing = true,
                    _ => is_stable = true,
                }
                if is_stable || (is_increasing && is_decreasing) {
                    problem_dumpener += 1;
                    report.remove(index + 1);
                    break;
                }
            }
            if !is_stable && !(is_increasing && is_decreasing) {
                is_safe = true;
                println!("{:?}", report)
            }
        }
        if is_safe {
            safe_reports += 1;
        }
    }
    safe_reports
}

fn check_increment(a: i32, b: i32) -> i32 {
    let diff = a - b;
    if diff == 0 || diff.abs() > 3 {
        return 0;
    }
    (diff >> 31) | (diff != 0) as i32
}
