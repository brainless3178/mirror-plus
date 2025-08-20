Mirror+ – Advanced Reflection Detection for Bug Bounty Hunters & Pentesters

Mirror+ is a powerful, lightweight tool built in Go that scans URLs for reflected input values across HTTP responses, headers, and more — helping you quickly identify potential injection points such as:

🧪 Reflected XSS

🔁 Open Redirects

🛡️ Authentication Bypass / Token Leak Surfaces

Inspired by @tomnomnom
's original mirror tool, this enhanced version adds support for:

🚀 Features

✅ Detect query parameter reflection in body & headers

✅ Detect URL path segment reflection

✅ Check POST data reflection

✅ Catch URL-encoded value reflections

✅ Skip short values to reduce false positives

✅ Custom wordlists for payload injection testing

✅ Colored CLI output + JSON export (coming soon)

⚙️ Fast & efficient — written in pure Go

🔧 Use Cases

Recon for bug bounties

Scanning large URL lists for XSS points

Testing reflected inputs in login, search, and redirect flows

Mapping vulnerable parameters across APIs and forms

📦 Installation
go install github.com/<your-username>/mirror-plus@latest

🕹️ Usage
Scan a Single URL:
mirror-plus "https://example.com/page?search=test"

Bulk Scan:
cat urls.txt | mirror-plus

POST Data Testing:
mirror-plus -method POST -data "username=test" https://example.com/login
