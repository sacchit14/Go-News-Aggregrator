## Go-News-Aggregrator

A news aggregator written in Go using Washington Post's sitemap

## Motivation

After taking a course in operating systems and concurrency in C, I was interested in learning more about concurrency patterns. I learned that Go had excellent concurrency resources and I decided to learn about them. By using waitgroups and go rountines in this project, the sitemap was loaded  **6x** faster

## Screenshots


![Alt text](/image1.png?raw=true "Full Sitemap")






![Alt text](/image2.png?raw=true "News Article")


## Getting Started

These instructions would help you download the project on your local machine for development and testing purposes


### Prerequsites

Have the Go compiler (precompiled Go binary package) installed on your local machine

### Installation

Download the precompiled Go binary package according to the requirements of your machine from here - https://golang.org/dl/

### Usage
After installing the Go Binary Package follow the steps below to run the project

1. Clone the repository at a desired location
2. `cd Go-News-Aggregrator/newsaggregator/`
3. `go run aggregator.go`
4. Open a web browser and type `localhost:8000/agg/`
5. See the news articles from Washington Post
