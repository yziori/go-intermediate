package apperrors

type MyApiError struct {
	ErrCode
	Message string
	Err     error // エラーチェーンのための内部エラー
}

// エラーメソッド
func (myErr *MyApiError) Error() string {
	return myErr.Err.Error()
}

// Unwrap メソッド
func (myErr *MyApiError) Unwrap() error {
	return myErr.Err
}
