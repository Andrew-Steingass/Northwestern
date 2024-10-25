package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

type ParagraphSection struct {
	Paragraphs []string `json:"paragraph"`
}

type Content struct {
	Sections []map[string]ParagraphSection `json:"sections"`
}

type WebsiteData struct {
	URL     string  `json:"url"` // Added URL field
	Title   string  `json:"title"`
	Content Content `json:"sections"`
}

type SectionInfo struct {
	Title      string
	Paragraphs []string
}

func main() {
	urls := []string{
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
		"https://en.wikipedia.org/wiki/Robot_Operating_System",
		"https://en.wikipedia.org/wiki/Intelligent_agent",
		"https://en.wikipedia.org/wiki/Software_agent",
		"https://en.wikipedia.org/wiki/Robotic_process_automation",
		"https://en.wikipedia.org/wiki/Chatbot",
		"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
		"https://en.wikipedia.org/wiki/Android_(robot)",
	}

	// Create and open the output file
	file, err := os.Create("wikipedia_data.jsonl")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	var mu sync.Mutex
	var wg sync.WaitGroup

	// Process each URL
	for _, url := range urls {
		wg.Add(1)
		time.Sleep(100 * time.Millisecond)

		go func(url string) {
			defer wg.Done()

			c := colly.NewCollector()

			var sections []SectionInfo
			currentSection := SectionInfo{
				Title:      "main_summary",
				Paragraphs: []string{},
			}
			var pageTitle string

			c.OnHTML(".mw-page-title-main", func(e *colly.HTMLElement) {
				pageTitle = e.Text
				fmt.Printf("\n=== Processing %s: Found main title: %s ===\n", url, pageTitle)
			})

			c.OnHTML("#mw-content-text", func(e *colly.HTMLElement) {
				e.ForEach("*", func(_ int, el *colly.HTMLElement) {
					if el.Name == "div" && el.Attr("class") == "mw-heading mw-heading2" {
						if len(currentSection.Paragraphs) > 0 {
							sections = append(sections, currentSection)
						}
						currentSection = SectionInfo{
							Title:      el.ChildText("h2"),
							Paragraphs: []string{},
						}
					}

					if el.Name == "p" {
						text := el.Text
						if text != "" && text != "\n" {
							currentSection.Paragraphs = append(currentSection.Paragraphs, text)
						}
					}
				})
			})

			c.OnScraped(func(r *colly.Response) {
				if len(currentSection.Paragraphs) > 0 {
					sections = append(sections, currentSection)
				}

				var finalSections []map[string]ParagraphSection
				for _, section := range sections {
					sectionMap := map[string]ParagraphSection{
						section.Title: {
							Paragraphs: section.Paragraphs,
						},
					}
					finalSections = append(finalSections, sectionMap)
				}

				// Create single website data object
				data := WebsiteData{
					URL:   url,
					Title: pageTitle,
					Content: Content{
						Sections: finalSections,
					},
				}

				// Marshal to JSON
				jsonData, err := json.Marshal(data)
				if err != nil {
					fmt.Printf("Error marshaling JSON for %s: %v\n", url, err)
					return
				}

				// Write to file with mutex
				mu.Lock()
				writer.Write(jsonData)
				writer.WriteString("\n")
				mu.Unlock()

				fmt.Printf("Completed processing %s\n", url)
			})

			c.Visit(url)
		}(url)
	}

	wg.Wait()

	fmt.Println("\nAll data written to wikipedia_data.jsonl")
	fmt.Println("\nSummary of completed scraping:")

	// Read and parse the file to show summary
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
		var data WebsiteData
		json.Unmarshal(scanner.Bytes(), &data)
		fmt.Printf("- %s: %s (Sections: %d)\n", data.URL, data.Title, len(data.Content.Sections))
	}
}
