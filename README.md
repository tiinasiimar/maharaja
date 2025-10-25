# Maharaja Restaurant Website

A restaurant website with Instagram feed integration using Go backend.

## Features

- Responsive restaurant website
- Instagram feed integration
- Go server for dynamic content
- CORS-enabled API endpoints

## Setup

### Prerequisites

- Go 1.21 or later
- Modern web browser

### Running the Server

1. Install Go dependencies:
```bash
go mod tidy
```

2. Set environment variables (optional):
```bash
export PORT=8080
export INSTAGRAM_TOKEN=your_instagram_token
export INSTAGRAM_USERNAME=your_username
```

3. Run the server:
```bash
go run main.go
```

4. Open your browser and navigate to `http://localhost:8080`

## API Endpoints

- `GET /api/instagram` - Fetch Instagram posts
- `GET /api/health` - Health check endpoint

## Instagram Integration

Currently uses mock data. To integrate with real Instagram:

1. Create an Instagram App at https://developers.facebook.com/
2. Get Instagram Basic Display API access
3. Implement user authentication flow
4. Update the `fetchRealInstagramPosts()` function in `main.go`

## File Structure

```
maharaja/
├── main.go              # Go server
├── go.mod              # Go dependencies
├── index.html          # Main website
├── js/script.js        # Frontend JavaScript
├── css/styles.css      # Styling
├── images/             # Static images
└── instagram.json      # Static Instagram URLs (fallback)
```

