# Skip403 - Bypass 403/401 Restrictions with Smart Header and Path Manipulation

Skip403 is a powerful and lightweight security testing tool written in Go, designed to help penetration testers, bug bounty hunters, and security researchers bypass `403 Forbidden` and `401 Unauthorized`. Instead of blindly guessing payloads, Skip403 uses smart and assertive manipulation of HTTP headers and URL paths to reveal improperly protected resources.

---

## 🔍 Overview

While many tools rely on brute-force techniques or wordlists to find access control issues, **Skip403** focuses on real-world behavior exploitation. It simulates requests that include trusted headers (such as `X-Forwarded-For`, `X-Original-URL`, etc.) and path obfuscations that are commonly misinterpreted by backend systems or reverse proxies.

This approach increases the chance of revealing misconfigurations without overwhelming the server with requests, making it faster, cleaner, and more precise than traditional fuzzing tools.

---

## ⚡ Features

- ✅ Tests multiple known bypass headers
- ✅ Manipulates URL paths with encoding, traversal, and obfuscation techniques
- ✅ Lightweight, fast, and written in pure Go
- ✅ Designed for automation and mass scanning

## 🛠 How to Use
```
$ skip403 -h
     _    _       _  _    ___ _____ 
 ___| | _(_)_ __ | || |  / _ \___ / 
/ __| |/ / | '_ \| || |_| | | ||_ \ 
\__ \   <| | |_) |__   _| |_| |__) |
|___/_|\_\_| .__/   |_|  \___/____/ 
           |_|                      

Created by: Geison Nunes | Version: 1.0

Usage:
  -u,  --url         Specify a single target URL
  -l,  --list        Specify a file with a list of target URLs
  -hf, --hide-fails  Hide failed bypass attempts
  -h,  --help        Show this help message
```
## 🚀 Installation
```go install github.com/geisonn/skip403@latest```
