# Stage 1: Build React app
FROM node:22-alpine3.20 as build

# Set working directory
WORKDIR /app

# Copy package.json and install dependencies
COPY package.json package-lock.json ./
RUN npm install

# Copy the rest of the app
COPY . .

# Build the React app
RUN npm run build

# Stage 2: Serve app with Nginx
FROM nginx:1.27.2-alpine3.20

# Copy built React app from the previous stage
COPY --from=build /app/build /usr/share/nginx/html

# # # Copy custom nginx config if needed
# COPY ./nginx.conf /etc/nginx/conf.d/default.conf

# Expose port 80
EXPOSE 80

# Start nginx
CMD ["nginx", "-g", "daemon off;"]
