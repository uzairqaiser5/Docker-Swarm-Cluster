# Go Application

## Application Name
Go HTTP Web Server.

## Base Image Selected
- **Builder Stage**: `golang:1.22-alpine` (Minimal Alpine distribution bundled with the official Go compiler and toolchain).
- **Final Runtime**: `scratch` (Docker's zero-byte reserved empty image).
- **Justification**: Go compiles down to a self-contained, statically linked machine binary. Using `scratch` in the final stage removes all operating system files, shell utilities, and package managers, producing a secure image with zero excess overhead.

## Port Used
- **Internal Container Port**: `8080`
- **Host Browser Port**: `8080` (`http://localhost:8080`)

## Dockerfile Explanation
- `FROM golang:1.22-alpine AS builder`: Loads the Go toolchain necessary to compile the source code.
- `WORKDIR /app`: Defines the build directory.
- `COPY go.mod ./` & `COPY . .`: Copies the module configuration and source files into the builder.
- `RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server .`:
  - `CGO_ENABLED=0`: Disables dynamic linking against the C standard library, ensuring the binary is 100% portable and standalone.
  - `-ldflags="-s -w"`: Strips out DWARF debug symbols and symbol tables to minimize binary footprint.
  - `-o server`: Outputs the compiled binary file.
- `FROM scratch`: Starts an empty layer containing no operating system files.
- `COPY --from=builder /app/server /server`: Copies strictly the standalone compiled binary from the builder layer into the runtime layer.
- `EXPOSE 8080`: Documents the listening port for incoming HTTP traffic.
- `ENTRYPOINT ["/server"]`: Executes the binary immediately when the container runs.

## Build and Run
1. Build the Docker image:
   ```bash
   docker build -t go-app .
   ```
2. Run the container:
   ```bash
   docker run -d -p 8080:8080 --name go-container go-app
   ```
3. Check runtime startup logs:
   ```bash
   docker logs go-container
   ```

## Screenshot
![Browser Screenshot](./screenshot.png)

## Image Size
- **Final Image Size**: *(Run `docker images go-app` and record the size here, typically ~10–15MB)*
- **Optimization Strategy**: By using a multi-stage build targeting Docker's `scratch` base image, eliminating CGO runtime dependencies, and stripping binary symbols with `-ldflags="-s -w"`, the entire runtime OS was omitted, keeping the final container footprint to just the raw executable.
