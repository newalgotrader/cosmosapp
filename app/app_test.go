package app

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

func TestMakeCodec(t *testing.T) {
	cdc := MakeCodec()

	require.Panics(t,
		func() { cdc.RegisterInterface((*interface{})(nil), nil) },
		"panics when new registrations are attempted (cdc is sealed)")
}

func TestApp(t *testing.T) {
	logger := log.NewNopLogger()
	db := dbm.NewMemDB()

	app := NewApp(logger, db)

	require.Equal(t, app.Name(), "app", "name is set to 'app'")
	//require.Equal(t, app.AppVersion(), "0.0.1", "version is 0.0.1")
}
