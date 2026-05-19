# ASCII Web Generator (Go)

A simple web application written in **Go** that converts user input text into ASCII art using selectable banner styles.

## Features

- Convert text into ASCII art
- Supports multiple banner styles:
  - `standard`
  - `shadow`
  - `thinkertoy`
- Handles multi-line input (`\n`)
- Clean web interface using HTML templates

---

## 📁 Project Structure

```
.
├── main.go
├── template/
│   └── index.html
├── asciiArt/
│   ├── generator.go
│   ├── reader/
│   │   └── filereading.go
│   └── banners/
│       ├── standard.txt
│       ├── shadow.txt
│       └── thinkertoy.txt
```

---

## Installation & Setup

### 1. Clone the repository

```bash
git clone https://github.com/your-username/ascii-web.git
cd ascii-web
```

### 2. Run the server

```bash
go run main.go
```

### 3. Open in browser

```
http://localhost:8001
```

---

## Usage

1. Enter text in the input box
2. Choose a banner style
3. Click **Generate ASCII**
4. View the output below the form

---

## How It Works

### Backend (`main.go`)

- Handles HTTP routes:
  - `/` → renders the homepage
  - `/ascii` → processes form submission
- Uses Go templates to display ASCII output

### ASCII Generator (`generator.go`)

- Converts each character into ASCII art
- Reads banner data and maps characters line by line
- Uses `strings.Builder` for efficient string concatenation

### File Reader (`filereading.go`)

- Reads banner files line by line
- Stores ASCII patterns in a slice
- Supports different banner styles dynamically

---

## Error Handling

- 404 for invalid routes
- 405 for non-POST requests to `/ascii`
- 400 for empty input
- Logs template execution errors

---

## Notes

- Only printable ASCII characters (32–126) are supported
- Empty lines in input are skipped
- Banner files must follow the correct ASCII format (8 lines per character)

---

## 💡 Future Improvements

- Add more banner styles
- Improve UI styling (CSS)
- Support file uploads
- Add download option for ASCII output
- Better error display on frontend

---

## Author

Staging, func and html

---

## License

This project is open-source and available under the MIT License.