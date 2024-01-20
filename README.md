# NBA Schedule XMLTV Server

This Go application serves NBA schedule data in XMLTV format. The schedule data is fetched from the NBA API, and the server generates an XMLTV feed that can be used with various media centers and TV guide applications.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Docker](#docker)
- [Configuration](#configuration)
- [Endpoints](#endpoints)
- [Contributing](#contributing)
- [License](#license)

## Prerequisites

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/nba-schedule-xmltv-server.git
   cd nba-schedule-xmltv-server
   ```

2. Build the application:

   ```bash
   go build -o main .
   ```

3. Run the application:

   ```bash
   ./main
   ```

   The server will be running at [http://localhost:8080/nba_schedule](http://localhost:8080/nba_schedule).

## Usage

The server exposes an XMLTV feed containing NBA schedule information. The feed can be accessed at the following endpoint:

- [http://localhost:8080/nba_schedule](http://localhost:8080/nba_schedule)

## Docker

You can also run the application using Docker. Build the Docker image and run the container:

```bash
docker pull registry.snry.xyz/sysadmin/nba-xmltv-schedule:main-latest
docker run -p 8080:8080 registry.snry.xyz/sysadmin/nba-xmltv-schedule:main-latest
```

## Configuration

The application fetches NBA schedule data from the NBA API. No additional configuration is required.

## Endpoints

- `/nba_schedule`: Returns the NBA schedule in XMLTV format.
- `/logos/{team_id}`: Returns a PNG logo of the specified team

## Contributing

Feel free to contribute to this project by opening issues or submitting pull requests. Your contributions are highly appreciated.

## License

This project is licensed under the [MIT License](LICENSE).