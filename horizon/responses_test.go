package horizon

var accountResponse = `{
  "_links": {
    "self": {
      "href": "https://horizon-testnet.stellar.org/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H"
    },
    "transactions": {
      "href": "https://horizon-testnet.stellar.org/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H/transactions{?cursor,limit,order}",
      "templated": true
    },
    "operations": {
      "href": "https://horizon-testnet.stellar.org/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H/operations{?cursor,limit,order}",
      "templated": true
    },
    "payments": {
      "href": "https://horizon-testnet.stellar.org/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H/payments{?cursor,limit,order}",
      "templated": true
    },
    "effects": {
      "href": "https://horizon-testnet.stellar.org/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H/effects{?cursor,limit,order}",
      "templated": true
    },
    "offers": {
      "href": "https://horizon-testnet.stellar.org/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H/Offers{?cursor,limit,order}",
      "templated": true
    }
  },
  "id": "GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H",
  "paging_token": "1",
  "account_id": "GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H",
  "sequence": "7384",
  "subentry_count": 0,
  "thresholds": {
    "low_threshold": 0,
    "med_threshold": 0,
    "high_threshold": 0
  },
  "flags": {
    "auth_required": false,
    "auth_revocable": false
  },
  "balances": [
    {
      "balance": "948522307.6146000",
      "asset_type": "native"
    }
  ],
  "signers": [
    {
      "public_key": "GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H",
      "weight": 1
    }
  ],
  "data": {
	"foo": "+xbxLlS9Exgb4n6tWSK6ruUmejywOykOHw1zCrotEgtNHeBzykVmdMhPipUOI839q1tybb9NUkrsteMoJas1/w=="
  }
}`

var accountsResponse = `{
  "_links": {
    "self": {
      "href": "https://horizon-testnet.stellar.org/accounts?order=asc\u0026limit=10\u0026cursor="
    },
    "next": {
      "href": "https://horizon-testnet.stellar.org/accounts?order=asc\u0026limit=10\u0026cursor=8388071133185"
    },
    "prev": {
      "href": "https://horizon-testnet.stellar.org/accounts?order=desc\u0026limit=10\u0026cursor=1"
    }
  },
  "_embedded": {
    "records": [
      {
        "id": "GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H",
        "paging_token": "1",
        "account_id": "GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H"
      },
      {
        "id": "GBS43BF24ENNS3KPACUZVKK2VYPOZVBQO2CISGZ777RYGOPYC2FT6S3K",
        "paging_token": "4672924422145",
        "account_id": "GBS43BF24ENNS3KPACUZVKK2VYPOZVBQO2CISGZ777RYGOPYC2FT6S3K"
      },
      {
        "id": "GDUPMP7J7FFPB5ACNLOOAFQ45GNDZH5IXKE5ID3YS5P4F26X7BKVXFCA",
        "paging_token": "5037996642305",
        "account_id": "GDUPMP7J7FFPB5ACNLOOAFQ45GNDZH5IXKE5ID3YS5P4F26X7BKVXFCA"
      },
      {
        "id": "GCQ6Y2L3KJ5DK3GQ2QYWNO2JFQLLKF5HEKP67PM3NANANESGXSJGG3VY",
        "paging_token": "6859062775809",
        "account_id": "GCQ6Y2L3KJ5DK3GQ2QYWNO2JFQLLKF5HEKP67PM3NANANESGXSJGG3VY"
      },
      {
        "id": "GCYRBZHGBN7A5HA5EFFPATCIZQ3OZXXXWXGY4QU7IMSAVH3FC6TMGSO6",
        "paging_token": "7189775257601",
        "account_id": "GCYRBZHGBN7A5HA5EFFPATCIZQ3OZXXXWXGY4QU7IMSAVH3FC6TMGSO6"
      },
      {
        "id": "GBWLLK6C5HF6PZQ7HKSWHCXH355AFF55SPEYYM7BMDHKO2PIT66QP7GJ",
        "paging_token": "7881264992257",
        "account_id": "GBWLLK6C5HF6PZQ7HKSWHCXH355AFF55SPEYYM7BMDHKO2PIT66QP7GJ"
      },
      {
        "id": "GC3AMRZPOWP2VA4JA3CNV23G3KW7FU6GRBYIGJC5HE5ZRET6ILBPP7TY",
        "paging_token": "7919919697921",
        "account_id": "GC3AMRZPOWP2VA4JA3CNV23G3KW7FU6GRBYIGJC5HE5ZRET6ILBPP7TY"
      },
      {
        "id": "GAB3OKD24EJ4VO22N7EEQTVDVKX6JX5SKLBJW63GLAW27XEHNP5L7AVC",
        "paging_token": "8315056689153",
        "account_id": "GAB3OKD24EJ4VO22N7EEQTVDVKX6JX5SKLBJW63GLAW27XEHNP5L7AVC"
      },
      {
        "id": "GBJFTW3H7FA7HZY2W4LX6PTU6W2HH3KXTUCDYIEPB373T4VROUQAK6CN",
        "paging_token": "8358006362113",
        "account_id": "GBJFTW3H7FA7HZY2W4LX6PTU6W2HH3KXTUCDYIEPB373T4VROUQAK6CN"
      },
      {
        "id": "GBLZWJINHGQ4YLBCQVVI6EGUH3OH63KPA2RMBF6RNCPS2IF5GCGF4AZO",
        "paging_token": "8388071133185",
        "account_id": "GBLZWJINHGQ4YLBCQVVI6EGUH3OH63KPA2RMBF6RNCPS2IF5GCGF4AZO"
      }
    ]
  }
}`

