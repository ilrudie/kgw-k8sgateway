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
var google_protobuf_wrappers_pb = require('google-protobuf/google/protobuf/wrappers_pb.js');
goog.exportSymbol('proto.core.solo.io.Metadata', null, global);
goog.exportSymbol('proto.core.solo.io.Metadata.OwnerReference', null, global);

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
proto.core.solo.io.Metadata = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.core.solo.io.Metadata.repeatedFields_, null);
};
goog.inherits(proto.core.solo.io.Metadata, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.core.solo.io.Metadata.displayName = 'proto.core.solo.io.Metadata';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.core.solo.io.Metadata.repeatedFields_ = [9];



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
proto.core.solo.io.Metadata.prototype.toObject = function(opt_includeInstance) {
  return proto.core.solo.io.Metadata.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.core.solo.io.Metadata} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.solo.io.Metadata.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 2, ""),
    namespace: jspb.Message.getFieldWithDefault(msg, 3, ""),
    cluster: jspb.Message.getFieldWithDefault(msg, 7, ""),
    resourceVersion: jspb.Message.getFieldWithDefault(msg, 4, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    annotationsMap: (f = msg.getAnnotationsMap()) ? f.toObject(includeInstance, undefined) : [],
    generation: jspb.Message.getFieldWithDefault(msg, 8, 0),
    ownerReferencesList: jspb.Message.toObjectList(msg.getOwnerReferencesList(),
    proto.core.solo.io.Metadata.OwnerReference.toObject, includeInstance)
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
 * @return {!proto.core.solo.io.Metadata}
 */
proto.core.solo.io.Metadata.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.core.solo.io.Metadata;
  return proto.core.solo.io.Metadata.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.core.solo.io.Metadata} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.core.solo.io.Metadata}
 */
