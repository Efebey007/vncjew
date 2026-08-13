# VNCJew Docker

This project contains Docker images for the VNCJew server and client.

The images are built by cloning the GitHub repository and keeping only the relevant directory:

- server: uses the repo's `server` folder
- client: uses the repo's `client` folder

## Requirements

- Docker
- Docker Compose (optional)

## Build the server image

From the `server` directory:

```bash
docker build -t vncjew-server .
```

## Run the server

```bash
docker run --rm -it \
  -p 8080:8080 \
  -e ADMIN_PASS=your_admin_password \
  -e CLIENT_PASS=your_client_password \
  vncjew-server
```

Environment variables:

- `ADMIN_PASS`: password for the admin account
- `CLIENT_PASS`: password for the client account

The server will start on port `8080`.

## Build the client image

From the `client` directory:

```bash
docker build -t vncjew-client .
```

## Run the client

```bash
docker run --rm -it \
  -e SERVER_ADDR=host.docker.internal:8080 \
  -e CLIENT_PASSWORD=your_client_password \
  vncjew-client
```

Environment variables:

- `SERVER_ADDR`: address of the VNCJew server, for example `localhost:8080` or `host.docker.internal:8080`
- `CLIENT_PASSWORD`: client password used by the client to connect

## Example with local host

If the server is running on your machine:

```bash
docker run --rm -it \
  -p 8080:8080 \
  -e ADMIN_PASS=admin123 \
  -e CLIENT_PASS=client123 \
  vncjew-server
```

Then run the client:

```bash
docker run --rm -it \
  -e SERVER_ADDR=host.docker.internal:8080 \
  -e CLIENT_PASSWORD=client123 \
  vncjew-client
```

## Notes

- The Dockerfiles clone `https://github.com/Efebey007/vncjew.git` and keep only the corresponding `server` or `client` directory.
- Other files and folders from the repository are removed before the app is built.
