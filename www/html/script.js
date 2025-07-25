console.log("js loaded")
fetch('https://cigarland.reneau.me/api/test')
  .then(response => {
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    return response.json(); 
  })
  .then(data => {
    console.log('Received:', data);
  })
  .catch(error => {
    console.error('Fetch error:', error);
  });

const cigars = [
  { origin: "Cuba",
  brand: "Arturo Fuente",
  name: "Hemingway Short Story",
  wrapper: "Cameroon",
  profile: "Medium",
  tasty_tip: false,
  pressed: false,
  binder: "string",
  spicy: 7,
  rating: 5,
  kyle_rating: 4,
  john_rating: 6,
  length: 60,
  ring: 50,
  review: "Smooth draw, nutty on the front, spice at the end. Perfect burn. Smoked with JHarmon while talking Go interfaces.",
  john_review: "",
  kyle_review: "",
  image_ref: "pictures/ArturoFuenteHemming.jpg",
  authentic_human_review: ""},
    { origin: "russiauba",
  brand: "Arturo Fuente",
  name: "Hemingway Short Story",
  wrapper: "Cameroon",
  profile: "Medium",
  tasty_tip: false,
  pressed: false,
  binder: "string",
  spicy: 7,
  rating: 5,
  kyle_rating: 4,
  john_rating: 6,
  length: 60,
  ring: 50,
  review: "Smooth draw, nutty on the front, spice at the end. Perfect burn. Smoked with JHarmon while talking Go interfaces.",
  john_review: "",
  kyle_review: "",
  image_ref: "pictures/ArturoFuenteHemming.2,,jpg.jpg",
  authentic_human_review: ""},
    { origin: "Cuba cigar land",
  brand: "Arturo Fuente",
  name: "Hemingway Short Story",
  wrapper: "Cameroon",
  profile: "Medium",
  tasty_tip: false,
  pressed: false,
  binder: "string",
  spicy: 7,
  rating: 5,
  kyle_rating: 4,
  john_rating: 6,
  length: 60,
  ring: 50,
  review: "Smooth draw, nutty on the front, spice at the end. Perfect burn. Smoked with JHarmon while talking Go interfaces.",
  john_review: "",
  kyle_review: "",
  image_ref: "pictures/cigar-00000001.webp",
  authentic_human_review: ""}
];

const app     = document.getElementById("app");
const section = document.createElement("section");
section.id    = "reviews";
app.appendChild(section);

cigars.forEach(cigar => {

  /* ── card container ────────────────── */
  const details = document.createElement("details");
  details.className = "cigar-card";      // keeps your card styling

  /* ── summary (visible when collapsed) ─*/
  const summary = document.createElement("summary");
  summary.innerHTML = `<h2>${cigar.brand} — ${cigar.name}</h2>`;
  details.appendChild(summary);

  /* ── body (revealed when open) ─────── */
  const body = document.createElement("div");
  body.className = "cigar-body";
  body.innerHTML = `
      <img class="cigar-image" src="${cigar.image_ref}" alt="${cigar.name}">
      <p><strong>Origin:</strong> ${cigar.origin}</p>
      <p><strong>Wrapper:</strong> ${cigar.wrapper}</p>
      <p><strong>Profile:</strong> ${cigar.profile} |
         <strong>Pressed:</strong> ${cigar.pressed}</p>
      <p><strong>Tasty Tip:</strong> ${cigar.tasty_tip}</p>
      <p><strong>Spicy Level:</strong> ${cigar.spicy}/10</p>
      <p><strong>Rating:</strong> ${cigar.rating}/10</p>
      <p><strong>Kyle:</strong> ${cigar.kyle_rating}/10 |
         <strong>John:</strong> ${cigar.john_rating}/10</p>
      <p><strong>Length:</strong> ${cigar.length} mm |
         <strong>Ring Gauge:</strong> ${cigar.ring}</p>
      <p><strong>Kyle’s Review:</strong>
         ${cigar.kyle_review || "No review yet."}</p>
      <p><strong>John’s Review:</strong>
         ${cigar.john_review || "No review yet."}</p>
      <p><strong>General Review:</strong>
         ${cigar.review || "No review yet."}</p>
      <p><strong>Authentic Human Review:</strong>
         ${cigar.authentic_human_review || "No review yet."}</p>
  `;
  details.appendChild(body);

  section.appendChild(details);
});

app.appendChild(section);


const themeSwitch   = document.getElementById('themeSwitch');
const prefersDarkMq = window.matchMedia('(prefers-color-scheme: dark)');
function applyTheme(theme) {
  document.documentElement.setAttribute('data-theme', theme);
  localStorage.setItem('theme', theme);
  themeSwitch.checked = (theme === 'dark');
}
(function initTheme() {
  const stored = localStorage.getItem('theme');
  const initial = stored || (prefersDarkMq.matches ? 'dark' : 'light');
  applyTheme(initial);
})();
themeSwitch.addEventListener('change', () => {
  applyTheme(themeSwitch.checked ? 'dark' : 'light');
});
prefersDarkMq.addEventListener('change', e => {
  if (!localStorage.getItem('theme')) applyTheme(e.matches ? 'dark' : 'light');
});



document.querySelectorAll('.spoiler-toggle').forEach(btn => {
  btn.addEventListener('click', () => {
    const box = document.getElementById(btn.dataset.target);
    const isHidden = box.classList.toggle('hidden');
    if (!isHidden) {
      /* add animation class only on show */
      box.classList.add('fade-in');
      setTimeout(() => box.classList.remove('fade-in'), 250);
    }
    btn.textContent = isHidden ? 'Show notes' : 'Hide notes';
  });
});
