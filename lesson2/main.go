package doc

import (
	"fmt"
	"os"
	"time"
)

//func main() {
//	err := thisIsPanic()
//	if err != nil {
//		fmt.Println(err.time.Format("Jan _2 15:04:05.00"), err.error)
//	}
//	createFile("dd.txt")
//}

//errorOutOfRange struct fir new Error
type ErrorOutOfRange struct {
	time  time.Time
	error interface{}
}

//NewError create new error
func NewError(s interface{}) *ErrorOutOfRange {
	t := time.Now()
	return &ErrorOutOfRange{
		time:  t,
		error: s,
	}
}

//thisIsPanic func is panic
func ThisIsPanic() (err *ErrorOutOfRange) {
	defer func() {
		if p := recover(); p != nil {
			err = NewError(p)
		}
	}()

	a := []int{1, 2, 3}
	fmt.Println(a[10])
	return nil
}

//createFile this func create new file
func CreateFile(name string) {
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
