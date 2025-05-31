# JSON Beautifier & Minifier Web App

A simple web application built with **Go (Gin framework)** and **HTML/CSS** to beautify and minify JSON data. It includes a user-friendly UI and supports deployment to platforms like **Render**, **EC2**, or **Cloud Run**.

---

## 🚀 Features

* JSON Beautifier: Paste minified JSON and get pretty-printed output.
* JSON Minifier: Paste beautified JSON and get compact output.
* Clean responsive UI with live formatting.
* Navigation bar to switch between tools.
* Supports deployment on multiple cloud platforms.

---

## 🧱 Project Structure

```
json-tool/
├── main.go              # Go backend (Gin web server)
├── go.mod/go.sum       # Go module files
├── static/
│   ├── beautify.html   # Frontend for JSON beautifier
│   └── minify.html     # Frontend for JSON minifier
├── Dockerfile          # For containerized deployment
└── render.yaml         # (Optional) Render deployment config
```

---

## 🛠️ Getting Started

### 1. Run Locally

```bash
go run main.go
```

Visit:

* `http://localhost:8080/beautify`
* `http://localhost:8080/minify`

---

## 🐳 Docker Deployment

### Build & Run Locally

```bash
docker build -t json-tool .
docker run -p 8080:8080 json-tool
```

---

## 📄 License

MIT License

---

## 🙋‍♂️ Author

Made by **Puvaan Raaj** — for quick and beautiful JSON formatting.

