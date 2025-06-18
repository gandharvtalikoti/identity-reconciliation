# 🔍 Identity Reconciliation API

> FluxKart’s time-traveling customer tracker, built in GoLang ⚙️🚀

---

## 🎯 Problem Statement

FluxKart shoppers use different emails/phones across orders. Bitespeed’s backend must reconcile contacts with common phone or email, and return a unified identity response.

---

## 📌 Tech Stack

- ✅ Golang
- ✅ Gin Web Framework
- ✅ PostgreSQL
- ✅ GORM ORM
- ✅ Hosted on Render.com

---

## 🧪 API: `/identify`

**POST** `/identify`

### 📥 Request Body

```json
{
  "email": "test@flux.com",
  "phoneNumber": "1234567890"
}
