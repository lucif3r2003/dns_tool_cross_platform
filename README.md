# 📌 DNS Tool Cross Platform

A simple cross-platform tool written in **Go + Fyne** that allows you to quickly switch DNS between **Google DNS** and your **default system DNS**.  
Supports **Windows** and **Linux** (tested on Ubuntu).  

---

## ✨ Features
- Switch DNS to **Google DNS (8.8.8.8, 8.8.4.4)**
- Revert back to system default DNS
- Simple and lightweight **GUI** built with [Fyne](https://fyne.io)

---

## 🛠️ Installation

### 🔹 Ubuntu / Debian (Linux)

1. **Install system dependencies**
   ```bash
   sudo apt update
   sudo apt install -y libgl1-mesa-dev xorg-dev libxcursor-dev \
   libxrandr-dev libxinerama-dev libxi-dev
2. Clone and build the project
    ```bash
    git clone https://github.com/your-username/dns_tool_cross_platform.git
    cd dns_tool_cross_platform
    go mod tidy
    go build -o dns_tool
3.Run (requires root to change DNS)
    ```bash
      sudo ./dns_tool

### 🔹 Window

1. Install Go from golang.org
2. . Clone and build the project
    ```bash
    git clone https://github.com/your-username/dns_tool_cross_platform.git
    cd dns_tool_cross_platform
    go mod tidy
    go build -o dns_tool.exe
