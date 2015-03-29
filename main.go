package stellarcore

//go:generate bundle exec xdrgen -l go -o ./xdr -n xdr xdr/Stellar-types.x
//go:generate bundle exec xdrgen -l go -o ./xdr -n xdr xdr/Stellar-ledger-entries.x
//go:generate bundle exec xdrgen -l go -o ./xdr -n xdr xdr/Stellar-overlay.x
//go:generate bundle exec xdrgen -l go -o ./xdr -n xdr xdr/SCPXDR.x
//go:generate go fmt ./xdr
