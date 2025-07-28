const form = document.getElementById('cigarForm');
form.addEventListener('submit', async e => {
e.preventDefault();
const data = new FormData(form);
let payload = {};
data.forEach((val, key) => {
    if (['pressed','tasty_tip'].includes(key)) {
    payload[key] = form.elements[key].checked;
    } else if (['spicy','rating','kyle_rating','john_rating','length','ring'].includes(key)) {
    payload[key] = val ? Number(val) : null;
    } else {
    payload[key] = val;
    }
});
console.log(data)
console.log(payload)
let cigar={}
cigar.cigar_list = [payload]
console.log(cigar)


try {
    const res = await fetch('https://cigarland.reneau.me/api/test', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(cigar)
    });
    if (!res.ok) throw new Error(`Server responded ${res.status}`);
    const created = await res.json();
    console.log('Cigar created:', created);

    form.reset();
} catch (err) {
    console.error('Submission failed:', err);
    alert('Failed to submit cigar: ' + err.message);
}
});