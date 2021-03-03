package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func SendToMail(user, pwd, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")

	auth := smtp.PlainAuth("", user, pwd, hp[0])

	var content_type string

	if mailtype == "html" {

		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"

	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"

	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err

}

func main() {

	user := "3023493319@qq.com" //  你开通smtp邮箱的 邮箱的地址
	pwd := "zhwtdyfhxiapdgci"   //这里填你自己的授权码
	host := "smtp.qq.com:25"
	to := "2449968976@qq.com" //  目标地址
	subject := "使用Golang发送邮件"
	body := `
      <html>
      <body>
      <h3>
      "你好小可爱"
      </h3>
      </body>
      </html>
      `
	fmt.Println("SEND EMAIL ...")

	err := SendToMail(user, pwd, host, to, subject, body, "html")

	if err != nil {
		fmt.Println("发送邮件错误！")
		fmt.Println(err)
	} else {
		fmt.Println("发送成功")
	}
}
