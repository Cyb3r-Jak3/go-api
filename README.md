# GO API Server

![Uptime 30 Days](https://img.shields.io/endpoint?url=https%3A%2F%2Fraw.githubusercontent.com%2FCyb3r-Jak3%2Fuptime-stats%2Fmaster%2Fapi%2Fapi%2Fuptime-month.json)
[![Build Docker](https://github.com/Cyb3r-Jak3/go-api/actions/workflows/docker.yml/badge.svg)](https://github.com/Cyb3r-Jak3/go-api/actions/workflows/docker.yml)
[![Test Go](https://github.com/Cyb3r-Jak3/go-api/actions/workflows/golang.yml/badge.svg)](https://github.com/Cyb3r-Jak3/go-api/actions/workflows/golang.yml)

This is go rewrite of my backend for my sites where I have all my api ends points for web interactions. [Python version](https://github.com/Cyb3r-Jak3/api_server)

## Endpoints

### /encrypted_resume

All this endpoint does is return my resume that has been encrypted with users uploaded public key.

### /git/user

Returns my Github API profile

### /git/repos

Returns my Github API repos

### /git/repos/list

Returns an array of my repos with the name and URL of them

### /misc/gravatar

Return a hash for a Gravatar from an email.

Request Format:

```json
{"email": "example@example.com"}
```

Response Format:

```json
{"hash": "<Gravatar ID Hash>"}
```

### /mics/string

Returns lower, upper or title case string.

Request Format:

```json
{"string": "String Here", "modification": "lower|upper|title"}
```

Response Format:

```json
{"string": "String here"}
```
