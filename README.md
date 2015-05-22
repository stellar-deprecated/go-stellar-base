# Go Stellar Base
[![Build Status](https://travis-ci.org/stellar/go-stellar-base.svg?branch=master)](https://travis-ci.org/stellar/go-stellar-base)

*STATUS:  This library is currently pre-alpha.  It has no support for reading/writing xdr, but can sign and hash byte slices in accordance with the stellar protocol*

The stellar-base library is the lowest-level stellar helper library.  It consists of classes
to read, write, hash, and sign the xdr structures that are used in stellar-core.

## Installation


```shell
go get github.com/stellar/go-stellar-base
```

## Usage

Let's first decode a transaction (taken from stellar-core's `txhistory` table):

```go
func ExampleDecodeTransaction() {
	data := "rqN6LeOagjxMaUP96Bzfs9e0corNZXzBWJkFoK7kvkwAAAAKAAAAAwAAAAEAAAAAAAA" +
		"AAAAAAAEAAAAAAAAAAW5oJtVdnYOVdZqtXpTHBtbcY0mCmfcBIKEgWnlvFIhaAAAAAA" +
		"AAAAAC+vCAAAAAAa6jei0gQGmrUfm+o2CMv/w32YzJgGYlmgG6CUW3FwyD6AZ/5TtPZ" +
		"qEt9kyC3GJeXfzoS667ZPhPUSNjSWgAeDPHFLcM"

	rawr := strings.NewReader(data)
	b64r := base64.NewDecoder(base64.StdEncoding, rawr)

	var tx xdr.TransactionEnvelope
	bytesRead, err := xdr.Unmarshal(b64r, &tx)

	fmt.Printf("read %d bytes\n", bytesRead)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("This tx has %d operations\n", len(tx.Tx.Operations))
	// Output: read 180 bytes
	// This tx has 1 operations
}
```

## Contributing

Please [see CONTRIBUTING.md for details](CONTRIBUTING.md).
