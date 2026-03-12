const content = document.getElementById('admin-content');

async function loadUsers() {
  const res = await fetch('/api/admin/users');
  if (!res.ok) { content.textContent = 'Access denied.'; return; }
  const users = await res.json();
  renderUsers(users);
}

function renderUsers(users) {
  const table = document.createElement('table');
  table.innerHTML = `
    <thead>
      <tr>
        <th>Email</th>
        <th>Can Add</th>
        <th>Can Edit</th>
        <th>Can Delete</th>
        <th>Can Admin</th>
        <th></th>
      </tr>
    </thead>
  `;
  const tbody = document.createElement('tbody');
  users.forEach(u => {
    const tr = document.createElement('tr');
    tr.innerHTML = `
      <td>${u.email}</td>
      <td><input type="checkbox" ${u.can_add ? 'checked' : ''} data-field="can_add"></td>
      <td><input type="checkbox" ${u.can_edit ? 'checked' : ''} data-field="can_edit"></td>
      <td><input type="checkbox" ${u.can_delete ? 'checked' : ''} data-field="can_delete"></td>
      <td><input type="checkbox" ${u.can_admin ? 'checked' : ''} data-field="can_admin"></td>
      <td><button>Save</button></td>
    `;
    tr.querySelector('button').onclick = () => saveUser(u.email, tr);
    tbody.appendChild(tr);
  });
  table.appendChild(tbody);

  const addForm = document.createElement('div');
  addForm.innerHTML = `
    <h3>Add user</h3>
    <input id="new-email" type="email" placeholder="email">
    <label><input type="checkbox" id="new-can-add"> Add</label>
    <label><input type="checkbox" id="new-can-edit"> Edit</label>
    <label><input type="checkbox" id="new-can-delete"> Delete</label>
    <label><input type="checkbox" id="new-can-admin"> Admin</label>
    <button id="add-user-btn">Add</button>
  `;
  addForm.querySelector('#add-user-btn').onclick = addUser;

  content.innerHTML = '';
  content.appendChild(table);
  content.appendChild(addForm);
}

async function saveUser(email, tr) {
  const checks = tr.querySelectorAll('input[type=checkbox]');
  const payload = { email };
  checks.forEach(c => { payload[c.dataset.field] = c.checked; });
  const res = await fetch('/api/admin/users', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload)
  });
  if (!res.ok) alert('Save failed');
}

async function addUser() {
  const email = document.getElementById('new-email').value.trim();
  if (!email) return;
  const payload = {
    email,
    can_add:    document.getElementById('new-can-add').checked,
    can_edit:   document.getElementById('new-can-edit').checked,
    can_delete: document.getElementById('new-can-delete').checked,
    can_admin:  document.getElementById('new-can-admin').checked,
  };
  const res = await fetch('/api/admin/users', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload)
  });
  if (res.ok) loadUsers();
  else alert('Failed to add user');
}

loadUsers();
