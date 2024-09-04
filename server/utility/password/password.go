package password

import (
	"math/rand/v2"
	"strings"
)

type Complexity struct {
	Digit     int // 数字
	Lowercase int // 小写字母
	Upper     int // 大写字母
	Other     int // 其他字符
}

// PasswordComplexity 校验密码格式，返回map[string]int,包含各个种类的长度
func PasswordComplexity(pwd string) *Complexity {
	password := &Complexity{
		Digit:     0,
		Lowercase: 0,
		Upper:     0,
		Other:     0,
	}

	for i := 0; i < len(pwd); i++ {
		switch {
		// 大写
		case 64 < pwd[i] && pwd[i] < 91:
			password.Upper += 1
		// 小写
		case 96 < pwd[i] && pwd[i] < 123:
			password.Lowercase += 1
		//数字
		case 47 < pwd[i] && pwd[i] < 58:
			password.Digit += 1
		// 空格
		case pwd[i] == 32:
			password.Other += 1
		//其他字符
		default:
			password.Other += 1
		}
	}

	return password
}

// PasswordStrength 密码强度
type PasswordStrength struct {
	Length    int //密码长度
	Digit     int //数字个数
	LowerCase int //小写字母个数
	Uppercase int //大写字母个数
	Special   int //特殊字符个数
}

// CreatePassword 随机生成符合要求长度的密码 length密码长度，最少个数：digit数字，special特殊字符 uppercase大写字母 lowerCase小写字母
func CreatePassword(length, digit, lowerCase, uppercase, special int, specialStr string) string {
	if strings.TrimSpace(specialStr) == "" {
		specialStr = "~=+%^*()[]{}!@#$?|"
	}
	//rand.Seed(time.Now().UnixNano()) //go1.20 被弃用了
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	digits := "0123456789"
	// specials := "~=+%^*()[]{}!@#$?|" //~!@#$%^&*()_+{}|[]\<>?=-
	uppercases := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerCases := "abcdefghijklmnopqrstuvwxyz"
	all := uppercases + lowerCases + digits + specialStr
	buf := make([]byte, length)
	//数字长度
	for i := 0; i < digit; i++ {
		buf[i] = digits[rand.IntN(len(digits))]
	}
	// 特殊字符
	for i := 0; i < special; i++ {
		buf[i+digit] = specialStr[rand.IntN(len(specialStr))]
	}
	// 大写字母
	for i := 0; i < uppercase; i++ {
		buf[i+digit+special] = uppercases[rand.IntN(len(uppercases))]
	}
	// 小写字母
	for i := 0; i < lowerCase; i++ {
		buf[i+digit+special+uppercase] = lowerCases[rand.IntN(len(lowerCases))]
	}
	// 剩下的随机生成
	for i := digit + special + uppercase + lowerCase - 1; i < length; i++ {
		buf[i] = all[rand.IntN(len(all))]
	}
	// 随机排列
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	return string(buf)
}

// PasswordFormat 校验密码强度 ,包含各个种类的长度
func PasswordFormat(pwd string) *PasswordStrength {
	var lower, s, d, upper int
	for i := 0; i < len(pwd); i++ {
		switch {
		// 大写
		case 64 < pwd[i] && pwd[i] < 91:
			upper += 1
		// 小写
		case 96 < pwd[i] && pwd[i] < 123:
			lower += 1
		//数字
		case 47 < pwd[i] && pwd[i] < 58:
			d += 1
		// 空格
		case pwd[i] == 32:
			s += 1
		//其他字符
		default:
			s += 1
		}
	}

	res := &PasswordStrength{
		Length:    len(pwd),
		Digit:     d,
		LowerCase: lower,
		Uppercase: upper,
		Special:   s,
	}
	return res
}
