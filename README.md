# mtlsc

Tests mTLS authentication between a client (this tool) and a server.

## Usage

You will need the following assets:

* Your own client certificate and key (that the server has been told to trust).
* Certificate(s) for the server, that you're supposed to trust.

With these in hand, run the command as follows:

```
$ mtlsc connect https://test-ingress.local \
--server minica.pem,test-ingress.local/cert.pem \
--cert client-cert/cert.pem --key client-cert/key.pem
```
Here, the server certificates have been specified as the CA cert (`minica.pem`) and the leaf certificate (`test-ingress.local/cert.pem`), which get read
into a temporary certificate store used by the code. You could equally bundle them together in a single file and just pass that in.

The output of the command will look as follows:

```
Connecting to https://test-ingress.local

[OK] Read client certificate and key
[OK] Created Certificate Pool
[OK] Added minica.pem to certificate pool
[OK] Added test-ingress.local/cert.pem to certificate pool
[OK] Initialized HTTP Client

Synchronous test initiated
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
The calls are made sequentially.

To make asynchronous calls, run the command as:

```
$ mtlsc connect https://test-ingress.local \
--server minica.pem,test-ingress.local/cert.pem \
--cert client-cert/cert.pem --key client-cert/key.pem \
--async
--threads 4
```

This will create the number of `chan`s defined by `threads` and will override anything passed to `count`, as it
will automatically send 4 requests per `chan`. Note that `--threads` in the example is only there for
illustrative purposes. It defaults to 4 threads, so only `--async` needs to be passed if you're happy
with 4 threads and 16 tests.

Output looks like:

```
Connecting to https://test-ingress.local

[OK] Read client certificate and key
[OK] Created Certificate Pool
[OK] Added minica.pem to certificate pool
[OK] Added test-ingress.local/cert.pem to certificate pool
[OK] Initialized HTTP Client

Asynchronous test initiated
Creating a pool of 4 threads
Overriding --count. Count will be 16 [4 x 4 threads]
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
[OK] Connected to https://test-ingress.local [HTTP 200 OK]
[OK] Connected to https://test-ingress.local [HTTP 200 OK]
[OK] Connected to https://test-ingress.local [HTTP 200 OK]
[OK] Connected to https://test-ingress.local [HTTP 200 OK]
[OK] Connected to https://test-ingress.local [HTTP 200 OK]

[OK] Test complete
```
