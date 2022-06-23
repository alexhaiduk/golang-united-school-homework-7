package coverage

import (
	"os"
	"testing"
	"time"
	"reflect"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestPeople_Len(t *testing.T) {
	peoples := People{
		{firstName: "Ivan", lastName: "Ivanov", birthDay: time.Now().UTC().Local()},
		{firstName: "Petr", lastName: "Petrov", birthDay: time.Now().UTC().Local()},
		{firstName: "Maxim", lastName: "Maximov", birthDay: time.Now().UTC().Local()}}

	if peoples.Len() != 3 {
		t.Errorf("the expected length is 3, but actual %d", peoples.Len())
	}
}

func TestPeople_Less(t *testing.T) {
	peoples := People{
		{firstName: "Ivan", lastName: "Ivanov", birthDay: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)},
		{firstName: "Ivan", lastName: "Petrov", birthDay: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)},
		{firstName: "Maxim", lastName: "Maximov", birthDay: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)},
		{firstName: "Ivan", lastName: "Maximov", birthDay: time.Date(2000, time.November, 10, 23, 0, 0, 0, time.UTC)}}

	if !peoples.Less(0, 1) {
		t.Errorf("the expected is True, but actual %t", peoples.Less(0, 1))
	}

	if peoples.Less(1, 0) {
		t.Errorf("the expected is False, but actual %t", peoples.Less(1, 0))
	}

	if !peoples.Less(1, 2) {
		t.Errorf("the expected is True, but actual %t", peoples.Less(1, 2))
	}

	if peoples.Less(2, 1) {
		t.Errorf("the expected is False, but actual %t", peoples.Less(2, 1))
	}

	if !peoples.Less(2, 3) {
		t.Errorf("the expected is True, but actual %t", peoples.Less(2, 3))
	}

	if peoples.Less(3, 2) {
		t.Errorf("the expected is False, but actual %t", peoples.Less(3, 2))
	}
}

func TestPeople_Swap(t *testing.T) {
	peoples := People{
		{firstName: "Ivan", lastName: "Ivanov", birthDay: time.Now().UTC().Local()},
		{firstName: "Petr", lastName: "Petrov", birthDay: time.Now().UTC().Local()},
		{firstName: "Maxim", lastName: "Maximov", birthDay: time.Now().UTC().Local()}}

	peoples.Swap(0, 1)

	if peoples[0].firstName != "Petr" {
		t.Errorf("the expected is Petr, but actual is %s", peoples[0].firstName)
	}

	peoples.Swap(1, 2)

	if peoples[1].firstName != "Maxim" {
		t.Errorf("the expected is Maxim, but actual is %s", peoples[1].firstName)
	}

	peoples.Swap(0, 2)

	if peoples[0].firstName != "Ivan" {
		t.Errorf("the expected is Ivan, but actual is %s", peoples[0].firstName)
	}
}

////////////////////////////////////////////////////////////////

func TestNew(t *testing.T) {
	_, err := New("123 456 789\n123")
	if err == nil {
		t.Errorf("the expected is \"Rows need to be the same length\", but actual is %v", err)
	}

	_, err = New("123 456 qwe\n123 456 789\n123 456 789")
	if err == nil {
		t.Errorf("the expected is Error, but actual is %v", err)
	}

	_, err = New("123 456 789\n789 456")
	if err == nil {
		t.Errorf("the expected is \"Rows need to be the same length\", but actual is %v", err)
	}
}

func TestMatrix_Rows(t *testing.T) {
	mat, _ := New("1 1 1\n2 2 2\n3 3 3")
	compare := mat.Rows()
	expvalue := [][]int{{1,1,1},{2,2,2},{3,3,3}}

	if !reflect.DeepEqual(compare, expvalue){
		t.Errorf("expected to get %v, but actual is %v", expvalue, compare)
	}
}

func TestMatrix_Cols(t *testing.T) {
	mat, _ := New("1 1 1\n2 2 2\n3 3 3")
	expvalue := [][]int{{1,2,3},{1,2,3},{1,2,3}}

	if !reflect.DeepEqual(expvalue, mat.Cols()){
		t.Errorf("expected to get %v, but actual is %v", expvalue, mat.Cols())
	}
}

func TestMatrix_Set(t *testing.T) {
	mat, _ := New("1 1 1\n2 2 2\n3 3 3")
	expvalue := &Matrix{3,3, []int{5,1,1,2,2,2,3,3,3}}
	result := mat.Set(0,0,5)
	if result {
		if !reflect.DeepEqual(mat, expvalue) {
			t.Errorf("expected to get %v, but actual is %v", expvalue, mat)
		}
	} else {
		t.Errorf("expected to get True, but actual is %v", result)
	}
	result = mat.Set(4,3,9)
	if result {
		t.Errorf("expected to get False, but actual is %v", result)
	}
}
