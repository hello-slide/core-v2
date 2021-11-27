package handler

import (
	"net/http"

	"github.com/hello-slide/core-v2/account_manager/oauth"
	"github.com/hello-slide/core-v2/utils"
)

func AccountLoginHandler(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	tokenOp, err := utils.NewTokenOp(URL)
	if err != nil {
		utils.ErrorResponse(w, 1, err)
		return
	}

	refreshToken, err := tokenOp.GetRefreshToken(r)

	// リフレッシュトークンがある場合、そのトークンを使用してログインする
	// ない場合はGoogle OAuthのログイン画面にリダイレクト
	if err != nil || len(refreshToken) == 0 {
		if err := redirectOAuth(w, r); err != nil {
			utils.ErrorResponse(w, 1, err)
			return
		}
	} else {
		// _, err := utils.Get
	}
}

// Google OAuthログイン画面にリダイレクトする
func redirectOAuth(w http.ResponseWriter, r *http.Request) error {
	url, err := oauth.GetAuthURL()
	if err != nil {
		return err
	}

	http.Redirect(w, r, url, http.StatusMovedPermanently)
	return nil
}
