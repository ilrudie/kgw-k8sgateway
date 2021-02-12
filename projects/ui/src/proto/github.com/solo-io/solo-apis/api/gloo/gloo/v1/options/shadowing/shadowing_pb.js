/* eslint-disable */
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var github_com_solo$io_solo$kit_api_v1_ref_pb = require('../../../../../../../../../github.com/solo-io/solo-kit/api/v1/ref_pb.js');
var extproto_ext_pb = require('../../../../../../../../../extproto/ext_pb.js');
goog.exportSymbol('proto.shadowing.options.gloo.solo.io.RouteShadowing', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.shadowing.options.gloo.solo.io.RouteShadowing = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.shadowing.options.gloo.solo.io.RouteShadowing, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.shadowing.options.gloo.solo.io.RouteShadowing.displayName = 'proto.shadowing.options.gloo.solo.io.RouteShadowing';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.shadowing.options.gloo.solo.io.RouteShadowing.prototype.toObject = function(opt_includeInstance) {
  return proto.shadowing.options.gloo.solo.io.RouteShadowing.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.shadowing.options.gloo.solo.io.RouteShadowing} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.shadowing.options.gloo.solo.io.RouteShadowing.toObject = function(includeInstance, msg) {
  var f, obj = {
    upstream: (f = msg.getUpstream()) && github_com_solo$io_solo$kit_api_v1_ref_pb.ResourceRef.toObject(includeInstance, f),
    percentage: +jspb.Message.getFieldWithDefault(msg, 2, 0.0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.shadowing.options.gloo.solo.io.RouteShadowing}
 */
proto.shadowing.options.gloo.solo.io.RouteShadowing.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.shadowing.options.gloo.solo.io.RouteShadowing;
  return proto.shadowing.options.gloo.solo.io.RouteShadowing.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.shadowing.options.gloo.solo.io.RouteShadowing} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.shadowing.options.gloo.solo.io.RouteShadowing}
 */
proto.shadowing.options.gloo.solo.io.RouteShadowing.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new github_com_solo$io_solo$kit_api_v1_ref_pb.ResourceRef;
      reader.readMessage(value,github_com_solo$io_solo$kit_api_v1_ref_pb.ResourceRef.deserializeBinaryFromReader);
      msg.setUpstream(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setPercentage(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.shadowing.options.gloo.solo.io.RouteShadowing.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.shadowing.options.gloo.solo.io.RouteShadowing.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.shadowing.options.gloo.solo.io.RouteShadowing} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.shadowing.options.gloo.solo.io.RouteShadowing.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUpstream();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      github_com_solo$io_solo$kit_api_v1_ref_pb.ResourceRef.serializeBinaryToWriter
    );
  }
  f = message.getPercentage();
  if (f !== 0.0) {
    writer.writeFloat(
      2,
      f
    );
  }
};


/**
 * optional core.solo.io.ResourceRef upstream = 1;
 * @return {?proto.core.solo.io.ResourceRef}
 */
proto.shadowing.options.gloo.solo.io.RouteShadowing.prototype.getUpstream = function() {
  return /** @type{?proto.core.solo.io.ResourceRef} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$kit_api_v1_ref_pb.ResourceRef, 1));
};


/** @param {?proto.core.solo.io.ResourceRef|undefined} value */
proto.shadowing.options.gloo.solo.io.RouteShadowing.prototype.setUpstream = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.shadowing.options.gloo.solo.io.RouteShadowing.prototype.clearUpstream = function() {
  this.setUpstream(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.shadowing.options.gloo.solo.io.RouteShadowing.prototype.hasUpstream = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional float percentage = 2;
 * @return {number}
 */
proto.shadowing.options.gloo.solo.io.RouteShadowing.prototype.getPercentage = function() {
  return /** @type {number} */ (+jspb.Message.getFieldWithDefault(this, 2, 0.0));
};


/** @param {number} value */
proto.shadowing.options.gloo.solo.io.RouteShadowing.prototype.setPercentage = function(value) {
  jspb.Message.setProto3FloatField(this, 2, value);
};


goog.object.extend(exports, proto.shadowing.options.gloo.solo.io);
