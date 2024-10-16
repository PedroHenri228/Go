package main

import (
	"fmt"
	"net/smtp"
)

func main() {

	from := "Seu Usuario"
	password := "Sua Senha"

	to := []string{"destinatario@exemple.com"}

	smtpHost := "Servidor SMPT"
	smtpPort := "587"

	message := []byte("Subject: Teste de envio de e-mail\n\nEste é um e-mail de teste enviado pelo Go!! ")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	if err != nil {
		fmt.Println("Erro ao enviar o e-mail:", err)
		return
	}

	fmt.Println("E-mail enviado com sucesso!")

}
