package testhelpers

import (
	"fmt"
	"testing"

	"github.com/gitznik/wazzup/ent"
	"github.com/gitznik/wazzup/ent/enttest"
)

func DBClient(t *testing.T) (*ent.Client, string) {
	dsn := fmt.Sprintf("file:%s?mode=memory&_fk=1", t.Name())
	return enttest.Open(t, "sqlite3", dsn), dsn
}
