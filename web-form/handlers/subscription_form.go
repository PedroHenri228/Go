package handlers

import (
	"fmt"
	"net/http"
	"web-form/controllers"
)

func SubscriptionHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Erro ao fazer o parse do form: %v", err)
			return
		}

		err :=  controllers.Create(r.PostForm.Get("name"), r.PostForm.Get("email"))

		if err != nil {
			fmt.Fprintf(w, "Erro ao salvar os dados: %v", err)
			return
		}
	}

	http.ServeFile(w, r, "handlers/templates/subscription_form.html")

}