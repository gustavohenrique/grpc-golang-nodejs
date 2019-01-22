FROM node:11-alpine
LABEL author="Gustavo Henrique <gustavo@gustavohenrique.net>"

COPY . /app
WORKDIR /app

ENV NODE_ENV=production

RUN npm install

CMD ["npm", "start"]