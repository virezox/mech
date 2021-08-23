# August 23 2021

we need to call `proto.Marshal`:

https://pkg.go.dev/google.golang.org/protobuf/proto#Marshal

to do that, we need any type that has a `ProtoReflect` method:

https://pkg.go.dev/google.golang.org/protobuf/reflect/protoreflect#ProtoMessage

what types have a `ProtoReflect` method?

https://github.com/protocolbuffers/protobuf-go

~~~
types\descriptorpb\descriptor.pb.go
496:func (x *FileDescriptorSet) ProtoReflect() protoreflect.Message {
566:func (x *FileDescriptorProto) ProtoReflect() protoreflect.Message {
702:func (x *DescriptorProto) ProtoReflect() protoreflect.Message {
814:func (x *ExtensionRangeOptions) ProtoReflect() protoreflect.Message {
913:func (x *FieldDescriptorProto) ProtoReflect() protoreflect.Message {
1032:func (x *OneofDescriptorProto) ProtoReflect() protoreflect.Message {
1096:func (x *EnumDescriptorProto) ProtoReflect() protoreflect.Message {
1174:func (x *EnumValueDescriptorProto) ProtoReflect() protoreflect.Message {
1238:func (x *ServiceDescriptorProto) ProtoReflect() protoreflect.Message {
1315:func (x *MethodDescriptorProto) ProtoReflect() protoreflect.Message {
1496:func (x *FileOptions) ProtoReflect() protoreflect.Message {
1743:func (x *MessageOptions) ProtoReflect() protoreflect.Message {
1888:func (x *FieldOptions) ProtoReflect() protoreflect.Message {
1979:func (x *OneofOptions) ProtoReflect() protoreflect.Message {
2041:func (x *EnumOptions) ProtoReflect() protoreflect.Message {
2114:func (x *EnumValueOptions) ProtoReflect() protoreflect.Message {
2180:func (x *ServiceOptions) ProtoReflect() protoreflect.Message {
2248:func (x *MethodOptions) ProtoReflect() protoreflect.Message {
2323:func (x *UninterpretedOption) ProtoReflect() protoreflect.Message {
2457:func (x *SourceCodeInfo) ProtoReflect() protoreflect.Message {
2509:func (x *GeneratedCodeInfo) ProtoReflect() protoreflect.Message {
2558:func (x *DescriptorProto_ExtensionRange) ProtoReflect() protoreflect.Message {
2623:func (x *DescriptorProto_ReservedRange) ProtoReflect() protoreflect.Message {
2684:func (x *EnumDescriptorProto_EnumReservedRange) ProtoReflect() protoreflect.Message {
2744:func (x *UninterpretedOption_NamePart) ProtoReflect() protoreflect.Message {
2877:func (x *SourceCodeInfo_Location) ProtoReflect() protoreflect.Message {
2963:func (x *GeneratedCodeInfo_Annotation) ProtoReflect() protoreflect.Message {

types\dynamicpb\dynamic.go
91:func (m *Message) ProtoReflect() pref.Message {

types\known\anypb\any.pb.go
386:func (x *Any) ProtoReflect() protoreflect.Message {

types\known\apipb\api.pb.go
112:func (x *Api) ProtoReflect() protoreflect.Message {
215:func (x *Method) ProtoReflect() protoreflect.Message {
386:func (x *Mixin) ProtoReflect() protoreflect.Message {

types\known\durationpb\duration.pb.go
266:func (x *Duration) ProtoReflect() protoreflect.Message {

types\known\emptypb\empty.pb.go
73:func (x *Empty) ProtoReflect() protoreflect.Message {

types\known\fieldmaskpb\field_mask.pb.go
486:func (x *FieldMask) ProtoReflect() protoreflect.Message {

types\known\sourcecontextpb\source_context.pb.go
70:func (x *SourceContext) ProtoReflect() protoreflect.Message {

types\known\structpb\struct.pb.go
251:func (x *Struct) ProtoReflect() protoreflect.Message {
458:func (x *Value) ProtoReflect() protoreflect.Message {
629:func (x *ListValue) ProtoReflect() protoreflect.Message {

types\known\timestamppb\timestamp.pb.go
277:func (x *Timestamp) ProtoReflect() protoreflect.Message {

types\known\typepb\type.pb.go
303:func (x *Type) ProtoReflect() protoreflect.Message {
407:func (x *Field) ProtoReflect() protoreflect.Message {
527:func (x *Enum) ProtoReflect() protoreflect.Message {
608:func (x *EnumValue) ProtoReflect() protoreflect.Message {
680:func (x *Option) ProtoReflect() protoreflect.Message {

types\known\wrapperspb\wrappers.pb.go
85:func (x *DoubleValue) ProtoReflect() protoreflect.Message {
141:func (x *FloatValue) ProtoReflect() protoreflect.Message {
197:func (x *Int64Value) ProtoReflect() protoreflect.Message {
253:func (x *UInt64Value) ProtoReflect() protoreflect.Message {
309:func (x *Int32Value) ProtoReflect() protoreflect.Message {
365:func (x *UInt32Value) ProtoReflect() protoreflect.Message {
421:func (x *BoolValue) ProtoReflect() protoreflect.Message {
477:func (x *StringValue) ProtoReflect() protoreflect.Message {
533:func (x *BytesValue) ProtoReflect() protoreflect.Message {

types\pluginpb\plugin.pb.go
146:func (x *Version) ProtoReflect() protoreflect.Message {
237:func (x *CodeGeneratorRequest) ProtoReflect() protoreflect.Message {
318:func (x *CodeGeneratorResponse) ProtoReflect() protoreflect.Message {
435:func (x *CodeGeneratorResponse_File) ProtoReflect() protoreflect.Message {
~~~

--------------------------------------------------------------------------------

~~~
next
request
continuationCommand

next
response
commentThreadRenderer commentRenderer -language:json
~~~

YouTube API for comments?

https://www.youtube.com/watch?v=q5UnT4Ik6KU

So we can search for a comment that identifies the end song.

~~~
The unbiased journalism we need
~~~

- https://github.com/DeDiS/protobuf
- https://github.com/golang/go/blob/master/src/runtime/pprof/protobuf.go
- https://github.com/iv-org/invidious/blob/master/src/invidious/comments.cr
- https://github.com/user234683/youtube-local/blob/master/youtube/comments.py
- https://github.com/user234683/youtube-local/blob/master/youtube/proto.py
- https://godocs.io/google.golang.org/protobuf/encoding/protowire

## size difference

next:

~~~
Content-Encoding: gzip
Content-Length: 34976
~~~

watch:

~~~
Content-Encoding: gzip
79,940 bytes
~~~

make sure "next next" actually works before deleting proto stuff.

## clients

Check different clients
