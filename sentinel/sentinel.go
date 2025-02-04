package sentinel

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/arkeonetwork/arkeo/common"
	"github.com/arkeonetwork/arkeo/sentinel/conf"

	"github.com/gorilla/handlers"
)

type Proxy struct {
	Metadata   Metadata
	Config     conf.Configuration
	MemStore   *MemStore
	ClaimStore *ClaimStore
	logger     log.Logger
}

func NewProxy(config conf.Configuration) Proxy {
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	claimStore, err := NewClaimStore(config.ClaimStoreLocation)
	if err != nil {
		panic(err)
	}
	return Proxy{
		Metadata:   NewMetadata(config),
		Config:     config,
		MemStore:   NewMemStore(config.SourceChain, logger),
		ClaimStore: claimStore,
		logger:     logger,
	}
}

// Given a request send it to the appropriate url
func (p Proxy) handleRequestAndRedirect(w http.ResponseWriter, r *http.Request) {
	// remove arkauth query arg
	values := r.URL.Query()
	values.Del(QueryArkAuth)
	r.URL.RawQuery = values.Encode()

	parts := strings.Split(r.URL.Path, "/")
	host := parts[1]
	parts = append(parts[:1], parts[1+1:]...)
	r.URL.Path = strings.Join(parts, "/")

	switch host { // nolint
	case "btc-mainnet-fullnode":
		// TODO
	case "eth-mainnet-fullnode":
		// TODO
	case "arkeo-mainnet-fullnode":
		// we forbid arkeo-mainnet-fullnode see chain.go:L50
		// TODO
	case "gaia-mainnet-rpc-archive":
		gaiaHost := p.Config.GaiaRpcArchiveHost
		gaiaHostUrl, err := url.Parse(gaiaHost)
		if err != nil {
			respondWithError(w, "failed to parse backend url", http.StatusInternalServerError)
			return
		}
		r.URL.Scheme = gaiaHostUrl.Scheme
		r.URL.Host = gaiaHostUrl.Host
		r.URL.Path = gaiaHostUrl.Path
	}

	// Serve a reverse proxy for a given url
	// create the reverse proxy
	proxy := common.NewSingleHostReverseProxy(r.URL)

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(w, r)
}

func (p Proxy) handleMetadata(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")

	d, _ := json.Marshal(p.Metadata)
	_, _ = w.Write(d)
}

func (p Proxy) handleOpenClaims(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")

	open_claims := make([]Claim, 0)
	for _, claim := range p.ClaimStore.List() {
		if claim.Claimed {
			p.logger.Info("already claimed")
			continue
		}
		contract, err := p.MemStore.Get(claim.Key())
		if err != nil {
			p.logger.Error("bad fetch")
			continue
		}

		if contract.IsClose(p.MemStore.GetHeight()) {
			_ = p.ClaimStore.Remove(claim.Key()) // clear expired
			p.logger.Info("claim expired")
			continue
		}

		open_claims = append(open_claims, claim)

	}

	d, _ := json.Marshal(open_claims)
	_, _ = w.Write(d)
}

func (p Proxy) handleContract(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	path := r.URL.Path

	parts := strings.Split(path, "/")
	if len(parts) < 5 {
		respondWithError(w, "not enough parameters", http.StatusBadRequest)
		return
	}

	providerPK, err := common.NewPubKey(parts[2])
	if err != nil {
		p.logger.Error("fail to parse provider pubkey", "error", err, "pubkey", parts[2])
		respondWithError(w, fmt.Sprintf("bad provider pubkey: %s", err), http.StatusBadRequest)
		return
	}

	chain, err := common.NewChain(parts[3])
	if err != nil {
		p.logger.Error("fail to parse chain", "error", err, "chain", parts[3])
		respondWithError(w, fmt.Sprintf("bad provider pubkey: %s", err), http.StatusBadRequest)
		return
	}

	spenderPK, err := common.NewPubKey(parts[4])
	if err != nil {
		p.logger.Error("fail to parse spender pubkey", "error", err, "chain", parts[4])
		respondWithError(w, "Invalid spender pubkey", http.StatusBadRequest)
		return
	}

	key := p.MemStore.Key(providerPK.String(), chain.String(), spenderPK.String())
	contract, err := p.MemStore.Get(key)
	if err != nil {
		p.logger.Error("fail to get contract from memstore", "error", err, "key", key)
		respondWithError(w, fmt.Sprintf("fetch contract error: %s", err), http.StatusBadRequest)
		return
	}

	d, _ := json.Marshal(contract)
	_, _ = w.Write(d)
}

func (p Proxy) handleClaim(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	path := r.URL.Path

	parts := strings.Split(path, "/")
	if len(parts) < 5 {
		respondWithError(w, "not enough parameters", http.StatusBadRequest)
		return
	}

	providerPK, err := common.NewPubKey(parts[2])
	if err != nil {
		p.logger.Error("fail to parse provider pubkey", "error", err, "pubkey", parts[2])
		respondWithError(w, fmt.Sprintf("bad provider pubkey: %s", err), http.StatusBadRequest)
		return
	}

	chain, err := common.NewChain(parts[3])
	if err != nil {
		p.logger.Error("fail to parse chain", "error", err, "chain", parts[3])
		respondWithError(w, fmt.Sprintf("bad provider pubkey: %s", err), http.StatusBadRequest)
		return
	}

	spenderPK, err := common.NewPubKey(parts[4])
	if err != nil {
		p.logger.Error("fail to parse spender pubkey", "error", err, "chain", parts[4])
		respondWithError(w, "Invalid spender pubkey", http.StatusBadRequest)
		return
	}

	claim := NewClaim(providerPK, chain, spenderPK, 0, 0, "")
	claim, err = p.ClaimStore.Get(claim.Key())
	if err != nil {
		p.logger.Error("fail to get contract from memstore", "error", err, "key", claim.Key())
		respondWithError(w, fmt.Sprintf("fetch contract error: %s", err), http.StatusBadRequest)
		return
	}

	d, _ := json.Marshal(claim)
	_, _ = w.Write(d)
}

func (p Proxy) Run() {
	p.logger.Info("Starting Sentinel (reverse proxy)....")
	p.Config.Print()

	go p.EventListener(p.Config.EventStreamHost)

	mux := http.NewServeMux()

	// start server
	mux.Handle("/metadata.json", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(p.handleMetadata)))
	mux.Handle("/contract/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(p.handleContract)))
	mux.Handle("/claim/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(p.handleClaim)))
	mux.Handle("/open_claims/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(p.handleOpenClaims)))
	mux.Handle("/", p.auth(
		handlers.LoggingHandler(
			os.Stdout,
			handlers.ProxyHeaders(
				http.HandlerFunc(p.handleRequestAndRedirect),
			),
		),
	))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", p.Config.Port), mux); err != nil {
		panic(err)
	}
}

func respondWithError(w http.ResponseWriter, message string, code int) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		response = []byte(`{"error": "failed to marshal response payload"}`)
		code = http.StatusInternalServerError
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}
