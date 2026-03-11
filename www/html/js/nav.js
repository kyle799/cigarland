fetch('/me')
  .then(r => r.ok ? r.json() : null)
  .then(data => {
    window.userPerms = data || {};
    const authLi = document.getElementById('nav-auth');
    if (authLi) {
      if (data && data.email) {
        authLi.innerHTML = `<span>${data.email}</span>&nbsp;|&nbsp;<a href="/logout">Logout</a>`;
      } else {
        authLi.innerHTML = `<a href="/login">Login</a>`;
      }
    }
    const adminLi = document.getElementById('nav-admin');
    if (adminLi && data && data.can_admin) {
      adminLi.style.display = '';
    }
    document.dispatchEvent(new Event('perms-loaded'));
  });