var effectsResponse = `{
  "_links": {
    "self": {
      "href": "https://horizon-testnet.stellar.org/effects?order=desc\u0026limit=3\u0026cursor="
    },
    "next": {
      "href": "https://horizon-testnet.stellar.org/effects?order=desc\u0026limit=3\u0026cursor=5987652562063361-2"
    },
    "prev": {
      "href": "https://horizon-testnet.stellar.org/effects?order=asc\u0026limit=3\u0026cursor=5987652562063361-4"
    }
  },
  "_embedded": {
    "records": [
      {
        "_links": {
          "operation": {
            "href": "https://horizon-testnet.stellar.org/operations/5987652562063361"
          },
          "succeeds": {
            "href": "https://horizon-testnet.stellar.org/effects?order=desc\u0026cursor=5987652562063361-4"
          },
          "precedes": {
            "href": "https://horizon-testnet.stellar.org/effects?order=asc\u0026cursor=5987652562063361-4"
          }
        },
        "id": "0005987652562063361-0000000004",
        "paging_token": "5987652562063361-4",
        "account": "GBLZWJINHGQ4YLBCQVVI6EGUH3OH63KPA2RMBF6RNCPS2IF5GCGF4AZO",
        "type": "trade",
        "type_i": 33,
        "seller": "GC3AMRZPOWP2VA4JA3CNV23G3KW7FU6GRBYIGJC5HE5ZRET6ILBPP7TY",
        "offer_id": 17,
        "sold_amount": "145.1700000",
        "sold_asset_type": "credit_alphanum4",
        "sold_asset_code": "ZAR",
        "sold_asset_issuer": "GBQMCYLX7NWGVRETHIO2P3U2UMFUBST3QSVTFM3FCZFDWZ7NYTFUHHLJ",
        "bought_amount": "10.0196334",
        "bought_asset_type": "credit_alphanum4",
        "bought_asset_code": "USD",
        "bought_asset_issuer": "GBWLLK6C5HF6PZQ7HKSWHCXH355AFF55SPEYYM7BMDHKO2PIT66QP7GJ"
      },
      {
        "_links": {
          "operation": {
            "href": "https://horizon-testnet.stellar.org/operations/5987652562063361"
          },
          "succeeds": {
            "href": "https://horizon-testnet.stellar.org/effects?order=desc\u0026cursor=5987652562063361-3"
          },
          "precedes": {
            "href": "https://horizon-testnet.stellar.org/effects?order=asc\u0026cursor=5987652562063361-3"
          }
        },
        "id": "0005987652562063361-0000000003",
        "paging_token": "5987652562063361-3",
        "account": "GC3AMRZPOWP2VA4JA3CNV23G3KW7FU6GRBYIGJC5HE5ZRET6ILBPP7TY",
        "type": "trade",
        "type_i": 33,
        "seller": "GBLZWJINHGQ4YLBCQVVI6EGUH3OH63KPA2RMBF6RNCPS2IF5GCGF4AZO",
        "offer_id": 17,
        "sold_amount": "10.0196334",
        "sold_asset_type": "credit_alphanum4",
        "sold_asset_code": "USD",
        "sold_asset_issuer": "GBWLLK6C5HF6PZQ7HKSWHCXH355AFF55SPEYYM7BMDHKO2PIT66QP7GJ",
        "bought_amount": "145.1700000",
        "bought_asset_type": "credit_alphanum4",
        "bought_asset_code": "ZAR",
        "bought_asset_issuer": "GBQMCYLX7NWGVRETHIO2P3U2UMFUBST3QSVTFM3FCZFDWZ7NYTFUHHLJ"
      },
      {
        "_links": {
          "operation": {
            "href": "https://horizon-testnet.stellar.org/operations/5987652562063361"
          },
          "succeeds": {
            "href": "https://horizon-testnet.stellar.org/effects?order=desc\u0026cursor=5987652562063361-2"
          },
          "precedes": {
            "href": "https://horizon-testnet.stellar.org/effects?order=asc\u0026cursor=5987652562063361-2"
          }
        },
        "id": "0005987652562063361-0000000002",
        "paging_token": "5987652562063361-2",
        "account": "GC3AMRZPOWP2VA4JA3CNV23G3KW7FU6GRBYIGJC5HE5ZRET6ILBPP7TY",
        "type": "account_debited",
        "type_i": 3,
        "asset_type": "credit_alphanum4",
        "asset_code": "USD",
        "asset_issuer": "GBWLLK6C5HF6PZQ7HKSWHCXH355AFF55SPEYYM7BMDHKO2PIT66QP7GJ",
        "amount": "10.0196334"
      }
    ]
  }
}`

var ledgerResponse = `{
  "_links": {
    "self": {
      "href": "https://horizon-testnet.stellar.org/ledgers/1"
    },
    "transactions": {
      "href": "https://horizon-testnet.stellar.org/ledgers/1/transactions{?cursor,limit,order}",
      "templated": true
    },
    "operations": {
      "href": "https://horizon-testnet.stellar.org/ledgers/1/operations{?cursor,limit,order}",
      "templated": true
    },
    "payments": {
      "href": "https://horizon-testnet.stellar.org/ledgers/1/payments{?cursor,limit,order}",
      "templated": true
    },
    "effects": {
      "href": "https://horizon-testnet.stellar.org/ledgers/1/effects{?cursor,limit,order}",
      "templated": true
    }
  },
  "id": "63d98f536ee68d1b27b5b89f23af5311b7569a24faf1403ad0b52b633b07be99",
  "paging_token": "4294967296",
  "hash": "63d98f536ee68d1b27b5b89f23af5311b7569a24faf1403ad0b52b633b07be99",
  "sequence": 1,
  "transaction_count": 0,
  "operation_count": 0,
  "closed_at": "1970-01-01T00:00:00Z",
  "total_coins": "100000000000.0000000",
  "fee_pool": "0.0000000",
  "base_fee": 100,
  "base_reserve": "10.0000000",
  "max_tx_set_size": 100
}`

