FROM node:20 AS builder

WORKDIR /app
COPY package.json .
COPY package-lock.json .
RUN npm install
COPY . .
RUN npm run build

FROM nginx:alpine
RUN mkdir /app
COPY --from=builder /app/dist /usr/share/nginx/html/chat
COPY nginx.conf /etc/nginx/conf.d/default.conf