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
    m.iter().filter(|report| is_safe(report)).count() as i32
}

fn part_two(m: &Vec<Vec<i32>>) -> i32 {
    let mut safe_reports = 0;
    for report in m {
        if is_safe(report) {
            safe_reports += 1;
            continue;
        }

        let mut is_safe_with_dampener = false;
        for i in 0..report.len() {
            let mut modified_report = report.clone();
            modified_report.remove(i);
            if is_safe(&modified_report) {
                is_safe_with_dampener = true;
                break;
            }
        }
        if is_safe_with_dampener {
            safe_reports += 1
        }
    }
    safe_reports
}

fn is_safe(report: &[i32]) -> bool {
    if report.len() < 2 {
        return false;
    }

    let mut is_increasing = true;
    let mut is_decreasing = true;

    for window in report.windows(2) {
        match check_increment(window[0], window[1]) {
            1 => is_decreasing = false,
            -1 => is_increasing = false,
            _ => return false, // Non valido
        }
    }

    is_increasing || is_decreasing
}

fn check_increment(a: i32, b: i32) -> i32 {
    let diff = a - b;
    if diff == 0 || diff.abs() > 3 {
        return 0;
    }
    (diff >> 31) | (diff != 0) as i32
}
