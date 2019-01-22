const grpc = require('grpc');
const services = require('./ecommerce_grpc_pb');
const Application = require('./src/application');


const startGRPCServer = () => {
  const host = '0.0.0.0:11443'
  const server = new grpc.Server();
  server.addService(services.DiscountService, { applyDiscount: Application.applyDiscount });
  server.bind(host, grpc.ServerCredentials.createInsecure());
  console.log(`Server running on ${host}`);
  server.start();
}

startGRPCServer();
