const form         = document.getElementById('cigarForm');
const status       = document.getElementById('form-status');
const wrapperSel   = document.getElementById('f-wrapper');
const wrapperOther = document.getElementById('f-wrapper-other');
const submitBtn    = form.querySelector('.submit-btn');
const pageTitle    = document.querySelector('.page-title');

const BASE_WRAPPERS = [
  'Double Claro / Candela',
  'Claro',
  'Colorado Claro',
  'Colorado',
  'Colorado Maduro',
  'Maduro',
  'Oscuro',
];

// Check if we're in edit mode via query params
const params    = new URLSearchParams(window.location.search);
const editBrand = params.get('brand');
const editName  = params.get('name');
const isEdit    = !!(editBrand && editName);

if (isEdit) {
  pageTitle.textContent = 'Edit Cigar';
  submitBtn.textContent = 'Save Changes';
}

// Populate wrapper dropdown, then pre-fill form if editing
fetch('/api/cigars/wrappers')
  .then(r => r.ok ? r.json() : [])
  .then(fromServer => {
    const all = [...new Set([...BASE_WRAPPERS, ...fromServer])].sort();
    wrapperSel.innerHTML = '';
    const blank = document.createElement('option');
    blank.value = '';
    blank.textContent = '— Select wrapper —';
    wrapperSel.appendChild(blank);
    all.forEach(w => {
      const opt = document.createElement('option');
      opt.value = w;
      opt.textContent = w;
      wrapperSel.appendChild(opt);
    });
    const other = document.createElement('option');
    other.value = '__other__';
    other.textContent = 'Other...';
    wrapperSel.appendChild(other);

    if (isEdit) prefillForm();
  })
  .catch(() => {
    wrapperSel.innerHTML = '<option value="">Unable to load</option>';
    if (isEdit) prefillForm();
  });

async function prefillForm() {
  try {
    const res = await fetch('/api/cigars');
    if (!res.ok) throw new Error();
    const cigars = await res.json();
    const cigar  = cigars.find(c => c.brand === editBrand && c.name === editName);
    if (!cigar) return;

    const textFields = ['brand', 'name', 'origin', 'image_ref', 'wrapper', 'binder', 'profile',
                        'kyle_review', 'john_review', 'review', 'authentic_human_review'];
    const numFields  = ['spicy', 'rating', 'kyle_rating', 'john_rating', 'length', 'ring'];

    textFields.forEach(k => {
      const el = form.elements[k];
      if (el) el.value = cigar[k] || '';
    });
    numFields.forEach(k => {
      const el = form.elements[k];
      if (el && cigar[k] != null) el.value = cigar[k];
    });
    ['pressed', 'tasty_tip'].forEach(k => {
      const el = form.elements[k];
      if (el) el.checked = !!cigar[k];
    });

    // Set wrapper select — fall back to Other if not in list
    const wrapperVal = cigar.wrapper || '';
    const match = [...wrapperSel.options].find(o => o.value === wrapperVal);
    if (match) {
      wrapperSel.value = wrapperVal;
    } else if (wrapperVal) {
      wrapperSel.value = '__other__';
      wrapperOther.style.display = '';
      wrapperOther.value = wrapperVal;
    }
  } catch {
    status.textContent = 'Failed to load cigar data.';
    status.className   = 'error';
  }
}

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
    if (key === 'wrapper_other') return;
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
  ['pressed', 'tasty_tip'].forEach(k => {
    if (!(k in payload)) payload[k] = false;
  });

  status.textContent = '';
  status.className   = '';

  try {
    const res = await fetch(isEdit ? '/api/cigars' : '/api/test', {
      method:  isEdit ? 'PUT' : 'POST',
      headers: { 'Content-Type': 'application/json' },
      body:    JSON.stringify(isEdit ? payload : { cigar_list: [payload] }),
    });
    if (!res.ok) throw new Error(`Server responded ${res.status}`);
    status.textContent = isEdit ? 'Changes saved.' : 'Cigar added successfully.';
    status.className   = 'success';
    if (!isEdit) form.reset();
  } catch (err) {
    status.textContent = 'Failed to submit: ' + err.message;
    status.className   = 'error';
  }
});
