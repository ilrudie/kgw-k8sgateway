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

var extproto_ext_pb = require('../../../../../../../../../../extproto/ext_pb.js');
var github_com_solo$io_solo$kit_api_v1_ref_pb = require('../../../../../../../../../../github.com/solo-io/solo-kit/api/v1/ref_pb.js');
goog.exportSymbol('proto.aws_ec2.options.gloo.solo.io.TagFilter', null, global);
goog.exportSymbol('proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair', null, global);
goog.exportSymbol('proto.aws_ec2.options.gloo.solo.io.UpstreamSpec', null, global);

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
proto.aws_ec2.options.gloo.solo.io.UpstreamSpec = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.repeatedFields_, null);
};
goog.inherits(proto.aws_ec2.options.gloo.solo.io.UpstreamSpec, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.displayName = 'proto.aws_ec2.options.gloo.solo.io.UpstreamSpec';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.repeatedFields_ = [3];



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
proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.prototype.toObject = function(opt_includeInstance) {
  return proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.aws_ec2.options.gloo.solo.io.UpstreamSpec} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.toObject = function(includeInstance, msg) {
  var f, obj = {
    region: jspb.Message.getFieldWithDefault(msg, 1, ""),
    secretRef: (f = msg.getSecretRef()) && github_com_solo$io_solo$kit_api_v1_ref_pb.ResourceRef.toObject(includeInstance, f),
    roleArn: jspb.Message.getFieldWithDefault(msg, 7, ""),
    filtersList: jspb.Message.toObjectList(msg.getFiltersList(),
    proto.aws_ec2.options.gloo.solo.io.TagFilter.toObject, includeInstance),
    publicIp: jspb.Message.getFieldWithDefault(msg, 4, false),
    port: jspb.Message.getFieldWithDefault(msg, 5, 0)
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
 * @return {!proto.aws_ec2.options.gloo.solo.io.UpstreamSpec}
 */
proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.aws_ec2.options.gloo.solo.io.UpstreamSpec;
  return proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.aws_ec2.options.gloo.solo.io.UpstreamSpec} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.aws_ec2.options.gloo.solo.io.UpstreamSpec}
 */
proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 2:
      var value = new github_com_solo$io_solo$kit_api_v1_ref_pb.ResourceRef;
      reader.readMessage(value,github_com_solo$io_solo$kit_api_v1_ref_pb.ResourceRef.deserializeBinaryFromReader);
      msg.setSecretRef(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setRoleArn(value);
      break;
    case 3:
      var value = new proto.aws_ec2.options.gloo.solo.io.TagFilter;
      reader.readMessage(value,proto.aws_ec2.options.gloo.solo.io.TagFilter.deserializeBinaryFromReader);
      msg.addFilters(value);
      break;
    case 4:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setPublicIp(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setPort(value);
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
proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.aws_ec2.options.gloo.solo.io.UpstreamSpec} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getSecretRef();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      github_com_solo$io_solo$kit_api_v1_ref_pb.ResourceRef.serializeBinaryToWriter
    );
  }
  f = message.getRoleArn();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getFiltersList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      3,
      f,
      proto.aws_ec2.options.gloo.solo.io.TagFilter.serializeBinaryToWriter
    );
  }
  f = message.getPublicIp();
  if (f) {
    writer.writeBool(
      4,
      f
    );
  }
  f = message.getPort();
  if (f !== 0) {
    writer.writeUint32(
      5,
      f
    );
  }
};


/**
 * optional string region = 1;
 * @return {string}
 */
proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.prototype.setRegion = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional core.solo.io.ResourceRef secret_ref = 2;
 * @return {?proto.core.solo.io.ResourceRef}
 */
proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.prototype.getSecretRef = function() {
  return /** @type{?proto.core.solo.io.ResourceRef} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$kit_api_v1_ref_pb.ResourceRef, 2));
};


/** @param {?proto.core.solo.io.ResourceRef|undefined} value */
proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.prototype.setSecretRef = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.prototype.clearSecretRef = function() {
  this.setSecretRef(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.prototype.hasSecretRef = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional string role_arn = 7;
 * @return {string}
 */
proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.prototype.getRoleArn = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/** @param {string} value */
proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.prototype.setRoleArn = function(value) {
  jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * repeated TagFilter filters = 3;
 * @return {!Array<!proto.aws_ec2.options.gloo.solo.io.TagFilter>}
 */
proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.prototype.getFiltersList = function() {
  return /** @type{!Array<!proto.aws_ec2.options.gloo.solo.io.TagFilter>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.aws_ec2.options.gloo.solo.io.TagFilter, 3));
};


/** @param {!Array<!proto.aws_ec2.options.gloo.solo.io.TagFilter>} value */
proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.prototype.setFiltersList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 3, value);
};


/**
 * @param {!proto.aws_ec2.options.gloo.solo.io.TagFilter=} opt_value
 * @param {number=} opt_index
 * @return {!proto.aws_ec2.options.gloo.solo.io.TagFilter}
 */
proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.prototype.addFilters = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 3, opt_value, proto.aws_ec2.options.gloo.solo.io.TagFilter, opt_index);
};


proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.prototype.clearFiltersList = function() {
  this.setFiltersList([]);
};


/**
 * optional bool public_ip = 4;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.prototype.getPublicIp = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 4, false));
};


/** @param {boolean} value */
proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.prototype.setPublicIp = function(value) {
  jspb.Message.setProto3BooleanField(this, 4, value);
};


/**
 * optional uint32 port = 5;
 * @return {number}
 */
proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.prototype.getPort = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/** @param {number} value */
proto.aws_ec2.options.gloo.solo.io.UpstreamSpec.prototype.setPort = function(value) {
  jspb.Message.setProto3IntField(this, 5, value);
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
proto.aws_ec2.options.gloo.solo.io.TagFilter = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.aws_ec2.options.gloo.solo.io.TagFilter.oneofGroups_);
};
goog.inherits(proto.aws_ec2.options.gloo.solo.io.TagFilter, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.aws_ec2.options.gloo.solo.io.TagFilter.displayName = 'proto.aws_ec2.options.gloo.solo.io.TagFilter';
}
/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.aws_ec2.options.gloo.solo.io.TagFilter.oneofGroups_ = [[1,2]];

