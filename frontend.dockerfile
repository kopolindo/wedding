# Use the official Node.js image as base
FROM node:latest
WORKDIR /app
COPY ./frontend/package*.json ./
RUN npm install --force
COPY ./frontend/ .
RUN npm run build
#EXPOSE 3000
#CMD ["npm", "start"]