var ledgersResponse = `{
  "_links": {
    "self": {
      "href": "https://horizon-testnet.stellar.org/ledgers?order=asc\u0026limit=10\u0026cursor="
    },
    "next": {
      "href": "https://horizon-testnet.stellar.org/ledgers?order=asc\u0026limit=10\u0026cursor=42949672960"
    },
    "prev": {
      "href": "https://horizon-testnet.stellar.org/ledgers?order=desc\u0026limit=10\u0026cursor=4294967296"
    }
  },
  "_embedded": {
    "records": [
      {
        "_links": {
          "self": {
            "href": "https://horizon-testnet.stellar.org/ledgers/1"
          },
          "transactions": {
            "href": "https://horizon-testnet.stellar.org/ledgers/1/transactions{?cursor,limit,order}",
            "templated": true
          },
          "operations": {
            "href": "https://horizon-testnet.stellar.org/ledgers/1/operations{?cursor,limit,order}",
            "templated": true
          },
          "payments": {
            "href": "https://horizon-testnet.stellar.org/ledgers/1/payments{?cursor,limit,order}",
            "templated": true
          },
          "effects": {
            "href": "https://horizon-testnet.stellar.org/ledgers/1/effects{?cursor,limit,order}",
            "templated": true
          }
        },
        "id": "63d98f536ee68d1b27b5b89f23af5311b7569a24faf1403ad0b52b633b07be99",
        "paging_token": "4294967296",
        "hash": "63d98f536ee68d1b27b5b89f23af5311b7569a24faf1403ad0b52b633b07be99",
        "sequence": 1,
        "transaction_count": 0,
        "operation_count": 0,
        "closed_at": "1970-01-01T00:00:00Z",
        "total_coins": "100000000000.0000000",
        "fee_pool": "0.0000000",
        "base_fee": 100,
        "base_reserve": "10.0000000",
        "max_tx_set_size": 100
      },
      {
        "_links": {
          "self": {
            "href": "https://horizon-testnet.stellar.org/ledgers/2"
          },
          "transactions": {
            "href": "https://horizon-testnet.stellar.org/ledgers/2/transactions{?cursor,limit,order}",
            "templated": true
          },
          "operations": {
            "href": "https://horizon-testnet.stellar.org/ledgers/2/operations{?cursor,limit,order}",
            "templated": true
          },
          "payments": {
            "href": "https://horizon-testnet.stellar.org/ledgers/2/payments{?cursor,limit,order}",
            "templated": true
          },
          "effects": {
            "href": "https://horizon-testnet.stellar.org/ledgers/2/effects{?cursor,limit,order}",
            "templated": true
          }
        },
        "id": "a08da49689fbe0f637063a75bdbb0e8f4a3ebc952769af8a1725d8577763860c",
        "paging_token": "8589934592",
        "hash": "a08da49689fbe0f637063a75bdbb0e8f4a3ebc952769af8a1725d8577763860c",
        "prev_hash": "63d98f536ee68d1b27b5b89f23af5311b7569a24faf1403ad0b52b633b07be99",
        "sequence": 2,
        "transaction_count": 0,
        "operation_count": 0,
        "closed_at": "2016-04-08T17:09:17Z",
        "total_coins": "100000000000.0000000",
        "fee_pool": "0.0000000",
        "base_fee": 100,
        "base_reserve": "10.0000000",
        "max_tx_set_size": 50
      },
      {
        "_links": {
          "self": {
            "href": "https://horizon-testnet.stellar.org/ledgers/3"
          },
          "transactions": {
            "href": "https://horizon-testnet.stellar.org/ledgers/3/transactions{?cursor,limit,order}",
            "templated": true
          },
          "operations": {
            "href": "https://horizon-testnet.stellar.org/ledgers/3/operations{?cursor,limit,order}",
            "templated": true
          },
          "payments": {
            "href": "https://horizon-testnet.stellar.org/ledgers/3/payments{?cursor,limit,order}",
            "templated": true
          },
          "effects": {
            "href": "https://horizon-testnet.stellar.org/ledgers/3/effects{?cursor,limit,order}",
            "templated": true
          }
        },
        "id": "37f31dd1ee38d6bd1c15a99ab8d8eab6de081631cd66dc2ebe5a89669af802c6",
        "paging_token": "12884901888",
        "hash": "37f31dd1ee38d6bd1c15a99ab8d8eab6de081631cd66dc2ebe5a89669af802c6",
        "prev_hash": "a08da49689fbe0f637063a75bdbb0e8f4a3ebc952769af8a1725d8577763860c",
        "sequence": 3,
        "transaction_count": 0,
        "operation_count": 0,
        "closed_at": "2016-04-08T17:09:41Z",
        "total_coins": "100000000000.0000000",
        "fee_pool": "0.0000000",
        "base_fee": 100,
        "base_reserve": "10.0000000",
        "max_tx_set_size": 50
      },
      {
        "_links": {
          "self": {
            "href": "https://horizon-testnet.stellar.org/ledgers/4"
          },
          "transactions": {
            "href": "https://horizon-testnet.stellar.org/ledgers/4/transactions{?cursor,limit,order}",
            "templated": true
          },
          "operations": {
            "href": "https://horizon-testnet.stellar.org/ledgers/4/operations{?cursor,limit,order}",
            "templated": true
          },
          "payments": {
            "href": "https://horizon-testnet.stellar.org/ledgers/4/payments{?cursor,limit,order}",
            "templated": true
          },
          "effects": {
            "href": "https://horizon-testnet.stellar.org/ledgers/4/effects{?cursor,limit,order}",
            "templated": true
          }
        },
        "id": "0211746e38fc1de84e5f26c965cc9ba5868671294657d90b1e6d320ab290905e",
        "paging_token": "17179869184",
        "hash": "0211746e38fc1de84e5f26c965cc9ba5868671294657d90b1e6d320ab290905e",
        "prev_hash": "37f31dd1ee38d6bd1c15a99ab8d8eab6de081631cd66dc2ebe5a89669af802c6",
        "sequence": 4,
        "transaction_count": 0,
        "operation_count": 0,
        "closed_at": "2016-04-08T17:09:43Z",
        "total_coins": "100000000000.0000000",
        "fee_pool": "0.0000000",
        "base_fee": 100,
        "base_reserve": "10.0000000",
        "max_tx_set_size": 50
      },
      {
        "_links": {
          "self": {
            "href": "https://horizon-testnet.stellar.org/ledgers/5"
          },
          "transactions": {
            "href": "https://horizon-testnet.stellar.org/ledgers/5/transactions{?cursor,limit,order}",
            "templated": true
          },
          "operations": {
            "href": "https://horizon-testnet.stellar.org/ledgers/5/operations{?cursor,limit,order}",
            "templated": true
          },
          "payments": {
            "href": "https://horizon-testnet.stellar.org/ledgers/5/payments{?cursor,limit,order}",
            "templated": true
          },
          "effects": {
            "href": "https://horizon-testnet.stellar.org/ledgers/5/effects{?cursor,limit,order}",
            "templated": true
          }
        },
        "id": "d7ff342606ce24dfaa1eae60cd9d0f4ddcda64e6c6ca231feecd88679697ba46",
        "paging_token": "21474836480",
        "hash": "d7ff342606ce24dfaa1eae60cd9d0f4ddcda64e6c6ca231feecd88679697ba46",
        "prev_hash": "0211746e38fc1de84e5f26c965cc9ba5868671294657d90b1e6d320ab290905e",
        "sequence": 5,
        "transaction_count": 0,
        "operation_count": 0,
        "closed_at": "2016-04-08T17:09:51Z",
        "total_coins": "100000000000.0000000",
        "fee_pool": "0.0000000",
        "base_fee": 100,
        "base_reserve": "10.0000000",
        "max_tx_set_size": 50
      },
      {
        "_links": {
          "self": {
            "href": "https://horizon-testnet.stellar.org/ledgers/6"
          },
          "transactions": {
            "href": "https://horizon-testnet.stellar.org/ledgers/6/transactions{?cursor,limit,order}",
            "templated": true
          },
          "operations": {
            "href": "https://horizon-testnet.stellar.org/ledgers/6/operations{?cursor,limit,order}",
            "templated": true
          },
          "payments": {
            "href": "https://horizon-testnet.stellar.org/ledgers/6/payments{?cursor,limit,order}",
            "templated": true
          },
          "effects": {
            "href": "https://horizon-testnet.stellar.org/ledgers/6/effects{?cursor,limit,order}",
            "templated": true
          }
        },
        "id": "6b346310f3fd9019652e7dd0795eb623369582f5f5631be2c0ed1c438e086d31",
        "paging_token": "25769803776",
        "hash": "6b346310f3fd9019652e7dd0795eb623369582f5f5631be2c0ed1c438e086d31",
        "prev_hash": "d7ff342606ce24dfaa1eae60cd9d0f4ddcda64e6c6ca231feecd88679697ba46",
        "sequence": 6,
        "transaction_count": 0,
        "operation_count": 0,
        "closed_at": "2016-04-08T17:09:56Z",
        "total_coins": "100000000000.0000000",
        "fee_pool": "0.0000000",
        "base_fee": 100,
        "base_reserve": "10.0000000",
        "max_tx_set_size": 50
      },
      {
        "_links": {
          "self": {
            "href": "https://horizon-testnet.stellar.org/ledgers/7"
          },
          "transactions": {
            "href": "https://horizon-testnet.stellar.org/ledgers/7/transactions{?cursor,limit,order}",
            "templated": true
          },
          "operations": {
            "href": "https://horizon-testnet.stellar.org/ledgers/7/operations{?cursor,limit,order}",
            "templated": true
          },
          "payments": {
            "href": "https://horizon-testnet.stellar.org/ledgers/7/payments{?cursor,limit,order}",
            "templated": true
          },
          "effects": {
            "href": "https://horizon-testnet.stellar.org/ledgers/7/effects{?cursor,limit,order}",
            "templated": true
          }
        },
        "id": "8136a689a22972756c3caffaef73751ea753f12a24b03b25269aa47fb5ffec6b",
        "paging_token": "30064771072",
        "hash": "8136a689a22972756c3caffaef73751ea753f12a24b03b25269aa47fb5ffec6b",
        "prev_hash": "6b346310f3fd9019652e7dd0795eb623369582f5f5631be2c0ed1c438e086d31",
        "sequence": 7,
        "transaction_count": 0,
        "operation_count": 0,
        "closed_at": "2016-04-08T17:10:06Z",
        "total_coins": "100000000000.0000000",
        "fee_pool": "0.0000000",
        "base_fee": 100,
        "base_reserve": "10.0000000",
        "max_tx_set_size": 50
      },
      {
        "_links": {
          "self": {
            "href": "https://horizon-testnet.stellar.org/ledgers/8"
          },
          "transactions": {
            "href": "https://horizon-testnet.stellar.org/ledgers/8/transactions{?cursor,limit,order}",
            "templated": true
          },
          "operations": {
            "href": "https://horizon-testnet.stellar.org/ledgers/8/operations{?cursor,limit,order}",
            "templated": true
          },
          "payments": {
            "href": "https://horizon-testnet.stellar.org/ledgers/8/payments{?cursor,limit,order}",
            "templated": true
          },
          "effects": {
            "href": "https://horizon-testnet.stellar.org/ledgers/8/effects{?cursor,limit,order}",
            "templated": true
          }
        },
        "id": "4a643175f46be5da25ed146d23ec6139ae1eca9f48a577f3ea53bc94633c6437",
        "paging_token": "34359738368",
        "hash": "4a643175f46be5da25ed146d23ec6139ae1eca9f48a577f3ea53bc94633c6437",
        "prev_hash": "8136a689a22972756c3caffaef73751ea753f12a24b03b25269aa47fb5ffec6b",
        "sequence": 8,
        "transaction_count": 0,
        "operation_count": 0,
        "closed_at": "2016-04-08T17:10:11Z",
        "total_coins": "100000000000.0000000",
        "fee_pool": "0.0000000",
        "base_fee": 100,
        "base_reserve": "10.0000000",
        "max_tx_set_size": 50
      },
      {
        "_links": {
          "self": {
            "href": "https://horizon-testnet.stellar.org/ledgers/9"
          },
          "transactions": {
            "href": "https://horizon-testnet.stellar.org/ledgers/9/transactions{?cursor,limit,order}",
            "templated": true
          },
          "operations": {
            "href": "https://horizon-testnet.stellar.org/ledgers/9/operations{?cursor,limit,order}",
            "templated": true
          },
          "payments": {
            "href": "https://horizon-testnet.stellar.org/ledgers/9/payments{?cursor,limit,order}",
            "templated": true
          },
          "effects": {
            "href": "https://horizon-testnet.stellar.org/ledgers/9/effects{?cursor,limit,order}",
            "templated": true
          }
        },
        "id": "5139da28d6d2f80d44e87bf37195277be3136db45c9853bec18bb55a673f423e",
        "paging_token": "38654705664",
        "hash": "5139da28d6d2f80d44e87bf37195277be3136db45c9853bec18bb55a673f423e",
        "prev_hash": "4a643175f46be5da25ed146d23ec6139ae1eca9f48a577f3ea53bc94633c6437",
        "sequence": 9,
        "transaction_count": 0,
        "operation_count": 0,
        "closed_at": "2016-04-08T17:10:16Z",
        "total_coins": "100000000000.0000000",
        "fee_pool": "0.0000000",
        "base_fee": 100,
        "base_reserve": "10.0000000",
        "max_tx_set_size": 50
      },
      {
        "_links": {
          "self": {
            "href": "https://horizon-testnet.stellar.org/ledgers/10"
          },
          "transactions": {
            "href": "https://horizon-testnet.stellar.org/ledgers/10/transactions{?cursor,limit,order}",
            "templated": true
          },
          "operations": {
            "href": "https://horizon-testnet.stellar.org/ledgers/10/operations{?cursor,limit,order}",
            "templated": true
          },
          "payments": {
            "href": "https://horizon-testnet.stellar.org/ledgers/10/payments{?cursor,limit,order}",
            "templated": true
          },
          "effects": {
            "href": "https://horizon-testnet.stellar.org/ledgers/10/effects{?cursor,limit,order}",
            "templated": true
          }
        },
        "id": "651a974af2c39bed94975269aa3f762d4e311728b8075a7a394fadf75db4a2e7",
        "paging_token": "42949672960",
        "hash": "651a974af2c39bed94975269aa3f762d4e311728b8075a7a394fadf75db4a2e7",
        "prev_hash": "5139da28d6d2f80d44e87bf37195277be3136db45c9853bec18bb55a673f423e",
        "sequence": 10,
        "transaction_count": 0,
        "operation_count": 0,
        "closed_at": "2016-04-08T17:10:21Z",
        "total_coins": "100000000000.0000000",
        "fee_pool": "0.0000000",
        "base_fee": 100,
        "base_reserve": "10.0000000",
        "max_tx_set_size": 50
      }
    ]
  }
}`

