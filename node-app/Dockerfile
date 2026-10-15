# 1. Use an official, lightweight Node.js LTS image
FROM node:20-alpine

# 2. Set the working directory inside the container
WORKDIR /usr/src/app

# 3. Copy dependency manifests first to leverage Docker layer caching
COPY package*.json ./

# 4. Install production dependencies
RUN npm install --omit=dev

# 5. Copy the rest of the application source code
COPY . .

# 6. Inform Docker that the container listens on port 3000
EXPOSE 3000

# 7. Start the application
CMD ["node", "server.js"]
