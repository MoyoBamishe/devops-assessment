# Build the app
FROM node:20-alpine AS build
WORKDIR /app

RUN apk add --no-cache python3 make g++ py3-setuptools

COPY package*.json ./
RUN npm install --legacy-peer-deps

COPY . .
RUN rm -rf dist
RUN npm run build --prod

# Serve with Nginx
FROM nginx:alpine

# Copy everything directly from the dist folder to Nginx
# COPY --from=build /app/dist/ /usr/share/nginx/html/
COPY --from=build /app/dist/ /usr/share/nginx/html/
# COPY --from=build /app/dist /usr/share/nginx/html

# COPY YOUR CUSTOM CONFIG HERE:
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
