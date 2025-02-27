FROM node:22.14

WORKDIR /frontend

COPY package.json package-lock.json ./

RUN npm install

CMD ["npm", "run", "dev"]