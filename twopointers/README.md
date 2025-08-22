# Two Pointers

## [Valid Palindrome](https://neetcode.io/solutions/valid-palindrome)
Given a string `s`, return true if it is a **palindrome**, otherwise return `false`.

A **palindrome** is a string that reads the same forward and backward. It is also case-insensitive and ignores all non-alphanumeric characters.

Note: Alphanumeric characters consist of letters `(A-Z, a-z)` and numbers `(0-9)`.

**Example 1**:
```text
Input: s = "Was it a car or a cat I saw?"

Output: true
```
Explanation: After considering only alphanumerical characters we have **"wasitacaroracatisaw"**, which is a palindrome.

**Example 2**:
```text
Input: s = "tab a cat"

Output: false
```
Explanation: **"tabacat"** is not a palindrome.

**Constraints**:
* 1 <= s.length <= 1000
* s is made up of only printable ASCII characters.

## [Two Sum II Input Array Is Sorted](https://neetcode.io/solutions/two-sum-ii-input-array-is-sorted)
Given an array of integers `numbers` that is sorted in non-decreasing order.

Return the indices **(1-indexed)** of two numbers, `[index1, index2]`, such that they add up to a given target number `target` and `index1 < index2`. Note that `index1` and `index2` cannot be equal, therefore you may not use the same element twice.

There will always be exactly one valid solution.

Your solution must use O(1) additional space.

**Example 1**:
```text
Input: numbers = [1,2,3,4], target = 3

Output: [1,2]
```

**Explanation**:
The sum of 1 and 2 is 3. Since we are assuming a 1-indexed array, `index1` = 1, `index2` = 2. We return `[1, 2]`.

**Constraints**:
* `2 <= numbers.length <= 1000`
* `-1000 <= numbers[i] <= 1000`
* `-1000 <= target <= 1000`

## [3Sum](https://neetcode.io/solutions/3sum)
Given an integer array `nums`, return all the triplets `[nums[i], nums[j], nums[k]]` where nums`[i] + nums[j] + nums[k] == 0`, and the indices `i`, `j` and `k` are all distinct.

The output should not contain any duplicate triplets. You may return the output and the triplets in any order.

**Example 1**:
```text
Input: nums = [-1,0,1,2,-1,-4]

Output: [[-1,-1,2],[-1,0,1]]
```
**Explanation**:
`nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0`.
`nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0`.
`nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0`.
The distinct triplets are `[-1,0,1]` and `[-1,-1,2]`.

**Example 2**:
```text
Input: nums = [0,1,1]

Output: []
```
**Explanation**: The only possible triplet does not sum up to 0.

**Example 3**:
```text
Input: nums = [0,0,0]

Output: [[0,0,0]]
```
**Explanation**: The only possible triplet sums up to 0.

**Constraints**:
* `3 <= nums.length <= 1000`
* `-10^5 <= nums[i] <= 10^5`

