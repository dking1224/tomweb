package web

import "fmt"

type CheckError struct {
	ErrorMsg string
}

func (err CheckError) Error() string {
	return fmt.Sprintf("check error:%s", err.ErrorMsg)
}

type BusinessError struct {
	ErrorMsg string
}

func (err BusinessError) Error() string {
	return fmt.Sprintf("business error:%s", err.ErrorMsg)
}