var transactionResponse = `{
  "_links": {
    "self": {
      "href": "https://horizon-testnet.stellar.org/transactions/e7791b6d3040a09216d4ed696ddc245f1e833e280d8cc2cf7d902c80e9e487f1"
    },
    "account": {
      "href": "https://horizon-testnet.stellar.org/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H"
    },
    "ledger": {
      "href": "https://horizon-testnet.stellar.org/ledgers/1088"
    },
    "operations": {
      "href": "https://horizon-testnet.stellar.org/transactions/e7791b6d3040a09216d4ed696ddc245f1e833e280d8cc2cf7d902c80e9e487f1/operations{?cursor,limit,order}",
      "templated": true
    },
    "effects": {
      "href": "https://horizon-testnet.stellar.org/transactions/e7791b6d3040a09216d4ed696ddc245f1e833e280d8cc2cf7d902c80e9e487f1/effects{?cursor,limit,order}",
      "templated": true
    },
    "precedes": {
      "href": "https://horizon-testnet.stellar.org/transactions?order=asc\u0026cursor=4672924422144"
    },
    "succeeds": {
      "href": "https://horizon-testnet.stellar.org/transactions?order=desc\u0026cursor=4672924422144"
    }
  },
  "id": "e7791b6d3040a09216d4ed696ddc245f1e833e280d8cc2cf7d902c80e9e487f1",
  "paging_token": "4672924422144",
  "hash": "e7791b6d3040a09216d4ed696ddc245f1e833e280d8cc2cf7d902c80e9e487f1",
  "ledger": 1088,
  "created_at": "2016-04-08T18:24:46Z",
  "source_account": "GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H",
  "source_account_sequence": "2",
  "fee_paid": 100,
  "operation_count": 1,
  "envelope_xdr": "AAAAAGL8HQvQkbK2HA3WVjRrKmjX00fG8sLI7m0ERwJW/AX3AAAAZAAAAAAAAAACAAAAAAAAAAEAAAAAAAAAAQAAAAAAAAAAAAAAAGXNhLrhGtltTwCpmqlarh7s1DB2hIkbP//jgzn4Fos/AABa8xB6QAAAAAAAAAAAAVb8BfcAAABAVxWYvw8pmgZLHi5ItJR0R1rzP+a25UFqxojdXhLSMZJgY/JR+oVszD8Dg+YDPeDUdO5GUtYRvlySEBXUskXTDA==",
  "result_xdr": "AAAAAAAAAGQAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAA=",
  "result_meta_xdr": "AAAAAAAAAAEAAAACAAAAAAAABEAAAAAAAAAAAGXNhLrhGtltTwCpmqlarh7s1DB2hIkbP//jgzn4Fos/AABa8xB6QAAAAARAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAQAABEAAAAAAAAAAAGL8HQvQkbK2HA3WVjRrKmjX00fG8sLI7m0ERwJW/AX3DeBbwJbpvzgAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAA",
  "fee_meta_xdr": "AAAAAgAAAAMAAAQUAAAAAAAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9w3gtrOnY/+cAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEAAARAAAAAAAAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9w3gtrOnY/84AAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==",
  "memo_type": "text",
  "signatures": [
    "VxWYvw8pmgZLHi5ItJR0R1rzP+a25UFqxojdXhLSMZJgY/JR+oVszD8Dg+YDPeDUdO5GUtYRvlySEBXUskXTDA=="
  ]
}`

