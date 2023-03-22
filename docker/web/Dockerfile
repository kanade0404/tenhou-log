FROM node:18 AS builder

WORKDIR /app
COPY . .
RUN yarn install --frozen-lockfile && yarn build && yarn install --frozen-lockfile --production

FROM gcr.io/distroless/nodejs18-debian11

WORKDIR /app
COPY --from=builder /app/next.config.js .
COPY --from=builder /app/.next .
COPY --from=builder /app/public .
CMD ["yarn", "start"]
