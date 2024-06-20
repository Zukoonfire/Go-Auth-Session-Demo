# Go-Auth-Session-Demo
# Go-Auth-Session-Demo

A demo project showcasing user authentication and session management using Go.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [License](#license)

## Introduction

This project demonstrates a basic user authentication system using Go, featuring password hashing, session management, and user registration. It is a great starting point for learning about web development with Go.

## Features

- User Registration
- User Login
- Password Hashing
- Session Management
- Secure User Authentication

## Installation

1. **Clone the repository:**
    ```sh
    git clone https://github.com/yourusername/Go-Auth-Session-Demo.git
    cd Go-Auth-Session-Demo
    ```

2. **Install dependencies:**
    Ensure you have [Go](https://golang.org/dl/) installed on your machine.

    ```sh
    go mod tidy
    ```

3. **Run the application:**
    ```sh
    go run main.go
    ```

4. **Open your browser:**
    Navigate to `http://localhost:8080/register` to register a new user.

## Usage

### Register a New User

1. Navigate to `http://localhost:8080/register`.
2. Fill out the registration form and submit.

### Login

1. Navigate to `http://localhost:8080/login`.
2. Enter your registered email and password and submit.

### Logout

1. Navigate to `http://localhost:8080/logout` to log out.

## Project Structure

```plaintext
Go-Auth-Session-Demo/
|-- main.go
|-- handlers/
|   |-- auth.go
|-- models/
|   |-- user.go
|-- templates/
|   |-- login.html
|   |-- register.html
|   |-- welcome.html
|-- static/
|   |-- css/
|       |-- styles.css
|-- go.mod
|-- go.sum
