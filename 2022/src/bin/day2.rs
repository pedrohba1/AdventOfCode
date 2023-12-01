use std::{
    collections::HashMap,
    fs::File,
    io::{BufRead, BufReader},
};

fn main() {
    let file = File::open("inputs/day2.txt").unwrap();
    let reader = BufReader::new(file);

    // scores.insert(String::from("A"), 10);
    // scores.insert(String::from("Paper"), 50);

    for line in reader.lines() {
        let line = line.unwrap();
        let splitted = line.split_whitespace();
        for play in splitted {
            print!("{} ", play);
        }
        println!("");
    }
}
