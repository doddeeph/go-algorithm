# Arrays & Hashing

## [Contains Duplicate](https://neetcode.io/solutions/contains-duplicate)
Given an integer array `nums`, return `true` if any value appears more than once in the array, otherwise return `false`.

Example 1:
```text
Input: nums = [1, 2, 3, 3]

Output: true
```

Example 2:
```text
Input: nums = [1, 2, 3, 4]

Output: false
```

## [Valid Anagram](https://neetcode.io/solutions/valid-anagram)
Given two strings `s` and `t`, return `true` if the two strings are anagrams of each other, otherwise return `false`.

An **anagram** is a string that contains the exact same characters as another string, but the order of the characters can be different.

Example 1:
```text
Input: s = "racecar", t = "carrace"

Output: true
```

Example 2:
```text
Input: s = "jar", t = "jam"

Output: false
```

**Constraints**:

`s` and `t` consist of lowercase English letters.

## [Two Sum](https://neetcode.io/solutions/two-sum)
Given an array of integers `nums` and an integer `target`, return the indices `i` and `j` such that `nums[i] + nums[j] == target` and `i != j`.

You may assume that every input has exactly one pair of indices `i` and `j` that satisfy the condition.

Return the answer with the smaller index first.

**Example 1**:
```text
Input: 
nums = [3,4,5,6], target = 7

Output: [0,1]
```
Explanation: `nums[0] + nums[1] == 7`, so we return `[0, 1]`.

**Example 2**:
```text
Input: nums = [4,5,6], target = 10

Output: [0,2]
```

**Example 3**:
```text
Input: nums = [5,5], target = 10

Output: [0,1]
```

**Constraints**:
* `2 <= nums.length <= 1000`
* `-10,000,000 <= nums[i] <= 10,000,000`
* `-10,000,000 <= target <= 10,000,000`

## [Group Anagrams](https://neetcode.io/solutions/group-anagrams)
Given an array of strings `strs`, group all anagrams together into sublists. You may return the output in **any order**.

An **anagram** is a string that contains the exact same characters as another string, but the order of the characters can be different.

**Example 1**:
```text
Input: strs = ["act","pots","tops","cat","stop","hat"]

Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]
```

**Example 2**:
```text
Input: strs = ["x"]

Output: [["x"]]
```

**Example 3**:
```text
Input: strs = [""]

Output: [[""]]
```

**Constraints**:
* `1 <= strs.length <= 1000.`
* `0 <= strs[i].length <= 100`
* `strs[i] is made up of lowercase English letters.`

## [Top K Frequent Elements](https://neetcode.io/solutions/top-k-frequent-elements)
Given an integer array `nums` and an integer `k`, return the `k` most frequent elements within the array.

The test cases are generated such that the answer is always **unique**.

You may return the output in **any order**.

**Example 1**:
```text
Input: nums = [1,2,2,3,3,3], k = 2

Output: [2,3]
```

**Example 2**:
```text
Input: nums = [7,7], k = 1

Output: [7]
```

**Constraints**:
* `1 <= nums.length <= 10^4.`
* `-1000 <= nums[i] <= 1000`
* `1 <= k <= number of distinct elements in nums.`

## [Product of Array Except Self](https://neetcode.io/solutions/product-of-array-except-self)
Given an integer array `nums`, return an array `output` where `output[i]` is the product of all the elements of `nums` except `nums[i]`.

Each product is **guaranteed** to fit in a **32-bit** integer.

Follow-up: Could you solve it in O(n) time without using the division operation?

**Example 1**:
```text
Input: nums = [1,2,4,6]

Output: [48,24,12,8]
```

**Example 2**:
```text
Input: nums = [-1,0,1,2,3]

Output: [0,-6,0,0,0]
```

**Constraints**:
* `2 <= nums.length <= 1000`
* `-20 <= nums[i] <= 20`

## [Valid Sudoku](https://neetcode.io/solutions/valid-sudoku)


## [Longest Consecutive Sequence]()
Given an array of integers `nums`, return the length of the longest consecutive sequence of elements that can be formed.

A consecutive sequence is a sequence of elements in which each element is exactly `1` greater than the previous element. The elements do not have to be consecutive in the original array.

You must write an algorithm that runs in `O(n)` time.

**Example 1**:
```test
Input: nums = [2,20,4,10,3,4,5]

Output: 4
```
Explanation: The longest consecutive sequence is `[2, 3, 4, 5]`.

**Example 2**:
```text
Input: nums = [0,3,2,5,4,6,1,1]

Output: 7
```

**Constraints**:
* `0 <= nums.length <= 1000`
* `-10^9 <= nums[i] <= 10^9`