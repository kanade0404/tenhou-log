FROM node:19 AS builder

# ビルドには devDependencies もインストールする必要があるため
ENV NODE_ENV=development

WORKDIR /app
COPY . .
RUN yarn install && yarn build

FROM node:alpine
ENV NODE_ENV=production

WORKDIR /app
COPY ./package.json ./yarn.lock ./
RUN yarn install
COPY --from=builder /app/next.config.js ./
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
CMD ["yarn", "start"]
