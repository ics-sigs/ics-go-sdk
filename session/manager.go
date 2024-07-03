package session

import (
	"context"
	"net/url"

	"github.com/ics-sigs/ics-go-sdk/client"
	"github.com/ics-sigs/ics-go-sdk/client/methods"
	"github.com/ics-sigs/ics-go-sdk/client/types"
)

var (
	Locale = "cn"
	Domain = "internal"
	sessions = make(map[string]types.UserSession)
)

func init() {
	Locale = "cn"
	Domain = "internal"
}

type Manager struct {
	client      *client.Client
	userSession *types.UserSession
}

func NewManager(client *client.Client) *Manager {
	m := Manager{
		client: client,
	}

	return &m
}

func (sm *Manager) Login(ctx context.Context, u *url.Userinfo) error {
	req := types.Login{
		Domain: Domain,
		Locale: Locale,
	}

	if u != nil {
		req.Username = u.Username()
		if pw, ok := u.Password(); ok {
			req.Password = pw
		}
	}

	login, err := methods.Login(ctx, sm.client, &req)
	if err != nil {
		return err
	}

	sm.client.SetToken(login.SessonId)
	session := types.UserSession{
		UserId: login.UserId,
		Username: login.Username,
		SessonId: login.SessonId,
		RoleType: login.RoleType,
		Locale: login.Locale,
		IP: login.IP,
		Themes: login.Themes,
		CreateDate: login.CreateDate,
		LoginTime: login.LoginTime,
	}
	sessions[login.SessonId] = session
	sm.userSession = &session
	return nil
}

// UserSession retrieves and returns the SessionManager's CurrentSession field.
// Nil is returned if the session is not authenticated.
func (sm *Manager) UserSession(ctx context.Context) (*types.UserSession, error) {
	if sm.client == nil {
		return nil, nil
	}
	session := sessions[sm.client.Authorization]
	if &session == nil {
		return nil, nil
	}
	sm.userSession = &session
	err := methods.ValidUserSession(ctx, sm.client, sm.userSession)
	if err != nil {
		delete(sessions, sm.client.Authorization)
		if f, ok := err.(*types.SDKError); ok {
			if f.Code == "401" {
				return nil, nil
			}
		}
		return nil, err
	}
	return sm.userSession, nil
}

func (sm *Manager) Logout(ctx context.Context) error {
	_, err := methods.Logout(ctx, sm.client)
	return err
}