/**
 * @enum {number}
 */
proto.aws_ec2.options.gloo.solo.io.TagFilter.SpecCase = {
  SPEC_NOT_SET: 0,
  KEY: 1,
  KV_PAIR: 2
};

/**
 * @return {proto.aws_ec2.options.gloo.solo.io.TagFilter.SpecCase}
 */
proto.aws_ec2.options.gloo.solo.io.TagFilter.prototype.getSpecCase = function() {
  return /** @type {proto.aws_ec2.options.gloo.solo.io.TagFilter.SpecCase} */(jspb.Message.computeOneofCase(this, proto.aws_ec2.options.gloo.solo.io.TagFilter.oneofGroups_[0]));
};



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
proto.aws_ec2.options.gloo.solo.io.TagFilter.prototype.toObject = function(opt_includeInstance) {
  return proto.aws_ec2.options.gloo.solo.io.TagFilter.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.aws_ec2.options.gloo.solo.io.TagFilter} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.aws_ec2.options.gloo.solo.io.TagFilter.toObject = function(includeInstance, msg) {
  var f, obj = {
    key: jspb.Message.getFieldWithDefault(msg, 1, ""),
    kvPair: (f = msg.getKvPair()) && proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair.toObject(includeInstance, f)
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
 * @return {!proto.aws_ec2.options.gloo.solo.io.TagFilter}
 */
proto.aws_ec2.options.gloo.solo.io.TagFilter.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.aws_ec2.options.gloo.solo.io.TagFilter;
  return proto.aws_ec2.options.gloo.solo.io.TagFilter.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.aws_ec2.options.gloo.solo.io.TagFilter} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.aws_ec2.options.gloo.solo.io.TagFilter}
 */
proto.aws_ec2.options.gloo.solo.io.TagFilter.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setKey(value);
      break;
    case 2:
      var value = new proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair;
      reader.readMessage(value,proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair.deserializeBinaryFromReader);
      msg.setKvPair(value);
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
proto.aws_ec2.options.gloo.solo.io.TagFilter.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.aws_ec2.options.gloo.solo.io.TagFilter.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.aws_ec2.options.gloo.solo.io.TagFilter} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.aws_ec2.options.gloo.solo.io.TagFilter.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {string} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getKvPair();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair.serializeBinaryToWriter
    );
  }
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
proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair.displayName = 'proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair';
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
proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair.prototype.toObject = function(opt_includeInstance) {
  return proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair.toObject = function(includeInstance, msg) {
  var f, obj = {
    key: jspb.Message.getFieldWithDefault(msg, 1, ""),
    value: jspb.Message.getFieldWithDefault(msg, 2, "")
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
 * @return {!proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair}
 */
proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair;
  return proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair}
 */
proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setKey(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setValue(value);
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
proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getKey();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getValue();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional string key = 1;
 * @return {string}
 */
proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair.prototype.getKey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair.prototype.setKey = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string value = 2;
 * @return {string}
 */
proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair.prototype.getValue = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair.prototype.setValue = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string key = 1;
 * @return {string}
 */
proto.aws_ec2.options.gloo.solo.io.TagFilter.prototype.getKey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.aws_ec2.options.gloo.solo.io.TagFilter.prototype.setKey = function(value) {
  jspb.Message.setOneofField(this, 1, proto.aws_ec2.options.gloo.solo.io.TagFilter.oneofGroups_[0], value);
};


proto.aws_ec2.options.gloo.solo.io.TagFilter.prototype.clearKey = function() {
  jspb.Message.setOneofField(this, 1, proto.aws_ec2.options.gloo.solo.io.TagFilter.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.aws_ec2.options.gloo.solo.io.TagFilter.prototype.hasKey = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional KvPair kv_pair = 2;
 * @return {?proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair}
 */
proto.aws_ec2.options.gloo.solo.io.TagFilter.prototype.getKvPair = function() {
  return /** @type{?proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair} */ (
    jspb.Message.getWrapperField(this, proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair, 2));
};


/** @param {?proto.aws_ec2.options.gloo.solo.io.TagFilter.KvPair|undefined} value */
proto.aws_ec2.options.gloo.solo.io.TagFilter.prototype.setKvPair = function(value) {
  jspb.Message.setOneofWrapperField(this, 2, proto.aws_ec2.options.gloo.solo.io.TagFilter.oneofGroups_[0], value);
};


proto.aws_ec2.options.gloo.solo.io.TagFilter.prototype.clearKvPair = function() {
  this.setKvPair(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.aws_ec2.options.gloo.solo.io.TagFilter.prototype.hasKvPair = function() {
  return jspb.Message.getField(this, 2) != null;
};


goog.object.extend(exports, proto.aws_ec2.options.gloo.solo.io);