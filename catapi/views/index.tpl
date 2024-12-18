<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Cat Voting</title>
  <link rel="stylesheet" href="/static/css/catapi.css">
  <script src="/static/js/catapi.js" defer></script>
</head>
<body>
  <div class="container">
    <h1>Cat Voting</h1>

    <!-- Buttons for switching between Voting and Breed -->
    <div class="button-container">
      <button id="votingButton" class="button" title="Voting" onclick="showVotingLayout()">Voting</button>
      <button id="breedButton" class="button" title="Breed" onclick="showBreedLayout()">Breed</button>
      <button id="favButton" class="button" title="Favorite" onclick="showFavoriteLayout()">❤️</button>
    </div>

    <!-- Main Card Container -->
    <div id="catCard" class="cat-card">
      <div id="votingLayout" class="layout">
        <div id="catImageContainer" class="image-container">
          {{if .CatImage}}
            <img id="catImage" src="{{.CatImage}}" alt="Cute Cat Image">
          {{else}}
            <p>No Image Available</p>
          {{end}}
        </div>
        <div class="vote-buttons">
          <button id="heartButton" class="button" title="Favorite">❤️</button>
          <button id="likeButton" class="button" title="Like">👍</button>
          <button id="dislikeButton" class="button" title="Dislike">👎</button>
        </div>
      </div>

      <!-- Breed Layout (Initially Hidden) -->
      <div id="breedLayout" class="layout" style="display:none;">
        <h2>Breed Information</h2>
        <p>Here is the breed information...</p>
      </div>

      <!-- Favorite Layout (Initially Hidden) -->
      <div id="favoriteLayout" class="layout" style="display:none;">
        <h2>Your Favorite Cats</h2>
        <div id="favoriteImagesContainer" class="image-container">
          <!-- All Favorite Images will be appended here -->
        </div>
      </div>
    </div>
  </div>
</body>
</html>
