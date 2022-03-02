#!/bin/sh

echo "call /helloworld.Greeter/SayHello"
curl "http://localhost:10000/helloworld.Greeter/SayHello" \
    -H 'Content-Type: application/json; charset=utf-8'
echo "\n----------------------------"

echo "call /helloworld.Greeter/Status"
curl "http://localhost:10000/helloworld.Greeter/Status" \
    -H 'Content-Type: application/json; charset=utf-8'
echo "\n----------------------------"

echo "call /status"
curl "http://localhost:10000/status" \
    -H 'Content-Type: application/json; charset=utf-8'
echo "\n----------------------------"

echo "call /bookstore.BookStore/Status"
curl "http://localhost:10000/bookstore.BookStore/Status" \
    -H 'Content-Type: application/json; charset=utf-8'
echo "\n----------------------------"
