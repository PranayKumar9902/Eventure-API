# 🎟️ TicketBooking Backend (Go)

Welcome to the TicketBooking Backend project! This repository contains the backend code for a ticket booking system, built using Go. 

## 🚀 Features

- User authentication and authorization
- Ticket booking and validation
- Event management
- Admin dashboard


## 📖 Usage

### 🔐 User Authentication

- **Register**: Users can register by sending a POST request to `/api/user/register` with their details.
- **Login**: Users can log in by sending a POST request to `/api/user/login` with their credentials.

### 🎫 Ticket Booking

- **Book Ticket**: Book tickets by sending a POST request to `/api/tickets` with event and ticket details.
- **Get Tickets**: Get all the booked tickets by sending a GET request to `/api/tickets/{user_id}` with their user_id.
- **Validate Ticket**: Validate tickets by sending a POST request to `/api/tickets/validate` with their user_id and ticket_id details.

### 📅 Event Management

- **Create Event**: Admins can create events by sending a POST request to `/api/events` with event details.
- **Update Event**: Admins can update events by sending a PUT request to `/api/events/{event_id}`.
- **Delete Event**: Admins can delete events by sending a DELETE request to `/api/events/{event_id}`.


## 🛠️ Tech Stack

- **Go**: The main programming language used for the backend.
- **Fiber**: Web framework for building the API.
- **GORM**: ORM library for database interactions.
- **PostgreSQL**: Database for storing data.

## 📦 Installation

1. **Clone the repository**
    ```sh
    git clone https://github.com/PranayKumar9902/Eventure-API.git
    cd TicketBooking/backend
    ```

2. **Set up environment variables**
    Create a `.env` file in the root directory and add the following:
    ```env
    dbstring = "host=your_host user=your_username password=your_password dbname=your_dbname port=your_port sslmode=disable TimeZone=your_timezone"
    JWT_SECRET = "your_secret_key"
    ```

3. **Install dependencies**
    ```sh
    go mod tidy
    ```

4. **Run the application**
    ```sh
    go run main.go
    ```

## ⚙️ Configuration

- **Database Configuration**: Ensure PostgreSQL is installed and running. Update the `.env` file with your database credentials.
- **JWT Configuration**: Set the `JWT_SECRET` in the `.env` file for token generation and validation.
