package utils

import (
	"fmt"
	"testing"
)

func TestCheckCertId(t *testing.T) {
	err := CheckCertId("51132119890624813x")
	fmt.Println(err)
}

func TestDecodePassword(t *testing.T) {
	s := DecodePassword("7762bfd5246b0892220932449b")
	fmt.Println(s)
}
func TestEncodePassword(t *testing.T) {
	s := EncodePassword("123456")
	fmt.Println(s)
}

func TestEncode(t *testing.T) {
	fmt.Println(CheckEmail("123@d2.cs"))
}

func TestSendToMail(t *testing.T) {
	err := SendToMail(
		"system@gkcdbc.com",
		"@Feng1024",
		"system",
		"smtp.gmail.com:465",
		"malfurion02@outlook.com",
		"test",
		"哈哈",
		false,
	)
	fmt.Println(err)

}
