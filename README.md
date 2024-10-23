# Pocketbase Blog
A simple blogging platform built with PocketBase and Vue.js 3. 
## Development

### Project Setup
Ensure you have Node.js 18+ and go 1.23.2+ installed on your machine.

#### Client
1. Install dependencies
```sh
cd client && npm install
```
2. Run the development server
```sh
npm run dev
```

#### Server
1. Install dependencies
```sh
cd server && go mod download
```
2. Run the development server
```sh
go run main.go serve
```

#### Release
1. Ensure you have [goreleaser installed](https://goreleaser.com/install/) on your machin

2. You can test the goreleaser configuration by running the following command
```sh
goreleaser release --snapshot --clean
```

## Requirements
Create a simple blogging platform using PocketBase as the backend and a frontend framework of your choice (e.g., Svelte, React, Vue, or plain JavaScript).

* Key Features
    - User authentication (signup, login, logout)
    - Create, read, update, and delete (CRUD) operations for blog posts
    - Comments system for blog posts
    - Tag system for categorizing posts
    - User profiles with author information
  
* Technical Components
    - ~~Backend (PocketBase~~) 
    - ~~Set up a PocketBase instance~~
    - Create collections for:
        - ~~Users (built-in)~~
        - ~~Posts (title, content, author, created_at, updated_at)~~
        - ~~Comments (content, author, post, created_at)~~
        - ~~Tags (name)~~

* Configure security rules to ensure appropriate access control

    - Frontend
      - ~~Implement user authentication flow~~
      - ~~Create a rich text editor for writing blog posts~~
      - ~~Display a list of blog posts with pagination~~
      - ~~Implement a commenting system for each post~~
      - ~~Add a tagging system for categorizing posts~~
      - ~~Create user profile pages displaying author information and their posts~~
* Bonus Features

    - Search functionality for posts
    - ~~Email notifications for new comments~~
    - ~~Social media sharing buttons~~
    - Basic analytics (e.g., view count for posts)
    - RSS feed for blog posts