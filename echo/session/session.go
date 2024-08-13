package session

import (
	"errors"

	"github.com/gorilla/sessions"
	"github.com/gucchi0421/gopkg/app"
	"github.com/labstack/echo/v4"
)

var store = sessions.NewCookieStore([]byte(app.GetEnv("SESSION_SECRET_KEY", "bqfUqp4Eh6CH")))

// 新しいセッションに値を設定する
// err := session.New(c, "セッション名", "キー", "値")
func New(c echo.Context, sessName, key, value string) error {
    sess, err := store.Get(c.Request(), sessName)
    if err != nil {
        return err
    }
    sess.Values[key] = value

    return sess.Save(c.Request(), c.Response())
}

// セッションから値を取得する
func Get(c echo.Context, sessName, key string) (string, error) {
    sess, err := store.Get(c.Request(), sessName)
    if err != nil {
        return "", err
    }

    val, ok := sess.Values[key].(string)
    if !ok || val == "" {
        return "", errors.New("value not found or invalid")
    }

    return val, nil
}

// セッションをクリアする
func Clear(c echo.Context, sessName string) error {
	sess, err := store.Get(c.Request(), sessName)
	if err != nil {
		return err
	}
	sess.Options.MaxAge = -1

	return sess.Save(c.Request(), c.Response())
}
