# 🧠 Identity Reconciliation – GoLang Backend

This project solves the **Identity Reconciliation** problem for [Bitespeed](https://www.bitespeed.co/). The goal is to identify and link users who may use different contact information (email/phone) across multiple orders on FluxKart.com.

## 🚀 Hosted API

🔗 **Live URL**: [https://identity-reconciliation-r68x.onrender.com](https://identity-reconciliation-r68x.onrender.com)

## 📬 API Endpoint

### `POST /identify`

Identifies and links user contacts by email and/or phone number.

#### 🔧 Request Body (JSON)
```json
{
  "email": "user@example.com",
  "phoneNumber": "1234567890"
}
