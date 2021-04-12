package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	err := thisIsPanic()
	if err != nil {
		fmt.Println(err.time.Format("Jan _2 15:04:05.00"), err.error)
	}
	createFile("dd.txt")
}

type errorOutOfRange struct {
	time  time.Time
	error interface{}
}

func NewError(s interface{}) *errorOutOfRange {
	t := time.Now()
	return &errorOutOfRange{
		time:  t,
		error: s,
	}
}

func thisIsPanic() (err *errorOutOfRange) {
	defer func() {
		if p := recover(); p != nil {
			err = NewError(p)
		}
	}()

	a := []int{1, 2, 3}
	fmt.Println(a[10])
	return nil
}

func createFile(name string) {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Write([]byte("Hello"))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

}
