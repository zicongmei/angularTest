if (!("token" in window.sessionStorage) || window.sessionStorage.token.length < 10) {
    window.location.href = '/login.html';
}