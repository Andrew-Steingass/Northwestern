# CSV to JSON Lines Converter

This Go application converts a local CSV file to a JSON Lines (.jl) file.

## Prerequisites

- Go installed on your system
- Access to the CSV file you want to convert
- Write permissions in the directory where you want to save the output file

## Usage

### Step 1: Navigate to the Repository

Open your command line interface and navigate to the repository directory (example below):
    cd C:\Users\{username}\Documents\Github\CMD_line_csv_reader

### Step 2: Run the Application

Use the following command to run the application, adjusting the paths as necessary:
go run main.go --input "C:\Users\Andy\Downloads\housesInput.csv" --output "C:\Users\Andy\Downloads\housesOutput.jl"
- Replace `"C:\Users\{username}\Downloads\housesInput.csv"` with the path to your input CSV file.
- Replace `"C:\Users\{username}\Downloads\housesOutput.jl"` with the desired path and filename for your output JSON Lines file.

## Notes

- Ensure that the input CSV file exists at the specified path.
- Make sure you have write permissions in the directory where you're saving the output file.
- The application will validate inputs, read the CSV, convert it to JSON Lines format, and save the output file.

## Error Handling

If you encounter any errors, the application will display relevant error messages. Common issues include:
- Missing input or output file paths
- Non-existent input file
- Non-existent output directory

For any issues, please check the error messages and ensure all paths are correct and accessible.
