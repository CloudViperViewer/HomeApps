FROM node:18-bullseye

# Install git
RUN apt-get update && apt-get install -y git

WORKDIR /app

# Install dependencies
COPY package.json package-lock.json ./
RUN npm install

# Copy configs first (ensure cache invalidation if configs change)
COPY postcss.config.mjs ./

# Copy the rest of the app
COPY . .

EXPOSE 3000

CMD ["npm", "run", "dev"]