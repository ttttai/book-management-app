FROM node:22.14

WORKDIR /go/src/app/frontend

COPY ./frontend/package.json ./frontend/package-lock.json ./

RUN npm install

CMD ["npm", "run", "dev"]