# Tel-Number Verify 📞

> Real-time global phone number verification, carrier lookup, and line type inspection powered by Go (Golang) and a modern monochrome web interface.

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![HTML5](https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white)
![CSS3](https://img.shields.io/badge/CSS3-1572B6?style=for-the-badge&logo=css3&logoColor=white)
![JavaScript](https://img.shields.io/badge/JavaScript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black)
![Vercel](https://img.shields.io/badge/Vercel-000000?style=for-the-badge&logo=vercel&logoColor=white)

---

## ✨ Features

- **⚡ Instant Validation**: Validates global phone numbers against official carrier databases.
- **📡 Carrier & Network Intelligence**: Retrieves telecom carrier name, line type (*Mobile*, *Landline*, *VoIP*, *Satellite*).
- **🌍 Geographical Specs**: Provides country name, ISO country code, dial prefix, and city/region location.
- **🔢 Multi-Format Parsing**: Displays international E.164 format, local national format, and raw query inputs.
- **🎨 Glassmorphic Monochrome UI**: Ultra-modern, dark-mode responsive interface built with clean CSS custom properties and Google Fonts (`Plus Jakarta Sans`).
- **⚡ Test Sample Presets**: One-click sample numbers (🇺🇸 USA, 🇬🇧 UK, 🇳🇬 Nigeria, 🇩🇪 Germany) for instant testing.
- **📜 Local Search History**: Remembers your recent lookups in `localStorage` for quick re-inspection.
- **📋 One-Click Copy & Toast Notifications**: Copy formatted numbers instantly with visual feedback.
- **🛠️ Raw JSON Inspector**: Collapsible developer drawer to inspect exact API payload responses.

---

## 🛠️ Tech Stack

### Backend
- **Go (Golang)**: High-performance HTTP server & API handler (`net/http`)
- **Numverify API**: External validation engine for telecommunication metadata

### Frontend
- **HTML5**: Semantic web architecture with accessibility metadata
- **Vanilla CSS3**: Custom dark glassmorphic design system with zero external framework dependencies
- **JavaScript (ES6+)**: Async/await fetch API, DOM manipulation, local storage management

### Deployment & Hosting
- **Vercel**: Configured for Vercel Serverless Functions (`vercel.json` rewrites)

---

## 📁 Project Structure

```text
num-verify/
├── api/
│   └── index.go            # Vercel serverless HTTP router & static asset handler
├── controller/
│   └── validate_controller.go # Request handler & CORS controller for /validate
├── service/
│   └── validate_service.go    # Numverify API HTTP client
├── model/                  # Go structs for request/response binding
├── response/               # Standardized JSON response envelope
├── assets/
│   └── css/
│       └── style.css       # Comprehensive glassmorphic monochrome design system
├── index.html              # Main web application user interface
├── vercel.json             # Vercel deployment rewrite rules
└── README.md               # Project documentation
```

---

## 🚀 Getting Started

### Prerequisites

- [Go 1.18+](https://golang.org/dl/)
- [Vercel CLI](https://vercel.com/docs/cli) (optional for local deployment testing)

### Environment Variables

Create a `.env` or `.env.local` file in the root directory and configure your Numverify access key:

```env
NUMVERIFY_KEY=your_apilayer_numverify_api_key
```

### Running Locally

1. **Using Vercel CLI (Recommended)**:
   ```bash
   vercel dev
   ```
   Open `http://localhost:3000` in your web browser.

2. **Using Go standard server**:
   ```bash
   go run main.go
   ```

---

## 🔌 API Endpoints

### 1. Web Frontend Interface
- **`GET /`**: Serves the main `index.html` single-page web app.
- **`GET /assets/*`**: Serves static CSS and asset files.

### 2. Validate Phone Number
- **`POST /validate`**

#### Request Body
```json
{
  "number": "14158586273"
}
```

#### Success Response (`200 OK`)
```json
{
  "success": true,
  "message": "Request success",
  "data": {
    "valid": true,
    "number": "14158586273",
    "local_format": "4158586273",
    "international_format": "+14158586273",
    "country_prefix": "+1",
    "country_code": "US",
    "country_name": "United States of America",
    "location": "Novato",
    "carrier": "AT&T Mobility LLC",
    "line_type": "mobile"
  }
}
```

---

## 📄 License

This project is open-source under the MIT License.
