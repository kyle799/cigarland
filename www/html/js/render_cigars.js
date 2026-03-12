console.log("js loaded")

const app     = document.getElementById("app");
const section = document.createElement("section");
section.id    = "reviews";
app.appendChild(section);

window.addEventListener('DOMContentLoaded', loadCigars);

async function loadCigars() {
  try {
    console.log("loading cigars...")
    const [cigarsRes, meRes] = await Promise.all([
      fetch('/api/cigars'),
      fetch('/me')
    ]);
    if (!cigarsRes.ok) throw new Error(`HTTP ${cigarsRes.status}`);
    const cigars = await cigarsRes.json();
    const me = meRes.ok ? await meRes.json() : {};
    console.log(cigars[0]);
    renderCigars(cigars, me.can_delete, me.can_edit);
  } catch (err) {
    console.error('Fetch failed:', err);
  }
}

function renderCigars(cigars, canDelete, canEdit) {
  section.innerHTML = '';
  cigars.forEach(cigar => {
    const details = document.createElement("details");
    details.className = "cigar-card";
    const summary = document.createElement("summary");
    summary.innerHTML = `<h2>${cigar.brand} — ${cigar.name}</h2>`;
    details.appendChild(summary);
    const body = document.createElement("div");
    body.className = "cigar-body";
    body.innerHTML = `
        <img class="cigar-image" src="${cigar.image_ref}" alt="${cigar.name}">
        <p><strong>Origin:</strong> ${cigar.origin}</p>
        <p><strong>Wrapper:</strong> ${cigar.wrapper}</p>
        <p><strong>Binder:</strong> ${cigar.binder}</p>
        <p><strong>Profile:</strong> ${cigar.profile} |
           <strong>Pressed:</strong> ${cigar.pressed}</p>
        <p><strong>Tasty Tip:</strong> ${cigar.tasty_tip}</p>
        <p><strong>Spicy Level:</strong> ${cigar.spicy}/10</p>
        <p><strong>Rating:</strong> ${cigar.rating}/10</p>
        <p><strong>Kyle:</strong> ${cigar.kyle_rating}/10 |
           <strong>John:</strong> ${cigar.john_rating}/10</p>
        <p><strong>Length:</strong> ${cigar.length} mm |
           <strong>Ring Gauge:</strong> ${cigar.ring}</p>
        <p><strong>Kyle's Review:</strong>
           ${cigar.kyle_review || "No review yet."}</p>
        <p><strong>John's Review:</strong>
           ${cigar.john_review || "No review yet."}</p>
        <p><strong>General Review:</strong>
           ${cigar.review || "No review yet."}</p>
        <p><strong>Authentic Human Review:</strong>
           ${cigar.authentic_human_review || "No review yet."}</p>
    `;
    if (canEdit) {
      const btn = document.createElement("button");
      btn.textContent = "Edit";
      btn.className = "edit-btn";
      btn.onclick = () => {
        window.location.href = `/add_cigar?brand=${encodeURIComponent(cigar.brand)}&name=${encodeURIComponent(cigar.name)}`;
      };
      body.appendChild(btn);
    }
    if (canDelete) {
      const btn = document.createElement("button");
      btn.textContent = "Delete";
      btn.className = "delete-btn";
      btn.onclick = () => deleteCigar(cigar.brand, cigar.name, details);
      body.appendChild(btn);
    }
    details.appendChild(body);
    section.appendChild(details);
  });
}

async function deleteCigar(brand, name, el) {
  if (!confirm(`Delete "${brand} — ${name}"?`)) return;
  const res = await fetch(`/api/cigars?brand=${encodeURIComponent(brand)}&name=${encodeURIComponent(name)}`, { method: 'DELETE' });
  if (res.ok) {
    el.remove();
  } else {
    alert('Delete failed');
  }
}

console.log("end")
