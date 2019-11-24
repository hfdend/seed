package utils

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"mime"
	"net"
	"net/mail"
	"net/smtp"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func CheckEmail(email string) error {
	if !regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`).MatchString(email) {
		return errors.New("请输入正确的邮箱")
	}
	return nil
}

func CheckUserName(userName string) error {
	if !regexp.MustCompile(`^\w{4,16}$`).MatchString(userName) {
		return errors.New("请输入正确的账号")
	}
	return nil
}

func CheckPhone(phone string) error {
	if !regexp.MustCompile(`^\d{11}$`).MatchString(phone) {
		return errors.New("请输入正确的手机号")
	}
	return nil
}

// 验证身份证号码
func CheckCertId(cartId string) error {
	if !regexp.MustCompile(`^\d{17}(\d|x|X)$`).MatchString(cartId) {
		return errors.New("请输入正确的身份证号码")
	}
	return nil
}

func Round(v float64, places int) float64 {
	s := fmt.Sprintf(fmt.Sprintf("%%.%df", places), v)
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func EncodePassword(password string) string {
	password = fmt.Sprintf("%s|%s", password, "")
	s := AesEncode(password)
	return s
}

func DecodePassword(password string) string {
	s := AesDecode(password)
	ary := strings.Split(s, "|")
	return ary[0]
}

func LRead(name string, level int) (raw []byte, err error) {
	var file *os.File
	for i := 0; i <= level; i++ {
		filePath := fmt.Sprintf("%s%s", strings.Repeat("../", i), name)
		file, err = os.OpenFile(filePath, os.O_RDONLY, 0600)
		if err != nil {
			continue
		} else {
			break
		}
	}
	if err != nil {
		return
	}
	raw, err = ioutil.ReadAll(file)
	return
}

func TimeLimit(start, end int) bool {
	h := time.Now().Hour()
	if h >= start && h < end {
		return true
	}
	return false
}

func RandomString(n int) string {
	var original = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	random := make([]rune, n)
	for i := range random {
		random[i] = original[rand.Intn(62)]
	}
	return string(random)
}

type Attachment struct {
	Name string
	Body []byte
}

func SendToMail(user, password, name, addr, to, subject, body string, isHtml bool, attachments ...Attachment) error {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return err
	}
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return err
	}
	// create new SMTP client
	smtpClient, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}
	auth := smtp.PlainAuth("", user, password, host)
	err = smtpClient.Auth(auth)
	if err != nil {
		return err
	}
	from := mail.Address{Name: name, Address: user}
	if err := smtpClient.Mail(from.Address); err != nil {
		return err
	}
	for _, v := range strings.Split(to, ";") {
		if err := smtpClient.Rcpt(strings.TrimSpace(v)); err != nil {
			return err
		}
	}

	writer, err := smtpClient.Data()
	if err != nil {
		return err
	}
	var contentType string
	if isHtml {
		contentType = "text/html;\r\n\tcharset=utf-8"
	} else {
		contentType = "text/plain;\r\n\tcharset=utf-8"
	}

	boundary := "----THIS_IS_BOUNDARY_JUST_MAKE_YOURS_MIXED"

	buffer := bytes.NewBuffer(nil)

	header := fmt.Sprintf("To: %s\r\n"+
		"From: %s\r\n"+
		"Subject: %s\r\n"+
		"Content-Type: multipart/mixed;\r\n\tBoundary=\"%s\"\r\n"+
		"Mime-Version: 1.0\r\n"+
		"Date: %s\r\n\r\n", to, user, subject, boundary, time.Now().String())
	buffer.WriteString(header)
	buffer.WriteString("This is a multi-part message in MIME format.\r\n\r\n")

	// 正文
	if len(body) > 0 {
		bodyBoundary := "----THIS_IS_BOUNDARY_JUST_MAKE_YOURS_BODY"
		buffer.WriteString(fmt.Sprintf("--%s\r\n", boundary))
		buffer.WriteString(fmt.Sprintf("Content-Type: multipart/alternative;\r\n\tBoundary=\"%s\"\r\n\r\n", bodyBoundary))

		buffer.WriteString(fmt.Sprintf("--%s\r\n", bodyBoundary))
		buffer.WriteString(fmt.Sprintf("Content-Type: %s\r\n", contentType))
		buffer.WriteString(fmt.Sprintf("Content-Transfer-Encoding: base64\r\n\r\n"))
		buffer.WriteString(fmt.Sprintf("%s\r\n\r\n", base64.StdEncoding.EncodeToString([]byte(body))))
		buffer.WriteString(fmt.Sprintf("--%s--\r\n", bodyBoundary))

	}
	for _, attachment := range attachments {
		t := mime.TypeByExtension(filepath.Ext(attachment.Name))
		if t == "" {
			t = "application/octet-stream"
		}
		buffer.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
		buffer.WriteString(fmt.Sprintf("Content-Transfer-Encoding: base64\r\n"))
		buffer.WriteString(fmt.Sprintf("Content-Type: %s; name=\"%s\"\r\n\r\n", t, attachment.Name))
		buffer.WriteString(fmt.Sprintf("%s\r\n\r\n", base64.StdEncoding.EncodeToString(attachment.Body)))
	}

	buffer.WriteString("\r\n\r\n--" + boundary + "--")
	_, err = writer.Write(buffer.Bytes())
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}
	return smtpClient.Quit()
}

func JSON(data interface{}) {
	bts, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(bts))
}

func Gzip(data []byte) ([]byte, error) {
	var res bytes.Buffer
	gz, _ := gzip.NewWriterLevel(&res, 7)
	_, err := gz.Write(data)
	if err != nil {
		return nil, err
	} else {
		gz.Close()
	}
	return res.Bytes(), nil
}

func GetImageSrc(domain, fieldID string) string {
	if strings.HasPrefix(fieldID, "data:image") {
		return fieldID
	}
	domain = strings.TrimRight(domain, "/") + "/upload/images/"
	return domain + fieldID
}
