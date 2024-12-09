
# How to Use the Go Movie Database Code

## Prerequisites
- Install Go from [https://golang.org/dl/](https://golang.org/dl/).
- Install SQLite (use the pure Go version from `modernc.org/sqlite`).

## Steps to Use the Code

1. **Clone or Download the Project**
   - Clone the repository or download the Go code to your local machine.

2. **Install Dependencies**
   - In the project directory, initialize Go modules:
     ```bash
     go mod init assignment_10
     ```
   - Install the SQLite dependency:
     ```bash
     go get modernc.org/sqlite
     ```

3. **Prepare SQLite Database**
   - Make sure the `movies.db` SQLite database exists in your project directory.
   - Import your CSV data into the `Movies` and `MovieGenres` tables using the SQLite CLI:
     ```sql
     .mode csv
     .import "path/to/IMDB-movies.csv" Movies
     .import "path/to/IMDB-movies_genres.csv" MovieGenres
     ```

4. **Run the Go Code**
   - In the project directory, run the Go program:
     ```bash
     go run main.go
     ```
   - The program will query the database and print the top 5 movies with the genre "Documentary" and rank >= 8.

5. **Modify the Code (Optional)**
   - Edit the SQL query in the `main.go` file to customize the query or database operations as needed.

# Assignment Write-up

### 1. Describe your efforts in setting up the SQLite relational database.

<details>
<summary>View Details</summary>

    I had some experience doing this in the past for an employer, but I usually did it in Python. So, doing it in the command line SQLite was interesting. I downloaded the SQLite files and ran some raw commands through the SQL CLI. The copy and paste functionality was a bit buggy, but it was an interesting experience.

<details>
---

### 2. Describe how you might add to this database by including a table showing the movies that you have in your personal collection, where those movies are located, and perhaps your personal ratings of the movies.

<details>
<summary>View Details</summary>

    Perhaps I could connect via an API to various services like Netflix, Hulu, Amazon Prime, etc. I could automatically pull in "watched" tags or maybe the percentage watched and whether I liked it or not. I could then store this information in another table or design a schema for it. If it were for a business, perhaps I would attach a SKU or a contract/royalty ID to each item in the database to indicate that I own the rights to a movie, along with a link to the contract.

<details>
---

### 3. Describe plans for drawing on the personal movie database. What purpose would it serve? Describe possible user interactions with the database (beyond what can be obtained from SQL queries alone). In other words, what would a useful Go movie application look like? What would be the advantages of this application over IMDb alone?

<details>
<summary>View Details</summary>

    I think I kind of answered most of these questions in item 2. I would design a phone/computer application where you could link your streaming platforms or even brick-and-mortar movie theater ticket purchases to your phone/computer application. A user could then keep track of what they have seen or not seen. A suggestion machine learning algorithm could be added to enhance the application and query the database. Perhaps others in your life could plan out movie dates based on what the two of you have in common and haven't seen. Maybe there could be free community viewing places, and you could tune in to see when and where they are. A good way to meet friends.

<details>
---

### 4. Describe possible database enhancement

<details>
<summary>View Details</summary>

    I think the rank could be better. Pulling in rotten tomato or other items like that could help provide a better ranking. I was thinking about breaking out the genres into a third table having Genre ID's with a many to many connector in the middle. That may be a good idea depending on a couple items, but if the need for speed is there for the use case, then I would say the current format is good.
    
<details>