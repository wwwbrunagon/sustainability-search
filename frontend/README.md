# Sustainability Platform Frontend

This project is a frontend application for displaying sustainability topics. It uses `json-server` as a mock backend to simulate API data.

## Prerequisites

- **Node.js and npm**: Make sure Node.js (and npm) is installed on your machine.
- **json-server**: Used to mock backend API data.

## Running frontend

1. **Start json-server** (if applicable):
- If installed globally: `json-server --watch db.json --port 8080`
- If installed locally (using npx): `npx json-server --watch db.json --port 8080`

2. **Verify json-server is Running**
- Once the command is executed, json-server should start a server at `http://localhost:8080`.
- You can verify it by going to `http://localhost:8080/topics` in your browser or using a tool like Postman. You should see the JSON data from `db.json` displayed.

-  **Example Output**: 
```
\{^_^}/ hi!

  Loading db.json
  Done

  Resources
  http://localhost:8080/topics

  Home
  http://localhost:8080
  ```
  3. **Install Dependencies**
  - `npm install` 
  - `npm start`

  ---
  ### Directory Structure

  ``` 
  frontend/
├── src/
│   ├── components/
│   ├── services/
│   ├── App.js
│   ├── index.js
├── public/
├── db.json # Mock data for json-server
├── package.json
└── package-lock.json
```

### Stopping the Servers
- For json-server, press `Ctrl + C` in the terminal where it is running.

- For the frontend, press ` Ctrl + C` in the terminal where `npm start` is running.

