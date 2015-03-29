package stellarcore

//go:generate bundle exec xdrgen -l go -o ./xdr -n xdr xdr/Stellar-types.x
//go:generate go fmt ./xdr
