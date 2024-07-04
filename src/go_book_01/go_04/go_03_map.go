package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {

	//内置的make函数可以创建一个map：
	ages := make(map[string]int) // mapping from strings to ints

	//初始化
	ages1 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	ages2 := make(map[string]int)
	ages2["alice"] = 31
	ages2["charlie"] = 34

	fmt.Println(ages)
	fmt.Println(ages1)
	fmt.Println(ages2)

	for name, age := range ages2 {
		fmt.Printf("%s\t%d\n", name, age)
	}

}

func dedup() {
	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

func charcount() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

/**


curd

1.初始化
ages2 := make(map[string]int)

2.修改+添加
ages["alice"] = 32

3.查询
fmt.Println(ages["alice"]) // "32"

4.删除
delete(ages, "alice") // remove element ages["alice"]

遍历
for name, age := range ages {
    fmt.Printf("%s\t%d\n", name, age)
}
在实践中，遍历的顺序是随机的，每一次遍历的顺序都不相同
如果要按顺序遍历key/value对，我们必须显式地对key进行排序，可以使用sort包的Strings函数对字符串slice进行排序
**/
