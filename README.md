# Task Manager Application

This is a task management application built with Go (Echo) for the backend, Vue.js (Nuxt.js) for the frontend, and PostgreSQL for the database.

## Features

- User authentication
- Workspace management
- Task creation and management
- Comments on tasks

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Docker and Docker Compose installed

## Getting Started

To get the application running, follow these steps:

1. Clone the repository:

    ``` shell
    $ git clone https://github.com/naoking158/taskmanager.git
    $ cd taskmanager
    ```

2. Copy the example environment file and edit it as needed:

    ``` shell
    $ cp .env.example .env
    ```

3. Build and start the Docker containers:

    ``` shell
    $ docker compose up --build
    ```


The application should now be running at:
- Frontend: http://localhost:3000
- Backend API: http://localhost:8888
