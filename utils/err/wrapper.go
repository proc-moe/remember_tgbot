package err

import "runtime"

func Wrap(err error) {
	runtime.Caller(0)
}
