use regex::Regex;
use std::fs;

fn main() {
    let muls = parse_input("input.txt");
    println!("part one:{}", part_one(&muls));
    println!("part two:{}", part_two(&muls));
}

fn parse_input(filename: &str) -> String {
    fs::read_to_string(filename).expect("Impossibile leggere il file")
}

fn part_one(memory: &str) -> i32 {
    let pattern = r"mul\((\d{1,3}),(\d{1,3})\)";
    let re = Regex::new(pattern).unwrap();
    re.captures_iter(memory)
        .map(|cap| {
            let x: i32 = cap[1].parse().unwrap();
            let y: i32 = cap[2].parse().unwrap();
            x * y
        })
        .sum()
}

fn part_two(memory: &str) -> i32 {
    let mul_re = Regex::new(r"mul\((\d{1,3}),(\d{1,3})\)").unwrap();
    let do_re = Regex::new(r"do\(\)").unwrap();
    let dont_re = Regex::new(r"don't\(\)").unwrap();

    let mut sum = 0;
    let mut pos = 0;
    let mut enabled = true;

    while pos < memory.len() {
        if let Some(m) = do_re.find(&memory[pos..]) {
            if m.start() == 0 {
                enabled = true;
                pos += m.end();
                continue;
            }
        }
        if let Some(m) = dont_re.find(&memory[pos..]) {
            if m.start() == 0 {
                enabled = false;
                pos += m.end();
                continue;
            }
        }
        if enabled {
            if let Some(cap) = mul_re.captures(&memory[pos..]) {
                if cap.get(0).unwrap().start() == 0 {
                    let n1: i32 = cap[1].parse().unwrap();
                    let n2: i32 = cap[2].parse().unwrap();
                    sum += n1 * n2;
                    pos += cap.get(0).unwrap().end();
                    continue;
                }
            }
        }
        pos += 1;
    }
    sum
}
