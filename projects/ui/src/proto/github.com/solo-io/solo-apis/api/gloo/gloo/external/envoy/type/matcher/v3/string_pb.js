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

var github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_type_matcher_v3_regex_pb = require('../../../../../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/type/matcher/v3/regex_pb.js');
var envoy_annotations_deprecation_pb = require('../../../../../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/annotations/deprecation_pb.js');
var udpa_annotations_status_pb = require('../../../../../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/external/udpa/annotations/status_pb.js');
var udpa_annotations_versioning_pb = require('../../../../../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/external/udpa/annotations/versioning_pb.js');
var validate_validate_pb = require('../../../../../../../../../../../validate/validate_pb.js');
var extproto_ext_pb = require('../../../../../../../../../../../extproto/ext_pb.js');
goog.exportSymbol('proto.solo.io.envoy.type.matcher.v3.ListStringMatcher', null, global);
goog.exportSymbol('proto.solo.io.envoy.type.matcher.v3.StringMatcher', null, global);

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
proto.solo.io.envoy.type.matcher.v3.StringMatcher = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.solo.io.envoy.type.matcher.v3.StringMatcher.oneofGroups_);
};
goog.inherits(proto.solo.io.envoy.type.matcher.v3.StringMatcher, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.solo.io.envoy.type.matcher.v3.StringMatcher.displayName = 'proto.solo.io.envoy.type.matcher.v3.StringMatcher';
}
/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.solo.io.envoy.type.matcher.v3.StringMatcher.oneofGroups_ = [[1,2,3,5]];

/**
 * @enum {number}
 */
proto.solo.io.envoy.type.matcher.v3.StringMatcher.MatchPatternCase = {
  MATCH_PATTERN_NOT_SET: 0,
  EXACT: 1,
  PREFIX: 2,
  SUFFIX: 3,
  SAFE_REGEX: 5
};

/**
 * @return {proto.solo.io.envoy.type.matcher.v3.StringMatcher.MatchPatternCase}
 */
proto.solo.io.envoy.type.matcher.v3.StringMatcher.prototype.getMatchPatternCase = function() {
  return /** @type {proto.solo.io.envoy.type.matcher.v3.StringMatcher.MatchPatternCase} */(jspb.Message.computeOneofCase(this, proto.solo.io.envoy.type.matcher.v3.StringMatcher.oneofGroups_[0]));
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
proto.solo.io.envoy.type.matcher.v3.StringMatcher.prototype.toObject = function(opt_includeInstance) {
  return proto.solo.io.envoy.type.matcher.v3.StringMatcher.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.solo.io.envoy.type.matcher.v3.StringMatcher} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.solo.io.envoy.type.matcher.v3.StringMatcher.toObject = function(includeInstance, msg) {
  var f, obj = {
    exact: jspb.Message.getFieldWithDefault(msg, 1, ""),
    prefix: jspb.Message.getFieldWithDefault(msg, 2, ""),
    suffix: jspb.Message.getFieldWithDefault(msg, 3, ""),
    safeRegex: (f = msg.getSafeRegex()) && github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_type_matcher_v3_regex_pb.RegexMatcher.toObject(includeInstance, f),
    ignoreCase: jspb.Message.getFieldWithDefault(msg, 6, false)
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
 * @return {!proto.solo.io.envoy.type.matcher.v3.StringMatcher}
 */
proto.solo.io.envoy.type.matcher.v3.StringMatcher.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.solo.io.envoy.type.matcher.v3.StringMatcher;
  return proto.solo.io.envoy.type.matcher.v3.StringMatcher.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.solo.io.envoy.type.matcher.v3.StringMatcher} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.solo.io.envoy.type.matcher.v3.StringMatcher}
 */
proto.solo.io.envoy.type.matcher.v3.StringMatcher.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setExact(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setPrefix(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setSuffix(value);
      break;
    case 5:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_type_matcher_v3_regex_pb.RegexMatcher;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_type_matcher_v3_regex_pb.RegexMatcher.deserializeBinaryFromReader);
      msg.setSafeRegex(value);
      break;
    case 6:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIgnoreCase(value);
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
proto.solo.io.envoy.type.matcher.v3.StringMatcher.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.solo.io.envoy.type.matcher.v3.StringMatcher.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.solo.io.envoy.type.matcher.v3.StringMatcher} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.solo.io.envoy.type.matcher.v3.StringMatcher.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {string} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeString(
      1,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeString(
      2,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 3));
  if (f != null) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getSafeRegex();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_type_matcher_v3_regex_pb.RegexMatcher.serializeBinaryToWriter
    );
  }
  f = message.getIgnoreCase();
  if (f) {
    writer.writeBool(
      6,
      f
    );
  }
};


