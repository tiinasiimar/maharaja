# Maharaja Restaurant - Deployment Guide

This project has TWO backend options:

## Option 1: Go Backend (Current - For VPS/Server)

**Files:**
- `main.go` - Go server
- `go.mod`, `go.sum` - Go dependencies

**To Run Locally:**
```bash
go run main.go
```

**To Deploy:**
- Requires VPS or server with Go installed
- Build: `go build -o maharaja main.go`
- Run: `./maharaja`

**API Endpoints:**
- `http://localhost:8080/api/instagram` - Get Instagram posts
- `http://localhost:8080/api/health` - Health check
- `http://localhost:8080/` - Static files (HTML, CSS, JS, images)

---

## Option 2: PHP Backend (For zone.ee / Shared Hosting)

**Files:**
- `api/instagram.php` - Instagram posts endpoint
- `api/health.php` - Health check endpoint
- `api/.htaccess` - URL rewriting (makes /api/instagram work)

**To Deploy on zone.ee:**

1. Upload all files to your web hosting root directory

2. Your file structure should be:
```
public_html/
├── index.html
├── index-ee.html
├── privaatsuspoliitika.html
├── instagram.json
├── css/
│   └── styles.css
├── js/
│   └── script.js
├── images/
│   └── (all images)
└── api/
    ├── .htaccess
    ├── instagram.php
    └── health.php
```

3. **The JavaScript will automatically work** - No changes needed!
   Your `js/script.js` calls `/api/instagram` which will be handled by PHP

**API Endpoints (after deployment):**
- `https://yourdomain.com/api/instagram` - Get Instagram posts
- `https://yourdomain.com/api/health` - Health check

---

## Comparison

| Feature | Go Backend | PHP Backend |
|---------|------------|-------------|
| **Hosting** | VPS/Server required | Standard shared hosting |
| **Performance** | Faster | Slightly slower |
| **Setup** | More complex | Simple upload |
| **Cost** | VPS hosting (~€5-20/month) | Shared hosting (~€2-5/month) |
| **zone.ee Compatible** | Only with VPS package | ✅ Yes, standard hosting |

---

## For zone.ee Deployment (Recommended)

**Use the PHP backend:**

1. **Upload via FTP or File Manager:**
   - All HTML, CSS, JS, images files
   - The `api/` folder with PHP files
   - `instagram.json`
   - `privaatsuspoliitika.html`

2. **No configuration needed** - It just works!

3. **Test:**
   - Visit: `https://yourdomain.com/`
   - Check API: `https://yourdomain.com/api/instagram`

---

## Instagram Posts Management

Both backends read from `instagram.json`:

```json
[
  "https://www.instagram.com/maharajarestoran/p/POST_ID_1",
  "https://www.instagram.com/maharajarestoran/p/POST_ID_2"
]
```

**To update posts:**
1. Edit `instagram.json`
2. Add/remove Instagram post URLs
3. Save and upload

---

## Privacy Policy

The privacy policy is available at:
- `https://yourdomain.com/privaatsuspoliitika.html`

Use this URL when setting up Facebook/Instagram API.
