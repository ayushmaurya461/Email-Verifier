# ✉️ Email Domain Verifier in Go

A small command-line utility written in Go to check if a given domain has valid **MX**, **SPF**, and **DMARC** records.

---

## ✅ Features

- Checks for presence of:
  - **MX records** (mail server configured)
  - **SPF record** (`v=spf1`)
  - **DMARC record** (`v=DMARC1`)
- Prints structured info for each domain
- Works via standard input (stdin), supports multiple domains

---

## 📦 Example Output

```bash
DOMAIN: example.com 
MX: true 
SPF: true 
SPF Record: v=spf1 include:_spf.google.com ~all 
DMARC: true 
DMARC Records: v=DMARC1; p=none; rua=mailto:dmarc@example.com
-------------------------------
