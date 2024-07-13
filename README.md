# 3.-URL-Shortener

URL Shortener
Overview
This project implements a URL shortening service in Go, designed to generate short URLs using MD5 hashing and provide efficient redirection to original URLs via RESTful APIs.

Features
URL Shortening: Converts long URLs into short, unique identifiers.
Redirection: Handles HTTP requests to redirect users from short URLs to their original destinations.
RESTful APIs: Supports endpoints for creating short URLs and resolving them to original URLs.
Dynamic Caching: Optimizes storage and retrieval using a dynamic caching mechanism.

Installation
Clone the Repository:
git clone https://github.com/your-username/url-shortener.git
cd url-shortener

Install Dependencies:
This project requires Go to be installed. Ensure you have Go installed on your system.

Build and Run:
go build
./url-shortener
The server will start running at http://localhost:8080.

Shortening a URL
Send a POST request to /shorten endpoint with JSON payload:
{
  "url": "https://example.com/very-long-url-path"
}
curl -X POST http://localhost:8080/shorten -d '{"url":"https://example.com/very-long-url-path"}' -H "Content-Type: application/json"
Redirecting to Original URL
Access the short URL generated to be redirected to the original URL.
