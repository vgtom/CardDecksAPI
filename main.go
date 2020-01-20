package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Resources struct {
	Name    string `json:"name"`
	URI     string `json:"uri"`
	Methods string `json:"methods"`
}

type Response struct {
	DeckId    string `json:"deck_id"`
	Shuffled  string `json:"shuffled"`
	Remaining int    `json:"remaining"`
	Cards     []Card `json:"cards"`
}

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

//deck

var cardsstd = []Card{{Value: "2", Suit: "Spades", Code: "2S"}, {Value: "3", Suit: "Spades", Code: "3S"}, {Value: "4", Suit: "Spades", Code: "4S"}, {Value: "5", Suit: "Spades", Code: "5S"},
	{Value: "6", Suit: "Spades", Code: "6S"}, {Value: "7", Suit: "Spades", Code: "7S"}, {Value: "8", Suit: "Spades", Code: "8S"},
	{Value: "9", Suit: "Spades", Code: "9S"}, {Value: "10", Suit: "Spades", Code: "10S"}, {Value: "JACK", Suit: "Spades", Code: "JS"}, {Value: "QUEEN", Suit: "Spades", Code: "QS"},
	{Value: "KING", Suit: "Spades", Code: "KS"}, {Value: "ACE", Suit: "Spades", Code: "AS"},
	{Value: "2", Suit: "Hearts", Code: "2H"}, {Value: "2", Suit: "Hearts", Code: "3H"}, {Value: "2", Suit: "Hearts", Code: "4H"}, {Value: "5", Suit: "Hearts", Code: "5H"},
	{Value: "6", Suit: "Hearts", Code: "6H"}, {Value: "7", Suit: "Hearts", Code: "7H"}, {Value: "8", Suit: "Hearts", Code: "8H"},
	{Value: "9", Suit: "Hearts", Code: "9H"}, {Value: "10", Suit: "Hearts", Code: "10H"}, {Value: "JACK", Suit: "Hearts", Code: "JH"}, {Value: "QUEEN", Suit: "Hearts", Code: "QH"},
	{Value: "KING", Suit: "Hearts", Code: "KH"}, {Value: "ACE", Suit: "Hearts", Code: "AH"},
	{Value: "2", Suit: "Diamonds", Code: "2D"}, {Value: "2", Suit: "Diamonds", Code: "3D"}, {Value: "2", Suit: "Diamonds", Code: "4D"}, {Value: "5", Suit: "Diamonds", Code: "5D"},
	{Value: "6", Suit: "Diamonds", Code: "6D"}, {Value: "7", Suit: "Diamonds", Code: "7D"}, {Value: "8", Suit: "Diamonds", Code: "8D"},
	{Value: "9", Suit: "Diamonds", Code: "9D"}, {Value: "10", Suit: "Diamonds", Code: "10D"}, {Value: "JACK", Suit: "Diamonds", Code: "JD"}, {Value: "QUEEN", Suit: "Diamonds", Code: "QD"},
	{Value: "KING", Suit: "Diamonds", Code: "KD"}, {Value: "ACE", Suit: "Diamonds", Code: "AD"},
	{Value: "2", Suit: "Clubs", Code: "2C"}, {Value: "2", Suit: "Clubs", Code: "3C"}, {Value: "2", Suit: "Clubs", Code: "4C"}, {Value: "5", Suit: "Clubs", Code: "5C"},
	{Value: "6", Suit: "Clubs", Code: "6C"}, {Value: "7", Suit: "Clubs", Code: "7C"}, {Value: "8", Suit: "Clubs", Code: "8C"},
	{Value: "9", Suit: "Clubs", Code: "9C"}, {Value: "10", Suit: "Clubs", Code: "10C"}, {Value: "JACK", Suit: "Clubs", Code: "JC"}, {Value: "QUEEN", Suit: "Clubs", Code: "QC"},
	{Value: "KING", Suit: "Clubs", Code: "KC"}, {Value: "ACE", Suit: "Clubs", Code: "AC"}}

var cardMap = map[string]Card{}

//maintain a list of all decks created in memory...instead of dynamically creating it everytime...this is much faster even though a little more memory
//and in a real card game at later points some of the Decks can be deleted to free up memory
var Decks = map[string]Response{}

//shuffle
func Shuffle(slc []Card) {
	for i := 1; i < len(slc); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			slc[r], slc[i] = slc[i], slc[r]
		}
	}
}

func remove(slice []Card, s int) []Card {
	return append(slice[:s], slice[s+1:]...)
}

func Create(r *gin.Context) {

	cards := append(cardsstd[:0:0], cardsstd...)

	bshuffle := r.Query("shuffle")
	if bshuffle == "true" {
		Shuffle(cards)
	} else {
		bshuffle = "false"
	}
	strcards := r.Query("cards")
	cardslst := strings.Split(strcards, ",")

	lncard := len(cardslst)

	if lncard > 1 {
		fmt.Println("Inside cards list")
		var fewcards []Card
		for _, element := range cardslst {
			//get that element from the global array
			if v, ok := cardMap[element]; ok {
				fewcards = append(fewcards, v)
			}
		}
		id, err := uuid.NewUUID()
		if err != nil {
			fmt.Println("Cannot create uuid")
		}
		jsonResp := Response{DeckId: id.String(), Shuffled: bshuffle, Remaining: 52 - lncard, Cards: fewcards}
		Decks[id.String()] = jsonResp
		r.JSON(200, jsonResp)
		return
	}

	id, err := uuid.NewUUID()
	if err != nil {
		fmt.Println("Cannot create uuid")
	}
	jsonResp := Response{DeckId: id.String(), Shuffled: bshuffle, Remaining: 52, Cards: cards}
	Decks[id.String()] = jsonResp // append to the gobal list of unique decks created
	r.JSON(200, jsonResp)
}

func Open(r *gin.Context) {
	szuuid := r.Query("uuid")
	v, ok := Decks[szuuid]
	if !ok {
		//return error
	} else {
		r.JSON(200, v)
	}
}

func Draw(r *gin.Context) {
	szuuid := r.Query("uuid")
	v, ok := Decks[szuuid]
	if !ok {
		//return error
	} else {
		szcount := r.Query("count")
		i, err := strconv.Atoi(szcount)
		if err != nil {
			fmt.Println("Invalid count")
		} else {
			var fewcards []Card
			for j := 0; j < i; j++ {
				// draw a card from the Deck and remove that card and add it to the return list
				tcards := v.Cards
				ln := len(tcards)
				fmt.Println("Number of cards", ln)
				r := rand.Intn(ln - 1)

				// get that random index from the deck and delete it from there
				fewcards = append(fewcards, tcards[r])
				//delete that card from the deck
				v.Cards = remove(v.Cards, r)
			}
			jsonResp := Response{DeckId: v.DeckId, Shuffled: v.Shuffled, Remaining: v.Remaining - i, Cards: fewcards}
			Decks[v.DeckId] = jsonResp // append to the gobal list of unique decks created
			r.JSON(200, jsonResp)
		}
	}
}

func SetupRouter() *gin.Engine {
	m := gin.Default()

	m.GET("/", func(r *gin.Context) {
		json := Resources{Name: "cards", URI: "/cards/", Methods: "GET"}
		r.JSON(200, json)
	})

	m.GET("/cards/create", Create)

	m.GET("/cards/open", Open)

	m.GET("/cards/draw", Draw)

	return m
}

func main() {

	for _, v := range cardsstd {
		cardMap[v.Code] = v
	}

	router := SetupRouter()
	router.Run()
}
