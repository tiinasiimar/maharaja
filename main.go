package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// InstagramPost represents a single Instagram post
type InstagramPost struct {
	ID          string `json:"id"`
	Permalink   string `json:"permalink"`
	Caption     string `json:"caption"`
	MediaType   string `json:"media_type"`
	MediaURL    string `json:"media_url"`
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	Timestamp   string `json:"timestamp"`
}

// InstagramResponse represents the API response
type InstagramResponse struct {
	Data []InstagramPost `json:"data"`
}

// Server configuration
type Config struct {
	Port        string
	InstagramToken string
	Username    string
}

var config Config

func main() {
	// Load configuration from environment variables
	config = Config{
		Port:        getEnv("PORT", "8080"),
		InstagramToken: getEnv("INSTAGRAM_TOKEN", ""),
		Username:    getEnv("INSTAGRAM_USERNAME", ""),
	}

	// Create router
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/instagram", getInstagramPosts).Methods("GET")
	api.HandleFunc("/health", healthCheck).Methods("GET")

	// Serve static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))

	// Add CORS middleware
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)(r)

	// Add logging middleware
	loggedRouter := handlers.LoggingHandler(os.Stdout, corsHandler)

	log.Printf("Server starting on port %s", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, loggedRouter))
}

func getInstagramPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Try to fetch real Instagram posts first
	posts, err := fetchRealInstagramPosts()
	if err != nil || len(posts) == 0 {
		if err != nil {
			log.Printf("Failed to fetch real Instagram posts: %v", err)
		} else {
			log.Printf("No posts found from scraping, using fallback data")
		}
		// Fallback to mock data
		posts = getMockPosts()
	}

	response := InstagramResponse{Data: posts}
	json.NewEncoder(w).Encode(response)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// fetchRealInstagramPosts attempts to scrape Instagram posts from @maharajarestoran
func fetchRealInstagramPosts() ([]InstagramPost, error) {
	// Instagram profile URL
	profileURL := "https://www.instagram.com/maharajarestoran/"
	
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	// Make request to Instagram profile
	req, err := http.NewRequest("GET", profileURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	
	// Set headers to mimic a real browser
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Instagram profile: %v", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Instagram returned status code: %d", resp.StatusCode)
	}
	
	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	
	// Parse Instagram posts from HTML
	posts, err := parseInstagramPosts(string(body))
	if err != nil {
		return nil, fmt.Errorf("failed to parse Instagram posts: %v", err)
	}
	
	return posts, nil
}

// parseInstagramPosts extracts post data from Instagram HTML
func parseInstagramPosts(html string) ([]InstagramPost, error) {
	var posts []InstagramPost
	
	// Regex to find Instagram post URLs in the HTML
	// This is a simplified approach - Instagram's actual structure is more complex
	postRegex := regexp.MustCompile(`https://www\.instagram\.com/p/[A-Za-z0-9_-]+/`)
	matches := postRegex.FindAllString(html, 6) // Limit to 6 posts
	
	for i, match := range matches {
		// Extract post ID from URL
		parts := strings.Split(match, "/")
		postID := parts[len(parts)-2]
		
		post := InstagramPost{
			ID:        postID,
			Permalink: match,
			Caption:   fmt.Sprintf("Maharaja Restaurant - Post %d", i+1),
			MediaType: "IMAGE",
			MediaURL:  "", // Would need additional parsing for actual media URLs
			Timestamp: time.Now().Add(-time.Duration(i*24) * time.Hour).Format(time.RFC3339),
		}
		
		posts = append(posts, post)
	}
	
	return posts, nil
}

// getMockPosts returns fallback mock data from instagram.json file
func getMockPosts() []InstagramPost {
	// Try to read from instagram.json file
	file, err := os.ReadFile("instagram.json")
	if err != nil {
		log.Printf("Failed to read instagram.json: %v", err)
		return getDefaultPosts()
	}

	var urls []string
	if err := json.Unmarshal(file, &urls); err != nil {
		log.Printf("Failed to parse instagram.json: %v", err)
		return getDefaultPosts()
	}

	var posts []InstagramPost
	for i, url := range urls {
		// Extract post ID from URL
		parts := strings.Split(strings.TrimSuffix(url, "/"), "/")
		postID := parts[len(parts)-1]

		post := InstagramPost{
			ID:        postID,
			Permalink: url,
			Caption:   fmt.Sprintf("Maharaja Restaurant - Post %d", i+1),
			MediaType: "IMAGE",
			MediaURL:  "",
			Timestamp: time.Now().Add(-time.Duration(i*24) * time.Hour).Format(time.RFC3339),
		}
		posts = append(posts, post)
	}

	if len(posts) == 0 {
		return getDefaultPosts()
	}

	return posts
}

// getDefaultPosts returns hardcoded fallback posts if instagram.json doesn't work
func getDefaultPosts() []InstagramPost {
	return []InstagramPost{
		{
			ID:        "DCr_gNaIlnf",
			Permalink: "https://www.instagram.com/p/DCr_gNaIlnf/",
			Caption:   "Welcome to Maharaja Restaurant! Experience authentic Indian cuisine in Tallinn Old Town",
			MediaType: "IMAGE",
			MediaURL:  "",
			Timestamp: time.Now().Add(-24 * time.Hour).Format(time.RFC3339),
		},
		{
			ID:        "DCpyT8VoMzm",
			Permalink: "https://www.instagram.com/p/DCpyT8VoMzm/",
			Caption:   "Fresh naan bread from our traditional tandoor oven! Perfect with our aromatic curries",
			MediaType: "IMAGE",
			MediaURL:  "",
			Timestamp: time.Now().Add(-48 * time.Hour).Format(time.RFC3339),
		},
		{
			ID:        "DCnmVEAo7aJ",
			Permalink: "https://www.instagram.com/p/DCnmVEAo7aJ/",
			Caption:   "Our signature tandoori chicken - marinated for 24 hours in authentic spices!",
			MediaType: "IMAGE",
			MediaURL:  "",
			Timestamp: time.Now().Add(-72 * time.Hour).Format(time.RFC3339),
		},
		{
			ID:        "DClbQBsIF0Q",
			Permalink: "https://www.instagram.com/p/DClbQBsIF0Q/",
			Caption:   "Biryani lovers, this one's for you! Fragrant basmati rice with tender meat and aromatic spices",
			MediaType: "IMAGE",
			MediaURL:  "https://example.com/biryani.jpg",
			Timestamp: time.Now().Add(-96 * time.Hour).Format(time.RFC3339),
		},
		{
			ID:        "5",
			Permalink: "https://www.instagram.com/p/example5/",
			Caption:   "Located in the heart of Tallinn Old Town - Harju tn 7. Come visit us! 🏰📍",
			MediaType: "IMAGE",
			MediaURL:  "https://example.com/location.jpg",
			Timestamp: time.Now().Add(-120 * time.Hour).Format(time.RFC3339),
		},
	}
}
