package universe

// A client with which to interact with the KeyManager service
type KeyManagerClient interface {
	KeyManager() KeyManager
}
	
// KeyManager service interface
type KeyManager interface {
	AddWallet(new_wallet *Wallet) error
	SignTx(id int, tx []byte) ([]byte, error)
}
