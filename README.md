# Eventure-API

## Ticket Booking App

### Setting Up the `.env` File

To properly configure the environment variables for the **Eventure API** (Ticket Booking App), follow these steps to create a `.env` file in your project directory.

### Step 1: Create a `.env` File

In the root directory of your project, create a `.env` file if it doesn't already exist. This file will store sensitive information such as your database credentials and JWT secret key, which should not be hardcoded in your source code.

### Step 2: Add PostgreSQL Database Credentials

Add the following variable `dbstring` in the `.env` file to store your PostgreSQL connection details. Replace the placeholders with your actual PostgreSQL credentials:

```env
dbstring="host=yourhost user=username password=password dbname=database port=yourport sslmode=yoursslmode TimeZone=your_timezone"
