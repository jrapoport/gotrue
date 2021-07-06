package audit

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jrapoport/gothic/core/context"
	"github.com/jrapoport/gothic/models/auditlog"
	"github.com/jrapoport/gothic/models/user"
	"github.com/jrapoport/gothic/store"
	"github.com/jrapoport/gothic/store/types"
	"github.com/jrapoport/gothic/store/types/key"
)

func TestLogLogin(t *testing.T) {
	testLogEntry(t, auditlog.Login, uuid.New(), nil,
		func(ctx context.Context, conn *store.Connection, uid uuid.UUID, _ types.Map) error {
			return LogLogin(ctx, conn, uid)
		})
}

func TestLogLogout(t *testing.T) {
	testLogEntry(t, auditlog.Logout, uuid.New(), nil,
		func(ctx context.Context, conn *store.Connection, uid uuid.UUID, _ types.Map) error {
			return LogLogout(ctx, conn, uid)
		})
}

func TestLogPasswordChange(t *testing.T) {
	testLogEntry(t, auditlog.Password, uuid.New(), nil,
		func(ctx context.Context, conn *store.Connection, uid uuid.UUID, _ types.Map) error {
			return LogPasswordChange(ctx, conn, uid)
		})
}

func TestLogEmailChange(t *testing.T) {
	testLogEntry(t, auditlog.Email, uuid.New(), nil,
		func(ctx context.Context, conn *store.Connection, uid uuid.UUID, _ types.Map) error {
			return LogEmailChange(ctx, conn, uid)
		})
}

func TestLogUpdate(t *testing.T) {
	testLogEntry(t, auditlog.Updated, uuid.New(), nil,
		func(ctx context.Context, conn *store.Connection, uid uuid.UUID, _ types.Map) error {
			return LogUserUpdated(ctx, conn, uid)
		})
}

func TestLogChangeRole(t *testing.T) {
	r := user.RoleAdmin
	testLogEntry(t, auditlog.ChangeRole, uuid.New(),
		types.Map{
			key.Role: r.String(),
		},
		func(ctx context.Context, conn *store.Connection, uid uuid.UUID, _ types.Map) error {
			return LogChangeRole(ctx, conn, uid, r)
		})
}
