package apperrors

type MyAppError struct {
	ErrCode
	Message string
	Err     error // エラーチェーンのための内部エラー
}

// エラーメソッド
func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

// Unwrap メソッド
func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}

// エラーをラップする
// 元のエラーに、エラーコードやメッセージをくっつけて新規のエラーを返す
// エラーの階層構造を作ることができる
func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}
