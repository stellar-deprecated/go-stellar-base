package stellarcore

//go:generate bundle exec xdrgen -l go -o . -n stellarcore xdr/Stellar-types.x
//go:generate bundle exec xdrgen -l go -o . -n stellarcore xdr/Stellar-ledger-entries.x
//go:generate go fmt
