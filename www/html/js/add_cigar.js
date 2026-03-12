const form   = document.getElementById('cigarForm');
const status = document.getElementById('form-status');

form.addEventListener('submit', async e => {
  e.preventDefault();

  const data    = new FormData(form);
  const payload = {};
  data.forEach((val, key) => {
    if (['pressed', 'tasty_tip'].includes(key)) {
      payload[key] = form.elements[key].checked;
    } else if (['spicy', 'rating', 'kyle_rating', 'john_rating', 'length', 'ring'].includes(key)) {
      payload[key] = val ? Number(val) : null;
    } else {
      payload[key] = val;
    }
  });

  // checkboxes default to unchecked (not in FormData) — ensure they're set
  ['pressed', 'tasty_tip'].forEach(k => {
    if (!(k in payload)) payload[k] = false;
  });

  status.textContent = '';
  status.className   = '';

  try {
    const res = await fetch('/api/test', {
      method:  'POST',
      headers: { 'Content-Type': 'application/json' },
      body:    JSON.stringify({ cigar_list: [payload] }),
    });
    if (!res.ok) throw new Error(`Server responded ${res.status}`);
    status.textContent = 'Cigar added successfully.';
    status.className   = 'success';
    form.reset();
  } catch (err) {
    status.textContent = 'Failed to submit: ' + err.message;
    status.className   = 'error';
  }
});
