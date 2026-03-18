# Build the app
FROM node:20-alpine AS build
WORKDIR /app

COPY package*.json ./
RUN npm install --legacy-peer-deps

COPY . .
RUN npm run build --prod

# Serve with Nginx
FROM nginx:alpine

# Copy everything directly from the dist folder to Nginx
COPY --from=build /app/dist /usr/share/nginx/html

# COPY YOUR CUSTOM CONFIG HERE:
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]