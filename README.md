# Sample Card of Decks API

There are 3 REST Endpoints

# /cards/create
For creating a deck of cards
# /cards/open
For opening a Deck of cards
# /cards/draw
for drawing cards from a deck of cards

# Create
You can choose to create a "suffled deck of cards"
for e.g.
curl -i -k -u token: "http://localhost:8080/cards/create?shuffled=true"

or create a specific set of cards
curl -i -k -u token: "http://localhost:8080/cards/create?cards=AS,KD,AC,2C,KH"

# Open
You can choose to open a deck of cards that has been created or from which cards have been drawn and see its status
for e.g.
curl -i -k -u token: curl -i -k -u token: "http://localhost:8080/cards/open?uuid=cfb1d142-3b43-11ea-a7e5-9cb6d0d5fe15"

# Draw
You can choose to draw a certain number(count) of cards from a deck at random
for e.g.
curl -i -k -u token: curl -i -k -u token: "http://localhost:8080/cards/draw?uuid=d4b61149-3b44-11ea-a9bf-9cb6d0d5fe15&count=2"


# Buildng and Testing with Makefile files
Make sure you have the make system installed on your machine
Then do,

make build 
to build the project

make run
to run the project

make test
to test the project


# BUILDING and RUNNING
Since everything is encapsulated in a single file, the build process is rather simple

Just do

go build
and that creates an executable ./main
and you can run ./main

or do 
go run main.go

# Testing

go test -v
will run all the test cases for you

