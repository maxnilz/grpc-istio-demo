/**
 * @fileoverview gRPC-Web generated client stub for proto
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.proto = require('./emoji-without-annotations_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.EmojiServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.EmojiServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.HelloRequest,
 *   !proto.proto.HelloResponse>}
 */
const methodInfo_EmojiService_SayHello = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.HelloResponse,
  /** @param {!proto.proto.HelloRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.HelloResponse.deserializeBinary
);


/**
 * @param {!proto.proto.HelloRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.HelloResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.HelloResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.EmojiServiceClient.prototype.sayHello =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.EmojiService/SayHello',
      request,
      metadata || {},
      methodInfo_EmojiService_SayHello,
      callback);
};


/**
 * @param {!proto.proto.HelloRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.HelloResponse>}
 *     A native promise that resolves to the response
 */
proto.proto.EmojiServicePromiseClient.prototype.sayHello =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.EmojiService/SayHello',
      request,
      metadata || {},
      methodInfo_EmojiService_SayHello);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.EmojiRequest,
 *   !proto.proto.EmojiResponse>}
 */
const methodInfo_EmojiService_InsertEmojis = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.EmojiResponse,
  /** @param {!proto.proto.EmojiRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.EmojiResponse.deserializeBinary
);


/**
 * @param {!proto.proto.EmojiRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.EmojiResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.EmojiResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.EmojiServiceClient.prototype.insertEmojis =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.EmojiService/InsertEmojis',
      request,
      metadata || {},
      methodInfo_EmojiService_InsertEmojis,
      callback);
};


/**
 * @param {!proto.proto.EmojiRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.EmojiResponse>}
 *     A native promise that resolves to the response
 */
proto.proto.EmojiServicePromiseClient.prototype.insertEmojis =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.EmojiService/InsertEmojis',
      request,
      metadata || {},
      methodInfo_EmojiService_InsertEmojis);
};


module.exports = proto.proto;

