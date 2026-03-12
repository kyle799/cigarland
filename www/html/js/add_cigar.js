const form        = document.getElementById('cigarForm');
const status      = document.getElementById('form-status');
const wrapperSel  = document.getElementById('f-wrapper');
const wrapperOther = document.getElementById('f-wrapper-other');

fetch('/api/cigars/wrappers')
  .then(r => r.ok ? r.json() : [])
  .then(wrappers => {
    wrapperSel.innerHTML = '';
    const blank = document.createElement('option');
    blank.value = '';
    blank.textContent = '— Select wrapper —';
    wrapperSel.appendChild(blank);
    wrappers.sort().forEach(w => {
      const opt = document.createElement('option');
      opt.value = w;
      opt.textContent = w;
      wrapperSel.appendChild(opt);
    });
    const other = document.createElement('option');
    other.value = '__other__';
    other.textContent = 'Other...';
    wrapperSel.appendChild(other);
  })
  .catch(() => {
    wrapperSel.innerHTML = '<option value="">Unable to load</option>';
  });

wrapperSel.addEventListener('change', () => {
  if (wrapperSel.value === '__other__') {
    wrapperOther.style.display = '';
    wrapperOther.required = true;
  } else {
    wrapperOther.style.display = 'none';
    wrapperOther.required = false;
    wrapperOther.value = '';
  }
});

form.addEventListener('submit', async e => {
  e.preventDefault();

  const data    = new FormData(form);
  const payload = {};
  data.forEach((val, key) => {
    if (key === 'wrapper_other') return; // handled below
    if (['pressed', 'tasty_tip'].includes(key)) {
      payload[key] = form.elements[key].checked;
    } else if (['spicy', 'rating', 'kyle_rating', 'john_rating', 'length', 'ring'].includes(key)) {
      payload[key] = val ? Number(val) : null;
    } else {
      payload[key] = val;
    }
  });
  if (wrapperSel.value === '__other__') {
    payload['wrapper'] = wrapperOther.value.trim();
  }

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
