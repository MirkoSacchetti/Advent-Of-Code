use std::{fs, usize};

fn main() {
    let map = parse_input_to_matrix("input.txt");
    println!("Part One: {}", part_one(&map));
    println!("Part Two: {}", part_two(&map));
}

fn parse_input_to_matrix(filename: &str) -> Vec<Vec<char>> {
    let content = fs::read_to_string(filename).expect("file not found");
    content.lines().map(|line| line.chars().collect()).collect()
}

fn part_one(map: &[Vec<char>]) -> i32 {
    let mut guardian = Guardian::new(map).unwrap();
    guardian.simulate()
}

// TODO too slow!!! > 18secs
fn part_two(map: &[Vec<char>]) -> i32 {
    let mut g = Guardian::new(&map).unwrap();
    g.simulate();

    // Raccogli le posizioni con 'X'
    let positions_with_x: Vec<(usize, usize)> = g
        .map
        .iter()
        .enumerate()
        .flat_map(|(y, row)| {
            row.iter().enumerate().filter_map(
                move |(x, &cell)| {
                    if cell == 'X' {
                        Some((y, x))
                    } else {
                        None
                    }
                },
            )
        })
        .collect();

    let mut counter_loop = 0;

    // Processa ogni posizione con 'X'
    for (y, x) in positions_with_x {
        let mut modified_map = g.map.clone(); // Clona solo una volta
        modified_map[y][x] = '#'; // Modifica temporaneamente la cella
        let mut guardian = Guardian::new(&modified_map).unwrap();
        if guardian.has_infinite_loop() {
            counter_loop += 1;
        }
    }

    counter_loop
}

struct Position {
    x: usize,
    y: usize,
    direction: usize,
}

struct Guardian {
    position: Position,
    directions: [(isize, isize); 4],
    map: Vec<Vec<char>>,
}

impl Guardian {
    fn new(map: &[Vec<char>]) -> Option<Self> {
        for (y, row) in map.iter().enumerate() {
            for (x, &cell) in row.iter().enumerate() {
                if cell == '^' {
                    return Some(Self {
                        position: Position { x, y, direction: 0 },
                        directions: [
                            (0, -1), // Up
                            (1, 0),  // Right
                            (0, 1),  // Down
                            (-1, 0), // Left
                        ],
                        map: map.to_vec(),
                    });
                }
            }
        }
        None
    }

    fn step(&mut self) -> bool {
        let (dx, dy) = self.directions[self.position.direction];
        let new_x = self.position.x as isize + dx;
        let new_y = self.position.y as isize + dy;
        // check boundaries
        if new_y < 0
            || new_x < 0
            || new_y as usize >= self.map.len()
            || new_x as usize >= self.map[0].len()
        {
            return false;
        }
        // change direction is there are an obstacle
        if self.map[new_y as usize][new_x as usize] == '#' {
            self.position.direction = (self.position.direction + 1) % self.directions.len();
            return true;
        }
        // set new position and mark it
        self.position.x = new_x as usize;
        self.position.y = new_y as usize;
        if self.map[self.position.y][self.position.x] != 'X' {
            self.map[self.position.y][self.position.x] = 'X';
        }
        true
    }

    fn simulate(&mut self) -> i32 {
        while self.step() {}
        self.map
            .iter()
            .flatten()
            .filter(|&&cell| cell == '^' || cell == 'X')
            .count() as i32
    }

    fn has_infinite_loop(&mut self) -> bool {
        let mut visited_positions = std::collections::HashSet::new();
        while self.step() {
            let current_state = (self.position.x, self.position.y, self.position.direction);
            if visited_positions.contains(&current_state) {
                return true;
            }
            visited_positions.insert(current_state);
        }
        false
    }
}
