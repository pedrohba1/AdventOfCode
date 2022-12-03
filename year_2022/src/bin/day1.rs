use std::{
    fs::File,
    io::{BufRead, BufReader},
    path::Path,
};

fn main() {
    let file = File::open("inputs/day1.txt").unwrap();
    let reader = BufReader::new(file);

    let mut maxCalories: i128 = 0;
    let mut currentCalories: i128 = 0;

    for line in reader.lines() {
        let x = line.unwrap();

        if x != "" {
            println!("current calorie {}", x);
            currentCalories += x.parse::<i128>().unwrap()
        } else {
            println!("no current calorie {}", x);
            if (currentCalories > maxCalories) {
                maxCalories = currentCalories
            };
            currentCalories = 0;
        }
    }
    println!("{}", maxCalories)
}
