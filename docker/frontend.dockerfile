FROM node:20-alpine3.18 as build

ENV NODE_ENV=production

WORKDIR /app

COPY frontend .
RUN npm install --include=dev
RUN npm run build


FROM node:20-alpine3.18

WORKDIR /app
COPY --from=build /app .


ENV HOST=0.0.0.0
EXPOSE 4173
CMD ["npm","run", "preview","--", "--host", "0.0.0.0"]
