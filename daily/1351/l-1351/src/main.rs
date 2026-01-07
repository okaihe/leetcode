fn main() {
    let grid1 = vec![
        vec![4, 3, 2, -1],
        vec![3, 2, 1, -1],
        vec![1, 1, -1, -2],
        vec![-1, -1, -2, -3],
    ];
    println!("{}", Solution::count_negatives(grid1));

    println!("Jetzt habe ich etwas geändert!");

    println!("----------------");

    let mut int_8: i8 = 0;
    for _ in 0..8 {
        int_8 = int_8 << 1;
        int_8 |= 0b0000_0001;
        println!("{int_8}");
        println!("{:08b}", int_8);
        println!();
    }

    int_8 += 1;
    println!("{int_8}");
    println!("{:08b}", int_8);
}

struct Solution;
impl Solution {
    pub fn count_negatives(grid: Vec<Vec<i32>>) -> i32 {
        grid.iter()
            .flat_map(|row| row.iter())
            .filter(|&&x| x < 0)
            .count() as i32
    }

    // pub fn count_negatives(grid: Vec<Vec<i32>>) -> i32 {
    //     let mut result: i32 = 0;
    //     for vector in grid.iter() {
    //         for row in vector.iter() {
    //             result += if *row < 0 { 1 } else { 0 };
    //         }
    //     }
    //     result
    // }
}
