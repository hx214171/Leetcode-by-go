									// 简单题
// # 剑指 Offer 15. 二进制中1的个数
// # 请实现一个函数，输入一个整数（以二进制串形式），输出该数二进制表示中 1 的个数。
// # 例如，把 9 表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2
func hammingWeight(num uint32) int {
    var res uint32
    for num != 0 {
        res += num&1
        num >>= 1
    }
    return int(res)
}

func hammingWeight(num uint32) int {
    res := 0
    for num != 0 {
        num &= num-1
        res++
    }
    return res
}

// # 剑指 Offer 17. 打印从1到最大的n位数
// # 输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
func printNumbers(n int) []int {
    res := []int{}
    max := 1
    for i := 0; i < n; i++{
        max *= 10
    }
    for i := 1; i < max; i++{
        res = append(res, i)
    }
    return res
}

// # 剑指 Offer 57. 和为s的两个数字
// # 输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。
// # 如果有多对数字的和等于s，则输出任意一对即可。
思路1：双指针
func twoSum(nums []int, target int) []int {
    i := 0
    j := len(nums)-1
    for i < j {
        if nums[i] + nums[j] == target {return []int{nums[i], nums[j]}}
        if nums[i] + nums[j] < target {
            i++
        } else {
            j--
        }
    }
    return nil
}

思路2：哈希表
func twoSum(nums []int, target int) []int {
    if len(nums) <= 1 {return nil}
    set := map[int]int{}
    for _, v := range nums {
        if k, ok := set[target-v]; ok {
            return []int{k, v}
        }
        set[v] = v
    }
    return nil
}

// # 剑指 Offer 57 - II. 和为s的连续正数序列
// # 输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
// # 序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
思路：滑动窗口
func findContinuousSequence(target int) [][]int {
    i, j, sum, res := 1,2,3,[][]int{}
    for i <= target/2 {
        if sum < target {
            j++
            sum += j
        } else if sum > target {
            i++
            sum -= i
        } else {
            res = append(res,list(i,j))
            sum -= i
            i += 1
        }
    }
    return res
}

func list(a, b int) []int {
    res := []int{}
    for i := a; i < b; i++ {
        res = append(res, i)
    }
    return res
}

// # 剑指 Offer 61. 扑克牌中的顺子
// # 从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，
// # A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。
思路1：集合判断是否存在重复扑克牌，若重复则返回false。5张牌的最大值和最小值小于5则返回true。
func isStraight(nums []int) bool {
    set := map[int]bool{}
    maxNum, minNum := 0, 14
    for _,val := range nums {
        if val == 0 {continue}
        maxNum = max(maxNum, val)
        minNum = min(minNum, val)
        if _,ok := set[val]; ok {return false}
        set[val] = true
    }
    return maxNum-minNum < 5
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a,b int) int {
    if a < b {
        return a
    } 
    return b
}

思路2：先排序，记录大小王个数，遍历数组，若相邻的扑克牌有相同数字则返回false，
最后若nums[4]-nums[joker]小于5则返回true。
func isStraight(nums []int) bool {
    sort.Ints(nums)
    joker := 0
    for i := 0; i < 4; i++ {
        if nums[i] == 0 {
            joker++
        } else if nums[i] == nums[i+1] {
            return false
        }
    }
    return nums[4]-nums[joker] < 5
}

// # 剑指 Offer 62. 圆圈中最后剩下的数字
// # 0,1,···,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字（删除后从下一个数字开始计数）。
// # 求出这个圆圈里剩下的最后一个数字。
// # 例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，
// # 因此最后剩下的数字是3
思路：数学公式，需要再看看。
func lastRemaining(n int, m int) int {
    res := 0
    for i := 2; i <= n; i++{
        res = (res + m) % i
    }
    return res
}

// # 剑指 Offer 65. 不用加减乘除做加法
// # 写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号
思路：位运算
func add(a int, b int) int {
    carry := 0
    for b != 0 {
        carry = (a & b) << 1 //carry为进位
        a ^= b  // a^b为非进位和
        b = carry 
    }
    return a
}


									// 中等题
// # 剑指Offer13.机器人的运动范围
// # 地上有一个m行n列的方格，从坐标[0, 0]到坐标[m - 1, n - 1] 。一个机器人从坐标[0, 0]
// # 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。
// # 例如，当k为18时，机器人能够进入方格[35, 37] ，因为3 + 5 + 3 + 7 = 18。但它不能进入方格[35, 38]，
// # 因为3 + 5 + 3 + 8 = 19。请问该机器人能够到达多少个格子？
思路：DFS+剪枝
func movingCount(m int, n int, k int) int {
    vis := make([][]bool, m)
    for i := range vis {
        vis[i] = make([]bool, n)
    }
    return dfs(0, 0, m, n, k, vis)
}

func dfs(i, j, m, n, k int, vis [][]bool) int {
    if i < 0 || i >= m || j < 0 || j >= n || numSum(i) + numSum(j) > k || vis[i][j] == true{
        return 0}
    vis[i][j] = true
    return 1 + dfs(i+1, j, m, n, k, vis) + dfs(i, j+1, m, n, k, vis)
}

