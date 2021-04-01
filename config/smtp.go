package config

import "net/smtp"

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-12-29 00:09:36
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type SmtpMail struct {
	Auth        smtp.Auth
	SMTPAddress string
	Email       string
}
