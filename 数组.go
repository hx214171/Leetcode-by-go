//package Algorithm
//简单题
//剑指 Offer 03. 数组中重复的数字
//找出数组中重复的数字。
//在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，
//但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

//# 思路1：集合
//# 遍历数组，若元素不在集合内则添加这个元素，若在集合内则返回这个元素。
func findRepeatNumber(nums []int) int {
	s := map[int]bool{}
	for _,num := range nums {
		if s[num] { 
			return num
		}
		s[num] = true
	}
	return -1
}

//思路2：原地交换
//# 遍历数组的同时通过交换操作将每个元素的索引和他的值对应。
func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] != nums[nums[i]] {
				nums[i],nums[nums[i]] = nums[nums[i]],nums[i]
			} else {
				return nums[i]
			}
		}
	}
	return -1
}

//# 剑指 Offer 11. 旋转数组的最小数字
//# 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
//# 输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，
//# 数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。
//思路：因为是递增序列的翻转，找到旋转数组的中间值，和数组末位值比较，更新末位值。
func minArray(numbers []int) int {
	low, high := 0, len(numbers)-1
	for low < high {
		mid := (low + high)/2
		if numbers[mid] > numbers[high] {
			low = mid + 1
		} else if numbers[mid] < numbers[high] {
			high = mid
		} else {
			high -= 1
		}
	}
	return numbers[low]
}

//# 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
//# 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。
//思路：分别从首位找到偶数和奇数，调换位置即可。
func exchange(nums []int) []int {
	for i,j := 0,len(nums)-1 ; i < j; {
		for i < j && nums[i] % 2 == 1 {
			i += 1
		}
		for i < j && nums[j] % 2 == 0 {
			j -= 1
		}
		nums[i],nums[j] = nums[j],nums[i]
	}
	return nums
}

//# 剑指 Offer 29. 顺时针打印矩阵
//# 输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
//# 思路：分别设置左右上下边界，然后从左到右打印完，t+1，从上到下打印，r-1，从右到左打印，b-1，从下到上打印，l+1，
//# 分别添加元素到res结果列表,l > r 或者 t > b时，则退出循环。
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0{
		return nil
	}
	l,r,t,b := 0,len(matrix[0])-1,0,len(matrix)-1
	var res []int
	for {
		for i := l; i <= r; i++ {
			res = append(res,matrix[t][i])
		}
		t += 1
		if t > b {
			break
		}
		for i := t; i <= b; i++ {
			res = append(res,matrix[i][r])
		}
		r -= 1
		if r < l {
			break
		}
		for i := r; i >= l; i-- {
			res = append(res,matrix[b][i])
		}
		b -= 1
		if b < t {
			break
		}
		for i := b; i >= t; i-- {
			res = append(res,matrix[i][l])
		}
		l += 1
		if l > r {
			break
		}
	}
	return res
}

//# 剑指 Offer 39.数组中出现次数超过一半的数字
//# 数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//# 思路1：投票法，遍历数组的同时，记录元素的投票数，如果遇到相同的元素，则投票+1，否则-1，返回投票最多的那个元素。
func majorityElement(nums []int) int {
	vote := 0
	var result int
	for _,num := range nums {
		if vote == 0 {
			result = num
		}
		if result == num {
			vote += 1
		} else {
			vote -= 1
		}
	}
	return result
}
//思路2：遍历数组的同时将每个元素加入hash表，并判断元素个数是否超过数组长度一半。
func majorityElement(nums []int) int {
	hash := make(map[int]int)
	for i := 0;i < len(nums); i++ {
		_, ok := hash[nums[i]]
		if ok {
			hash[nums[i]] += 1
		} else {
			hash[nums[i]] = 1
		}
		if hash[nums[i]] > len(nums)/2 {
			return nums[i]
		}
	}
	return 0
}
//思路3：排序，取中间值。

//# 剑指 Offer 40.最小的k个数
//# 输入整数数组arr ，找出其中最小的k个数。例如，输入4、5、1、6、2、7、3、8这8个数字，
//# 则最小的4个数字是1、2、3、4。
有bug下次再做




//# 剑指 Offer 42. 连续子数组的最大和
//# 输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
//# 要求时间复杂度为O(n)。
//# 输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
//# 输出: 6
//# 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
//思路：动态规划
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0{
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

//# 剑指 Offer 53 - I. 在排序数组中查找数字 I
//# 统计一个数字在排序数组中出现的次数。
//思路1：遍历数组统计target次数
func search(nums []int, target int) int {
	res := 0
	for _, num := range nums {
		if num == target {
			res++
		}
	}
	return res
}

//需要再看一下
func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i+j)/2
		if nums[mid] <= target {       //考虑边界问题
			i = mid + 1
		} else {
			j = mid -1
		}
	}
	right := i
	i = 0
	for i <= j {
		mid := (i+j)/2
		if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid -1
		}
	}
	left := j
	return right-left-1
}

//思路2：两次二分查找，查找target的右边界以及target-1的右边界，相减得到次数
func search(nums []int, target int) int {
	return help(nums, target)-help(nums, target-1)
}

func help(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i+j)/2
		if nums[mid] <= target {
			i = mid + 1
		} else {
			j = mid -1
		}
	}
	return i
}

