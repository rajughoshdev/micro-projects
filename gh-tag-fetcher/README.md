
# GitHub Latest Release Tag Fetcher

This is a simple Go program that fetches the latest release tag of GitHub repositories using the GitHub API. It demonstrates how to make HTTP requests, handle JSON responses, and extract the latest release tag from the API data.

## Prerequisites

- Go programming language (https://golang.org/doc/install)

## Usage

1. Clone the repository or copy the code snippet from this file.

2. Replace `"fluxcd"` with your GitHub repository owner's username.

3. Modify the `repositories` array to include the names of the repositories you want to fetch the latest release tags for.

4. Run the program using the following command:

   ```bash
   go run main.go
   ```

   The program will display the latest release tags for the specified repositories.
