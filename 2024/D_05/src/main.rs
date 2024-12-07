use std::fs;

struct PrintQueue {
    rules: Vec<(i32, i32)>,
    updates: Vec<Vec<i32>>,
}

fn main() {
    let q = parse_input("input.txt");
    println!("part one:{}", part_one(&q));
    println!("part two:{}", part_two(&q));
}

fn parse_input(filename: &str) -> PrintQueue {
    let content = fs::read_to_string(filename).expect("Impossibile leggere il file");
    let mut updates = Vec::new();
    let mut rules = Vec::new();
    for line in &mut content.lines() {
        if line.trim().is_empty() {
            continue;
        }
        let parts: Vec<&str> = line.split('|').collect();
        if parts.len() == 2 {
            let a = parts[0].trim().parse::<i32>().expect("Numero non valido");
            let b = parts[1].trim().parse::<i32>().expect("Numero non valido");
            rules.push((a, b));
        } else {
            let rule: Vec<i32> = line
                .split(',')
                .map(|s| s.trim().parse::<i32>().expect("Numero non valido"))
                .collect();
            updates.push(rule);
        }
    }
    PrintQueue { updates, rules }
}

fn part_one(pq: &PrintQueue) -> i32 {
    let mut correctly_ordered = 0;
    for update in &pq.updates {
        if validate(&pq.rules, update) {
            correctly_ordered += update[update.len() / 2]
        }
    }
    correctly_ordered
}

fn part_two(pq: &PrintQueue) -> i32 {
    let mut correctly_ordered = 0;
    for update in &pq.updates {
        if !validate(&pq.rules, update) {
            let fixed = fixing(&pq.rules, update.clone());
            correctly_ordered += fixed[fixed.len() / 2]
        }
    }
    correctly_ordered
}

fn validate(rules: &[(i32, i32)], update: &[i32]) -> bool {
    for &(a, b) in rules {
        let pos_a = update.iter().position(|&x| x == a);
        let pos_b = update.iter().position(|&x| x == b);
        if let (Some(pos_a), Some(pos_b)) = (pos_a, pos_b) {
            if pos_a > pos_b {
                return false;
            }
        }
    }
    true
}

fn fixing(rules: &[(i32, i32)], mut update: Vec<i32>) -> Vec<i32> {
    while !validate(rules, &update) {
        for &(a, b) in rules {
            let pos_a = update.iter().position(|&x| x == a);
            let pos_b = update.iter().position(|&x| x == b);
            if let (Some(pos_a), Some(pos_b)) = (pos_a, pos_b) {
                if pos_a > pos_b {
                    update.swap(pos_a, pos_b);
                }
            }
        }
    }
    update.to_vec()
}
