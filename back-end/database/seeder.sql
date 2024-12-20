-- Insert users
INSERT INTO users (username, email, password) VALUES
('john_doe', 'john@example.com', 'password123'),
('jane_smith', 'jane@example.com', 'password456'),
('alice_wonder', 'alice@example.com', 'password789'),
('yoofahdagchill', 'faguenou@gmail.com', 'fahd12345');


-- Insert categories
INSERT INTO categories (name) VALUES
('Technology'),
('Health'),
('Travel'),
('Education'),
('Entertainment');

-- Insert posts
INSERT INTO posts (user_id, title, content) VALUES
(1, 'How to Learn Go', 'Go is a powerful language for web development. Here are some tips to get started.'),
(2, 'Healthy Eating Tips', 'Discover how to improve your health with better eating habits.'),
(3, 'Top 10 Travel Destinations', 'Check out these amazing travel destinations for 2024.');

-- Insert post categories
INSERT INTO post_categories (post_id, category_id) VALUES
(1, 1), -- Post 1 belongs to "Technology"
(2, 2), -- Post 2 belongs to "Health"
(3, 3); -- Post 3 belongs to "Travel"

-- Insert comments
INSERT INTO comments (post_id, user_id, content) VALUES
(1, 2, 'Thanks for the tips on Go! Very helpful.'),
(1, 3, 'I started learning Go recently and love it!'),
(2, 1, 'Great advice! I will try to follow these tips.'),
(3, 2, 'I have been to 5 of these destinations, and they are amazing!');

-- Insert likes (posts and comments)
INSERT INTO likes (user_id, post_id, is_like) VALUES
(2, 1, TRUE), -- User 2 likes Post 1
(3, 1, TRUE), -- User 3 likes Post 1
(1, 2, TRUE); -- User 1 likes Post 2

