ARG DOCKERHUB_URL
FROM ${DOCKERHUB_URL}builder as build


WORKDIR /build/
COPY web3-gateway web3-gateway 
COPY proto proto

RUN cd web3-gateway; make web3-gateway-server 

EXPOSE 8090

#########################################################
FROM node:16.14-alpine

WORKDIR /app
COPY --from=build /build/web3-gateway /app/web3-gateway

WORKDIR /app/web3-gateway
CMD ["node", "dist/app/main"]