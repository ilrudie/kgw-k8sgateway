// package: fed.rpc.solo.io
// file: github.com/solo-io/solo-projects/projects/apiserver/api/fed.rpc/v1/ratelimit_resources.proto

var github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_ratelimit_resources_pb = require("../../../../../../../../github.com/solo-io/solo-projects/projects/apiserver/api/fed.rpc/v1/ratelimit_resources_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var RatelimitResourceApi = (function () {
  function RatelimitResourceApi() {}
  RatelimitResourceApi.serviceName = "fed.rpc.solo.io.RatelimitResourceApi";
  return RatelimitResourceApi;
}());

RatelimitResourceApi.ListRateLimitConfigs = {
  methodName: "ListRateLimitConfigs",
  service: RatelimitResourceApi,
  requestStream: false,
  responseStream: false,
  requestType: github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_ratelimit_resources_pb.ListRateLimitConfigsRequest,
  responseType: github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_ratelimit_resources_pb.ListRateLimitConfigsResponse
};

RatelimitResourceApi.GetRateLimitConfigYaml = {
  methodName: "GetRateLimitConfigYaml",
  service: RatelimitResourceApi,
  requestStream: false,
  responseStream: false,
  requestType: github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_ratelimit_resources_pb.GetRateLimitConfigYamlRequest,
  responseType: github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_ratelimit_resources_pb.GetRateLimitConfigYamlResponse
};

exports.RatelimitResourceApi = RatelimitResourceApi;

function RatelimitResourceApiClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

RatelimitResourceApiClient.prototype.listRateLimitConfigs = function listRateLimitConfigs(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RatelimitResourceApi.ListRateLimitConfigs, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

RatelimitResourceApiClient.prototype.getRateLimitConfigYaml = function getRateLimitConfigYaml(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(RatelimitResourceApi.GetRateLimitConfigYaml, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.RatelimitResourceApiClient = RatelimitResourceApiClient;

