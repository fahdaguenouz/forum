function toggleComments(postID) {
    const commentSection = document.getElementById('comments-section-' + postID);
    
    if (commentSection.style.display === "none") {
        commentSection.style.display = "block";  // Show comments
    } else {
        commentSection.style.display = "none";   // Hide comments
    }
}