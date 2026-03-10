fetch('/me')
  .then(r => r.ok ? r.json() : null)
  .then(data => {
    const authLi = document.getElementById('nav-auth');
    if (!authLi) return;
    if (data && data.email) {
      authLi.innerHTML = `<span>${data.email}</span>&nbsp;|&nbsp;<a href="/logout">Logout</a>`;
    } else {
      authLi.innerHTML = `<a href="/login">Login</a>`;
    }
  });
