function handleReaction(postID, reactionType) {
    // Prepare data to be sent
    const data = {
        post_id:  parseInt(postID),
        reaction: reactionType
    };
    console.log("data sent : " + JSON.stringify(data));

    // Use Fetch API to send a POST request
    fetch('/reaction', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    
    .then(response => {
        console.log("response : " + JSON.stringify(response));
        
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(result => {
        // Update the UI with new like/dislike counts
        document.querySelector(`.likes[data-post-id="${postID}"]`).textContent = "👍"+result.likes;
        document.querySelector(`.dislikes[data-post-id="${postID}"]`).textContent = "👎"+result.dislikes;
    })
    .catch(error => {
        console.error('Error:', error);
        alert("There was an error processing your reaction. Please try again.");
    });
}


function toggleComments(postID) {
    const commentSection = document.getElementById('comments-section-' + postID);
    
    if (commentSection.style.display === "none") {
        commentSection.style.display = "block";  // Show comments
        document.getElementById('new-comment-' + postID).style.display = "block"; // Show input for new comment
    } else {
        commentSection.style.display = "none";   // Hide comments
    }
}

function addComment(postID) {
    const commentInput = document.getElementById('comment-input-' + postID);
    const commentContent = commentInput.value;

    if (!commentContent) {
        alert("Please enter a comment.");
        return;
    }

    // Prepare data to be sent
    const data = {
        post_id: parseInt(postID),
        content: commentContent
    };

    // Use Fetch API to send a POST request
    fetch('/add-comment', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(result => {
        // Update the UI with the new comment
        const commentsSection = document.getElementById('comments-section-' + postID);
        
        // Create a new comment card element
        const newCommentCard = document.createElement('div');
        newCommentCard.className = 'comment-card';
        
        newCommentCard.innerHTML = `<p><strong>${result.username}</strong>: ${result.content}</p><em>Posted just now</em>`;
        
        commentsSection.prepend(newCommentCard);
        
        // Clear the input field
        commentInput.value = '';
        
    })
    .catch(error => {
        console.error('Error:', error);
        alert("There was an error adding your comment. Please try again.");
    });
}
