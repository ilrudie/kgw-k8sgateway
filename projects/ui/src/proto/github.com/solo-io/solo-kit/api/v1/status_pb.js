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

var extproto_ext_pb = require('../../../../../extproto/ext_pb.js');
var google_protobuf_struct_pb = require('google-protobuf/google/protobuf/struct_pb.js');
goog.exportSymbol('proto.core.solo.io.NamespacedStatuses', null, global);
goog.exportSymbol('proto.core.solo.io.Status', null, global);
goog.exportSymbol('proto.core.solo.io.Status.State', null, global);

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
proto.core.solo.io.NamespacedStatuses = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.core.solo.io.NamespacedStatuses, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.core.solo.io.NamespacedStatuses.displayName = 'proto.core.solo.io.NamespacedStatuses';
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
proto.core.solo.io.NamespacedStatuses.prototype.toObject = function(opt_includeInstance) {
  return proto.core.solo.io.NamespacedStatuses.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.core.solo.io.NamespacedStatuses} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.solo.io.NamespacedStatuses.toObject = function(includeInstance, msg) {
  var f, obj = {
    statusesMap: (f = msg.getStatusesMap()) ? f.toObject(includeInstance, proto.core.solo.io.Status.toObject) : []
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
 * @return {!proto.core.solo.io.NamespacedStatuses}
 */
proto.core.solo.io.NamespacedStatuses.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.core.solo.io.NamespacedStatuses;
  return proto.core.solo.io.NamespacedStatuses.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.core.solo.io.NamespacedStatuses} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.core.solo.io.NamespacedStatuses}
 */
proto.core.solo.io.NamespacedStatuses.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = msg.getStatusesMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readMessage, proto.core.solo.io.Status.deserializeBinaryFromReader, "");
         });
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
proto.core.solo.io.NamespacedStatuses.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.core.solo.io.NamespacedStatuses.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.core.solo.io.NamespacedStatuses} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.solo.io.NamespacedStatuses.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getStatusesMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(1, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeMessage, proto.core.solo.io.Status.serializeBinaryToWriter);
  }
};


/**
 * map<string, Status> statuses = 1;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,!proto.core.solo.io.Status>}
 */
proto.core.solo.io.NamespacedStatuses.prototype.getStatusesMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,!proto.core.solo.io.Status>} */ (
      jspb.Message.getMapField(this, 1, opt_noLazyCreate,
      proto.core.solo.io.Status));
};


proto.core.solo.io.NamespacedStatuses.prototype.clearStatusesMap = function() {
  this.getStatusesMap().clear();
};



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
proto.core.solo.io.Status = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.core.solo.io.Status.repeatedFields_, null);
};
goog.inherits(proto.core.solo.io.Status, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.core.solo.io.Status.displayName = 'proto.core.solo.io.Status';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.core.solo.io.Status.repeatedFields_ = [6];



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
proto.core.solo.io.Status.prototype.toObject = function(opt_includeInstance) {
  return proto.core.solo.io.Status.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.core.solo.io.Status} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.solo.io.Status.toObject = function(includeInstance, msg) {
  var f, obj = {
    state: jspb.Message.getFieldWithDefault(msg, 1, 0),
    reason: jspb.Message.getFieldWithDefault(msg, 2, ""),
    reportedBy: jspb.Message.getFieldWithDefault(msg, 3, ""),
    subresourceStatusesMap: (f = msg.getSubresourceStatusesMap()) ? f.toObject(includeInstance, proto.core.solo.io.Status.toObject) : [],
    details: (f = msg.getDetails()) && google_protobuf_struct_pb.Struct.toObject(includeInstance, f),
    messagesList: jspb.Message.getRepeatedField(msg, 6)
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
 * @return {!proto.core.solo.io.Status}
 */
proto.core.solo.io.Status.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.core.solo.io.Status;
  return proto.core.solo.io.Status.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.core.solo.io.Status} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.core.solo.io.Status}
 */
proto.core.solo.io.Status.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.core.solo.io.Status.State} */ (reader.readEnum());
      msg.setState(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setReason(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setReportedBy(value);
      break;
    case 4:
      var value = msg.getSubresourceStatusesMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readMessage, proto.core.solo.io.Status.deserializeBinaryFromReader, "");
         });
      break;
    case 5:
      var value = new google_protobuf_struct_pb.Struct;
      reader.readMessage(value,google_protobuf_struct_pb.Struct.deserializeBinaryFromReader);
      msg.setDetails(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.addMessages(value);
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
proto.core.solo.io.Status.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.core.solo.io.Status.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.core.solo.io.Status} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.solo.io.Status.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getState();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getReason();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getReportedBy();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getSubresourceStatusesMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(4, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeMessage, proto.core.solo.io.Status.serializeBinaryToWriter);
  }
  f = message.getDetails();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      google_protobuf_struct_pb.Struct.serializeBinaryToWriter
    );
  }
  f = message.getMessagesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      6,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.core.solo.io.Status.State = {
  PENDING: 0,
  ACCEPTED: 1,
  REJECTED: 2,
  WARNING: 3
};

/**
 * optional State state = 1;
 * @return {!proto.core.solo.io.Status.State}
 */
proto.core.solo.io.Status.prototype.getState = function() {
  return /** @type {!proto.core.solo.io.Status.State} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {!proto.core.solo.io.Status.State} value */
proto.core.solo.io.Status.prototype.setState = function(value) {
  jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional string reason = 2;
 * @return {string}
 */
proto.core.solo.io.Status.prototype.getReason = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.core.solo.io.Status.prototype.setReason = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string reported_by = 3;
 * @return {string}
 */
proto.core.solo.io.Status.prototype.getReportedBy = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.core.solo.io.Status.prototype.setReportedBy = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * map<string, Status> subresource_statuses = 4;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,!proto.core.solo.io.Status>}
 */
proto.core.solo.io.Status.prototype.getSubresourceStatusesMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,!proto.core.solo.io.Status>} */ (
      jspb.Message.getMapField(this, 4, opt_noLazyCreate,
      proto.core.solo.io.Status));
};


proto.core.solo.io.Status.prototype.clearSubresourceStatusesMap = function() {
  this.getSubresourceStatusesMap().clear();
};


/**
 * optional google.protobuf.Struct details = 5;
 * @return {?proto.google.protobuf.Struct}
 */
proto.core.solo.io.Status.prototype.getDetails = function() {
  return /** @type{?proto.google.protobuf.Struct} */ (
    jspb.Message.getWrapperField(this, google_protobuf_struct_pb.Struct, 5));
};


/** @param {?proto.google.protobuf.Struct|undefined} value */
proto.core.solo.io.Status.prototype.setDetails = function(value) {
  jspb.Message.setWrapperField(this, 5, value);
};


proto.core.solo.io.Status.prototype.clearDetails = function() {
  this.setDetails(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.core.solo.io.Status.prototype.hasDetails = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * repeated string Messages = 6;
 * @return {!Array<string>}
 */
proto.core.solo.io.Status.prototype.getMessagesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 6));
};


/** @param {!Array<string>} value */
proto.core.solo.io.Status.prototype.setMessagesList = function(value) {
  jspb.Message.setField(this, 6, value || []);
};


/**
 * @param {!string} value
 * @param {number=} opt_index
 */
proto.core.solo.io.Status.prototype.addMessages = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 6, value, opt_index);
};


proto.core.solo.io.Status.prototype.clearMessagesList = function() {
  this.setMessagesList([]);
};


goog.object.extend(exports, proto.core.solo.io);