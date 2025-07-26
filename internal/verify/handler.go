package verify

import (
	"API-email-confirmation/configs"
	"API-email-confirmation/pkg/email"
	"API-email-confirmation/pkg/hash"
	"API-email-confirmation/pkg/request"
	"API-email-confirmation/pkg/response"
	"API-email-confirmation/pkg/storage"
	"fmt"
	"net/http"
)

type VerifyHandlerDeps struct {
	*configs.Config
}

type VerifyHandler struct {
	*configs.Config
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[EmailRequest](&w, req)
		if err != nil {
			return
		}

		hash, err := hash.GenerateHash()
		if err != nil {
			return
		}

		link := fmt.Sprintf("http://localhost:8081/verify/%s", hash)

		err = email.SendEmail(
			handler.Config.Verify.Email,
			body.Email,
			"Confirm your email",
			fmt.Sprintf("Click to verify: %s", link),
			handler.Config.Verify.Password,
			handler.Config.Verify.Address,
		)
		if err != nil {
			http.Error(w, "Failed to send email", http.StatusInternalServerError)
			return
		}

		storage.Save("data.json", map[string]any{
			"email": body.Email,
			"hash":  hash,
		})

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Email sent")
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		hash := req.PathValue("hash")

		found, err := storage.DeleteByHash("data.json", hash)
		if err != nil {
			response.JSONResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if found {
			fmt.Fprint(w, "true")
		} else {
			fmt.Fprint(w, "false")
		}
	}
}
