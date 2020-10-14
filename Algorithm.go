package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

// 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4
type ListNode struct {
	Val  int
	Next *ListNode
}

// func main() {
// 	num1 := []int{1, 2, 3, 0, 0, 0}
// 	// num2 := nil
// 	// var num2 []int

// 	num2 := []int{2, 5, 6}
// 	merge(num1, num2)
// }

// 因为这俩个是有序的，那么就只需要判断当前的首个值和另一个的比较即可
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var L3 *ListNode
	if l1.Val > l2.Val {
		L3 = l2
		L3.Next = mergeTwoLists(l1, l2.Next)
	} else {
		L3 = l1
		L3.Next = mergeTwoLists(l1.Next, l2)
	}

	return L3

}

// 2给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		return right
	} else if right == nil {
		return left
	} else {
		return root
	}
}

// 3给定一个二叉树，判断其是否是一个有效的二叉搜索树。

// 一个二叉搜索树具有如下特征：

// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树

func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}
func isValid(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min {
		return false
	}
	if root.Val >= max {
		return false
	}
	return isValid(root.Left, min, root.Val) && isValid(root.Right, root.Val, max)
}

// 4给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？
// 找出所有满足条件且不重复的三元组
// 例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

// 满足要求的三元组集合为：
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]
func threeSum(nums []int) [][]int {
	result := [][]int{}
	mmap := make(map[int]int)
	for x := 0; x < len(nums)-2; x++ {
		for y := x + 1; y < len(nums)-1; y++ {
			for z := y + 1; z < len(nums); z++ {
				if nums[x]+nums[y]+nums[z] == 0 {
					// 这样可以是可以但是不能去重（元素去重），这里的去重指的是 元素去重 并不是 索引去重
					// 可以引用计算，将之前的值存取进来，然后然后进行判断，那么这样就时间复杂度有很高了
					// captial, ok := countryCapitalMap [ "美国" ]
					// if nums[x]

					_, ok1 := mmap[nums[x]]
					_, ok2 := mmap[nums[y]]
					_, ok3 := mmap[nums[z]]
					if !(ok1 && ok2 && ok3) { //这样就可以进行通过了与预期结果一致
						arr := []int{nums[x], nums[y], nums[z]}
						result = append(result, arr)
						mmap[nums[x]] = 0
						mmap[nums[y]] = 0
						mmap[nums[z]] = 0
					}

				}

			}
		}
	}
	return result

}

//这是最优解但是实现比价麻烦，但是时间复杂度比较低，利用了和俩数之和一样的，都采用了双指针游走+二分法
func threeSum2(nums []int) [][]int {
	//先对数组排序
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums)-1; i++ { //这是起始位置类似于冒泡排序  nlg(n)
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		z := len(nums) - 1
		for z > j { //内部采用二分法时间复杂度为lg(n)
			b := nums[j]
			c := nums[z]
			if nums[i]+b+c > 0 {
				z--
			} else if nums[i]+b+c < 0 {
				j++
			} else {
				item := []int{nums[i], b, c}
				result = append(result, item)
				for j < z && nums[j] == nums[j+1] {
					//如果去重的话，就推荐的是里面的数字进行内部for循环遍历，当然因为不走外部循环那个就要包含外部的循环条件
					j++
				}
				for j < z && nums[z] == nums[z-1] { //因为是双指针，那么俩个都必须判断是否去重
					z--
				}
				j++
				z--
			}
		}
	}
	return result
}

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的一个字母异位词。
// 输入: s = "anagram", t = "nagaram"
// 输出: true
// 输入: s = "rat", t = "car"
// 输出: false
// 那我们可以通过对两个字符串里面的字符进行排序，如果排序后的两个字符串是一样的，那么就可以说这两个字符串是有效的
// 还有另外一种方法就是，使用map把字符串的字符出现个数保存起来
//方法一
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if s == "" {
		return true
	}
	var SArr []string
	var TArr []string

	for _, v := range s {
		SArr = append(SArr, string(v))
	}
	for _, v := range t {
		TArr = append(TArr, string(v))
	}
	sort.Strings(SArr)
	sort.Strings(TArr)
	// return reflect.DeepEqual(SArr, TArr)//这个也可以判断内容
	for k, v := range SArr {
		if v != TArr[k] {
			return false
		}
	}

	return true
}

//方法二
func isAnagram1(s string, t string) bool {
	var Smap = make(map[rune]int)
	var Tmap = make(map[rune]int)
	for _, v := range s {
		Smap[v] = Smap[v] + 1
	}
	for _, v := range t {
		Tmap[v] = Tmap[v] + 1
	}
	return reflect.DeepEqual(Smap, Tmap)
}

// 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口 k 内的数字。滑动窗口每次只向右移动一位。

// 返回滑动窗口最大值

// 复制代码
// 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
// 输出: [3,3,5,5,6,7]

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	length := len(nums)
	if length == 0 {
		return res
	} else {
		for i := 0; i <= length-k; i++ {
			res = append(res, GetMax(nums[i:i+k]))
		}

	}
	return res
}

