package main

import (
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
	Title   string  `json:"title"`
	Content Content `json:"sections"`
}

type SectionInfo struct {
	Title      string
	Paragraphs []string
}

type AllWebsiteData struct {
	Websites map[string]WebsiteData `json:"websites"`
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

	// Initialize the data store with mutex
	var mu sync.Mutex
	allData := AllWebsiteData{
		Websites: make(map[string]WebsiteData),
	}

	// Create a WaitGroup for synchronization
	var wg sync.WaitGroup

	// Process each URL
	for _, url := range urls {
		wg.Add(1)
		// Add delay between requests
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

			// Get the page title
			c.OnHTML(".mw-page-title-main", func(e *colly.HTMLElement) {
				pageTitle = e.Text
				fmt.Printf("\n=== Processing %s: Found main title: %s ===\n", url, pageTitle)
			})

			// Process the content
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

			// After scraping, save the data
			c.OnScraped(func(r *colly.Response) {
				// Add the last section if it has content
				if len(currentSection.Paragraphs) > 0 {
					sections = append(sections, currentSection)
				}

				// Convert sections to final format
				var finalSections []map[string]ParagraphSection
				for _, section := range sections {
					sectionMap := map[string]ParagraphSection{
						section.Title: {
							Paragraphs: section.Paragraphs,
						},
					}
					finalSections = append(finalSections, sectionMap)
				}

				// Safely add to the combined data
				mu.Lock()
				allData.Websites[url] = WebsiteData{
					Title: pageTitle,
					Content: Content{
						Sections: finalSections,
					},
				}
				mu.Unlock()

				fmt.Printf("Completed processing %s\n", url)
			})

			c.Visit(url)
		}(url)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Write the combined data to a single file
	jsonData, err := json.MarshalIndent(allData, "", "    ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	err = os.WriteFile("wikipedia_data.json", jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Println("\nAll data written to wikipedia_data.json")
	fmt.Println("\nSummary of scraped data:")
	for url, data := range allData.Websites {
		fmt.Printf("- %s: %s (Sections: %d)\n", url, data.Title, len(data.Content.Sections))
	}
}
