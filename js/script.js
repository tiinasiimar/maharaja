// Gallery lightbox functionality with swipe support
document.addEventListener('DOMContentLoaded', function() {
    const galleryItems = document.querySelectorAll('.gallery-item');

    // Create lightbox modal
    const lightbox = document.createElement('div');
    lightbox.className = 'lightbox';
    lightbox.innerHTML = `
        <div class="lightbox-content">
            <span class="lightbox-close">&times;</span>
            <img class="lightbox-image" src="" alt="">
            <div class="lightbox-nav">
                <button class="lightbox-prev">‹</button>
                <button class="lightbox-next">›</button>
            </div>
        </div>
    `;
    document.body.appendChild(lightbox);

    let currentImageIndex = 0;
    const images = Array.from(galleryItems).map(item => ({
        src: item.dataset.src,
        alt: item.querySelector('img').alt
    }));

    // Touch/swipe variables
    let startX = 0;
    let startY = 0;
    let endX = 0;
    let endY = 0;
    let isSwipe = false;

    // Open lightbox
    galleryItems.forEach((item, index) => {
        item.addEventListener('click', () => {
            currentImageIndex = index;
            openLightbox();
        });
    });

    // Close lightbox
    lightbox.addEventListener('click', (e) => {
        if (e.target === lightbox || e.target.classList.contains('lightbox-close')) {
            closeLightbox();
        }
    });

    // Navigation
    const prevBtn = lightbox.querySelector('.lightbox-prev');
    const nextBtn = lightbox.querySelector('.lightbox-next');

    prevBtn.addEventListener('click', (e) => {
        e.stopPropagation();
        currentImageIndex = (currentImageIndex - 1 + images.length) % images.length;
        updateLightboxImage();
    });

    nextBtn.addEventListener('click', (e) => {
        e.stopPropagation();
        currentImageIndex = (currentImageIndex + 1) % images.length;
        updateLightboxImage();
    });

    // Touch events for swiping
    lightbox.addEventListener('touchstart', (e) => {
        startX = e.touches[0].clientX;
        startY = e.touches[0].clientY;
        isSwipe = false;
    }, { passive: true });

    lightbox.addEventListener('touchmove', (e) => {
        if (!isSwipe) {
            endX = e.touches[0].clientX;
            endY = e.touches[0].clientY;

            const diffX = Math.abs(startX - endX);
            const diffY = Math.abs(startY - endY);

            // Determine if it's a horizontal swipe
            if (diffX > diffY && diffX > 50) {
                isSwipe = true;
            }
        }
    }, { passive: true });

    lightbox.addEventListener('touchend', (e) => {
        if (isSwipe) {
            const diffX = startX - endX;
            const minSwipeDistance = 50;

            if (Math.abs(diffX) > minSwipeDistance) {
                if (diffX > 0) {
                    // Swipe left - next image
                    nextBtn.click();
                } else {
                    // Swipe right - previous image
                    prevBtn.click();
                }
            }
        }
        isSwipe = false;
    }, { passive: true });

    // Keyboard navigation
    document.addEventListener('keydown', (e) => {
        if (lightbox.classList.contains('active')) {
            if (e.key === 'Escape') closeLightbox();
            if (e.key === 'ArrowLeft') prevBtn.click();
            if (e.key === 'ArrowRight') nextBtn.click();
        }
    });

    function openLightbox() {
        updateLightboxImage();
        lightbox.classList.add('active');
        document.body.style.overflow = 'hidden';
    }

    function closeLightbox() {
        lightbox.classList.remove('active');
        document.body.style.overflow = '';
    }

    function updateLightboxImage() {
        const img = lightbox.querySelector('.lightbox-image');
        img.src = images[currentImageIndex].src;
        img.alt = images[currentImageIndex].alt;
    }
});