var transactionsResponse = `{
  "_links": {
    "self": {
      "href": "https://horizon-testnet.stellar.org/transactions?order=asc\u0026limit=3\u0026cursor="
    },
    "next": {
      "href": "https://horizon-testnet.stellar.org/transactions?order=asc\u0026limit=3\u0026cursor=6859062775808"
    },
    "prev": {
      "href": "https://horizon-testnet.stellar.org/transactions?order=desc\u0026limit=3\u0026cursor=4672924422144"
    }
  },
  "_embedded": {
    "records": [
      {
        "_links": {
          "self": {
            "href": "https://horizon-testnet.stellar.org/transactions/e7791b6d3040a09216d4ed696ddc245f1e833e280d8cc2cf7d902c80e9e487f1"
          },
          "account": {
            "href": "https://horizon-testnet.stellar.org/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H"
          },
          "ledger": {
            "href": "https://horizon-testnet.stellar.org/ledgers/1088"
          },
          "operations": {
            "href": "https://horizon-testnet.stellar.org/transactions/e7791b6d3040a09216d4ed696ddc245f1e833e280d8cc2cf7d902c80e9e487f1/operations{?cursor,limit,order}",
            "templated": true
          },
          "effects": {
            "href": "https://horizon-testnet.stellar.org/transactions/e7791b6d3040a09216d4ed696ddc245f1e833e280d8cc2cf7d902c80e9e487f1/effects{?cursor,limit,order}",
            "templated": true
          },
          "precedes": {
            "href": "https://horizon-testnet.stellar.org/transactions?order=asc\u0026cursor=4672924422144"
          },
          "succeeds": {
            "href": "https://horizon-testnet.stellar.org/transactions?order=desc\u0026cursor=4672924422144"
          }
        },
        "id": "e7791b6d3040a09216d4ed696ddc245f1e833e280d8cc2cf7d902c80e9e487f1",
        "paging_token": "4672924422144",
        "hash": "e7791b6d3040a09216d4ed696ddc245f1e833e280d8cc2cf7d902c80e9e487f1",
        "ledger": 1088,
        "created_at": "2016-04-08T18:24:46Z",
        "source_account": "GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H",
        "source_account_sequence": "2",
        "fee_paid": 100,
        "operation_count": 1,
        "envelope_xdr": "AAAAAGL8HQvQkbK2HA3WVjRrKmjX00fG8sLI7m0ERwJW/AX3AAAAZAAAAAAAAAACAAAAAAAAAAEAAAAAAAAAAQAAAAAAAAAAAAAAAGXNhLrhGtltTwCpmqlarh7s1DB2hIkbP//jgzn4Fos/AABa8xB6QAAAAAAAAAAAAVb8BfcAAABAVxWYvw8pmgZLHi5ItJR0R1rzP+a25UFqxojdXhLSMZJgY/JR+oVszD8Dg+YDPeDUdO5GUtYRvlySEBXUskXTDA==",
        "result_xdr": "AAAAAAAAAGQAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAA=",
        "result_meta_xdr": "AAAAAAAAAAEAAAACAAAAAAAABEAAAAAAAAAAAGXNhLrhGtltTwCpmqlarh7s1DB2hIkbP//jgzn4Fos/AABa8xB6QAAAAARAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAQAABEAAAAAAAAAAAGL8HQvQkbK2HA3WVjRrKmjX00fG8sLI7m0ERwJW/AX3DeBbwJbpvzgAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAA",
        "fee_meta_xdr": "AAAAAgAAAAMAAAQUAAAAAAAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9w3gtrOnY/+cAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEAAARAAAAAAAAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9w3gtrOnY/84AAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==",
        "memo_type": "text",
        "signatures": [
          "VxWYvw8pmgZLHi5ItJR0R1rzP+a25UFqxojdXhLSMZJgY/JR+oVszD8Dg+YDPeDUdO5GUtYRvlySEBXUskXTDA=="
        ]
      },
      {
        "_links": {
          "self": {
            "href": "https://horizon-testnet.stellar.org/transactions/2d907f006a08b2529b89d9530edbc1765dbed8acede3eeae665ea08afa3effc7"
          },
          "account": {
            "href": "https://horizon-testnet.stellar.org/accounts/GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H"
          },
          "ledger": {
            "href": "https://horizon-testnet.stellar.org/ledgers/1173"
          },
          "operations": {
            "href": "https://horizon-testnet.stellar.org/transactions/2d907f006a08b2529b89d9530edbc1765dbed8acede3eeae665ea08afa3effc7/operations{?cursor,limit,order}",
            "templated": true
          },
          "effects": {
            "href": "https://horizon-testnet.stellar.org/transactions/2d907f006a08b2529b89d9530edbc1765dbed8acede3eeae665ea08afa3effc7/effects{?cursor,limit,order}",
            "templated": true
          },
          "precedes": {
            "href": "https://horizon-testnet.stellar.org/transactions?order=asc\u0026cursor=5037996642304"
          },
          "succeeds": {
            "href": "https://horizon-testnet.stellar.org/transactions?order=desc\u0026cursor=5037996642304"
          }
        },
        "id": "2d907f006a08b2529b89d9530edbc1765dbed8acede3eeae665ea08afa3effc7",
        "paging_token": "5037996642304",
        "hash": "2d907f006a08b2529b89d9530edbc1765dbed8acede3eeae665ea08afa3effc7",
        "ledger": 1173,
        "created_at": "2016-04-08T18:30:15Z",
        "source_account": "GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H",
        "source_account_sequence": "3",
        "fee_paid": 100,
        "operation_count": 1,
        "envelope_xdr": "AAAAAGL8HQvQkbK2HA3WVjRrKmjX00fG8sLI7m0ERwJW/AX3AAAAZAAAAAAAAAADAAAAAAAAAAEAAAAAAAAAAQAAAAAAAAAAAAAAAOj2P+n5SvD0Amrc4BYc6Zo8n6i6idQPeJdfwuvX+FVbAABa8xB6QAAAAAAAAAAAAVb8BfcAAABAgTy8cCVJgEC5duYf8N86owH7O4qUdlAM8KpgZEF8HYN1d4e2aBNlnae/pyHzOlvmk4N2Lh3BD69nubCbaMLfAg==",
        "result_xdr": "AAAAAAAAAGQAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAA=",
        "result_meta_xdr": "AAAAAAAAAAEAAAACAAAAAAAABJUAAAAAAAAAAOj2P+n5SvD0Amrc4BYc6Zo8n6i6idQPeJdfwuvX+FVbAABa8xB6QAAAAASVAAAAAAAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAQAABJUAAAAAAAAAAGL8HQvQkbK2HA3WVjRrKmjX00fG8sLI7m0ERwJW/AX3DeAAzYZvftQAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAA",
        "fee_meta_xdr": "AAAAAgAAAAMAAARAAAAAAAAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9w3gW8CW6b84AAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEAAASVAAAAAAAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9w3gW8CW6b7UAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==",
        "memo_type": "text",
        "signatures": [
          "gTy8cCVJgEC5duYf8N86owH7O4qUdlAM8KpgZEF8HYN1d4e2aBNlnae/pyHzOlvmk4N2Lh3BD69nubCbaMLfAg=="
        ]
      },
      {
        "_links": {
          "self": {
            "href": "https://horizon-testnet.stellar.org/transactions/02a142294acde3e8920fc1e3879b34763df56b9c387a6dcf16d613b21cb9b87a"
          },
          "account": {
            "href": "https://horizon-testnet.stellar.org/accounts/GBS43BF24ENNS3KPACUZVKK2VYPOZVBQO2CISGZ777RYGOPYC2FT6S3K"
          },
          "ledger": {
            "href": "https://horizon-testnet.stellar.org/ledgers/1597"
          },
          "operations": {
            "href": "https://horizon-testnet.stellar.org/transactions/02a142294acde3e8920fc1e3879b34763df56b9c387a6dcf16d613b21cb9b87a/operations{?cursor,limit,order}",
            "templated": true
          },
          "effects": {
            "href": "https://horizon-testnet.stellar.org/transactions/02a142294acde3e8920fc1e3879b34763df56b9c387a6dcf16d613b21cb9b87a/effects{?cursor,limit,order}",
            "templated": true
          },
          "precedes": {
            "href": "https://horizon-testnet.stellar.org/transactions?order=asc\u0026cursor=6859062775808"
          },
          "succeeds": {
            "href": "https://horizon-testnet.stellar.org/transactions?order=desc\u0026cursor=6859062775808"
          }
        },
        "id": "02a142294acde3e8920fc1e3879b34763df56b9c387a6dcf16d613b21cb9b87a",
        "paging_token": "6859062775808",
        "hash": "02a142294acde3e8920fc1e3879b34763df56b9c387a6dcf16d613b21cb9b87a",
        "ledger": 1597,
        "created_at": "2016-04-08T19:01:08Z",
        "source_account": "GBS43BF24ENNS3KPACUZVKK2VYPOZVBQO2CISGZ777RYGOPYC2FT6S3K",
        "source_account_sequence": "4672924418049",
        "fee_paid": 100,
        "operation_count": 1,
        "envelope_xdr": "AAAAAGXNhLrhGtltTwCpmqlarh7s1DB2hIkbP//jgzn4Fos/AAAAZAAABEAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAoexpe1J6NWzQ1DFmu0ksFrUXpyKf772baBoGkka8kmMAAAAXSHboAAAAAAAAAAAB+BaLPwAAAEAtHpu3C1iaUHhCmpui9L60pb+waQfo5cJjUhzPQBNn+AG0G8K7zLqDtDFe2fTm5CGZ23nOJ64ae2VhtCibJwUD",
        "result_xdr": "AAAAAAAAAGQAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAA=",
        "result_meta_xdr": "AAAAAAAAAAEAAAACAAAAAAAABj0AAAAAAAAAAKHsaXtSejVs0NQxZrtJLBa1F6cin++9m2gaBpJGvJJjAAAAF0h26AAAAAY9AAAAAAAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAQAABj0AAAAAAAAAAGXNhLrhGtltTwCpmqlarh7s1DB2hIkbP//jgzn4Fos/AABa28gDV5wAAARAAAAAAQAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAA",
        "fee_meta_xdr": "AAAAAgAAAAMAAARAAAAAAAAAAABlzYS64RrZbU8AqZqpWq4e7NQwdoSJGz//44M5+BaLPwAAWvMQekAAAAAEQAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEAAAY9AAAAAAAAAABlzYS64RrZbU8AqZqpWq4e7NQwdoSJGz//44M5+BaLPwAAWvMQej+cAAAEQAAAAAEAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==",
        "memo_type": "none",
        "signatures": [
          "LR6btwtYmlB4QpqbovS+tKW/sGkH6OXCY1Icz0ATZ/gBtBvCu8y6g7QxXtn05uQhmdt5zieuGntlYbQomycFAw=="
        ]
      }
    ]
  }
}`

