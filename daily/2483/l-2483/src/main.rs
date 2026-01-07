fn main() {
    let mut customers1 = "YYNY";

    let s2 = &customers1;

    println!("{}", best_closing_time(customers1));

    let mut s1 = String::from("Hallo was geht ab!");

    println!("{}", s2);

    s1.push_str("Noch ein String!");
}

fn best_closing_time(customers: &str) -> i32 {
    let mut max_score = 0;
    let mut best_hour = 0;
    let mut score = 0;

    for (i, &byte) in customers.as_bytes().iter().enumerate() {
        score += if byte == b'Y' { 1 } else { -1 };
        if score > max_score {
            max_score = score;
            best_hour = (i + 1) as i32;
        }
    }

    best_hour
}

// Irgendein verrückter Rust-
// fn best_closing_time(customers: &str) -> i32 {
//     customers
//         .as_bytes()
//         .iter()
//         .map(|&b| if b == b'Y' { 1 } else { -1 })
//         .scan(0, |score, delta| {
//             *score += delta;
//             Some(*score)
//         })
//         .enumerate()
//         .chain(std::iter::once((usize::MAX, 0)))
//         .max_by_key(|&(i, score)| (score, -(i as i32 + 1)))
//         .map(|(i, _)| if i == usize::MAX { 0 } else { (i + 1) as i32 })
//         .unwrap_or(0)
// }
