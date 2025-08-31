FROM node:latest as builder

WORKDIR /build

COPY . .

RUN npm install -g pnpm
RUN pnpm install && pnpm run build

FROM nginx:alpine as app

COPY --from=builder /build/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