var notFoundResponse = `{
  "type": "https://stellar.org/horizon-errors/not_found",
  "title": "Resource Missing",
  "status": 404,
  "detail": "The resource at the url requested was not found.  This is usually occurs for one of two reasons:  The url requested is not valid, or no data in our database could be found with the parameters provided.",
  "instance": "horizon-live-001/61KdRW8tKi-18408110"
}`

var submitResponse = `{
  "_links": {
    "transaction": {
      "href": "https://horizon-testnet.stellar.org/transactions/ee14b93fcd31d4cfe835b941a0a8744e23a6677097db1fafe0552d8657bed940"
    }
  },
  "hash": "ee14b93fcd31d4cfe835b941a0a8744e23a6677097db1fafe0552d8657bed940",
  "ledger": 3128812,
  "envelope_xdr": "AAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAZAAT3TUAAAAwAAAAAAAAAAAAAAABAAAAAAAAAAMAAAABSU5SAAAAAAA0jDEZkBgx+hCc5IIv+z6CoaYTB8jRkIA6drZUv3YRlwAAAAFVU0QAAAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAAAX14QAAAAAKAAAAAQAAAAAAAAAAAAAAAAAAAAG/dhGXAAAAQLuStfImg0OeeGAQmvLkJSZ1MPSkCzCYNbGqX5oYNuuOqZ5SmWhEsC7uOD9ha4V7KengiwNlc0oMNqBVo22S7gk=",
  "result_xdr": "AAAAAAAAAGQAAAAAAAAAAQAAAAAAAAADAAAAAAAAAAAAAAAAAAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAAAAAAPEAAAABSU5SAAAAAAA0jDEZkBgx+hCc5IIv+z6CoaYTB8jRkIA6drZUv3YRlwAAAAFVU0QAAAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAAAX14QAAAAAKAAAAAQAAAAAAAAAAAAAAAA==",
  "result_meta_xdr": "AAAAAAAAAAEAAAACAAAAAAAvoHwAAAACAAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAAAAAAPEAAAABSU5SAAAAAAA0jDEZkBgx+hCc5IIv+z6CoaYTB8jRkIA6drZUv3YRlwAAAAFVU0QAAAAAADSMMRmQGDH6EJzkgi/7PoKhphMHyNGQgDp2tlS/dhGXAAAAAAX14QAAAAAKAAAAAQAAAAAAAAAAAAAAAAAAAAEAL6B8AAAAAAAAAAA0jDEZkBgx+hCc5IIv+z6CoaYTB8jRkIA6drZUv3YRlwAAABZ9zvNAABPdNQAAADAAAAAEAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA=="
}`

var transactionFailure = `{
  "type": "https://stellar.org/horizon-errors/transaction_failed",
  "title": "Transaction Failed",
  "status": 400,
  "detail": "The transaction failed when submitted to the stellar network. The extras.result_codes field on this response contains further details.  Descriptions of each code can be found at: https://www.stellar.org/developers/learn/concepts/list-of-operations.html",
  "instance": "horizon-testnet-001.prd.stellar001.internal.stellar-ops.com/4elYz2fHhC-528285",
  "extras": {
    "envelope_xdr": "AAAAAKpmDL6Z4hvZmkTBkYpHftan4ogzTaO4XTB7joLgQnYYAAAAZAAAAAAABeoyAAAAAAAAAAEAAAAAAAAAAQAAAAAAAAABAAAAAD3sEVVGZGi/NoC3ta/8f/YZKMzyi9ZJpOi0H47x7IqYAAAAAAAAAAAF9eEAAAAAAAAAAAA=",
    "result_codes": {
      "transaction": "tx_no_source_account"
    },
    "result_xdr": "AAAAAAAAAAD////4AAAAAA=="
  }
}`
