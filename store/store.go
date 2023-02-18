package store

import (
	"github.com/ngoduykhanh/wireguard-ui/model"
)

type IStore interface {
	Init() error
	GetUser() (model.User, error)
	SaveUser(user model.User) error
	GetGlobalSettings() (model.GlobalSetting, error)
	GetServer() (model.Server, error)
	GetClients(hasQRCode bool) ([]model.ClientData, error)
	GetClientByID(clientID string, qrCode model.QRCodeSettings) (model.ClientData, error)
	SaveClient(client model.Client) error
	DeleteClient(clientID string) error
	SaveServerInterface(serverInterface model.ServerInterface) error
	SaveServerKeyPair(serverKeyPair model.ServerKeypair) error
	SaveGlobalSettings(globalSettings model.GlobalSetting) error
	GetWakeOnLanHosts() ([]model.WakeOnLanHost, error)
	GetWakeOnLanHost(macAddress string) (*model.WakeOnLanHost, error)
	DeleteWakeOnHostLanHost(macAddress string) error
	SaveWakeOnLanHost(host model.WakeOnLanHost) error
	DeleteWakeOnHost(host model.WakeOnLanHost) error
	GetBrandName() string
	SetBrandName(brandName string) error
	GetPath() string
}
