// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var ecommerce_pb = require('./ecommerce_pb.js');

function serialize_ecommerce_DiscountRequest(arg) {
  if (!(arg instanceof ecommerce_pb.DiscountRequest)) {
    throw new Error('Expected argument of type ecommerce.DiscountRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_ecommerce_DiscountRequest(buffer_arg) {
  return ecommerce_pb.DiscountRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ecommerce_DiscountResponse(arg) {
  if (!(arg instanceof ecommerce_pb.DiscountResponse)) {
    throw new Error('Expected argument of type ecommerce.DiscountResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_ecommerce_DiscountResponse(buffer_arg) {
  return ecommerce_pb.DiscountResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var DiscountService = exports.DiscountService = {
  applyDiscount: {
    path: '/ecommerce.Discount/ApplyDiscount',
    requestStream: false,
    responseStream: false,
    requestType: ecommerce_pb.DiscountRequest,
    responseType: ecommerce_pb.DiscountResponse,
    requestSerialize: serialize_ecommerce_DiscountRequest,
    requestDeserialize: deserialize_ecommerce_DiscountRequest,
    responseSerialize: serialize_ecommerce_DiscountResponse,
    responseDeserialize: deserialize_ecommerce_DiscountResponse,
  },
};

exports.DiscountClient = grpc.makeGenericClientConstructor(DiscountService);
