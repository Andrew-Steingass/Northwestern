package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite" // Use the cgo-free SQLite library
)

func main() {
	// Connect to the SQLite database
	db, err := sql.Open("sqlite", "C:/Users/Andy/Documents/Github/Northwestern/msds_431_intro_golang/assignment_10/movies.db")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Define the query
	query := `
	SELECT m.Name, m.Year, m.Rank, mg.GenreName
	FROM Movies m
	LEFT JOIN MovieGenres mg ON m.MovieID = mg.MovieID
	WHERE mg.GenreName = 'Documentary' AND m.Rank >= 8 AND m.Year >= 2000
	ORDER BY m.Rank DESC, m.Year DESC;
	`

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}
	defer rows.Close()

	// Variables for tracking results
	var results []string
	count := 0

	// Iterate through rows
	for rows.Next() {
		var name string
		var year int
		var rank float64
		var genreName string

		err := rows.Scan(&name, &year, &rank, &genreName)
		if err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}

		// Collect result
		results = append(results, fmt.Sprintf("Name: %s, Year: %d, Rank: %.1f, Genre: %s", name, year, rank, genreName))
		count++
	}

	// Check for iteration errors
	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating through rows: %v", err)
	}

	// Print results
	fmt.Printf("\nTotal Results Found: %d (Genre: Documentary, Rank >= 8, Year >= 2000)\n", count)
	fmt.Println("\nTop 5 Movies:")
	for i, result := range results {
		if i >= 5 {
			break
		}
		fmt.Println(result)
	}
}
