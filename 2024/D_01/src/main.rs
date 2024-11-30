use std::fs;

fn parse_input_to_matrix(filename: &str) -> Vec<Vec<char>> {
    let content = fs::read_to_string(filename).expect("Failed to read the input file");
    content.lines().map(|line| line.chars().collect()).collect()
}

fn main() {
    let matrix = parse_input_to_matrix("input.txt");
    for row in &matrix {
        println!("{:?}", row);
    }
}
