pub struct Solution {}

use std::collections::HashMap;

impl Solution {
    pub fn letter_combinations(digits: String) -> Vec<String> {
        let mut ans = Vec::new();
        let mut map: HashMap<char, Vec<String>> = HashMap::new();
        map.insert(
            '2',
            vec![String::from("a"), String::from("b"), String::from("c")],
        );
        map.insert(
            '3',
            vec![String::from("d"), String::from("e"), String::from("f")],
        );
        map.insert(
            '4',
            vec![String::from("g"), String::from("h"), String::from("i")],
        );
        map.insert(
            '5',
            vec![String::from("j"), String::from("k"), String::from("l")],
        );
        map.insert(
            '6',
            vec![String::from("m"), String::from("n"), String::from("o")],
        );
        map.insert(
            '7',
            vec![
                String::from("p"),
                String::from("q"),
                String::from("r"),
                String::from("s"),
            ],
        );
        map.insert(
            '8',
            vec![String::from("t"), String::from("u"), String::from("v")],
        );
        map.insert(
            '9',
            vec![
                String::from("w"),
                String::from("x"),
                String::from("y"),
                String::from("z"),
            ],
        );

        if digits.len() > 0 {
            ans.push(String::from(""));
            ans = Solution::recur_letter_combinations(digits, ans, &map);
        }
        return ans;
    }

    pub fn recur_letter_combinations(
        digits: String,
        ans: Vec<String>,
        map: &HashMap<char, Vec<String>>,
    ) -> Vec<String> {
        if digits.len() == 0 {
            return ans;
        }
        let mut ret: Vec<String> = Vec::new();
        let strs = &map[&digits.chars().next().unwrap()];
        //println!("{:?}", strs);
        for i in ans.iter() {
            for j in strs.iter() {
                ret.push(format!("{}{}", i, j));
            }
        }
        ret = Solution::recur_letter_combinations(digits[1..].to_string(), ret, &map);
        return ret;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    struct Case {
        i: &'static str,
        o: Vec<&'static str>,
    }

    #[test]
    fn cases() {
        let cases = [
            Case {
                i: "23",
                o: vec!["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"],
            },
            Case {
                i: "39",
                o: vec![
                    "dw", "dx", "dy", "dz", "ew", "ex", "ey", "ez", "fw", "fx", "fy", "fz",
                ],
            },
            Case { i: "", o: vec![] },
        ];

        for c in cases.iter() {
            let output = Solution::letter_combinations(String::from(c.i));
            println!("{:?}", output);
            assert_eq!(output, c.o);
        }
    }
}