func numSum(x int) int {
    res := 0
    for x != 0 {
        res += x%10
        x /= 10
    }
    return res
}

// # 剑指 Offer 14- I. 剪绳子
// # 给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为
// # k[0],k[1]...k[m-1]。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？
// # 例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18
思路1：公式法
func cuttingRope(n int) int {
    if n <= 3 {return n-1}
    a, b := n/3, n%3
    if b == 0 {return int(math.Pow(3, float64(a)))}
    if b == 1 {return int(math.Pow(3, float64(a-1))*4)}
    return int(math.Pow(3, float64(a))*2)
}

思路2：动态规划
func cuttingRope(n int) int {
    dp := make([]int, n+1)
    dp[2] = 1
    for i := 3; i < n+1; i++ {
        for j := 2; j < i; j++ {
             tmp := max(j*dp[i-j], j*(i-j))
             dp[i] = max(tmp,dp[i])
        }
    }
    return dp[n]
}

func max(a, b int) int {
    if a < b { return b}
    return a
}


// # 剑指 Offer 14- II. 剪绳子 II
// # 给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为
// # k[0],k[1]...k[m - 1] 。请问 k[0]*k[1]*...*k[m - 1]可能的最大乘积是多少？例如，当绳子的长度是8时，
// # 我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。
// # 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
思路：贪心算法
func cuttingRope(n int) int {
    if n <= 3 {return n-1}
    res, mod := 1, 1000000007
    for n > 4 {
        res = res * 3 % mod
        n -= 3
    }
    return res * n % mod
}

// # 剑指 Offer 16. 数值的整数次方
// # 实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。
func myPow(x float64, n int) float64 {
    if x == 0 {return 0}
    if n < 0 { 
        x, n = 1/x, -n
    }
    res := 1.0
    for n != 0 {
        if n & 1 == 1 {
            res *= x
        }
        x *= x
        n >>= 1
    }
    return res
}

// # 剑指 Offer 47. 礼物的最大价值
// # 在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。
// # 你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。
// # 给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
func maxValue(grid [][]int) int {
    for i := 0; i < len(grid); i++ {
        for j:= 0; j < len(grid[0]); j++ {
            if i == 0 && j == 0{ continue }
            if i == 0 {
                grid[i][j] = grid[i][j-1] + grid[i][j]
            } else if j == 0 {
                grid[i][j] = grid[i-1][j] + grid[i][j] 
            } else {
                grid[i][j]=max(grid[i-1][j],grid[i][j-1])+grid[i][j]
            }
        }
    }
    return grid[len(grid)-1][len(grid[0])-1]
}

func max(a, b int ) int {
    if a < b {return b}
    return a
}

// # 剑指 Offer 49. 丑数
// # 我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数
func nthUglyNumber(n int) int {
    dp := make([]int, n)
    var  a, b, c int
    dp[0] = 1
    for i := 1; i < n; i++ {
        dp[i] = min(dp[a]*2, dp[b]*3, dp[c]*5)
        if dp[i] == dp[a]*2 {a += 1}
        if dp[i] == dp[b]*3 {b += 1}
        if dp[i] == dp[c]*5 {c += 1}
    }
    return dp[n-1]
}

func min(a, b, c int) int {
    if a <= b && a <= c {return a}
    if b <= a && b <= c {return b}
    return c
}

// # 剑指 Offer 60. n个骰子的点数
// # 把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率
// # 你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 
// 小的那个的概率。
思路： 还有问题
func dicesProbability(n int) []float64 {
    dp := make([][]float64, n+1)
    for i := range dp {
        dp[i] = make([]float64, 6*n+1)
    }
    for i := 1; i < 7; i++ {
        dp[1][i] = 1
    }
    for i := 2; i < n+1; i++ {
        for j := i; j < 6*i+1; j++ {
            for k := 1; k < 7; j++{
                if j >= k+1 {
                    dp[i][j] += dp[i-1][j-k]
                }
            }
        }
    }
    res := make([]float64, 6*n+1)
    for i := n; i <= 6*n; i++ {
        res[i] = float64(dp[n][i]) / math.Pow(6,float64(n))
    }
    return res[n:]
}


// # 剑指 Offer 63. 股票的最大利润
// # 假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少
思路：动态规划
func maxProfit(prices []int) int {
    if len(prices) <= 1 {return 0}
    minPrice, profit := prices[0], 0
    for i := 1; i < len(prices); i++ {
        if minPrice > prices[i] {
            minPrice = prices[i]
        }
        if prices[i]-minPrice > profit {
            profit = prices[i]-minPrice
        }
    }
    return profit
}

// # 剑指 Offer 64. 求1+2+…+n
// # 求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）
func sumNums(n int) int {
    res := 0
    var sum func(n int) bool
    sum = func(n int) bool {
        res += n
        return n > 0 && sum(n-1)
    }
    sum(n)
    return res
}