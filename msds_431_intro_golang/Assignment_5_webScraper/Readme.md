# Wikipedia Web Crawler

This project is a Go-based web crawler and scraper built using the [Colly framework](https://github.com/gocolly/colly). The application retrieves text content from Wikipedia pages related to intelligent systems and robotics, strips away the HTML markup, and saves the extracted content into a JSON lines file (`.jsonl` format).

## Table of Contents

- [Project Overview](#project-overview)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Output Format](#output-format)
- [Performance Comparison](#performance-comparison)
- [Future Improvements](#future-improvements)
- [Contributing](#contributing)
- [License](#license)

## Project Overview

This project converts a previous Python/Scrapy solution into a Go-based web scraper that takes advantage of Go's concurrency features to perform fast, efficient scraping of multiple web pages concurrently. The goal is to collect text information from Wikipedia pages relating to intelligent systems and robotics, and save the results in a `.jsonl` file.

## Features

- Concurrent scraping of multiple web pages using Go's goroutines.
- Collection of text content from Wikipedia pages, including headings and paragraphs.
- Storage of the scraped data in JSON lines format (`.jsonl`), making it suitable for large datasets and easy import into databases.
- Detailed logging to monitor the scraping progress.

## Installation

### Prerequisites

Before running this project, you will need to have the following installed:

- [Go](https://golang.org/doc/install) (1.16 or later)
- Git

### Setup

1. Clone this repository:
   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. Install the required Go dependencies:
   ```bash
   go mod tidy
   ```

3. run the Go program:
   ```bash
   go run main.go
   ```

## Usage

1. Run the program using the compiled executable:
   ```bash
   ./wikipedia_crawler
   ```

   The program will scrape the text from the specified Wikipedia URLs and save the extracted content to a file named `wikipedia_data.jsonl`.

2. Adjust the list of Wikipedia URLs in the source code (`main.go`) under the `urls` variable if you want to scrape other pages.

## Output Format

The output is written to a file in [JSON lines](https://jsonlines.org/) format. Each line represents a separate web page's scraped data in JSON format. Below is an example of one entry:

```json
{
  "url": "https://en.wikipedia.org/wiki/Robotics",
  "title": "Robotics",
  "sections": [
    {
      "main_summary": {
        "paragraph": [
          "Robotics is an interdisciplinary branch of engineering and science that includes mechanical engineering, electrical engineering, computer science, and others."
        ]
      }
    }
  ]
}
```
