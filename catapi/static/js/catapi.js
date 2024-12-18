document.addEventListener("DOMContentLoaded", () => {
    const votingButton = document.getElementById("votingButton");
    const breedButton = document.getElementById("breedButton");
    const favButton = document.getElementById("favButton");
    const heartButton = document.getElementById("heartButton");
    const catImageContainer = document.getElementById("catImageContainer");
    const favoriteImagesContainer = document.getElementById("favoriteImagesContainer");
    const likeButton = document.getElementById("likeButton");
    const dislikeButton = document.getElementById("dislikeButton");

    // Retrieve favorite images from localStorage
    let favoriteImages = JSON.parse(localStorage.getItem("favoriteImages")) || [];

    // Function to fetch a new cat image
    function fetchNewCatImage() {
        fetch("/")
            .then((response) => response.text())
            .then((html) => {
                const parser = new DOMParser();
                const doc = parser.parseFromString(html, "text/html");
                const newImage = doc.getElementById("catImage").src;

                // Smooth fade-out and fade-in effect
                catImageContainer.classList.add("fade-out");
                setTimeout(() => {
                    document.getElementById("catImage").src = newImage;
                    catImageContainer.classList.remove("fade-out");
                }, 500);
            })
            .catch((error) => console.error("Error fetching image:", error));
    }

    // Event listener for the heart button (add to favorite images list)
    heartButton.addEventListener("click", () => {
        const imageUrl = document.getElementById("catImage").src;
        if (imageUrl && !favoriteImages.includes(imageUrl)) {
            favoriteImages.push(imageUrl);
            localStorage.setItem("favoriteImages", JSON.stringify(favoriteImages)); // Store in localStorage
            console.log("Added to favorites:", imageUrl);
        }

        // Fetch a new cat image after adding to favorites
        fetchNewCatImage();
    });

    // Event listeners for other buttons (like and dislike)
    likeButton.addEventListener("click", fetchNewCatImage);
    dislikeButton.addEventListener("click", fetchNewCatImage);

    // Function to show voting layout
    window.showVotingLayout = function () {
        document.getElementById("votingLayout").style.display = "block";
        document.getElementById("breedLayout").style.display = "none";
        document.getElementById("favoriteLayout").style.display = "none";
    };

    // Function to show breed layout
    window.showBreedLayout = function () {
        document.getElementById("votingLayout").style.display = "none";
        document.getElementById("breedLayout").style.display = "block";
        document.getElementById("favoriteLayout").style.display = "none";
    };

    // Function to show favorite layout (all favorite images)
    window.showFavoriteLayout = function () {
        if (favoriteImages.length > 0) {
            favoriteImagesContainer.innerHTML = ""; // Clear previous content
            favoriteImages.forEach((imageUrl) => {
                const img = document.createElement("img");
                img.src = imageUrl;
                img.alt = "Favorite Cat Image";
                img.classList.add("favorite-image");
                favoriteImagesContainer.appendChild(img);
            });
            document.getElementById("favoriteLayout").style.display = "block";
            document.getElementById("votingLayout").style.display = "none";
            document.getElementById("breedLayout").style.display = "none";
        } else {
            alert("No favorite images yet.");
        }
    };

    // Show the voting layout initially
    showVotingLayout();
});
