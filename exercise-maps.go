package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m_arr := strings.Fields(s)
	m_map := make(map[string]int)
	for i := range m_arr {
		if m_map[m_arr[i]] == 0 {
			m_map[m_arr[i]] = 1
		} else {
			m_map[m_arr[i]]++
		}
	}
	//for i := 0; i < len(m_arr); i++ {
	//
	//}
	return m_map
}
func main() {
	s := "I am learning Go!"
	res := WordCount(s)
	fmt.Println(res)
}
