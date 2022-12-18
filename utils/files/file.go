package files

import (
	"io/ioutil"
	"os"

	"github.com/proc-moe/remember_tgbot/utils/klog"
)

func CheckExist(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return true, err
	}
	return true, nil
}

func ReadString(path string) (string, error) {
	bytedStr, err := ioutil.ReadFile(path)
	if err != nil {
		klog.E("[Files ReadStr] file open failed, err=%v", err)
		return "", err
	}
	return string(bytedStr), nil
}

func Read(path string) ([]byte, error) {
	bytedStr, err := ioutil.ReadFile(path)
	if err != nil {
		klog.E("[Files ReadStr] file open failed, err=%v", err)
		return nil, err
	}
	return bytedStr, nil
}
