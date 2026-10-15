# Node.js (Express) Application

## Application Name
Node.js Express Demo API.

## Base Image Selected
- **Image**: `node:20-alpine`
- **Justification**: The `alpine` variant provides a minimal Linux distribution bundled with Node.js LTS (Long Term Support). It dramatically reduces the image footprint compared to standard Debian-based images while remaining fast, secure, and containing all essentials required to run an Express server.

## Port Used
- **Internal Container Port**: `3000`
- **Host Browser Port**: `3000` (`http://localhost:3000`)

## Dockerfile Explanation
- `FROM node:20-alpine`: Sets the minimal Node.js runtime environment.
- `WORKDIR /usr/src/app`: Creates and sets the active working directory inside the container for subsequent commands.
- `COPY package*.json ./`: Copies package manifest files before the source code. This allows Docker to cache the dependency layer and skip running `npm install` on subsequent builds if dependencies have not changed.
- `RUN npm install --omit=dev`: Installs only production dependencies to keep the image lean.
- `COPY . .`: Copies the remaining application files (excluding those in `.dockerignore`) into the container.
- `EXPOSE 3000`: Documents that the application inside the container intends to listen on port 3000.
- `CMD ["node", "server.js"]`: Specifies the command executed when the container starts.

## Build and Run
1. Build the Docker image:
   ```bash
   docker build -t node-express-app .
   ```
2. Run the container:
   ```bash
   docker run -d -p 3000:3000 --name node-container node-express-app
   ```
3. Verify logs:
   ```bash
   docker logs node-container
   ```

## Screenshot
![Browser Screenshot](./screenshot.png)

## Image Size
- **Final Image Size**: *(Run `docker images node-express-app` 202MB)*