//# 剑指 Offer 53 - II. 0～n-1中缺失的数字
//# 一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。
//# 在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。
//思路1：遍历数组，如果索引不等于值则返回索引，最后返回数组长度。
func missingNumber(nums []int) int {
	for i, v := range nums {
		if i != v {
			return i
		}
	}
	return len(nums)
}
//思路2：有序数组想到二分法，左子数组元素与索引号相等，右子数组元素与索引号不等，二分查找右子数组的首位元素即为缺失数字。
func missingNumber(nums []int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i+j)/2
		if nums[mid] == mid {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return i
}

//#                                       中等题
//# 剑指 Offer 4. 二维数组中的查找
//# 在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
//# 请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
//# 思路1：从右上角（左下角)出发，矩阵空时 返回False，target > matrix[i][j]则i+=1，
//target < matrix[i][j]则j-=1。
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) < 1 {
		return false
	}
	i, j := 0, len(matrix[0])-1
	for i <= len(matrix)-1 && j >= 0 {
		if target > matrix[i][j] {
			i++
		} else if target < matrix[i][j]{
			j--
		} else {
			return true
		}
	}
	return false
}

//# 剑指 Offer 12. 矩阵中的路径
//# 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；
//# 否则，返回 false 。单词必须按照字母顺序，通过相邻的单元格内的字母构成，
//# 其中“相邻”单思路：DFS+剪枝。元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
func exist(board [][]byte, word string) bool {
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {    //k= n-1时表示字符串全都匹配了
			return true
		}
		board[i][j] = ' '
		res := dfs(i+1,j,k+1) || dfs(i-1,j,k+1) || dfs(i,j+1,k+1) ||dfs(i,j-1,k+1)
		board[i][j] = word[k]
		return res
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++{
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

//# 剑指 Offer 44. 数字序列中某一位的数字
//# 数字以0123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，第5位（从下标0开始计数）是5，
//# 第13位是1，第19位是4，等等。
//# 请写一个函数，求任意第n位对应的数字。
//思路：找规律。找这个数字是几位数，找这个位数的起始位置。
func findNthDigit(n int) int {
	start, digit, count := 1, 1, 9
	for n > count {
		n -= count
		start *= 10
		digit += 1
		count = 9*start*digit
	}
	num := start + (n-1)/digit
	res := strconv.Itoa(num)[(n-1)%digit]-'0'
	return int(res)
}

//# 剑指 Offer 45. 把数组排成最小的数
//# 输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
//思路1：利用sort.Slice方法自定义排序
func minNumber(nums []int) string {
    if len(nums) == 0 {
        return ""
    }
    strs := make([]string, len(nums))
    for i := range nums {
        strs[i] = strconv.Itoa(nums[i])
    }
    sort.Slice(strs, func(i, j int)bool{
        str12 := strs[i] + strs[j]
        str21 := strs[j] + strs[i]
        return str12 < str21
    })
    res := ""
    for _,str := range strs{
        res += str
    }
    return res
}

//思路2：定义一个自定义排序函数，放在快速排序里比较

***
//# 剑指Offer 48.最长不含重复字符的子字符串
//# 请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
//思路1：动态规划
func lengthOfLongestSubstring(s string) int {
    dic := map[byte]int{}
    res, dp := 0, 0
    for i, str := range s {
        j, ok := dic[str]
        if !ok || dp < i-j {
            dp += 1
        } else {
            dp = i - j
        }
        dic[str] = i
        if res < dp {
            res = dp
        }
    }
    return res
}
***
思路2：滑动窗口

//# 剑指 Offer 56 - I. 数组中数字出现的次数
//# 一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。
//# 要求时间复杂度是O(n)，空间复杂度是O(1)。
// 思路：利用相同数字异或为0找到那两个出现一次的数字，接下来找到这两个数字不同的二进制位，
// 再次遍历数组通过这个二进制位将两个数字进行分组异或计算，最终得到两个值。
func singleNumbers(nums []int) []int {
    x, y, n, m := 0,0,0,1
    for _,num := range nums {
        n ^= num
    }
    for n&m == 0 {
        m <<= 1
    }
    for _,num := range nums {
        if num & m == 0 {
            x ^= num
        } else {
            y ^= num
        }
    }
    arr := []int{x, y}
    return arr
}

// # 剑指 Offer 56 - II. 数组中数字出现的次数 II
// # 在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。
// 思路：将每个数组的32个二进制位累加，和三除余，不为0的就是出现一次的那个数字的二进制位。
func singleNumber(nums []int) int {
    res := 0
    for i := 0; i < 32; i++ {
        bit := 0
        for _,num := range nums {
            bit += ((num >> i)& 1)
        }
        res += ((bit%3) << i)
    }
    return res
}

// # 剑指 Offer 66. 构建乘积数组
// # 给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B[i] 的值是数组 A 中除了下标 i 以外的元素的积,
// # 即 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。
// 思路：数组B[i]的每个值可以看成a[i]前面元素的乘积与a[i]后面元素的乘积。所以两次遍历数组，
// 一次顺序一次逆序，分别求得a[i]左边的乘积与a[i]后面的乘积。
func constructArr(a []int) []int {
    tmp := 1
    b := make([]int, len(a))
    for i := range a {
        b[i] = 1
    }
    for i := 1; i < len(a); i++ {
        b[i] = b[i-1]*a[i-1]
    }
    for i := len(a)-2; i >= 0; i-- {
        tmp *=a[i+1]
        b[i] *= tmp
    }
    return b
}



