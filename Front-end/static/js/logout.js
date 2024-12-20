
    document.getElementById('logoutButton').addEventListener('click', async () => {
        try {
            const response = await fetch('/logout', {
                method: 'POST',
                credentials: 'include' // Ensure cookies are included
            });

            if (response.ok) {
                // Redirect to the login page on successful logout
                window.location.href = '/authentification';
            } else {
                alert('Logout failed');
            }
        } catch (error) {
            console.error('Error during logout:', error);
        }
    });
