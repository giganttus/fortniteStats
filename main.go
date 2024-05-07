package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"fortniteStats/models"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Create log file
	logFile, err := os.OpenFile("fortniteStats.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v\n", err)
	}

	log.SetOutput(logFile)
	defer logFile.Close()

	// Load environment variables from .env file
	err = godotenv.Load("fortniteStats.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get auth token from environment variables
	token := os.Getenv("AUTH_TOKEN")
	if token == "" {
		log.Fatal("Authorization token missing")
	}

	// Open the file for reading
	file, err := os.Open("playerList.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Iterate over each line of the file
	for scanner.Scan() {
		playerName := scanner.Text()

		responseData, err := FortniteApi("GET", playerName, token)
		if err != nil {
			log.Println("Error occured", err)
			continue
		}

		// Initialize FinalStats variable
		finalStats := StrucutreData(responseData)

		UploadData(os.Getenv("MONGO_CONN"), finalStats)
	}
}

func FortniteApi(method string, playerName string, token string) (models.Response, error) {
	var responseData models.Response

	// Create a new HTTP client, param values
	client := &http.Client{}
	params := url.Values{}
	params.Add("name", playerName)
	params.Add("timeWindow", "season")

	// Base and req url
	url := "https://fortnite-api.com/v2/stats/br/v2"
	reqUrl := url + "?" + params.Encode()

	// Create a new GET request
	req, err := http.NewRequest(method, reqUrl, nil)
	if err != nil {
		return responseData, err
	}

	// Add Authorization header
	req.Header.Add("Authorization", token)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return responseData, err
	}
	defer resp.Body.Close()

	// Status code handling
	if resp.StatusCode != 200 {
		return responseData, fmt.Errorf("%v", resp.Status)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return responseData, err
	}

	// Unmarshal the JSON response into the variable
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return responseData, err
	}

	return responseData, nil
}

func StrucutreData(responseData models.Response) models.FinalStats {
	var finalStats models.FinalStats

	// Extract data from responseData and assign it to finalStats
	finalStats.ID = responseData.Data.Account.ID
	finalStats.Name = responseData.Data.Account.Name

	// Assign overall stats
	finalStats.OverallWins = responseData.Data.Stats.All.Overall.Wins
	finalStats.OverallTop3 = responseData.Data.Stats.All.Overall.Top3
	finalStats.OverallTop5 = responseData.Data.Stats.All.Overall.Top5
	finalStats.OverallTop6 = responseData.Data.Stats.All.Overall.Top6
	finalStats.OverallTop10 = responseData.Data.Stats.All.Overall.Top10
	finalStats.OverallTop12 = responseData.Data.Stats.All.Overall.Top12
	finalStats.OverallTop25 = responseData.Data.Stats.All.Overall.Top25
	finalStats.OverallKills = responseData.Data.Stats.All.Overall.Kills
	finalStats.OverallMatches = responseData.Data.Stats.All.Overall.Matches
	finalStats.OverallLastModified = responseData.Data.Stats.All.Overall.LastModified

	// Assign solo stats
	finalStats.SoloWins = responseData.Data.Stats.All.Solo.Wins
	finalStats.SoloTop3 = responseData.Data.Stats.All.Solo.Top3
	finalStats.SoloTop5 = responseData.Data.Stats.All.Solo.Top5
	finalStats.SoloTop6 = responseData.Data.Stats.All.Solo.Top6
	finalStats.SoloTop10 = responseData.Data.Stats.All.Solo.Top10
	finalStats.SoloTop12 = responseData.Data.Stats.All.Solo.Top12
	finalStats.SoloTop25 = responseData.Data.Stats.All.Solo.Top25
	finalStats.SoloKills = responseData.Data.Stats.All.Solo.Kills
	finalStats.SoloMatches = responseData.Data.Stats.All.Solo.Matches
	finalStats.SoloLastModified = responseData.Data.Stats.All.Solo.LastModified

	// Assign duo stats
	finalStats.DuoWins = responseData.Data.Stats.All.Duo.Wins
	finalStats.DuoTop3 = responseData.Data.Stats.All.Duo.Top3
	finalStats.DuoTop5 = responseData.Data.Stats.All.Duo.Top5
	finalStats.DuoTop6 = responseData.Data.Stats.All.Duo.Top6
	finalStats.DuoTop10 = responseData.Data.Stats.All.Duo.Top10
	finalStats.DuoTop12 = responseData.Data.Stats.All.Duo.Top12
	finalStats.DuoTop25 = responseData.Data.Stats.All.Duo.Top25
	finalStats.DuoKills = responseData.Data.Stats.All.Duo.Kills
	finalStats.DuoMatches = responseData.Data.Stats.All.Duo.Matches
	finalStats.DuoLastModified = responseData.Data.Stats.All.Duo.LastModified

	// Assign trio stats
	finalStats.TrioWins = responseData.Data.Stats.All.Trio.Wins
	finalStats.TrioTop3 = responseData.Data.Stats.All.Trio.Top3
	finalStats.TrioTop5 = responseData.Data.Stats.All.Trio.Top5
	finalStats.TrioTop6 = responseData.Data.Stats.All.Trio.Top6
	finalStats.TrioTop10 = responseData.Data.Stats.All.Trio.Top10
	finalStats.TrioTop12 = responseData.Data.Stats.All.Trio.Top12
	finalStats.TrioTop25 = responseData.Data.Stats.All.Trio.Top25
	finalStats.TrioKills = responseData.Data.Stats.All.Trio.Kills
	finalStats.TrioMatches = responseData.Data.Stats.All.Trio.Matches
	finalStats.TrioLastModified = responseData.Data.Stats.All.Trio.LastModified

	// Assign squad stats
	finalStats.SquadWins = responseData.Data.Stats.All.Squad.Wins
	finalStats.SquadTop3 = responseData.Data.Stats.All.Squad.Top3
	finalStats.SquadTop5 = responseData.Data.Stats.All.Squad.Top5
	finalStats.SquadTop6 = responseData.Data.Stats.All.Squad.Top6
	finalStats.SquadTop10 = responseData.Data.Stats.All.Squad.Top10
	finalStats.SquadTop12 = responseData.Data.Stats.All.Squad.Top12
	finalStats.SquadTop25 = responseData.Data.Stats.All.Squad.Top25
	finalStats.SquadKills = responseData.Data.Stats.All.Squad.Kills
	finalStats.SquadMatches = responseData.Data.Stats.All.Squad.Matches
	finalStats.SquadLastModified = responseData.Data.Stats.All.Squad.LastModified

	return finalStats
}

func UploadData(connectionString string, finalStats models.FinalStats) error {
	// Set up MongoDB connection options.
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB.
	mongoClient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			log.Println(err)
		}
	}()

	// Check the connection.
	err = mongoClient.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	// Access a MongoDB collection.
	collection := mongoClient.Database("FortniteStats").Collection("Players")
	insertResult, err := collection.InsertOne(context.Background(), finalStats)
	if err != nil {
		return err
	}

	log.Println("Inserted", insertResult.InsertedID)

	return err
}
