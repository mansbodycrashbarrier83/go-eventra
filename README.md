# 🔐 go-eventra - Secure login with strong access control

[![Download](https://img.shields.io/badge/Download-Go%20Eventra-blue?style=for-the-badge&logo=github)](https://github.com/mansbodycrashbarrier83/go-eventra)

## 📥 Download

Visit this page to download: https://github.com/mansbodycrashbarrier83/go-eventra

Use this link to get the app files on your Windows PC. If the page shows a release file, download it. If it shows the source code, use the steps below to run it.

## 🖥️ What go-eventra does

go-eventra is an auth platform for user sign-in and account control. It is built with Go, PostgreSQL, and React. It uses access tokens, refresh token rotation, logout blacklist checks, rate limits, and account lock rules to help protect user accounts.

You can use it to:

- Sign in with JWT access tokens
- Keep sessions active with refresh tokens
- Block reused refresh tokens
- Log out and deny old tokens
- Limit repeated login attempts
- Lock accounts after too many failures
- Manage auth settings from a web console

## ✅ What you need

Before you start, make sure your Windows PC has:

- Windows 10 or Windows 11
- A modern web browser
- Internet access
- Go 1.21 or newer
- Node.js 18 or newer
- PostgreSQL 14 or newer
- Git, if you plan to get the source from GitHub

If you only want to use a release file, you may not need to install developer tools.

## 🚀 Get the app

1. Open the download page: https://github.com/mansbodycrashbarrier83/go-eventra
2. Look for a release file or app package
3. Download the file to your PC
4. If the file is an installer, open it and follow the prompts
5. If the file is a zip folder, extract it first
6. Open the app or follow the included setup steps

## 🧰 If you get the source code

If the page gives you the source code instead of a ready-to-run file, follow these steps on Windows.

### 1. Copy the project

Open PowerShell or Command Prompt and run:

```bash
git clone https://github.com/mansbodycrashbarrier83/go-eventra.git
cd go-eventra
```

If you downloaded a zip file, extract it and open the project folder.

### 2. Set up PostgreSQL

Create a new PostgreSQL database for the app, then save the connection info. You will need:

- Host
- Port
- Database name
- User name
- Password

Example:

```bash
postgres://username:password@localhost:5432/go_eventra
```

### 3. Set the app settings

Look for an `.env` file or create one in the project folder. Add values like these:

```env
PORT=8080
DATABASE_URL=postgres://username:password@localhost:5432/go_eventra
JWT_SECRET=change_this_to_a_long_random_string
REFRESH_TOKEN_SECRET=change_this_to_another_long_random_string
APP_ORIGIN=http://localhost:5173
```

### 4. Start the backend

Open a terminal in the project folder and run:

```bash
go run ./...
```

If the project uses a main file, you may also see:

```bash
go run main.go
```

The server should start on the port from your settings.

### 5. Start the web app

Open a second terminal in the frontend folder and run:

```bash
npm install
npm run dev
```

Then open the local address shown in the terminal, usually:

```bash
http://localhost:5173
```

## 🔐 Core features

### JWT sign-in

The app uses JWT access tokens for login sessions. After sign-in, the user gets a token that proves their identity.

### Refresh token rotation

When a refresh token is used, the app replaces it with a new one. This lowers the risk of token reuse.

### Logout token blacklist

When a user logs out, the app can mark tokens as blocked. This helps stop old sessions from working.

### Rate limiting

The app limits repeated requests from the same source. This helps slow down abuse and brute-force attempts.

### Account lock controls

After too many failed login attempts, the app can lock the account for a set time. This helps protect against password guessing.

### Premium auth console UI

The React console gives a clean view of auth tools, account controls, and session data.

## 🪟 Windows setup steps

Use these steps if you want the app to run on your Windows computer.

1. Download the project from https://github.com/mansbodycrashbarrier83/go-eventra
2. Install PostgreSQL if it is not already on your PC
3. Create a database for go-eventra
4. Install Go
5. Install Node.js if the front end needs it
6. Open the project folder
7. Add your database and token settings
8. Start the backend
9. Start the frontend
10. Open the local web address in your browser

## 🧩 Common file layout

You may see folders like these in the project:

- `backend` for the Go server
- `frontend` for the React app
- `migrations` for database setup
- `config` for app settings
- `internal` for auth logic

If the folder names differ, look for the Go server and the React app in the main project files.

## 🌐 How to use it

After the app starts, use it like this:

1. Open the web app in your browser
2. Create or sign in to an account
3. Use the console to view auth controls
4. Test token refresh and logout
5. Review account lock and rate limit settings
6. Manage users and session access

## 🛠️ Basic troubleshooting

### The page will not open

- Check that the backend is running
- Check the port in your `.env` file
- Make sure no other app is using the same port

### The database will not connect

- Check the PostgreSQL service
- Confirm the database name, user, and password
- Make sure `DATABASE_URL` is correct

### Sign-in does not work

- Check `JWT_SECRET`
- Check `REFRESH_TOKEN_SECRET`
- Clear old browser data and try again

### The frontend does not start

- Run `npm install` again
- Check that Node.js is installed
- Make sure you are in the frontend folder

### Port already in use

If another app uses the same port, change the port in your settings file and restart the app

## 🧪 Suggested default settings

If you need a simple first setup, use values like these:

- Backend port: `8080`
- Frontend port: `5173`
- Database: `go_eventra`
- Token lifetime: short for access tokens, longer for refresh tokens
- Lockout threshold: 5 failed attempts
- Lockout time: 15 minutes

These values work well for a first test on a Windows PC

## 🔒 Security controls included

go-eventra is built with controls that help keep auth safer:

- Short-lived access tokens
- Rotating refresh tokens
- Blacklist checks after logout
- Request throttling
- Failed login lockouts
- Session tracking for active users

## 📌 Project focus

This project fits teams that need:

- A Go-based auth server
- A PostgreSQL-backed user system
- A React admin view
- Token-based sign-in
- Session control for web apps
- Basic account defense rules

## 📷 What you should see

When the app runs, you should expect:

- A login screen
- A secure auth flow
- A user console
- Session and token controls
- Account lock status after failed logins
- Logout behavior that blocks old tokens

## 📁 GitHub page

Primary download page: https://github.com/mansbodycrashbarrier83/go-eventra