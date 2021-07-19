package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"os"
)

type CalibrationUtil struct {
}

func (c CalibrationUtil) Md5Check(b []byte, hash string) bool {
	s := md5.New()
	s.Write(b)
	sh := fmt.Sprintf("%x", s.Sum(nil))
	if sh == hash {
		return true
	} else {
		return false
	}
}

func (c CalibrationUtil) Sha1Check(b []byte, hash string) bool {
	s := sha1.New()
	s.Write(b)
	sh := fmt.Sprintf("%x", s.Sum(nil))
	if sh == hash {
		return true
	} else {
		return false
	}
}

func (c CalibrationUtil) Sha256Check(b []byte, hash string) bool {
	s := sha256.New()
	s.Write(b)
	sh := fmt.Sprintf("%x", s.Sum(nil))
	if sh == hash {
		return true
	} else {
		return false
	}
}

func (c CalibrationUtil) Sha512Check(b []byte, hash string) bool {
	s := sha512.New()
	s.Write(b)
	sh := fmt.Sprintf("%x", s.Sum(nil))
	if sh == hash {
		return true
	} else {
		return false
	}
}

func (c CalibrationUtil) Md5FileCheck(f *os.File, hash string) bool {
	s := md5.New()
	b := readFile(f)
	s.Write(b)
	sh := fmt.Sprintf("%x", s.Sum(nil))
	if sh == hash {
		return true
	} else {
		return false
	}
}

func (c CalibrationUtil) Sha1FileCheck(f *os.File, hash string) bool {
	s := sha1.New()
	b := readFile(f)
	s.Write(b)
	sh := fmt.Sprintf("%x", s.Sum(nil))
	if sh == hash {
		return true
	} else {
		return false
	}
}

func (c CalibrationUtil) Sha256FileCheck(f *os.File, hash string) bool {
	s := sha256.New()
	b := readFile(f)
	s.Write(b)
	sh := fmt.Sprintf("%x", s.Sum(nil))
	if sh == hash {
		return true
	} else {
		return false
	}
}

func (c CalibrationUtil) Sha512FileCheck(f *os.File, hash string) bool {
	s := sha512.New()
	b := readFile(f)
	s.Write(b)
	sh := fmt.Sprintf("%x", s.Sum(nil))
	if sh == hash {
		return true
	} else {
		return false
	}
}

func readFile(f *os.File) []byte {
	b, _ := ioutil.ReadAll(f)
	return b
}
