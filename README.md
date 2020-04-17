# mtlsc

Tests mTLS authentication between a client (this tool) and a server.

## Usage

You will need the following assets:

* Your own client certificate and key (that the server has been told to trust).
* Certificate(s) for the server, that you're supposed to trust.

With these in hand, run the command as follows:

```
$ mtlsc connect https://test-ingress.local --server minica.pem,test-ingress.local/cert.pem --cert client-cert/cert.pem --key client-cert/key.pem
```
Here, the server certificates have been specified as the CA cert (`minica.pem`) and the leaf certificate (`test-ingress.local/cert.pem`), which get read
into a temporary certificate store used by the code. You could equally bundle them together in a single file and just pass that in.

The output of the command will look as follows:

```
Connecting to https://test-ingress.local [10 connections]

[OK] Read client certificate and key
[OK] Created Certificate Pool
[OK] Added minica.pem to certificate pool
[OK] Added test-ingress.local/cert.pem to certificate pool
[OK] Initialized HTTP Client

[OK] Connected to https://test-ingress.local [HTTP 200 OK]
[OK] Connected to https://test-ingress.local [HTTP 200 OK]
[OK] Connected to https://test-ingress.local [HTTP 200 OK]
[OK] Connected to https://test-ingress.local [HTTP 200 OK]
[OK] Connected to https://test-ingress.local [HTTP 200 OK]
[OK] Connected to https://test-ingress.local [HTTP 200 OK]
[OK] Connected to https://test-ingress.local [HTTP 200 OK]
[OK] Connected to https://test-ingress.local [HTTP 200 OK]
[OK] Connected to https://test-ingress.local [HTTP 200 OK]
[OK] Connected to https://test-ingress.local [HTTP 200 OK]
[OK] Connected to https://test-ingress.local [HTTP 200 OK]

[OK] Test complete
```
The calls are made sequentially, as one of the aspects I was testing initially was connection reuse with ingress-nginx and mTLS. A method to
test asynchronously will be added later.
