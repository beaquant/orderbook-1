# Orderbook

This package contains a basic implementation of an [orderbook](https://en.wikipedia.org/wiki/Order_book_(trading)). To run this orderbook against some example orders, use cmd/main.go:

```
$ cd cmd
$ go run main.go

```

Running this command should give you output that looks something like this:

```
Placed orders
|     limit-buy-order|     100.000|  1.00000000| Mon, 06 Aug 2018 19:26:26 +0000|  35c1f4cd-ab9b-473c-9bdb-3a47a729fd08|
|    limit-sell-order|     104.000|  1.32350000| Mon, 06 Aug 2018 19:27:31 +0000|  33f339b7-67de-4615-8e00-7da4e6ed674a|
|    limit-sell-order|     103.500|  4.20000000| Mon, 06 Aug 2018 19:30:34 +0000|  2809a42f-3e1f-4b34-9436-1ad59cb3a19c|
|     limit-buy-order|     101.000|  2.25000000| Mon, 06 Aug 2018 19:33:37 +0000|  d06490a0-ad01-49cb-9542-033fac3127ca|
|    limit-sell-order|     100.750|  5.00000000| Mon, 06 Aug 2018 19:35:12 +0000|  4d6a6438-8fd8-4175-87e1-6d4318b98b72|
|     limit-buy-order|     100.550|  3.45000000| Mon, 06 Aug 2018 19:40:07 +0000|  7d90e333-0c61-44ea-8435-8ffe6d1edc32|
|    limit-sell-order|     100.450|  2.73400000| Mon, 06 Aug 2018 19:45:55 +0000|  53002162-f2a8-4028-8565-6ad4125bc085|
|    limit-sell-order|     103.500|  2.20000000| Mon, 06 Aug 2018 19:48:55 +0000|  c6704e45-3e7e-4788-a647-08bf7dec67d4|
|     limit-buy-order|     100.750|  0.50000000| Mon, 06 Aug 2018 19:48:55 +0000|  705e9d5e-1bf5-44b5-b914-4ade6e1f4022|

Executed trades
Quantity: 2.25000000, Price: 101.000, Orders: [d06490a0-ad01-49cb-9542-033fac3127ca 4d6a6438-8fd8-4175-87e1-6d4318b98b72]
Quantity: 2.73400000, Price: 100.550, Orders: [7d90e333-0c61-44ea-8435-8ffe6d1edc32 53002162-f2a8-4028-8565-6ad4125bc085]
Quantity: 0.50000000, Price: 100.750, Orders: [705e9d5e-1bf5-44b5-b914-4ade6e1f4022 4d6a6438-8fd8-4175-87e1-6d4318b98b72]

Final orderbook state
|                Type|       Price|    Quantity|                            Time|                                  UUID|
|     limit-buy-order|     100.000|  1.00000000| Mon, 06 Aug 2018 19:26:26 +0000|  35c1f4cd-ab9b-473c-9bdb-3a47a729fd08|
|     limit-buy-order|     100.550|  0.71600000| Mon, 06 Aug 2018 19:40:07 +0000|  7d90e333-0c61-44ea-8435-8ffe6d1edc32|
|                    |            |            |                                |                                      |
|    limit-sell-order|     100.750|  2.25000000| Mon, 06 Aug 2018 19:35:12 +0000|  4d6a6438-8fd8-4175-87e1-6d4318b98b72|
|    limit-sell-order|     103.500|  4.20000000| Mon, 06 Aug 2018 19:30:34 +0000|  2809a42f-3e1f-4b34-9436-1ad59cb3a19c|
|    limit-sell-order|     103.500|  2.20000000| Mon, 06 Aug 2018 19:48:55 +0000|  c6704e45-3e7e-4788-a647-08bf7dec67d4|
|    limit-sell-order|     104.000|  1.32350000| Mon, 06 Aug 2018 19:27:31 +0000|  33f339b7-67de-4615-8e00-7da4e6ed674a|
```
