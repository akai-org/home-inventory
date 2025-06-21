FROM node:24-alpine3.21 AS base

WORKDIR /app
COPY .npmrc package.json pnpm-lock.yaml ./
RUN npm install --silent

COPY . .

FROM base AS dev
RUN echo "Dev mode"
ENTRYPOINT ["npm", "run", "dev", "--", "--host"]
