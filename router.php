<?php
/**
 * Router script for PHP built-in development server
 * This simulates Apache's .htaccess URL rewriting
 */

$uri = parse_url($_SERVER['REQUEST_URI'], PHP_URL_PATH);

// Route /api/instagram to /api/instagram.php
if ($uri === '/api/instagram') {
    require __DIR__ . '/api/instagram.php';
    return true;
}

// Route /api/health to /api/health.php
if ($uri === '/api/health') {
    require __DIR__ . '/api/health.php';
    return true;
}

// For everything else, let PHP serve the file normally
return false;
