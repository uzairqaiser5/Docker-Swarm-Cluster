This is a [Next.js](https://nextjs.org) project bootstrapped with [`create-next-app`](https://nextjs.org/docs/app/api-reference/cli/create-next-app).

## Getting Started

# Next.js Application

## Application Name
Next.js Production Web Application (App Router).

## Base Image Selected
- **Image**: `node:20-alpine`
- **Justification**: Alpine Linux is an ultra-compact distribution (~5MB base). Using `node:20-alpine` provides a minimal Node.js LTS environment without bulky OS utilities, significantly cutting image size while maintaining stability for multi-stage building and serving.

## Port Used
- **Internal Container Port**: `3000`
- **Host Browser Port**: `3000` (`http://localhost:3000`)

## Dockerfile Explanation
- `AS deps`: Sets up an isolated stage to install dependencies using `npm ci` based on `package*.json`.
- `AS builder`: Copies the dependencies from the `deps` stage and application source files to execute `npm run build`, producing production build artifacts.
- `AS runner`: A completely clean Alpine stage that serves as the final runtime environment.
- `output: "standalone"`: Instructed Next.js to isolate only the required node modules and assets into `.next/standalone`.
- `COPY --from=builder`: Copies only the standalone runtime server, public assets, and `.next/static` folder into the final runner image, discarding heavy devDependencies and source code.
- `ENV HOSTNAME="0.0.0.0"`: Binds the server to all network interfaces so it can receive traffic from outside the container.
- `USER nextjs`: Runs the application as a non-privileged user rather than `root` for enhanced container security.
- `CMD ["node", "server.js"]`: Executes the generated standalone server directly.

## Build and Run
1. Build the Docker image:
   ```bash
   docker build -t nextjs-app .
   ```
2. Run the container:
   ```bash
   docker run -d -p 3000:3000 --name next-container nextjs-app
   ```
3. Check container logs:
   ```bash
   docker logs next-container
   ```

## Screenshot
![Browser Screenshot](./screenshot.png)

## Image Size
- **Final Image Size**: *(Run `docker images nextjs-app` 289MB)*
- **Optimization Strategy**: By utilizing Next.js's standalone build feature and a 3-stage multi-stage Docker build, all compilers, development dependencies, and non-essential source files were omitted from the final runner layer.







