package internal

import (
	"fmt"
)

// 学生を表すオブジェクト。出席番号と身長をもつ
type Student struct {
	Id   int
	Name string
}

func (s Student) String() string {
	return fmt.Sprintf("出席番号： %d, 氏名： %s", s.Id, s.Name)
}

/*
Aggregator
・オブジェクトの集合（今回だと[]*Student）
をもつ。
オブジェクト集合の各要素を見ていく処理は、iteratorの方に切り出されている。
これにより、StudentList（Agregator）を利用する人は、中身の実装を気にせず、 for sl.HasNext {} 的な形で中身をなめていける。
あとから[]*Student から別のクラスに変えたくなっても、利用する方では改修を加えないでOKなどの利点がある。
*/
type StudentList struct {
	students []*Student
	iterator *StudentsIterator
}

func NewStudentList() *StudentList {
	s := &StudentList{
		iterator: &StudentsIterator{},
	}
	s.iterator.studentList = s
	return s
}

func (s *StudentList) GetStudentAt(idx int) *Student {
	if s.GetSize() <= idx {
		return nil
	}
	return s.students[idx]
}

func (s *StudentList) HasNext() bool {
	return s.iterator.HasNext()
}

func (s *StudentList) Scan() *Student {
	return s.iterator.Scan()
}

func (s *StudentList) Append(student *Student) {
	s.students = append(s.students, student)
}

func (s *StudentList) GetSize() int {
	return len(s.students)
}

/*
iterator
あるオブジェクトのまとまりを順番に操作していく場合に役立つ。
・次の要素があるか
・次の要素を戻す
上記二つのメソッドを持つ。
*/
type StudentsIterator struct {
	studentList *StudentList
	idx         int
}

func (s *StudentsIterator) HasNext() bool {
	// aggregatorのGetSizeで、集合の要素数を得る
	if s.idx < s.studentList.GetSize() {
		return true
	}
	return false
}

func (s *StudentsIterator) Scan() *Student {
	student := s.studentList.GetStudentAt(s.idx)
	s.idx += 1
	return student
}
