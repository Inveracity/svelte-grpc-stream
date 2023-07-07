FROM node:20-alpine3.18 as build

ENV NODE_ENV=production

WORKDIR /app

COPY frontend .
RUN npm install --include=dev
RUN npm run build

ENV HOST=0.0.0.0
EXPOSE 3000
CMD ["node", "build/index.js"]

FROM node:20-alpine3.18

WORKDIR /app
COPY --from=build /app .


ENV HOST=0.0.0.0
EXPOSE 3000
CMD ["node","build/index.js"]
