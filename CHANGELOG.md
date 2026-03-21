# Changelog

All notable changes to this project will be documented in this file.

## [1.0.0] - 2026-03-21

### Added

- Initial public release of Eventra Auth Platform.
- Go backend auth service with PostgreSQL integration.
- Endpoints: register, login, refresh, logout, me, and health.
- JWT access token support with refresh token rotation.
- Token blacklist support for logout and rotation security.
- Login protection with failed-attempt tracking and lock escalation.
- HTTP security middleware: headers, CORS allowlist, body size limits, panic recovery.
- Security audit logging repository support.
- Frontend auth console built with React, TypeScript, and Vite.
- Premium UI enhancements: custom favicon, richer background, refined typography, footer credits.
- Comprehensive root README with architecture diagrams, badges, widget visuals, and operational docs.
- CI workflows for backend tests and frontend lint/build.

### Notes

- Repository owner: CodeByPinar.
- Default branch: main.
