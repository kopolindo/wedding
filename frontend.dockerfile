# Use the official Node.js image as base
FROM node:latest
WORKDIR /app
COPY ./frontend/package*.json ./
RUN npm install
COPY ./frontend/ .
EXPOSE 3000
CMD ["npm", "start"]
