### OneTimeDownload

**OneTimeDownload** is a free and fast **Golang + HTMX** web application that allows users to safely download videos and audios from multiple platforms — including **YouTube, LinkedIn, Facebook, Twitter (X), Instagram**, and more.  
The system is built for performance and simplicity, powered by **Redis caching**, **Docker**, and a clean **Go backend**.

---

#### Features

- Download videos or audios from YouTube, LinkedIn, Facebook, Instagram, Twitter, etc.  
- Smart media extraction using `yt-dlp`  
- High-performance backend with **Golang**  
- Reactive UI built using **HTMX**  
- ast caching with **Redis**  
- Fully containerized using **Docker Compose**

---

#### Prerequisites

Ensure you have these installed before starting:

- [Git](https://git-scm.com/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- (Optional) [Go 1.25+](https://go.dev/dl/) for local testing

---

### Setup, Build & Run (All in One)

Copy and run the commands below 👇

```bash
# Clone the repository
git clone https://github.com/jimmymuthoni/onetimedownload.git
cd onetimedownload

# 2. Build and run using Docker Compose
docker compose up --build

# 3️Access the application
# http://localhost:8080
