package users

import (
	"github.com/google/uuid"
	"github.com/jrapoport/gothic/models/user"
	"github.com/jrapoport/gothic/store"
	"github.com/jrapoport/gothic/store/types/key"
)

// SearchUsers search the audit log entries.
func SearchUsers(conn *store.Connection, s store.Sort, f store.Filters, p *store.Pagination) ([]*user.User, error) {
	tx := conn.Model(new(user.User))
	filters := make(store.Filters, len(f))
	for k, v := range f {
		filters[k] = v
	}
	if v, ok := filters[key.UserID]; ok {
		filters[key.ID] = v
		delete(filters, key.UserID)
	}
	if uid, ok := f[key.ID].(string); ok && uid != "" {
		id, err := uuid.Parse(uid)
		if err != nil {
			return nil, err
		}
		f[key.ID] = id
	}
	flt := store.Filter{
		Filters:   filters,
		DataField: key.Data,
		Fields: []string{
			key.Email,
			key.ID,
			key.Provider,
			key.Role,
			key.Username,
		},
	}
	var users []*user.User
	err := store.Search(tx, &users, s, flt, p)
	if err != nil {
		return nil, err
	}
	return users, err
}
