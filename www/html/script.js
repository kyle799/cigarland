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

const app = document.getElementById("app");

const section = document.createElement("section");
section.id = "reviews";

cigars.forEach(cigar =>{
  const card = document.createElement("div");

  card.className = "cigar-card";
  const img = document.createElement("img");
  img.src = cigar.image;

card.appendChild(img);
  card.innerHTML = `
    <h2>${cigar.brand} -- ${cigar.name}</h2>
    <img class="cigar-image" src="${cigar.image_ref}" alt="${cigar.name}">
    <p><strong>Origin:</strong> ${cigar.origin}</p>
    <p><strong>Wrapper:</strong> ${cigar.wrapper}</p>
    <p><strong>Profile:</strong> ${cigar.profile}</p>
    <p><strong>Spicy Level:</strong> ${cigar.spicy}/10</p>
    <p><strong>Rating:</strong> ${cigar.rating}/10</p>
    <p><strong>Kyle:</strong> ${cigar.kyle_rating}/10 | <strong>John:</strong> ${cigar.john_rating}/10</p>
    <p><strong>Length:</strong> ${cigar.length} min | <strong>Ring Gauge:</strong> ${cigar.ring}</p>
    <p><strong>Review:</strong> ${cigar.review || "No review yet."}</p>
  `;

  app.appendChild(card);
  // const name = document.createElement("h2");
  // name.textContent = cigar.name;

  // const rating = document.createElement("p");
  // rating.innerHTML = `<strong>John's Rating:</strong> ${cigar.rating}`;  

  // const johns_rating = document.createElement("p");
  // johns_rating.innerHTML = `<strong>John's Rating:</strong> ${cigar.john_review}`;  
  
  // const kyles_rating = document.createElement("p");
  // kyles_rating.innerHTML = `<strong>Kyle's rating:</strong> ${cigar.kyle_rating}`;

  // const notes = document.createElement("p");
  // notes.innerHTML = `<strong>Tasting Notes:</strong> ${cigar.notes}`;

  // card.appendChild(name)
  // card.append(johns_rating)
  // card.append(kyles_rating)
  // card.append(rating)
  // card.append(notes)
  // section.appendChild(card)
});

app.appendChild(section);

/* ============ DARK‑MODE HANDLER ============ */
const themeSwitch   = document.getElementById('themeSwitch');
const prefersDarkMq = window.matchMedia('(prefers-color-scheme: dark)');

/* Apply a theme and reflect it on the switch */
function applyTheme(theme) {
  document.documentElement.setAttribute('data-theme', theme);
  localStorage.setItem('theme', theme);
  themeSwitch.checked = (theme === 'dark');
}

/* Initialize on load */
(function initTheme() {
  const stored = localStorage.getItem('theme');
  const initial = stored || (prefersDarkMq.matches ? 'dark' : 'light');
  applyTheme(initial);
})();

/* React to user toggle */
themeSwitch.addEventListener('change', () => {
  applyTheme(themeSwitch.checked ? 'dark' : 'light');
});

/* React to OS theme changes live */
prefersDarkMq.addEventListener('change', e => {
  if (!localStorage.getItem('theme')) applyTheme(e.matches ? 'dark' : 'light');
});
