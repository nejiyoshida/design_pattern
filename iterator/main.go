package main

import (
	"fmt"

	"github.com/nejiyoshida/design_pattern/iterator/internal"
)

func main() {
	studentlist := internal.NewStudentList()
	studentlist.Append(&internal.Student{Id: 1, Name: "田中太郎"})
	studentlist.Append(&internal.Student{Id: 2, Name: "山田花子"})
	studentlist.Append(&internal.Student{Id: 3, Name: "佐藤翔太"})
	studentlist.Append(&internal.Student{Id: 4, Name: "松本英雄"})

	// Iteratorパターンにより、中で集合がどのように保持されてるかを気にせず順番に調べることができる
	for studentlist.HasNext() {
		s := studentlist.Scan()
		fmt.Println(s)
	}
}