func GetMax(nums []int) int {
	var maxNum = nums[0]
	for _, v := range nums {
		if maxNum < v {
			maxNum = v
		}

	}
	return maxNum
}

// 给定一个链表，判断链表中是否有环。
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//思路一个快移动一个慢移动，那么俩者终将会相遇

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	//为了防止painc所以在上层加入了判断，然后才会取
	for fast != nil && slow != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}

// 模拟栈
type MyStack struct {
	queue []int
	len   int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{make([]int, 0), 0}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
	this.len++
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	val := this.queue[this.len-1]
	this.queue = this.queue[0 : this.len-1]
	this.len--
	return val
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.queue[this.len-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.len == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

//给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

//腾讯面试我的问题
// func merge(nums1 []int, m int, nums2 []int, n int) {
// }
func merge(nums1 []int, nums2 []int) {
	nums3 := make([]int, 0)
	nums1 = nums1[:3]
	nums2 = nums2[:3]
	// copy(nums3, nums1)
	//copy 是将后者复制到前者，且任意一个改变不影响另外一个，前提切片必须足够大，不够大的话不会扩容，只会复制部分
	k1, k2 := 0, 0
	for k1 < len(nums1) && k2 < len(nums2) {
		if nums1[k1] < nums2[k2] {
			nums3 = append(nums3, nums1[k1])
			k1++
		} else {
			nums3 = append(nums3, nums2[k2])
			k2++
		}
	}
	if k1 == len(nums1) {
		nums3 = append(nums3, nums2[k2:]...)
	} else if k2 == len(nums2) {
		nums3 = append(nums3, nums1[k1:]...)
	}

	fmt.Println(nums3)
}

//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

// 输入: 1->1->2
// 输出: 1->2
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//不用递归
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	curr := head //如果不用递归那么就要用一个变量记录当前循环的节点
	for curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}

//递归方式
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val == head.Next.Val {
		// deleteDuplicates(head)
		for head.Next == nil && head.Val == head.Next.Val { //这里for循环就是去除多个重复的元素,因为他是指针所以每次结果都是影响判断的
			head.Next = head.Next.Next
		} //这里加for循环其实本质和递归是一样的,只是减少递归的层级,没太多必要,而且如果使用了for循环那么就走不了初始的边界值了,那么在这里就要重新判断==nil
		return deleteDuplicates1(head.Next) //这是在这里for循环删除连续的多个节点,然后再返回
	} else {
		head.Next = deleteDuplicates1(head.Next)
		return head
	}
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val == head.Next.Val {
		//这里for循环就是去除多个重复的元素,因为他是指针所以每次结果都是影响判断的
		head.Next = head.Next.Next
		return deleteDuplicates2(head)
	} else {
		head.Next = deleteDuplicates2(head.Next)
		return head
	}
}

// 输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。
func twoSum(nums []int, target int) []int {
	res := []int{}
frist:
	for x := 0; x < len(nums)-1; x++ {
		for y := x + 1; y < len(nums); y++ {
			if nums[x]+nums[y] == target {
				res = append(res, nums[x], nums[y])
				break frist
			}
		}
	}
	return res
} //这种超出时间限制了，时间复杂度为O(N^2)

//双指针+二分法  前提必须是有序的   ，如果是无序的可以先进行排序
func twoSum1(nums []int, target int) []int {
	//   因为他是有序的，可以利用双指针来进行游走//如果是无序的可以将他先排序，然后再进行二分游走法
	res := []int{}

	x, y := 0, len(nums)-1
	for x < y {
		sum := nums[x] + nums[y]
		if sum < target { //因为是最大的，所以他的和如果还比那个数小，就只能将x+1了因为y上已经是最大的数字了

			x++
		} else if sum > target {
			y--
		} else {
			res = append(res, nums[x], nums[y])
			break
		}
	}
	return res
}

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// 1     1
// 2     2
// 3     3
// 4     5
// 5     8
// 动态规划，根据规律可以发现是一个斐波那契数列，就是从第三个开始后面的是前俩个之和
func climbStairs(n int) int {
	res := []int{1, 2} //但是扩容会占用时间，优化初始化时的
	for i := 2; i < n; i++ {
		res = append(res, res[i-1]+res[i-2])
	}
	return res[n-1]
}

//leetcode解法,避免扩容问题，注意判断
func climbStairs1(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

// 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

// 你可以假设除了整数 0 之外，这个整数不会以零开头。

// 示例 1:

// 输入: [1,2,3]
// 输出: [1,2,4]
// 解释: 输入数组表示数字 123。
func plusOne(digits []int) []int {
	// 注意如果是 9999那么位数也要变化
	//所以从最后一位开始遍历
	arrlen := len(digits)

	for i := arrlen - 1; i >= 0; i-- {

		if digits[i] == 9 { //等于9就要进位
			digits[i] = 0
			if i == 0 {
				digits[i] = 1
				digits = append(digits, 0)
			}
		} else {
			digits[i] = digits[i] + 1
			break
		}

	}
	return digits
}
