## Go-News-Aggregrator

A news aggregator written in Go using Washington Post's sitemap. Reduced idle server time by *6x* by using concurrency patterns such as channels and waitgroups

## Usage:
1. Clone the repository at a desired location
2. `cd Go-News-Aggregrator/newsaggregator/`
3. `go run aggregator.go`
4. Open a web browser and type `localhost:8000/agg/`
5. See the news articles from Washington Post
