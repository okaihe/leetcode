fn main() {
    let bottom1: String = String::from("BCD");
    let allowed1: Vec<String> = vec![
        String::from("BCC"),
        String::from("CDE"),
        String::from("CEA"),
        String::from("FFF"),
    ];

    let bottom2: String = String::from("AAAA");
    let allowed2: Vec<String> = vec![
        String::from("AAB"),
        String::from("AAC"),
        String::from("BCD"),
        String::from("BBE"),
        String::from("DEF"),
    ];

    println!("Example 1");
    println!("{}", Solution::pyramid_transition(bottom1, allowed1));
    println!();
    println!("Example 2");
    println!("{}", Solution::pyramid_transition(bottom2, allowed2));
}

struct Solution;
impl Solution {
    pub fn pyramid_transition(bottom: String, allowed: Vec<String>) -> bool {
        return true;
    }
}
