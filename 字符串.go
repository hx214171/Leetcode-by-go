                                        // 简单题
// # 剑指 Offer 05. 替换空格
// # 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
// golang的string是以UTF-8编码的,而UTF-8是一种1-4字节的可变长字符集，每个字符可用1-4字节 来表示
// 使用下标方式s[i]访问字符串s，s[i]是UTF-8编码后的一个字节(uint8)，即按字节遍历
// 使用for i,v := range s 方式访问s，i是字符串下标编号，v是对应的字符值(int32=rune)，即按字符遍历
// 使用fmt.Printf打印时，%c占位符打印的是字符,%+v占位符打印的是这个类型自身，
// 如fmt.Printf(“%+v”,s[i])打印的就是字节一个十进制的无符号整数s[i]
func replaceSpace(s string) string {
	str := ""
	for _, v := range s {
		if v == ' ' {
			str += "%20"
		} else {
			str += string(v)
		}
	}
	return str
}

// # 剑指 Offer 50. 第一个只出现一次的字符
// # 在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
// 思路：两次遍历字符串，第一次遍历用数组或者字典保存只出现一次的字符，
// 第二次遍历找到只出现一次的字符。
func firstUniqChar(s string) byte {
	res := [26]int{}
	for _, ch := range s {
		res[ch-'a']++
	}
	for i, ch := range s {
		if res[ch-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

func firstUniqChar(s string) byte {
	dic := map[rune]bool{}
	for _, str := range s {
		if _, ok := dic[str]; ok {
			dic[str] = false
		} else {
			dic[str] = true
		}
	}
	for _, str := range s {
		if dic[str] {
			return byte(str)
		}
	}
	return ' '
}

// # 剑指 Offer 58 - I. 翻转单词顺序
// # 输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。
// # 例如输入字符串"I am a student. "，则输出"student. a am I"。
// 思路1；双指针法，从字符串后面往前找单词，添加到字符串里面。
// 思路2：调包，先用strings.split(s, " ")将字符串按空格分出来，然后添加到新的字符串。
func reverseWords(s string) string {
	i, j, res := len(s)-1, 0, ""
	for i >= 0 {
		for i >= 0 && s[i] == ' ' {
			i--
		}
		j = i
		for i >= 0 && s[i] != ' ' {
			i--
		}
		res += s[i+1:j+1] + " "
	}
	return strings.TrimRight(res, " ")
}

func reverseWords(s string) string {
	sub := strings.Split(s, " ")
	res := []string{}
	for i := len(sub) - 1; i >= 0; i-- {
		if sub[i] != "" {
			res = append(res, sub[i])
		}

	}
	return strings.Join(res, " ")
}

// // # 剑指 Offer 58 - II. 左旋转字符串
// # 字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。
// # 比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。
思路1：切片
func reverseLeftWords(s string, n int) string {
    return s[n:] + s[:n]
}

思路2：遍历数组(n,n+len(s))，添加s(i%len(s))。
func reverseLeftWords(s string, n int) string {
    res := []byte{}
    for i := n; i < n + len(s); i++ {
        res = append(res, s[i%len(s)])
    } 
    return string(res)
}

                                        // 中等题
// # 剑指 Offer 20. 表示数值的字符串
// # 请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。
// # 例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"-1E-16"、"0123"都表示数值，
// # 但"12e"、"1a3.14"、"1.2.3"、"+-5"及"12e+5.4"都不是。
思路：正则表达式
建议放弃

// # 剑指 Offer 38. 字符串的排列
// # 输入一个字符串，打印出该字符串中字符的所有排列。
// # 你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。
// 思路：dfs+剪枝，x为固定位，先固定0位有abc，再分别固定1位，2位。
func permutation(s string) []string {
    res := []string{}
    bytes := []byte(s)
    var dfs func(x int)
    dfs = func(x int) {
        if x == len(bytes)-1 {
            res = append(res, string(bytes))
            return
        }
        dic := map[byte]bool{}
        for i := x; i < len(bytes); i++ {
            if dic[bytes[i]] {
                continue
            }
            dic[bytes[i]] = true
            bytes[i],bytes[x] = bytes[x],bytes[i]
            dfs(x+1)
            bytes[i],bytes[x] = bytes[x],bytes[i]
        }
    }
    dfs(0)
    return res
}





// # 剑指 Offer 46. 把数字翻译成字符串
// # 给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，
// # 1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。
// # 请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。



// # 剑指Offer 48.
// # 最长不含重复字符的子字符串
// # 请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
思路1：动态规划
func lengthOfLongestSubstring(s string) int {
    dic := map[rune]int{}
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

思路2：滑动窗口
func lengthOfLongestSubstring(s string) int {
    i,res := -1, 0
    dic := map[rune]int{}
    for j,chr := range s {
        _,ok := dic[chr]
        if ok {
            i = max(i,dic[chr])
        }
        dic[chr] = j
        res = max(res,j-i)
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}


// # 剑指 Offer 67. 把字符串转换成整数 写一个函数 StrToInt，实现把字符串转换成整数这个功能。
// # 不能使用 atoi 或者其他类似的库函数。 首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。
// # 当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；
// # 假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。
// # 该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。
// # 注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，
// # 则你的函数不需要进行转换。 在任何情况下，若函数不能进行有效的转换时，请返回 0。
// # 说明： 假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 1]。
// # 如果数值超过这个范围，请返回  INT_MAX (231 − 1) 或 INT_MIN (−231) 。
