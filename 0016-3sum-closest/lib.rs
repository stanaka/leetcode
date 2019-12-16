pub struct Solution {}

impl Solution {
    pub fn three_sum_closest(_nums: Vec<i32>, target: i32) -> i32 {
        let mut nums = _nums;
        let mut ans = 0;
        let mut ans_abs = std::i32::MAX;
        if nums.len() < 3 {
            return ans;
        }
        nums.sort();

        for i in 0..(nums.len() - 2) {
            let mut j = i + 1;
            let mut k = nums.len() - 1;
            while j < k {
                let score = nums[i] + nums[j] + nums[k];
                let score_abs = if score < target {
                    &target - score
                } else {
                    score - &target
                };
                if ans_abs > score_abs {
                    ans_abs = score_abs;
                    ans = score
                }
                //println!("i:{:?}, j:{:?}, k:{:?}, score_abs:{:?}", i, j, k, score_abs);

                if score < target {
                    j += 1;
                } else {
                    k -= 1;
                }
            }
        }

        return ans;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    struct Case {
        i1: Vec<i32>,
        i2: i32,
        o: i32,
    }

    #[test]
    fn cases() {
        let cases = [Case {
            i1: vec![-1, 2, 1, -4],
            i2: 1,
            o: 2,
        }];

        for c in cases.iter() {
            let output = Solution::three_sum_closest((&c.i1).to_vec(), c.i2);
            println!("{:?}", output);
            assert_eq!(output, c.o);
        }
    }
}
