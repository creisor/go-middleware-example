go-middleware-example
---------------------

To start:

    go run main.go

The home route requires no authentication:

    curl -v http://localhost:3000/

To reach the `/authorized` route, you must be authorized with your super cool token: `notahacker`:

    curl -v -H "auth: notahacker" http://localhost:3000/authorized

To be denied to `/authorized`:

    curl -v -H "auth: totalhacker" http://localhost:3000/authorized

or

    curl -v http://localhost:3000/authorized
