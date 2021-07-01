# GO API Server

[![Build Docker](https://github.com/Cyb3r-Jak3/go-api/actions/workflows/docker.yml/badge.svg)](https://github.com/Cyb3r-Jak3/go-api/actions/workflows/docker.yml)
[![Test Go](https://github.com/Cyb3r-Jak3/go-api/actions/workflows/golang.yml/badge.svg)](https://github.com/Cyb3r-Jak3/go-api/actions/workflows/golang.yml)

This is go rewrite of my backend for my sites where I have all my api ends points for web interactions. [Python version](https://github.com/Cyb3r-Jak3/api_server)

## Endpoints

### encrypted_resume

All this endpoint does is return my resume that has been encrypted with users uploaded public key.

### git/user

Returns my Github API profile

### git/repos

Returns my Github API repos

### git/repos/list

Returns an array of my repos with the name and URL of them
