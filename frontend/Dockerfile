FROM node:8-alpine as builder

ADD . /opt/frontend
WORKDIR /opt/frontend
RUN npm install
RUN npm run build

FROM nginx:alpine

COPY --from=builder /opt/frontend/dist /usr/share/nginx/html