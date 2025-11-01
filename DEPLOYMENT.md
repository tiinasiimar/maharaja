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
├── .htaccess                    (Security headers & performance)
├── .gitignore                   (For git - optional)
├── index.html                   (English homepage)
├── index-ee.html                (Estonian homepage)
├── privaatsuspoliitika.html     (Privacy policy)
├── instagram.json               (Instagram post URLs)
├── css/
│   └── styles.css
├── js/
│   └── script.js
├── images/
│   ├── new_maharaja_logo.png
│   └── (other images)
└── api/
    ├── .htaccess                (API URL rewriting)
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

## Security Configuration

The website includes comprehensive security measures via `.htaccess`:

### Security Headers Implemented

1. **Content Security Policy (CSP)**
   - Restricts content sources to prevent XSS attacks
   - Allows Google Fonts and Instagram embeds only

2. **Clickjacking Protection**
   - X-Frame-Options: SAMEORIGIN
   - Prevents site from being embedded in iframes on other domains

3. **MIME-Sniffing Prevention**
   - X-Content-Type-Options: nosniff
   - Prevents browsers from interpreting files as different MIME type

4. **XSS Protection**
   - X-XSS-Protection: 1; mode=block
   - Enables browser's built-in XSS filter

5. **HTTPS Enforcement**
   - Strict-Transport-Security (HSTS)
   - Forces HTTPS for 1 year
   - Automatic HTTP → HTTPS redirect

6. **Privacy Controls**
   - Referrer-Policy: strict-origin-when-cross-origin
   - Permissions-Policy: Disables geolocation, microphone, camera, payment APIs

7. **Server Signature Hiding**
   - Removes Apache version information
   - Hides X-Powered-By header

### SEO Optimization

Both HTML files include:
- **Meta descriptions** for search engines
- **Open Graph tags** for social media sharing (Facebook, LinkedIn)
- **Twitter Card tags** for Twitter sharing
- **Schema.org structured data** (Restaurant type with address, hours, phone)
- **Canonical URLs** to prevent duplicate content issues
- **Bilingual support** with proper hreflang tags

### Performance Optimizations

The `.htaccess` file includes:
- **Gzip compression** for text files (HTML, CSS, JS)
- **Browser caching** with appropriate expiry times:
  - Images: 1 year
  - CSS/JS: 1 month
  - HTML: 1 hour
  - JSON: 1 day

### File Protection

Sensitive files are blocked from direct access:
- `.htaccess`, `.gitignore`, `.env`
- `.git/`, `.idea/`, `.DS_Store`
- `go.mod`, `go.sum`, `main.go`, `router.php`
- `README.md`, `DEPLOYMENT.md`

### Testing Security

After deployment, test your security headers:
1. Visit: https://securityheaders.com/
2. Enter: `https://www.maharaja.ee`
3. Should get A or A+ rating

Test SEO and structured data:
1. Google Rich Results Test: https://search.google.com/test/rich-results
2. Enter your URL to verify Schema.org markup

### GDPR Compliance

✅ **No cookie banner needed** - The website:
- Does not set any cookies
- Does not use analytics or tracking
- Google Fonts doesn't set cookies
- Instagram feed is static (no embedded iframes)
- Privacy policy available at `/privaatsuspoliitika.html`

---

## For zone.ee Deployment (Recommended)

**Use the PHP backend:**

1. **Upload via FTP or File Manager:**
   - `.htaccess` (root directory - security headers)
   - All HTML files (index.html, index-ee.html, privaatsuspoliitika.html)
   - `instagram.json`
   - CSS, JS, and images folders
   - The `api/` folder with PHP files and its .htaccess

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
