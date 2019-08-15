package http

import (
	// stdlib
	"os"
	"log"
	"ioutil"
	"net/http"
	"encoding/json"
	// universe
	"github.com/universelabs/universe-core/universe"
	// deps
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Handles HTTP requests for the REST API to KeyManager
type KeyManagerHandler struct {
	// router
	*chi.Mux
	// services
	KeyManager universe.KeyManager
	// utilities
	Logger *log.Logger
}

// Returns a new instance of KeyManagerHandler
func NewKeyManagerHandler(km universe.KeyManager) *KeyManagerHandler {
	kmh := &KeyManagerHandler{
		Mux: chi.NewRouter(),
		KeyManager: km,
		Logger: log.New(os.Stderr, "[Key Manager] ", log.LstdFlags),
	}
	kmh.Post("/addwallet", kmh.AddWallet)
	kmh.Get("/signtx/{walletID}", kmh.SignTx)

	// print all routes
	walkFunc := func(method, route string, handler http.Handler, 
		middlewares ...func(http.Handler) http.Handler) error {
			log.Printf("[Key Manager Handler] %s -> %s\n", route, method)
			return nil
	}
	if err := chi.Walk(kmh.Mux, walkFunc); err != nil {
		log.Panicf("[Key Manager Handler] Logging error: %s\n", err.Error()) 
	}

	return kmh	
}

// Handle AddWallet request
func (kmh *KeyManagerHandler) AddWallet(w http.ResponseWriter, r *http.Request) {
	// unmarshal from json
	wallet := universe.Wallet{}
	err := json.NewDecoder(r.Body).Decode(&wallet)
	if err != nil {
		kmh.Error(w, r, err)
		return
	} 

	// add to KeyManager
	err = ksh.KeyManager.AddWallet(&wallet)
	// return err if failed, else confirmation
	if err != nil {
		kmh.Error(w, r, err)
	} else {
		log.Printf("AddWallet: %v", wallet)
		render.JSON(w, r, wallet)
	}
}

// Handle SignTx request. Sign the body with the key specified by {walletID}
// TODO: Make sure the requesting app has access to WalletID
func (kmh *KeyManagerHandler) SignTx(w http.ResponseWriter, r *http.Request) {
	// get wallet id from url
	walletID, urlerr := strconv.Atoi(chi.URLParam(r, "walletID"))
	if urlerr != nil {
		kmh.Error(w, r, urlerr)
		return
	}
	// check that wallet exists & app has permission to it
	 
	// get tx from http request
	tx, txerr := ioutil.ReadAll(r.Body)
	if txerr != nil {
		kmh.Error(w, r, txerr)
		return
	}

	// sign the tx
	signedTx, kmerr = kmh.KeyManager.SignTx(walletID, tx)
	if kmerr != nil {
		kmh.Error(w, r, kmerr)
	} else {
		log.Printf("Signed TX using wallet %v", walletID)
		render.JSON(w, r, signedTx)
	}
}

// Reports errors
func (kmh *KeyManagerHandler) Error(w http.ResponseWriter, r *http.Request, err Error) {
	log.Println(err)
	render.JSON(w, r, err)
}