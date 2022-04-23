package utils

import (
	"association/global"
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"math/rand"
	"net/smtp"
	"time"
)

func Send(mail string) string {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	em := email.NewEmail()
	em.From = "1609499603@qq.com"
	em.To = []string{mail}
	em.Subject = "Verification Code"
	code := verificationCode()
	em.HTML = []byte(fmt.Sprintf("验证码：" + code + "，感谢您对本站的支持</hr><b><i>by:Tan</i></b>"))
	err := em.Send("smtp.qq.com:25", smtp.PlainAuth("", "1609499603@qq.com", "qpzlvegkcztuhgdh", "smtp.qq.com"))
	if err != nil {
		global.ASS_LOG.Error("邮件发送失败")
	}
	return code
}

func verificationCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return code
}
