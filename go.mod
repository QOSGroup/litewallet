module github.com/QOSGroup/litewallet

go 1.12

require (
	github.com/allegro/bigcache v1.2.1 // indirect
	github.com/aristanetworks/goarista v0.0.0-20190704150520-f44d68189fd7 // indirect
	github.com/btcsuite/btcd v0.0.0-20190629003639-c26ffa870fd8
	github.com/btcsuite/btcutil v0.0.0-20190425235716-9e5f4b9a998d
	github.com/cespare/cp v1.1.1 // indirect
	github.com/cosmos/cosmos-sdk v0.34.1
	github.com/cosmos/go-bip39 v0.0.0-20180819234021-555e2067c45d
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/edsrzf/mmap-go v1.0.0 // indirect
	github.com/ethereum/go-ethereum v1.9.18
	github.com/fjl/memsize v0.0.0-20180929194037-2a09253e352a // indirect
	github.com/google/uuid v1.0.0
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/olekukonko/tablewriter v0.0.2-0.20190409134802-7e037d187b0c
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.4.0
	github.com/status-im/keycard-go v0.0.0-20190424133014-d95853db0f48 // indirect
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15
	github.com/tendermint/ed25519 v0.0.0-20171027050219-d8387025d2b9
	github.com/tendermint/go-amino v0.14.1
	github.com/tendermint/tendermint v0.31.4
	github.com/tyler-smith/go-bip32 v1.0.0
	github.com/tyler-smith/go-bip39 v1.0.1-0.20181017060643-dbb3b84ba2ef
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
)

replace golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5