/**
 * optional string exact = 1;
 * @return {string}
 */
proto.solo.io.envoy.type.matcher.v3.StringMatcher.prototype.getExact = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.solo.io.envoy.type.matcher.v3.StringMatcher.prototype.setExact = function(value) {
  jspb.Message.setOneofField(this, 1, proto.solo.io.envoy.type.matcher.v3.StringMatcher.oneofGroups_[0], value);
};


proto.solo.io.envoy.type.matcher.v3.StringMatcher.prototype.clearExact = function() {
  jspb.Message.setOneofField(this, 1, proto.solo.io.envoy.type.matcher.v3.StringMatcher.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.solo.io.envoy.type.matcher.v3.StringMatcher.prototype.hasExact = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional string prefix = 2;
 * @return {string}
 */
proto.solo.io.envoy.type.matcher.v3.StringMatcher.prototype.getPrefix = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.solo.io.envoy.type.matcher.v3.StringMatcher.prototype.setPrefix = function(value) {
  jspb.Message.setOneofField(this, 2, proto.solo.io.envoy.type.matcher.v3.StringMatcher.oneofGroups_[0], value);
};


proto.solo.io.envoy.type.matcher.v3.StringMatcher.prototype.clearPrefix = function() {
  jspb.Message.setOneofField(this, 2, proto.solo.io.envoy.type.matcher.v3.StringMatcher.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.solo.io.envoy.type.matcher.v3.StringMatcher.prototype.hasPrefix = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional string suffix = 3;
 * @return {string}
 */
proto.solo.io.envoy.type.matcher.v3.StringMatcher.prototype.getSuffix = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.solo.io.envoy.type.matcher.v3.StringMatcher.prototype.setSuffix = function(value) {
  jspb.Message.setOneofField(this, 3, proto.solo.io.envoy.type.matcher.v3.StringMatcher.oneofGroups_[0], value);
};


proto.solo.io.envoy.type.matcher.v3.StringMatcher.prototype.clearSuffix = function() {
  jspb.Message.setOneofField(this, 3, proto.solo.io.envoy.type.matcher.v3.StringMatcher.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.solo.io.envoy.type.matcher.v3.StringMatcher.prototype.hasSuffix = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional RegexMatcher safe_regex = 5;
 * @return {?proto.solo.io.envoy.type.matcher.v3.RegexMatcher}
 */
proto.solo.io.envoy.type.matcher.v3.StringMatcher.prototype.getSafeRegex = function() {
  return /** @type{?proto.solo.io.envoy.type.matcher.v3.RegexMatcher} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_type_matcher_v3_regex_pb.RegexMatcher, 5));
};


/** @param {?proto.solo.io.envoy.type.matcher.v3.RegexMatcher|undefined} value */
proto.solo.io.envoy.type.matcher.v3.StringMatcher.prototype.setSafeRegex = function(value) {
  jspb.Message.setOneofWrapperField(this, 5, proto.solo.io.envoy.type.matcher.v3.StringMatcher.oneofGroups_[0], value);
};


proto.solo.io.envoy.type.matcher.v3.StringMatcher.prototype.clearSafeRegex = function() {
  this.setSafeRegex(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.solo.io.envoy.type.matcher.v3.StringMatcher.prototype.hasSafeRegex = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional bool ignore_case = 6;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.solo.io.envoy.type.matcher.v3.StringMatcher.prototype.getIgnoreCase = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 6, false));
};


/** @param {boolean} value */
proto.solo.io.envoy.type.matcher.v3.StringMatcher.prototype.setIgnoreCase = function(value) {
  jspb.Message.setProto3BooleanField(this, 6, value);
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
proto.solo.io.envoy.type.matcher.v3.ListStringMatcher = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.solo.io.envoy.type.matcher.v3.ListStringMatcher.repeatedFields_, null);
};
goog.inherits(proto.solo.io.envoy.type.matcher.v3.ListStringMatcher, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.solo.io.envoy.type.matcher.v3.ListStringMatcher.displayName = 'proto.solo.io.envoy.type.matcher.v3.ListStringMatcher';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.solo.io.envoy.type.matcher.v3.ListStringMatcher.repeatedFields_ = [1];



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
proto.solo.io.envoy.type.matcher.v3.ListStringMatcher.prototype.toObject = function(opt_includeInstance) {
  return proto.solo.io.envoy.type.matcher.v3.ListStringMatcher.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.solo.io.envoy.type.matcher.v3.ListStringMatcher} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.solo.io.envoy.type.matcher.v3.ListStringMatcher.toObject = function(includeInstance, msg) {
  var f, obj = {
    patternsList: jspb.Message.toObjectList(msg.getPatternsList(),
    proto.solo.io.envoy.type.matcher.v3.StringMatcher.toObject, includeInstance)
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
 * @return {!proto.solo.io.envoy.type.matcher.v3.ListStringMatcher}
 */
proto.solo.io.envoy.type.matcher.v3.ListStringMatcher.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.solo.io.envoy.type.matcher.v3.ListStringMatcher;
  return proto.solo.io.envoy.type.matcher.v3.ListStringMatcher.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.solo.io.envoy.type.matcher.v3.ListStringMatcher} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.solo.io.envoy.type.matcher.v3.ListStringMatcher}
 */
proto.solo.io.envoy.type.matcher.v3.ListStringMatcher.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.solo.io.envoy.type.matcher.v3.StringMatcher;
      reader.readMessage(value,proto.solo.io.envoy.type.matcher.v3.StringMatcher.deserializeBinaryFromReader);
      msg.addPatterns(value);
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
proto.solo.io.envoy.type.matcher.v3.ListStringMatcher.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.solo.io.envoy.type.matcher.v3.ListStringMatcher.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.solo.io.envoy.type.matcher.v3.ListStringMatcher} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.solo.io.envoy.type.matcher.v3.ListStringMatcher.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPatternsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.solo.io.envoy.type.matcher.v3.StringMatcher.serializeBinaryToWriter
    );
  }
};


/**
 * repeated StringMatcher patterns = 1;
 * @return {!Array<!proto.solo.io.envoy.type.matcher.v3.StringMatcher>}
 */
proto.solo.io.envoy.type.matcher.v3.ListStringMatcher.prototype.getPatternsList = function() {
  return /** @type{!Array<!proto.solo.io.envoy.type.matcher.v3.StringMatcher>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.solo.io.envoy.type.matcher.v3.StringMatcher, 1));
};


/** @param {!Array<!proto.solo.io.envoy.type.matcher.v3.StringMatcher>} value */
proto.solo.io.envoy.type.matcher.v3.ListStringMatcher.prototype.setPatternsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.solo.io.envoy.type.matcher.v3.StringMatcher=} opt_value
 * @param {number=} opt_index
 * @return {!proto.solo.io.envoy.type.matcher.v3.StringMatcher}
 */
proto.solo.io.envoy.type.matcher.v3.ListStringMatcher.prototype.addPatterns = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.solo.io.envoy.type.matcher.v3.StringMatcher, opt_index);
};


proto.solo.io.envoy.type.matcher.v3.ListStringMatcher.prototype.clearPatternsList = function() {
  this.setPatternsList([]);
};


goog.object.extend(exports, proto.solo.io.envoy.type.matcher.v3);
