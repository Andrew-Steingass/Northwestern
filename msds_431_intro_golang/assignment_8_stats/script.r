# Load necessary libraries
if (!requireNamespace("readr", quietly = TRUE)) {
    install.packages("readr")
}
library(readr)

# Function to generate data
generate_data <- function(n, shape) {
  if (shape == "positively_skewed") {
    return(rexp(n, rate = 1))  # Exponential distribution (positive skew)
  } else if (shape == "symmetric") {
    return(rnorm(n, mean = 0, sd = 1))  # Normal distribution (symmetric)
  } else if (shape == "negatively_skewed") {
    return(-rexp(n, rate = 1))  # Negative exponential distribution
  } else {
    stop("Unknown distribution shape.")
  }
}

# Set seed for reproducibility
set.seed(123)

# Parameters
sample_size <- 25  # Only size 25
distributions <- c("positively_skewed", "symmetric", "negatively_skewed")
B <- 100  # Number of bootstrap samples

# Generate raw datasets and save them
for (shape in distributions) {
  # Generate data
  data <- generate_data(sample_size, shape)
  
  # Replace any NA values with 0 (unlikely but defensive programming)
  data[is.na(data)] <- 0
  
  # Save the dataset
  data_file <- paste0(shape, "_data_", sample_size, ".csv")
  write.csv(data, file = data_file, row.names = FALSE, col.names = FALSE)
}

# Generate and save resampling indices
indices <- replicate(B, sample(1:sample_size, sample_size, replace = TRUE))  # Resampling indices
indices[is.na(indices)] <- 1  # Replace NA indices with the first row (defensive programming)
indices_file <- paste0("resampling_indices_", sample_size, ".csv")
write.csv(indices, file = indices_file, row.names = FALSE)

# Compute and save medians and SEMedian
comparison_results <- data.frame()

# Load resampling indices
indices <- as.matrix(read_csv(indices_file, show_col_types = FALSE))

for (shape in distributions) {
  # Load the dataset
  data_file <- paste0(shape, "_data_", sample_size, ".csv")
  data_vec <- as.numeric(read_csv(data_file, col_names = FALSE)[[1]])
  
  # Replace NA values in the data vector with 0 (defensive programming)
  data_vec[is.na(data_vec)] <- 0
  
  # Compute bootstrap medians
  bootstrap_medians <- apply(indices, 2, function(col_idx) {
    resampled <- data_vec[col_idx]
    resampled[is.na(resampled)] <- 0  # Replace any NA in resampled data
    median(resampled)
  })
  
  # Save the medians
  median_file <- paste0(shape, "_medians_", sample_size, ".csv")
  write.csv(bootstrap_medians, file = median_file, row.names = FALSE, col.names = FALSE)
  
  # Compute SEMedian
  se_median <- sd(bootstrap_medians, na.rm = TRUE)
  
  # Append results to comparison table
  comparison_results <- rbind(comparison_results, data.frame(
    Distribution = shape,
    SampleSize = sample_size,
    SEMedian = se_median
  ))
}

# Save comparison results
comparison_file <- "r_standard_errors.csv"
write.csv(comparison_results, file = comparison_file, row.names = FALSE)

# Print results in a tabular format
cat("\nSEMedian Results:\n")
print(comparison_results)

cat("\nAll datasets, indices, medians, and comparisons for size 25 have been generated without NA values.\n")

