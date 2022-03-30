// # 简单题
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
