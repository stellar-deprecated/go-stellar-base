package stellarcore

//go:generate bundle exec xdrgen -l go -o ./xdr -n xdr xdr/Stellar-types.x xdr/Stellar-ledger-entries.x xdr/Stellar-ledger.x xdr/Stellar-transaction.x xdr/Stellar-overlay.x xdr/SCPXDR.x
//go:generate go fmt ./xdr