proto.core.solo.io.Metadata.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setNamespace(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setCluster(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setResourceVersion(value);
      break;
    case 5:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "");
         });
      break;
    case 6:
      var value = msg.getAnnotationsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "");
         });
      break;
    case 8:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setGeneration(value);
      break;
    case 9:
      var value = new proto.core.solo.io.Metadata.OwnerReference;
      reader.readMessage(value,proto.core.solo.io.Metadata.OwnerReference.deserializeBinaryFromReader);
      msg.addOwnerReferences(value);
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
proto.core.solo.io.Metadata.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.core.solo.io.Metadata.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.core.solo.io.Metadata} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.solo.io.Metadata.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getNamespace();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getCluster();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getResourceVersion();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(5, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getAnnotationsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(6, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getGeneration();
  if (f !== 0) {
    writer.writeInt64(
      8,
      f
    );
  }
  f = message.getOwnerReferencesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      9,
      f,
      proto.core.solo.io.Metadata.OwnerReference.serializeBinaryToWriter
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
proto.core.solo.io.Metadata.OwnerReference = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.core.solo.io.Metadata.OwnerReference, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.core.solo.io.Metadata.OwnerReference.displayName = 'proto.core.solo.io.Metadata.OwnerReference';
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
proto.core.solo.io.Metadata.OwnerReference.prototype.toObject = function(opt_includeInstance) {
  return proto.core.solo.io.Metadata.OwnerReference.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.core.solo.io.Metadata.OwnerReference} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.solo.io.Metadata.OwnerReference.toObject = function(includeInstance, msg) {
  var f, obj = {
    apiVersion: jspb.Message.getFieldWithDefault(msg, 1, ""),
    blockOwnerDeletion: (f = msg.getBlockOwnerDeletion()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    controller: (f = msg.getController()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    kind: jspb.Message.getFieldWithDefault(msg, 4, ""),
    name: jspb.Message.getFieldWithDefault(msg, 5, ""),
    uid: jspb.Message.getFieldWithDefault(msg, 6, "")
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
 * @return {!proto.core.solo.io.Metadata.OwnerReference}
 */
proto.core.solo.io.Metadata.OwnerReference.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.core.solo.io.Metadata.OwnerReference;
  return proto.core.solo.io.Metadata.OwnerReference.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.core.solo.io.Metadata.OwnerReference} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.core.solo.io.Metadata.OwnerReference}
 */
proto.core.solo.io.Metadata.OwnerReference.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setApiVersion(value);
      break;
    case 2:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setBlockOwnerDeletion(value);
      break;
    case 3:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setController(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setKind(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setUid(value);
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
proto.core.solo.io.Metadata.OwnerReference.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.core.solo.io.Metadata.OwnerReference.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.core.solo.io.Metadata.OwnerReference} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.core.solo.io.Metadata.OwnerReference.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getApiVersion();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getBlockOwnerDeletion();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getController();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getKind();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getUid();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
};


/**
 * optional string api_version = 1;
 * @return {string}
 */
proto.core.solo.io.Metadata.OwnerReference.prototype.getApiVersion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.core.solo.io.Metadata.OwnerReference.prototype.setApiVersion = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional google.protobuf.BoolValue block_owner_deletion = 2;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.core.solo.io.Metadata.OwnerReference.prototype.getBlockOwnerDeletion = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 2));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.core.solo.io.Metadata.OwnerReference.prototype.setBlockOwnerDeletion = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.core.solo.io.Metadata.OwnerReference.prototype.clearBlockOwnerDeletion = function() {
  this.setBlockOwnerDeletion(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.core.solo.io.Metadata.OwnerReference.prototype.hasBlockOwnerDeletion = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional google.protobuf.BoolValue controller = 3;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.core.solo.io.Metadata.OwnerReference.prototype.getController = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 3));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.core.solo.io.Metadata.OwnerReference.prototype.setController = function(value) {
  jspb.Message.setWrapperField(this, 3, value);
};


proto.core.solo.io.Metadata.OwnerReference.prototype.clearController = function() {
  this.setController(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.core.solo.io.Metadata.OwnerReference.prototype.hasController = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional string kind = 4;
 * @return {string}
 */
proto.core.solo.io.Metadata.OwnerReference.prototype.getKind = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.core.solo.io.Metadata.OwnerReference.prototype.setKind = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string name = 5;
 * @return {string}
 */
proto.core.solo.io.Metadata.OwnerReference.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.core.solo.io.Metadata.OwnerReference.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string uid = 6;
 * @return {string}
 */
proto.core.solo.io.Metadata.OwnerReference.prototype.getUid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/** @param {string} value */
proto.core.solo.io.Metadata.OwnerReference.prototype.setUid = function(value) {
  jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional string name = 2;
 * @return {string}
 */
proto.core.solo.io.Metadata.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.core.solo.io.Metadata.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string namespace = 3;
 * @return {string}
 */
proto.core.solo.io.Metadata.prototype.getNamespace = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.core.solo.io.Metadata.prototype.setNamespace = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string cluster = 7;
 * @return {string}
 */
proto.core.solo.io.Metadata.prototype.getCluster = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/** @param {string} value */
proto.core.solo.io.Metadata.prototype.setCluster = function(value) {
  jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional string resource_version = 4;
 * @return {string}
 */
proto.core.solo.io.Metadata.prototype.getResourceVersion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.core.solo.io.Metadata.prototype.setResourceVersion = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * map<string, string> labels = 5;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.core.solo.io.Metadata.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 5, opt_noLazyCreate,
      null));
};


proto.core.solo.io.Metadata.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
};


/**
 * map<string, string> annotations = 6;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.core.solo.io.Metadata.prototype.getAnnotationsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 6, opt_noLazyCreate,
      null));
};


proto.core.solo.io.Metadata.prototype.clearAnnotationsMap = function() {
  this.getAnnotationsMap().clear();
};


/**
 * optional int64 generation = 8;
 * @return {number}
 */
proto.core.solo.io.Metadata.prototype.getGeneration = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 8, 0));
};


/** @param {number} value */
proto.core.solo.io.Metadata.prototype.setGeneration = function(value) {
  jspb.Message.setProto3IntField(this, 8, value);
};


/**
 * repeated OwnerReference owner_references = 9;
 * @return {!Array<!proto.core.solo.io.Metadata.OwnerReference>}
 */
proto.core.solo.io.Metadata.prototype.getOwnerReferencesList = function() {
  return /** @type{!Array<!proto.core.solo.io.Metadata.OwnerReference>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.core.solo.io.Metadata.OwnerReference, 9));
};


/** @param {!Array<!proto.core.solo.io.Metadata.OwnerReference>} value */
proto.core.solo.io.Metadata.prototype.setOwnerReferencesList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 9, value);
};


/**
 * @param {!proto.core.solo.io.Metadata.OwnerReference=} opt_value
 * @param {number=} opt_index
 * @return {!proto.core.solo.io.Metadata.OwnerReference}
 */
proto.core.solo.io.Metadata.prototype.addOwnerReferences = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 9, opt_value, proto.core.solo.io.Metadata.OwnerReference, opt_index);
};


proto.core.solo.io.Metadata.prototype.clearOwnerReferencesList = function() {
  this.setOwnerReferencesList([]);
};


goog.object.extend(exports, proto.core.solo.io);
