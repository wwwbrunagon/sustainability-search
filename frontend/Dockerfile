# Step 1: Use a Node.js image to build the React app
FROM node:16 AS builder

# Step 2: Set the working directory
WORKDIR /app

# Step 3: Copy package.json and install dependencies
COPY package*.json ./
RUN npm install

# Step 4: Copy the app source code and build the application
COPY . .
RUN npm run build

# Step 5: Use an Nginx server to serve the static files
FROM nginx:alpine
COPY --from=builder /app/build /usr/share/nginx/html

# Step 6: Expose port 80 and start Nginx
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
