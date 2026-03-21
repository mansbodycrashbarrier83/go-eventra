# Eventra v1.0.0

Initial public release of Eventra Auth Platform.

## Highlights

- Production-oriented Go auth API with PostgreSQL persistence.
- Secure token lifecycle with JWT access tokens and rotating refresh tokens.
- Login protection with failed-attempt tracking and lock escalation.
- Access token blacklist support for stronger logout and refresh security.
- Premium React + TypeScript + Vite auth console.
- CI workflows for backend tests and frontend lint/build.

## Included Features

### Backend (Go)

- Register, login, refresh, logout, me, and health endpoints.
- Password hashing with bcrypt.
- Token blacklist integration.
- Security middleware:
  - CORS allowlist
  - request body limit
  - hardening headers
  - panic recovery
- Security audit repository support.

### Frontend (React)

- Real-time auth flow UI for register/login/refresh/logout/me.
- API health indicator and endpoint coverage view.
- Improved visual system:
  - custom favicon
  - premium background and typography
  - developer footer signature

### Documentation & DevEx

- Comprehensive root README:
  - architecture and lifecycle diagrams (Mermaid)
  - badges and widgets
  - quick start and operational defaults
  - troubleshooting and roadmap
- Changelog for v1.0.0.

## CI

- Backend workflow: .github/workflows/ci-backend.yml
- Frontend workflow: .github/workflows/ci-frontend.yml

## Tag

- v1.0.0

## Notes

- License file is not yet defined. Add a LICENSE file to publish license metadata.
