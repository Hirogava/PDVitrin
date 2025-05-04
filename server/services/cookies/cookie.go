package cookies

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

type Manager struct {
	Session *sessions.Session
}

var store *sessions.CookieStore

func Init(key string) {
	store = sessions.NewCookieStore([]byte(key))
	store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   86400 * 7,
		SameSite: http.SameSiteLaxMode,
	}
	log.Println("Cookie store инициализирован с ключом:", key[:5]+"...")
}

func NewCookieManager(r *http.Request) *Manager {
	session, err := store.Get(r, "street-store")
	if err != nil {
		log.Printf("Ошибка при получении сессии: %v", err)
	}
	return &Manager{
		Session: session,
	}
}