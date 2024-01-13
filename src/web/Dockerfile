# Build frontend.
FROM node:21 as build

WORKDIR /app

COPY ./frontend ./

RUN npm install
RUN npm run build

# Start nginx.
FROM nginx:alpine

COPY --from=build /app/dist /usr/share/nginx/html
COPY --from=build /app/public /usr/share/nginx/html/public

COPY ./config/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80 443

CMD ["nginx", "-g", "daemon off;"]