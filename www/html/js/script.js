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
  let cigars = [
      {
        origin: 'Cuba',
        brand: 'Arturo Fuente',
        name: 'Hemingway Short Story',
        wrapper: 'Cameroon',
        profile: 'Medium',
        tasty_tip: false,
        pressed: false,
        binder: 'string',
        spicy: 7,
        rating: 5,
        kyle_rating: 4,
        john_rating: 6,
        length: 60,
        ring: 50,
        review:' Smooth draw, nutty on the front, spice at the end. Perfect burn. Smoked with JHarmon while talking Go interfaces.',
        john_review: '',
        kyle_review: '',
        image_ref: 'pictures/ArturoFuenteHemming.jpg',
        authentic_human_review: ''
      }
    ];
  window.addEventListener('DOMContentLoaded', loadCigars);
  async function loadCigars() {
    try {
      console.log("loading cigars...")
      renderCigars(cigars)
      const res = await fetch('https://cigarland.reneau.me/api/cigars');
      if (!res.ok) throw new Error(`HTTP ${res.status}`);
      const data = await res.json();
      cigars = data;
      console.log(data[0])
      renderCigars(cigars);
    } catch (err) {
      console.error('Fetch failed:', err);
      renderError();
    }
  }
function renderCigars(cigars){
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
}
const app     = document.getElementById("app");
const section = document.createElement("section");
section.id    = "reviews";
app.appendChild(section);

console.log("end")