# Bootstrap Statistics Program
** Order matters
### Step 1
Run r.script   (this generates source files to compare with GoLang)
### Step 2
Run main.go with   go run .
### Notes
Writeup is at bottom

## Project Structure
```
assignment_8_stats/
├── main.go         # Main program code
├── logger.go       # Logging functionality
├── main_test.go    # Test suite
├── bootstrap.log   # Generated log file (after running)
├── cpu.prof        # CPU profile (when enabled)
└── mem.prof        # Memory profile (when enabled)
```

## Basic Commands

### Build Program (Required First Step)
```bash
go build -o bootstrap.exe
```

### Run Program
After building, run with:
```bash
.\bootstrap.exe
```

Or run without building:
```bash
go run .
```

## Testing Commands

### Run All Tests
```bash
go test
```

### Run Tests with Verbose Output
```bash
go test -v
```

### Run Specific Test
```bash
go test -run TestMedian
```

## Benchmarking Commands

### Run All Benchmarks
```bash
go test -bench=.
```

### Run Benchmarks with Memory Stats
```bash
go test -bench=. -benchmem
```

## Profiling Commands

### CPU Profiling
1. Build the program first (required):
```bash
go build -o bootstrap.exe
```

2. Run with CPU profiling enabled (use full path to avoid issues):
```bash
.\bootstrap.exe -cpuprofile="C:\your\full\path\cpu.prof"
```

3. Verify profile was created:
```bash
dir cpu.prof
```

4. Analyze CPU profile (interactive mode):
```bash
go tool pprof cpu.prof
```

5. Common pprof commands:
```
top         # Show top CPU users
top10       # Show top 10 CPU users
quit        # Exit pprof
```

### Memory Profiling
1. **Build the program first** (required):
```bash
go build -o bootstrap.exe
```

2. **Run with memory profiling enabled** (use full path to avoid issues):
```bash
.\bootstrap.exe -memprofile="C:\your\full\path\mem.prof"
```

3. **Verify profile was created**:
```bash
dir mem.prof
```

5. **Common pprof commands for memory profiling**:
   ```
   top         # Show top memory users
   top10       # Show top 10 memory users
   quit        # Exit pprof
   ```

## Program Output Files
- `bootstrap.log`: Program execution logs
- `cpu.prof`: CPU profile (when profiling enabled)
- `mem.prof`: Memory profile (when profiling enabled)


## Notes
- Always build the program before running profiling commands
- Use full paths for profile files to avoid creation issues
- The .exe extension is for Windows systems
- On Unix-based systems, omit the .exe extension



## Writeup

# Key R Packages for Bootstrap Resampling

- **`boot` Package**:
    - Provides extensive tools for bootstrap resampling and calculating statistics.
    - URL: [https://cran.r-project.org/web/packages/boot/index.html](https://cran.r-project.org/web/packages/boot/index.html)

- **`readr` Package**:
    - Used for efficient reading and writing of CSV files, especially for large datasets.
    - URL: [https://cran.r-project.org/web/packages/readr/index.html](https://cran.r-project.org/web/packages/readr/index.html)


### 1. What did you do to improve the performance of the Go implementation of the selected statistical method?

    I improved the performance of my Go implementation by using profiling (CPU and memory) to identify bottlenecks and guide optimization efforts. I added benchmark tests for key functions, enabling me to quantitatively measure performance. I optimized data handling by minimizing memory allocations and using Go’s efficient sorting algorithms. I also considered parallelism for tasks like bootstrap resampling and reduced the impact of logging by ensuring it only records important information. Additionally, I optimized statistical calculations, ensuring faster execution by avoiding unnecessary computations. These improvements make the code more efficient, scalable, and capable of handling larger datasets.

### 2. Describe your efforts in finding R and Go packages for the method. Review your process of building the Go implementation. Review your experiences with testing, benchmarking, software profiling, and logging.

    In my efforts to implement the method using both R and Go, I started by reviewing the available packages in R that could help with the statistical computations and resampling processes. For R, I used built-in packages like readr for reading CSV files and generating the data, and relied on simple functions for bootstrapping and calculating medians. While searching for Go packages, I realized that Go's standard library had sufficient tools for most tasks, such as math for statistical calculations and encoding/csv for reading and writing CSV files. I also explored Go’s capabilities for handling more complex operations like skewness calculation and bootstrapping, which I eventually implemented using custom functions. The most challenging part was ensuring that the Go implementation properly mirrored the statistical behavior and output of the R code, especially when dealing with data manipulation and the generation of resampling indices.

    Building the Go implementation involved translating the statistical concepts from R into Go code, and ensuring that the output structure matched the expected format. I focused on implementing the core functionality step-by-step, verifying that each part worked as expected before moving to the next. Testing the Go implementation involved comparing the generated results with those from R to ensure consistency in the medians and standard errors. I performed basic benchmarking to ensure that the Go code ran efficiently.

### 3. Under what circumstances would it make sense for the firm to use Go in place of R for the selected statistical method? Select a cloud provider of infrastructure as a service (IaS). Note the cloud costs for virtual machine (compute engine) services. What percentage of cloud computing costs might be saved with a move from R to Go?

    The research consultancy should carefully evaluate their options based on the size of the project, the need for speed, accuracy, and quick development. For smaller projects with a focus on rapid prototyping and in-depth statistical analysis, R remains a strong choice due to its vast ecosystem of packages and ease of use. However, for larger-scale projects or those requiring high performance and scalability, Go could offer significant advantages in speed and cloud efficiency. I would recommend initially developing a toolkit of statistical methods in GoLang that mimics R. Once the toolkit is mature, the consultancy can consider transitioning to Go, ensuring that core functionalities are re-implemented with performance optimizations, allowing them to meet both scalability and efficiency demands while maintaining accuracy.

    When comparing the cost savings between using R and Go on Google Cloud Platform (GCP), the key difference lies in Go's efficiency in resource usage, which directly translates to cost savings. GCP’s n1-standard-4 VM, which includes 4 vCPUs and 15 GB of RAM, costs approximately $0.15 per hour. Running this instance continuously for a year costs $1,296 (24 hours x 365 days x $0.15). Since Go is a compiled language and typically runs faster and more efficiently than R, it could reduce the required compute time and resources by 10% to 30%. For example, if Go reduces the compute time by 20%, the annual cloud cost would drop from $1,296 to $907.20, saving approximately $388.80 per year. Therefore, switching from R to Go could lead to 10% to 30% savings in cloud computing costs, making Go a more cost-effective choice for high-performance, scalable projects on GCP.