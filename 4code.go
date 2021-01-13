package main

import "fmt"

/*
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。

进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？
示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
示例 3：

输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000
示例 4：

输入：nums1 = [], nums2 = [1]
输出：1.00000
示例 5：

输入：nums1 = [2], nums2 = []
输出：2.00000


提示：

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106
*/

func main() {
	a := []int{1}
	b := []int{2, 3, 4, 5, 6}
	fmt.Println(findMedianSortedArrays(a, b))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	left, right := (m+n+1)/2, (m+n+2)/2
	a1 := process(nums1, 0, nums2, 0, left)
	fmt.Println(a1)
	a2 := process(nums1, 0, nums2, 0, right)
	fmt.Println(a2)
	return (float64(a1) + float64(a2)) / 2
}

const (
	INT_MAX = int(^uint((0)) >> 1)
)

func process(nums1 []int, i int, nums2 []int, j int, k int) int {

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	if i >= len(nums1) {
		return nums2[j+k-1]
	}
	if j >= len(nums2) {
		return nums1[i+k-1]
	}
	if k == 1 {
		return min(nums1[i], nums2[j])
	}
	midVal1, midVal2 := 0, 0
	if len(nums1) > (i + k/2 - 1) {
		midVal1 = nums1[i+k/2-1]
	} else {
		midVal1 = INT_MAX
	}
	if len(nums2) > (j + k/2 - 1) {
		midVal2 = nums2[j+k/2-1]
	} else {
		midVal2 = INT_MAX
	}
	if midVal1 < midVal2 {
		return process(nums1, i+k/2, nums2, j, k-k/2)
	} else {
		return process(nums1, i, nums2, j+k/2, k-k/2)
	}
}

/*
class Solution {
  public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int m = nums1.length;
        int n = nums2.length;
        int left = (m + n + 1) / 2;
        int right = (m + n + 2) / 2;
        return (findKth(nums1, 0, nums2, 0, left) + findKth(nums1, 0, nums2, 0, right)) / 2.0;
    }
    //i: nums1的起始位置 j: nums2的起始位置
    public int findKth(int[] nums1, int i, int[] nums2, int j, int k){
        if( i >= nums1.length) return nums2[j + k - 1];//nums1为空数组
        if( j >= nums2.length) return nums1[i + k - 1];//nums2为空数组
        if(k == 1){
            return Math.min(nums1[i], nums2[j]);
        }
        int midVal1 = (i + k / 2 - 1 < nums1.length) ? nums1[i + k / 2 - 1] : Integer.MAX_VALUE;
        int midVal2 = (j + k / 2 - 1 < nums2.length) ? nums2[j + k / 2 - 1] : Integer.MAX_VALUE;
        if(midVal1 < midVal2){
            return findKth(nums1, i + k / 2, nums2, j , k - k / 2);
        }else{
            return findKth(nums1, i, nums2, j + k / 2 , k - k / 2);
        }
    }
}
*/
