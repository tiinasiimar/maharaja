// Instagram feed
document.addEventListener('DOMContentLoaded', function() {
    console.log('Script loaded! Looking for instagram-feed element...');
    const feedContainer = document.getElementById('instagram-feed');
    if (!feedContainer) {
        console.error('ERROR: instagram-feed element not found!');
        return;
    }
    console.log('Found instagram-feed element, fetching posts...');

    // Helper to create a card for Instagram post
    function createInstaCard(post) {
        const card = document.createElement('a');
        card.href = post.permalink;
        card.target = '_blank';
        card.rel = 'noopener noreferrer';
        card.className = 'instagram-card';

        const content = document.createElement('div');
        content.className = 'instagram-card-content';

        const icon = document.createElement('div');
        icon.className = 'instagram-icon';
        icon.innerHTML = '📷';

        const caption = document.createElement('p');
        caption.textContent = post.caption || 'View on Instagram';

        const viewLink = document.createElement('span');
        viewLink.className = 'view-link';
        viewLink.textContent = 'View Post →';

        content.appendChild(icon);
        content.appendChild(caption);
        content.appendChild(viewLink);
        card.appendChild(content);

        return card;
    }

    // Load posts from Go server API
    fetch('/api/instagram', {
        cache: 'no-store',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        }
    })
        .then(function(res) {
            console.log('Instagram API response status:', res.status);
            if (!res.ok) throw new Error('Failed to fetch Instagram posts');
            return res.json();
        })
        .then(function(data) {
            console.log('Instagram API data:', data);
            const posts = data.data || [];
            console.log('Number of posts:', posts.length);

            if (posts.length === 0) {
                feedContainer.innerHTML = '<p>Follow us on Instagram for the latest updates.</p>';
                return;
            }

            // Render up to 6 posts as cards
            posts.slice(0, 6).forEach(function(post) {
                console.log('Adding post:', post.permalink);
                const card = createInstaCard(post);
                feedContainer.appendChild(card);
            });
        })
        .catch(function(err) {
            console.error('Error loading Instagram posts:', err);
            feedContainer.innerHTML = '<p>Unable to load Instagram posts right now.</p>';
        });
});