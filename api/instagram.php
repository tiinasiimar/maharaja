<?php
/**
 * Instagram API Endpoint (PHP version)
 * Returns Instagram posts from instagram.json
 */

// Enable CORS
header('Access-Control-Allow-Origin: *');
header('Access-Control-Allow-Methods: GET, OPTIONS');
header('Access-Control-Allow-Headers: X-Requested-With, Content-Type, Authorization');
header('Content-Type: application/json');

// Handle preflight OPTIONS request
if ($_SERVER['REQUEST_METHOD'] === 'OPTIONS') {
    http_response_code(200);
    exit;
}

/**
 * Get mock posts from instagram.json file
 */
function getMockPosts() {
    $jsonFile = __DIR__ . '/../instagram.json';

    if (!file_exists($jsonFile)) {
        error_log("instagram.json not found");
        return getDefaultPosts();
    }

    $jsonContent = file_get_contents($jsonFile);
    $urls = json_decode($jsonContent, true);

    if (!$urls || !is_array($urls)) {
        error_log("Failed to parse instagram.json");
        return getDefaultPosts();
    }

    $posts = [];
    foreach ($urls as $index => $url) {
        // Extract post ID from URL
        $url = rtrim($url, '/');
        $parts = explode('/', $url);
        $postID = end($parts);

        $posts[] = [
            'id' => $postID,
            'permalink' => $url,
            'caption' => 'Maharaja Restaurant - Post ' . ($index + 1),
            'media_type' => 'IMAGE',
            'media_url' => '',
            'timestamp' => date('c', strtotime('-' . ($index * 24) . ' hours'))
        ];
    }

    return !empty($posts) ? $posts : getDefaultPosts();
}

/**
 * Get default hardcoded posts as fallback
 */
function getDefaultPosts() {
    return [
        [
            'id' => 'DCr_gNaIlnf',
            'permalink' => 'https://www.instagram.com/p/DCr_gNaIlnf/',
            'caption' => 'Welcome to Maharaja Restaurant! Experience authentic Indian cuisine in Tallinn Old Town',
            'media_type' => 'IMAGE',
            'media_url' => '',
            'timestamp' => date('c', strtotime('-24 hours'))
        ],
        [
            'id' => 'DCpyT8VoMzm',
            'permalink' => 'https://www.instagram.com/p/DCpyT8VoMzm/',
            'caption' => 'Fresh naan bread from our traditional tandoor oven! Perfect with our aromatic curries',
            'media_type' => 'IMAGE',
            'media_url' => '',
            'timestamp' => date('c', strtotime('-48 hours'))
        ],
        [
            'id' => 'DCnmVEAo7aJ',
            'permalink' => 'https://www.instagram.com/p/DCnmVEAo7aJ/',
            'caption' => 'Our signature tandoori chicken - marinated for 24 hours in authentic spices!',
            'media_type' => 'IMAGE',
            'media_url' => '',
            'timestamp' => date('c', strtotime('-72 hours'))
        ],
        [
            'id' => 'DClbQBsIF0Q',
            'permalink' => 'https://www.instagram.com/p/DClbQBsIF0Q/',
            'caption' => 'Biryani lovers, this one\'s for you! Fragrant basmati rice with tender meat and aromatic spices',
            'media_type' => 'IMAGE',
            'media_url' => '',
            'timestamp' => date('c', strtotime('-96 hours'))
        ]
    ];
}

// Main execution
try {
    $posts = getMockPosts();

    $response = [
        'data' => $posts
    ];

    echo json_encode($response);

} catch (Exception $e) {
    error_log("Error in instagram.php: " . $e->getMessage());

    http_response_code(500);
    echo json_encode([
        'error' => 'Internal server error',
        'data' => []
    ]);
}
