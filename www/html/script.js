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