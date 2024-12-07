use std::{char, fs, i32};

fn main() {
    let the_matrix = parse_input_to_matrix("input.txt");
    println!("part one:{}", part_one(&the_matrix));
    println!("part two:{}", part_two(&the_matrix));
}

fn parse_input_to_matrix(filename: &str) -> Vec<Vec<char>> {
    let content = fs::read_to_string(filename).expect("Impossibile leggere il file");
    content
        .lines()
        .map(|line| line.chars().collect::<Vec<char>>())
        .collect()
}

fn part_one(m: &Vec<Vec<char>>) -> i32 {
    const WORD: &str = "XMAS";
    let word_chars: Vec<char> = WORD.chars().collect();
    let directions = [
        (0, 1),   // Right
        (1, 0),   // Down
        (0, -1),  // Left
        (-1, 0),  // Up
        (1, 1),   // Down-right
        (-1, -1), // Up-left
        (1, -1),  // Down-left
        (-1, 1),  // Up-right
    ];

    let mut counter = 0;

    for row_index in 0..m.len() {
        for col_index in 0..m[row_index].len() {
            for &(dx, dy) in &directions {
                if check_word_in_direction(
                    m,
                    &word_chars,
                    row_index as i32,
                    col_index as i32,
                    dx,
                    dy,
                ) {
                    counter += 1;
                }
            }
        }
    }
    counter
}

fn part_two(m: &Vec<Vec<char>>) -> i32 {
    const WORD: &str = "MAS";
    let word_chars: Vec<char> = WORD.chars().collect();
    let directions_left = [
        (1, -1), // Down-left ↙
        (-1, 1), // Up-right ↗
    ];

    let directions_right = [
        (1, 1),   // Down-right ↘
        (-1, -1), // Up-left ↖
    ];
    let mut counter = 0;
    for row_index in 0..m.len() {
        for col_index in 0..m[row_index].len() {
            for &(dx_l, dy_l) in &directions_left {
                if check_word_in_direction(
                    m,
                    &word_chars,
                    row_index as i32,
                    col_index as i32,
                    dx_l,
                    dy_l,
                ) {
                    for &(dx_r, dy_r) in &directions_right {
                        let (shift_x, shift_y) = match (dx_l, dx_r) {
                            (-1, 1) => (
                                col_index as i32,
                                row_index as i32 - (word_chars.len() as i32 - 1),
                            ), // ↗ ↘
                            (-1, -1) => (
                                col_index as i32 + (word_chars.len() as i32 - 1),
                                row_index as i32,
                            ), // ↗ ↖
                            (1, 1) => (
                                col_index as i32 - (word_chars.len() as i32 - 1),
                                row_index as i32,
                            ), // ↙ ↘
                            (1, -1) => (
                                col_index as i32,
                                row_index as i32 + (word_chars.len() as i32 - 1),
                            ), // ↙ ↖
                            _ => unreachable!(),
                        };
                        if check_word_in_direction(m, &word_chars, shift_y, shift_x, dx_r, dy_r) {
                            counter += 1;
                        }
                    }
                }
            }
        }
    }
    counter
}

fn check_word_in_direction(
    m: &Vec<Vec<char>>,
    word: &[char],
    start_row: i32,
    start_col: i32,
    dx: i32,
    dy: i32,
) -> bool {
    for (i, &ch) in word.iter().enumerate() {
        let row = start_row + i as i32 * dx;
        let col = start_col + i as i32 * dy;
        if row < 0 || col < 0 || row as usize >= m.len() || col as usize >= m[i].len() {
            return false;
        }
        if m[row as usize][col as usize] != ch {
            return false;
        }
    }
    true
}
