#!/bin/bash
# Usage: ./seed_cigars.sh <session_cookie>
# Get your session cookie from browser DevTools → Application → Cookies → cigarland_session

SESSION="${1:?Usage: $0 <session_cookie>}"
API="${2:-https://cigarland.reneau.me}"

curl -s -X POST "$API/api/test" \
  -H "Content-Type: application/json" \
  -H "Cookie: cigarland_session=$SESSION" \
  -d '{
  "cigar_list": [
    {
      "brand": "Arturo Fuente",
      "name": "Opus X Robusto",
      "wrapper": "Dominican Rosado",
      "profile": "Full",
      "binder": "Dominican",
      "pressed": false,
      "tasty_tip": false,
      "spicy": 6,
      "rating": 9,
      "length": 140,
      "ring": 50,
      "review": "Legendary Dominican puro. Rich cedar, leather, and dark fruit with a creamy, lingering finish. The Rosado wrapper adds a unique sweetness that sets it apart.",
      "kyle_rating": 0,
      "kyle_review": "",
      "john_rating": 0,
      "john_review": "",
      "image_ref": "",
      "authentic_human_review": ""
    },
    {
      "brand": "Padron",
      "name": "1964 Anniversary Exclusivo Maduro",
      "wrapper": "Nicaraguan Maduro",
      "profile": "Full",
      "binder": "Nicaraguan",
      "pressed": true,
      "tasty_tip": false,
      "spicy": 5,
      "rating": 9,
      "length": 143,
      "ring": 50,
      "review": "Box-pressed perfection. Deep cocoa, espresso, and earth with a long, satisfying finish. One of the benchmarks of the industry — consistently excellent.",
      "kyle_rating": 0,
      "kyle_review": "",
      "john_rating": 0,
      "john_review": "",
      "image_ref": "",
      "authentic_human_review": ""
    },
    {
      "brand": "My Father",
      "name": "Le Bijou 1922 Torpedo",
      "wrapper": "San Andres Mexican Maduro",
      "profile": "Full",
      "binder": "Nicaraguan",
      "pressed": false,
      "tasty_tip": false,
      "spicy": 7,
      "rating": 9,
      "length": 152,
      "ring": 52,
      "review": "Bold and complex from start to finish. Dark chocolate, black pepper, and leather dominate with a rich, oily draw. The torpedo format focuses the strength beautifully.",
      "kyle_rating": 0,
      "kyle_review": "",
      "john_rating": 0,
      "john_review": "",
      "image_ref": "",
      "authentic_human_review": ""
    },
    {
      "brand": "Cohiba",
      "name": "Siglo VI",
      "wrapper": "Cuban",
      "profile": "Medium-Full",
      "binder": "Cuban",
      "pressed": false,
      "tasty_tip": false,
      "spicy": 4,
      "rating": 9,
      "length": 150,
      "ring": 52,
      "review": "The flagship of Cuban cigars. Creaminess and complexity in perfect balance — cedar, honey, and subtle spice. Burns slow and even with a tight white ash.",
      "kyle_rating": 0,
      "kyle_review": "",
      "john_rating": 0,
      "john_review": "",
      "image_ref": "",
      "authentic_human_review": ""
    },
    {
      "brand": "Rocky Patel",
      "name": "Vintage 1990 Robusto",
      "wrapper": "Connecticut Broadleaf Maduro",
      "profile": "Medium-Full",
      "binder": "Honduran",
      "pressed": false,
      "tasty_tip": false,
      "spicy": 4,
      "rating": 8,
      "length": 127,
      "ring": 50,
      "review": "Approachable and smooth for a maduro. Notes of sweet cedar, cream, and a hint of cocoa. Great construction and an easy draw make this a reliable everyday smoke.",
      "kyle_rating": 0,
      "kyle_review": "",
      "john_rating": 0,
      "john_review": "",
      "image_ref": "",
      "authentic_human_review": ""
    }
  ]
}'

echo ""
echo "Done."
