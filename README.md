# mtlsc

Tests mTLS authentication between a client (this tool) and a server.

## Usage

You will need the following assets:

* Your own client certificate and key (that the server has been told to trust).
* Certificate(s) for the server, that you're supposed to trust.

With these in hand, you can invoke the `connect` command to start running tests, using `--concurrency` to set the number of
concurrent requests and `--count` to control the number of requests sent per connection.
