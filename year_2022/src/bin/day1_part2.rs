use std::{
    fs::File,
    io::{BufRead, BufReader},
    path::Path,
    vec,
};

fn main() {
    let file = File::open("inputs/day1.txt").unwrap();
    let reader = BufReader::new(file);

    let mut maxCalories: i128 = 0;
    let mut currentCalories: i128 = 0;
    let mut v = vec![];

    for line in reader.lines() {
        let x = line.unwrap();

        if x != "" {
            currentCalories += x.parse::<i128>().unwrap()
        } else {
            v.push(currentCalories);
            currentCalories = 0;
        }
    }
    v.sort();
    let topCal: i128 = v.iter().rev().take(3).sum();
    print!("{}", topCal);
}
