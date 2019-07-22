package web

func CheckObjIsNil(data interface{}, msg string) error {
	if IsNil(data) {
		return &BusinessError{ErrorMsg: msg}
	}
	return nil
}

func CheckString(data string, msg string) error {
	if StringIsEmpty(data) {
		return &BusinessError{ErrorMsg: msg}
	}
	return nil
}

func CheckIsTrue(flag bool, msg string) error {
	if !flag {
		return &BusinessError{ErrorMsg: msg}
	}
	return nil
}
