pub struct Solution {}

impl Solution {
    pub fn three_sum(_nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut nums = _nums;
        let mut ans = Vec::new();
        if nums.len() < 3 {
            return ans;
        }

        nums.sort();

        let mut i = 0;

        while i < nums.len() - 2 {
            let mut j = i + 1;
            let mut k = nums.len() - 1;
            while j < nums.len() - 1 && j < k {
                while j < k && nums[i] + nums[j] + nums[k] > 0 {
                    k -= 1;
                }
                if j != k && nums[i] + nums[j] + nums[k] == 0 {
                    let mut flag = false;
                    for t in ans.iter().rev() {
                        if t[0] == nums[i] {
                            if t[1] == nums[j] && t[2] == nums[k] {
                                flag = true;
                                break;
                            }
                        } else {
                            break;
                        }
                    }
                    if !flag {
                        //println!("i:{:?}, j:{:?}, k:{:?}", i, j, k);
                        ans.push(vec![nums[i], nums[j], nums[k]]);
                    }
                }
                j += 1;
            }
            i += 1;
        }

        return ans;
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }

    #[test]
    fn case1() {
        let nums = vec![-1, 0, 1, 2, -1, -4];
        let output = Solution::three_sum(nums);
        println!("{:?}", output);
        assert_eq!(output.len(), 2);
    }

    #[test]
    fn case2() {
        let nums = vec![1, 2, -2, -1];
        let output = Solution::three_sum(nums);
        println!("{:?}", output);
        assert_eq!(output.len(), 0);
    }

    #[test]
    fn case3() {
        let nums = vec![-2, 0, 0, 2, 2];
        let output = Solution::three_sum(nums);
        println!("{:?}", output);
        assert_eq!(output.len(), 1);
    }
}
