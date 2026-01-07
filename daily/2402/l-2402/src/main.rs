fn main() {
    println!("Hello, world!");

    let n1: i32 = 2;
    let meetings1: Vec<Vec<i32>> = vec![vec![0, 10], vec![1, 5], vec![2, 7], vec![3, 4]];

    let result = Solution::most_booked(n1, meetings1);
    println!("{result}");
}

struct Solution;
impl Solution {
    pub fn most_booked(n: i32, meetings: Vec<Vec<i32>>) -> i32 {
        69
    }
}
