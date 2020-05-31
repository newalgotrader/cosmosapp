package app

import (
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

// passed to baseapp, returned by app.Name()
const appName = "app"

// cosmos app built on baseapp
type App struct {
	*bam.BaseApp
}

// return new app
func NewApp(logger log.Logger, db dbm.DB) *App {
	cdc := MakeCodec()
	txDecoder := auth.DefaultTxDecoder(cdc)

	bApp := bam.NewBaseApp(appName, logger, db, txDecoder)

	app := &App{BaseApp: bApp}

	return app
}

// create a codec will all required registrations for the application
func MakeCodec() *codec.Codec {
	var cdc = codec.New()

	//ModuleBasics.RegisterCodec(cdc)
	//vesting.RegisterCodec(cdc)
	//sdk.RegisterCodec(cdc)
	//codec.RegisterCrypto(cdc)
	return cdc.Seal()

}
