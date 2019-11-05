go-middleware-example
---------------------

To start:

    go run main.go

To be authorized:

    curl -v -H "auth: notahacker" http://localhost:3000/ 

To be denied:

    curl -v -H "auth: imahacker" http://localhost:3000/

or

    curl -v http://localhost:3000/
