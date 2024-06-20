# golang-gin-gorm-dotenv-starter-app

This repository contains a starter app built using the following technologies:

Gin: A web framework written in Go (Golang).
GORM: The fantastic ORM library for Golang.
dotenv: A Go port of the Ruby dotenv library, which loads environment variables from a .env file.
Air: Live reload for Go apps.
Features
Gin: Fast, minimalistic web framework for building Go applications.
GORM: ORM library for database interaction.
Dotenv: Manage environment variables efficiently.
Air: Hot reloading for development.
Getting Started
Prerequisites
Go (1.16+)
Git
Air for live reloading
Installation
Clone the repository:

sh
Copy code
git clone https://github.com/your-username/your-repo-name.git
cd your-repo-name
Set up environment variables:
Create a .env file in the root directory and add your environment variables:

sh
Copy code
touch .env
echo 'DB_HOST=localhost' >> .env
echo 'DB_USER=username' >> .env
echo 'DB_PASSWORD=password' >> .env
echo 'DB_NAME=your_db_name' >> .env
Install dependencies:

sh
Copy code
go mod tidy
Running the Application
To start the project with live reloading using Air:

sh
Copy code
air
Air will monitor your files for changes and automatically reload the server.

Contributing
Contributions are welcome! Please open an issue or submit a pull request for any changes.
