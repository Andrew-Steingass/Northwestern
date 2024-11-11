
# Anscombe Regression Analysis in Go

## Overview
This project implements linear regression on Anscombe's Quartet using Go. It compares regression results with Python/R and benchmarks Go's performance.

## Requirements
- Go installed ([download here](https://golang.org/dl/))
- Install dependencies:
  ```bash
        go get -u gonum.org/v1/gonum/...
        go get gonum.org/v1/gonum/stat
        go get gonum.org/v1/plot/...
  ```

## Running the Program
To run the regression on all datasets (`x1`, `x2`, `x3`, `x4`):
```bash
go run main_copilot.go

go run main_LLM_browser.go
```
The co-pilot uses github co-pilot to adjust the code.
The LLM_Browser uses Chat GPT browser interface to adjust the code.   (ChatGPT_Conversation.txt) for conversation

## Running Unit Tests
To verify regression results:
```bash
go test
```

## Running Benchmarks
To benchmark the performance:
```bash
go test -bench=.
```



## Recommendation to Management
Using Go for regression analysis showed comparable results to Python and R, with slight differences due to rounding and floating-point precision. Go's execution time was faster, making it suitable for performance-critical tasks.

However, data scientists may find Go's statistical ecosystem less mature compared to Python's `pandas`, `statsmodels`, and R's vast library of statistical packages. Python and R have richer support for advanced statistical modeling and better community support. For routine statistical tasks, Python or R is recommended, but Go can be an efficient choice for high-performance applications. If there is a need for heavy processing, Go is recommended

## Challenges Faced
During development, some differences were noted in how calculations like regression coefficients, p-values, and rounding were handled between Go and Python/R. These differences stem from how floating-point precision is handled across languages and their respective libraries. Careful handling of these discrepancies was necessary to ensure comparable results across different platforms.



## Assignment 7 Write-up

### Co-Pilot (GPT 4o)
I didnt really like copilot and AI directly itegrated into my code. Atleast for this assignment. There is too much is going on in a single tool. I would rather have items segmented out. Although, when playing around with some of the inline auto suggestions I liked it. If I were to use something in a business setting, I probably would try cursor. My associates have good things to say about it. One of the problems I saw with this method, is that it changed the calculations for F, T, and P.

### Browser Chat (GPT 4o) https://chatgpt.com
(ChatGPT_Conversation.txt) for conversation

For me, I think splitting IDE and AI into seperate windows and interfaces helped me get a better results. I learned from the co-pilot attempt as well, and made adjustments from the get-go.
It produced a number of other metrics while keeping to my strict guidelines.

### Manual Coding and General Thoughts
Manual coding is essential for thoroughly learning a language and understanding its intricacies. However, once a developer becomes proficient, leveraging AI tools becomes invaluable. Routine code functions that would typically take 30 minutes can be completed by an AI agent in seconds, saving time and reducing mental strain. Using natural language prompts with AI agents can enhance code readability through better variable naming, inline comments, and alternative coding approaches. These agents also assist in learning, refactoring, and quickly implementing ideas for new use cases. I recommend a hybrid approach for all developers—manual coding to deepen knowledge and AI assistance to boost productivity. Just as mathematicians now use advanced tools like Excel or Python instead of solely relying on calculators, developers should embrace AI agents to streamline their workflow and enhance output.

### Future Assignment Thoughts
Students need to be warned to sign up for github Co-Pilot 2 weeks in advanced. Github takes up to 9 days to process the free student version.
Guidelines need to be be moved from copy paste, to specific to the assignment. I would think I am going to be graded on other items other than the guidelines provided for this assignment