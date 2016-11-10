// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/genomics/v1/references.proto
// DO NOT EDIT!

package google_genomics_v1 // import "google.golang.org/genproto/googleapis/genomics/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A reference is a canonical assembled DNA sequence, intended to act as a
// reference coordinate space for other genomic annotations. A single reference
// might represent the human chromosome 1 or mitochandrial DNA, for instance. A
// reference belongs to one or more reference sets.
//
// For more genomics resource definitions, see [Fundamentals of Google
// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
type Reference struct {
	// The server-generated reference ID, unique across all references.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// The length of this reference's sequence.
	Length int64 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	// MD5 of the upper-case sequence excluding all whitespace characters (this
	// is equivalent to SQ:M5 in SAM). This value is represented in lower case
	// hexadecimal format.
	Md5Checksum string `protobuf:"bytes,3,opt,name=md5checksum" json:"md5checksum,omitempty"`
	// The name of this reference, for example `22`.
	Name string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	// The URI from which the sequence was obtained. Typically specifies a FASTA
	// format file.
	SourceUri string `protobuf:"bytes,5,opt,name=source_uri,json=sourceUri" json:"source_uri,omitempty"`
	// All known corresponding accession IDs in INSDC (GenBank/ENA/DDBJ) ideally
	// with a version number, for example `GCF_000001405.26`.
	SourceAccessions []string `protobuf:"bytes,6,rep,name=source_accessions,json=sourceAccessions" json:"source_accessions,omitempty"`
	// ID from http://www.ncbi.nlm.nih.gov/taxonomy. For example, 9606 for human.
	NcbiTaxonId int32 `protobuf:"varint,7,opt,name=ncbi_taxon_id,json=ncbiTaxonId" json:"ncbi_taxon_id,omitempty"`
}

func (m *Reference) Reset()                    { *m = Reference{} }
func (m *Reference) String() string            { return proto.CompactTextString(m) }
func (*Reference) ProtoMessage()               {}
func (*Reference) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

// A reference set is a set of references which typically comprise a reference
// assembly for a species, such as `GRCh38` which is representative
// of the human genome. A reference set defines a common coordinate space for
// comparing reference-aligned experimental data. A reference set contains 1 or
// more references.
//
// For more genomics resource definitions, see [Fundamentals of Google
// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
type ReferenceSet struct {
	// The server-generated reference set ID, unique across all reference sets.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// The IDs of the reference objects that are part of this set.
	// `Reference.md5checksum` must be unique within this set.
	ReferenceIds []string `protobuf:"bytes,2,rep,name=reference_ids,json=referenceIds" json:"reference_ids,omitempty"`
	// Order-independent MD5 checksum which identifies this reference set. The
	// checksum is computed by sorting all lower case hexidecimal string
	// `reference.md5checksum` (for all reference in this set) in
	// ascending lexicographic order, concatenating, and taking the MD5 of that
	// value. The resulting value is represented in lower case hexadecimal format.
	Md5Checksum string `protobuf:"bytes,3,opt,name=md5checksum" json:"md5checksum,omitempty"`
	// ID from http://www.ncbi.nlm.nih.gov/taxonomy (for example, 9606 for human)
	// indicating the species which this reference set is intended to model. Note
	// that contained references may specify a different `ncbiTaxonId`, as
	// assemblies may contain reference sequences which do not belong to the
	// modeled species, for example EBV in a human reference genome.
	NcbiTaxonId int32 `protobuf:"varint,4,opt,name=ncbi_taxon_id,json=ncbiTaxonId" json:"ncbi_taxon_id,omitempty"`
	// Free text description of this reference set.
	Description string `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	// Public id of this reference set, such as `GRCh37`.
	AssemblyId string `protobuf:"bytes,6,opt,name=assembly_id,json=assemblyId" json:"assembly_id,omitempty"`
	// The URI from which the references were obtained.
	SourceUri string `protobuf:"bytes,7,opt,name=source_uri,json=sourceUri" json:"source_uri,omitempty"`
	// All known corresponding accession IDs in INSDC (GenBank/ENA/DDBJ) ideally
	// with a version number, for example `NC_000001.11`.
	SourceAccessions []string `protobuf:"bytes,8,rep,name=source_accessions,json=sourceAccessions" json:"source_accessions,omitempty"`
}

func (m *ReferenceSet) Reset()                    { *m = ReferenceSet{} }
func (m *ReferenceSet) String() string            { return proto.CompactTextString(m) }
func (*ReferenceSet) ProtoMessage()               {}
func (*ReferenceSet) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

type SearchReferenceSetsRequest struct {
	// If present, return reference sets for which the
	// [md5checksum][google.genomics.v1.ReferenceSet.md5checksum] matches exactly.
	Md5Checksums []string `protobuf:"bytes,1,rep,name=md5checksums" json:"md5checksums,omitempty"`
	// If present, return reference sets for which a prefix of any of
	// [sourceAccessions][google.genomics.v1.ReferenceSet.source_accessions]
	// match any of these strings. Accession numbers typically have a main number
	// and a version, for example `NC_000001.11`.
	Accessions []string `protobuf:"bytes,2,rep,name=accessions" json:"accessions,omitempty"`
	// If present, return reference sets for which a substring of their
	// `assemblyId` matches this string (case insensitive).
	AssemblyId string `protobuf:"bytes,3,opt,name=assembly_id,json=assemblyId" json:"assembly_id,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// To get the next page of results, set this parameter to the value of
	// `nextPageToken` from the previous response.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	// The maximum number of results to return in a single page. If unspecified,
	// defaults to 1024. The maximum value is 4096.
	PageSize int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *SearchReferenceSetsRequest) Reset()                    { *m = SearchReferenceSetsRequest{} }
func (m *SearchReferenceSetsRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchReferenceSetsRequest) ProtoMessage()               {}
func (*SearchReferenceSetsRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

type SearchReferenceSetsResponse struct {
	// The matching references sets.
	ReferenceSets []*ReferenceSet `protobuf:"bytes,1,rep,name=reference_sets,json=referenceSets" json:"reference_sets,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// Provide this value in a subsequent request to return the next page of
	// results. This field will be empty if there aren't any additional results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *SearchReferenceSetsResponse) Reset()                    { *m = SearchReferenceSetsResponse{} }
func (m *SearchReferenceSetsResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchReferenceSetsResponse) ProtoMessage()               {}
func (*SearchReferenceSetsResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

func (m *SearchReferenceSetsResponse) GetReferenceSets() []*ReferenceSet {
	if m != nil {
		return m.ReferenceSets
	}
	return nil
}

type GetReferenceSetRequest struct {
	// The ID of the reference set.
	ReferenceSetId string `protobuf:"bytes,1,opt,name=reference_set_id,json=referenceSetId" json:"reference_set_id,omitempty"`
}

func (m *GetReferenceSetRequest) Reset()                    { *m = GetReferenceSetRequest{} }
func (m *GetReferenceSetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetReferenceSetRequest) ProtoMessage()               {}
func (*GetReferenceSetRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

type SearchReferencesRequest struct {
	// If present, return references for which the
	// [md5checksum][google.genomics.v1.Reference.md5checksum] matches exactly.
	Md5Checksums []string `protobuf:"bytes,1,rep,name=md5checksums" json:"md5checksums,omitempty"`
	// If present, return references for which a prefix of any of
	// [sourceAccessions][google.genomics.v1.Reference.source_accessions] match
	// any of these strings. Accession numbers typically have a main number and a
	// version, for example `GCF_000001405.26`.
	Accessions []string `protobuf:"bytes,2,rep,name=accessions" json:"accessions,omitempty"`
	// If present, return only references which belong to this reference set.
	ReferenceSetId string `protobuf:"bytes,3,opt,name=reference_set_id,json=referenceSetId" json:"reference_set_id,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// To get the next page of results, set this parameter to the value of
	// `nextPageToken` from the previous response.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	// The maximum number of results to return in a single page. If unspecified,
	// defaults to 1024. The maximum value is 4096.
	PageSize int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *SearchReferencesRequest) Reset()                    { *m = SearchReferencesRequest{} }
func (m *SearchReferencesRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchReferencesRequest) ProtoMessage()               {}
func (*SearchReferencesRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{5} }

type SearchReferencesResponse struct {
	// The matching references.
	References []*Reference `protobuf:"bytes,1,rep,name=references" json:"references,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// Provide this value in a subsequent request to return the next page of
	// results. This field will be empty if there aren't any additional results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *SearchReferencesResponse) Reset()                    { *m = SearchReferencesResponse{} }
func (m *SearchReferencesResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchReferencesResponse) ProtoMessage()               {}
func (*SearchReferencesResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{6} }

func (m *SearchReferencesResponse) GetReferences() []*Reference {
	if m != nil {
		return m.References
	}
	return nil
}

type GetReferenceRequest struct {
	// The ID of the reference.
	ReferenceId string `protobuf:"bytes,1,opt,name=reference_id,json=referenceId" json:"reference_id,omitempty"`
}

func (m *GetReferenceRequest) Reset()                    { *m = GetReferenceRequest{} }
func (m *GetReferenceRequest) String() string            { return proto.CompactTextString(m) }
func (*GetReferenceRequest) ProtoMessage()               {}
func (*GetReferenceRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{7} }

type ListBasesRequest struct {
	// The ID of the reference.
	ReferenceId string `protobuf:"bytes,1,opt,name=reference_id,json=referenceId" json:"reference_id,omitempty"`
	// The start position (0-based) of this query. Defaults to 0.
	Start int64 `protobuf:"varint,2,opt,name=start" json:"start,omitempty"`
	// The end position (0-based, exclusive) of this query. Defaults to the length
	// of this reference.
	End int64 `protobuf:"varint,3,opt,name=end" json:"end,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// To get the next page of results, set this parameter to the value of
	// `nextPageToken` from the previous response.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	// The maximum number of bases to return in a single page. If unspecified,
	// defaults to 200Kbp (kilo base pairs). The maximum value is 10Mbp (mega base
	// pairs).
	PageSize int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *ListBasesRequest) Reset()                    { *m = ListBasesRequest{} }
func (m *ListBasesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListBasesRequest) ProtoMessage()               {}
func (*ListBasesRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{8} }

type ListBasesResponse struct {
	// The offset position (0-based) of the given `sequence` from the
	// start of this `Reference`. This value will differ for each page
	// in a paginated request.
	Offset int64 `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	// A substring of the bases that make up this reference.
	Sequence string `protobuf:"bytes,2,opt,name=sequence" json:"sequence,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// Provide this value in a subsequent request to return the next page of
	// results. This field will be empty if there aren't any additional results.
	NextPageToken string `protobuf:"bytes,3,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListBasesResponse) Reset()                    { *m = ListBasesResponse{} }
func (m *ListBasesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListBasesResponse) ProtoMessage()               {}
func (*ListBasesResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{9} }

func init() {
	proto.RegisterType((*Reference)(nil), "google.genomics.v1.Reference")
	proto.RegisterType((*ReferenceSet)(nil), "google.genomics.v1.ReferenceSet")
	proto.RegisterType((*SearchReferenceSetsRequest)(nil), "google.genomics.v1.SearchReferenceSetsRequest")
	proto.RegisterType((*SearchReferenceSetsResponse)(nil), "google.genomics.v1.SearchReferenceSetsResponse")
	proto.RegisterType((*GetReferenceSetRequest)(nil), "google.genomics.v1.GetReferenceSetRequest")
	proto.RegisterType((*SearchReferencesRequest)(nil), "google.genomics.v1.SearchReferencesRequest")
	proto.RegisterType((*SearchReferencesResponse)(nil), "google.genomics.v1.SearchReferencesResponse")
	proto.RegisterType((*GetReferenceRequest)(nil), "google.genomics.v1.GetReferenceRequest")
	proto.RegisterType((*ListBasesRequest)(nil), "google.genomics.v1.ListBasesRequest")
	proto.RegisterType((*ListBasesResponse)(nil), "google.genomics.v1.ListBasesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ReferenceServiceV1 service

type ReferenceServiceV1Client interface {
	// Searches for reference sets which match the given criteria.
	//
	// For the definitions of references and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	//
	// Implements
	// [GlobalAllianceApi.searchReferenceSets](https://github.com/ga4gh/schemas/blob/v0.5.1/src/main/resources/avro/referencemethods.avdl#L71)
	SearchReferenceSets(ctx context.Context, in *SearchReferenceSetsRequest, opts ...grpc.CallOption) (*SearchReferenceSetsResponse, error)
	// Gets a reference set.
	//
	// For the definitions of references and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	//
	// Implements
	// [GlobalAllianceApi.getReferenceSet](https://github.com/ga4gh/schemas/blob/v0.5.1/src/main/resources/avro/referencemethods.avdl#L83).
	GetReferenceSet(ctx context.Context, in *GetReferenceSetRequest, opts ...grpc.CallOption) (*ReferenceSet, error)
	// Searches for references which match the given criteria.
	//
	// For the definitions of references and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	//
	// Implements
	// [GlobalAllianceApi.searchReferences](https://github.com/ga4gh/schemas/blob/v0.5.1/src/main/resources/avro/referencemethods.avdl#L146).
	SearchReferences(ctx context.Context, in *SearchReferencesRequest, opts ...grpc.CallOption) (*SearchReferencesResponse, error)
	// Gets a reference.
	//
	// For the definitions of references and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	//
	// Implements
	// [GlobalAllianceApi.getReference](https://github.com/ga4gh/schemas/blob/v0.5.1/src/main/resources/avro/referencemethods.avdl#L158).
	GetReference(ctx context.Context, in *GetReferenceRequest, opts ...grpc.CallOption) (*Reference, error)
	// Lists the bases in a reference, optionally restricted to a range.
	//
	// For the definitions of references and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	//
	// Implements
	// [GlobalAllianceApi.getReferenceBases](https://github.com/ga4gh/schemas/blob/v0.5.1/src/main/resources/avro/referencemethods.avdl#L221).
	ListBases(ctx context.Context, in *ListBasesRequest, opts ...grpc.CallOption) (*ListBasesResponse, error)
}

type referenceServiceV1Client struct {
	cc *grpc.ClientConn
}

func NewReferenceServiceV1Client(cc *grpc.ClientConn) ReferenceServiceV1Client {
	return &referenceServiceV1Client{cc}
}

func (c *referenceServiceV1Client) SearchReferenceSets(ctx context.Context, in *SearchReferenceSetsRequest, opts ...grpc.CallOption) (*SearchReferenceSetsResponse, error) {
	out := new(SearchReferenceSetsResponse)
	err := grpc.Invoke(ctx, "/google.genomics.v1.ReferenceServiceV1/SearchReferenceSets", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *referenceServiceV1Client) GetReferenceSet(ctx context.Context, in *GetReferenceSetRequest, opts ...grpc.CallOption) (*ReferenceSet, error) {
	out := new(ReferenceSet)
	err := grpc.Invoke(ctx, "/google.genomics.v1.ReferenceServiceV1/GetReferenceSet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *referenceServiceV1Client) SearchReferences(ctx context.Context, in *SearchReferencesRequest, opts ...grpc.CallOption) (*SearchReferencesResponse, error) {
	out := new(SearchReferencesResponse)
	err := grpc.Invoke(ctx, "/google.genomics.v1.ReferenceServiceV1/SearchReferences", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *referenceServiceV1Client) GetReference(ctx context.Context, in *GetReferenceRequest, opts ...grpc.CallOption) (*Reference, error) {
	out := new(Reference)
	err := grpc.Invoke(ctx, "/google.genomics.v1.ReferenceServiceV1/GetReference", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *referenceServiceV1Client) ListBases(ctx context.Context, in *ListBasesRequest, opts ...grpc.CallOption) (*ListBasesResponse, error) {
	out := new(ListBasesResponse)
	err := grpc.Invoke(ctx, "/google.genomics.v1.ReferenceServiceV1/ListBases", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ReferenceServiceV1 service

type ReferenceServiceV1Server interface {
	// Searches for reference sets which match the given criteria.
	//
	// For the definitions of references and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	//
	// Implements
	// [GlobalAllianceApi.searchReferenceSets](https://github.com/ga4gh/schemas/blob/v0.5.1/src/main/resources/avro/referencemethods.avdl#L71)
	SearchReferenceSets(context.Context, *SearchReferenceSetsRequest) (*SearchReferenceSetsResponse, error)
	// Gets a reference set.
	//
	// For the definitions of references and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	//
	// Implements
	// [GlobalAllianceApi.getReferenceSet](https://github.com/ga4gh/schemas/blob/v0.5.1/src/main/resources/avro/referencemethods.avdl#L83).
	GetReferenceSet(context.Context, *GetReferenceSetRequest) (*ReferenceSet, error)
	// Searches for references which match the given criteria.
	//
	// For the definitions of references and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	//
	// Implements
	// [GlobalAllianceApi.searchReferences](https://github.com/ga4gh/schemas/blob/v0.5.1/src/main/resources/avro/referencemethods.avdl#L146).
	SearchReferences(context.Context, *SearchReferencesRequest) (*SearchReferencesResponse, error)
	// Gets a reference.
	//
	// For the definitions of references and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	//
	// Implements
	// [GlobalAllianceApi.getReference](https://github.com/ga4gh/schemas/blob/v0.5.1/src/main/resources/avro/referencemethods.avdl#L158).
	GetReference(context.Context, *GetReferenceRequest) (*Reference, error)
	// Lists the bases in a reference, optionally restricted to a range.
	//
	// For the definitions of references and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	//
	// Implements
	// [GlobalAllianceApi.getReferenceBases](https://github.com/ga4gh/schemas/blob/v0.5.1/src/main/resources/avro/referencemethods.avdl#L221).
	ListBases(context.Context, *ListBasesRequest) (*ListBasesResponse, error)
}

func RegisterReferenceServiceV1Server(s *grpc.Server, srv ReferenceServiceV1Server) {
	s.RegisterService(&_ReferenceServiceV1_serviceDesc, srv)
}

func _ReferenceServiceV1_SearchReferenceSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchReferenceSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReferenceServiceV1Server).SearchReferenceSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.genomics.v1.ReferenceServiceV1/SearchReferenceSets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReferenceServiceV1Server).SearchReferenceSets(ctx, req.(*SearchReferenceSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReferenceServiceV1_GetReferenceSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReferenceSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReferenceServiceV1Server).GetReferenceSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.genomics.v1.ReferenceServiceV1/GetReferenceSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReferenceServiceV1Server).GetReferenceSet(ctx, req.(*GetReferenceSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReferenceServiceV1_SearchReferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchReferencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReferenceServiceV1Server).SearchReferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.genomics.v1.ReferenceServiceV1/SearchReferences",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReferenceServiceV1Server).SearchReferences(ctx, req.(*SearchReferencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReferenceServiceV1_GetReference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReferenceServiceV1Server).GetReference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.genomics.v1.ReferenceServiceV1/GetReference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReferenceServiceV1Server).GetReference(ctx, req.(*GetReferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReferenceServiceV1_ListBases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReferenceServiceV1Server).ListBases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.genomics.v1.ReferenceServiceV1/ListBases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReferenceServiceV1Server).ListBases(ctx, req.(*ListBasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReferenceServiceV1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.genomics.v1.ReferenceServiceV1",
	HandlerType: (*ReferenceServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchReferenceSets",
			Handler:    _ReferenceServiceV1_SearchReferenceSets_Handler,
		},
		{
			MethodName: "GetReferenceSet",
			Handler:    _ReferenceServiceV1_GetReferenceSet_Handler,
		},
		{
			MethodName: "SearchReferences",
			Handler:    _ReferenceServiceV1_SearchReferences_Handler,
		},
		{
			MethodName: "GetReference",
			Handler:    _ReferenceServiceV1_GetReference_Handler,
		},
		{
			MethodName: "ListBases",
			Handler:    _ReferenceServiceV1_ListBases_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google.golang.org/genproto/googleapis/genomics/v1/references.proto",
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/genomics/v1/references.proto", fileDescriptor10)
}

var fileDescriptor10 = []byte{
	// 847 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x41, 0x6f, 0x1b, 0x45,
	0x14, 0xd6, 0x78, 0x63, 0x37, 0x7e, 0x76, 0x12, 0xf7, 0x15, 0xc2, 0xca, 0x55, 0xa8, 0xd9, 0x34,
	0xc5, 0x6a, 0x2a, 0xaf, 0x52, 0x84, 0x84, 0x90, 0x38, 0xe0, 0x4b, 0x65, 0x89, 0x43, 0xb4, 0x29,
	0x5c, 0xad, 0xcd, 0xee, 0x64, 0x33, 0x6a, 0x3c, 0x63, 0x76, 0x26, 0x51, 0x69, 0x95, 0x03, 0x48,
	0x1c, 0x81, 0x03, 0x17, 0x10, 0xbf, 0x85, 0x13, 0x3f, 0x81, 0x13, 0xe2, 0xca, 0x8f, 0xe0, 0x88,
	0x66, 0x76, 0x6c, 0x4f, 0xd6, 0x4b, 0x6c, 0xa9, 0xbd, 0x79, 0xbe, 0x79, 0xf3, 0xde, 0xf7, 0x7d,
	0x6f, 0xde, 0x78, 0x61, 0x98, 0x09, 0x91, 0x5d, 0xd0, 0x41, 0x26, 0x2e, 0x62, 0x9e, 0x0d, 0x44,
	0x9e, 0x85, 0x19, 0xe5, 0xd3, 0x5c, 0x28, 0x11, 0x16, 0x5b, 0xf1, 0x94, 0x49, 0x8d, 0x89, 0x09,
	0x4b, 0x64, 0x78, 0x75, 0x14, 0xe6, 0xf4, 0x8c, 0xe6, 0x94, 0x27, 0x54, 0x0e, 0x4c, 0x1c, 0xe2,
	0x2c, 0x87, 0x0d, 0x1a, 0x5c, 0x1d, 0x75, 0x47, 0xeb, 0xe5, 0x8d, 0xa7, 0x2c, 0x94, 0x34, 0xbf,
	0x62, 0x09, 0x4d, 0x04, 0x3f, 0x63, 0x59, 0x18, 0x73, 0x2e, 0x54, 0xac, 0x98, 0xe0, 0x36, 0x7d,
	0xf0, 0x17, 0x81, 0x66, 0x34, 0xab, 0x89, 0xdb, 0x50, 0x63, 0xa9, 0x4f, 0x7a, 0xa4, 0xdf, 0x8c,
	0x6a, 0x2c, 0xc5, 0x5d, 0x68, 0x5c, 0x50, 0x9e, 0xa9, 0x73, 0xbf, 0xd6, 0x23, 0x7d, 0x2f, 0xb2,
	0x2b, 0xec, 0x41, 0x6b, 0x92, 0x7e, 0x9c, 0x9c, 0xd3, 0xe4, 0x85, 0xbc, 0x9c, 0xf8, 0x9e, 0x39,
	0xe0, 0x42, 0x88, 0xb0, 0xc1, 0xe3, 0x09, 0xf5, 0x37, 0xcc, 0x96, 0xf9, 0x8d, 0x7b, 0x00, 0x52,
	0x5c, 0xe6, 0x09, 0x1d, 0x5f, 0xe6, 0xcc, 0xaf, 0x9b, 0x9d, 0x66, 0x81, 0x7c, 0x99, 0x33, 0x3c,
	0x84, 0xbb, 0x76, 0x3b, 0x4e, 0x12, 0x2a, 0xa5, 0x66, 0xe9, 0x37, 0x7a, 0x5e, 0xbf, 0x19, 0x75,
	0x8a, 0x8d, 0xcf, 0xe7, 0x38, 0x06, 0xb0, 0xc5, 0x93, 0x53, 0x36, 0x56, 0xf1, 0x4b, 0xc1, 0xc7,
	0x2c, 0xf5, 0xef, 0xf4, 0x48, 0xbf, 0x1e, 0xb5, 0x34, 0xf8, 0x5c, 0x63, 0xa3, 0x34, 0xf8, 0xa5,
	0x06, 0xed, 0xb9, 0xb6, 0x13, 0xaa, 0x96, 0xe4, 0xed, 0xc3, 0xd6, 0xdc, 0xef, 0x31, 0x4b, 0xa5,
	0x5f, 0x33, 0xd5, 0xda, 0x73, 0x70, 0x94, 0xca, 0x35, 0xb4, 0x2e, 0x71, 0xd9, 0x58, 0xe2, 0xa2,
	0xb3, 0xa4, 0x54, 0x26, 0x39, 0x9b, 0x6a, 0xf7, 0xad, 0x78, 0x17, 0xc2, 0x07, 0xd0, 0x8a, 0xa5,
	0xa4, 0x93, 0xd3, 0x8b, 0x6f, 0x74, 0x8e, 0x86, 0x89, 0x80, 0x19, 0x34, 0x4a, 0x4b, 0xf6, 0xdd,
	0x59, 0xcb, 0xbe, 0xcd, 0x6a, 0xfb, 0x82, 0xdf, 0x09, 0x74, 0x4f, 0x68, 0x9c, 0x27, 0xe7, 0xae,
	0x41, 0x32, 0xa2, 0x5f, 0x5f, 0x52, 0xa9, 0x30, 0x80, 0xb6, 0x23, 0x50, 0xfa, 0xa4, 0xf0, 0xc5,
	0xc5, 0xf0, 0x7d, 0x00, 0xa7, 0x50, 0xe1, 0x9c, 0x83, 0x94, 0xf5, 0x78, 0x55, 0x7a, 0xa6, 0x71,
	0x46, 0xc7, 0x4a, 0xbc, 0xa0, 0xdc, 0x5e, 0x94, 0xa6, 0x46, 0x9e, 0x6b, 0x00, 0xef, 0x83, 0x59,
	0x8c, 0x25, 0x7b, 0x45, 0x8d, 0x5f, 0xf5, 0x68, 0x53, 0x03, 0x27, 0xec, 0x15, 0x0d, 0x7e, 0x24,
	0x70, 0xbf, 0x92, 0xbf, 0x9c, 0x0a, 0x2e, 0x29, 0x3e, 0x83, 0xed, 0x45, 0x67, 0x25, 0x55, 0x85,
	0x84, 0xd6, 0xd3, 0xde, 0x60, 0x79, 0x9c, 0x06, 0x6e, 0x8a, 0x68, 0x71, 0x23, 0x74, 0x42, 0x7c,
	0x04, 0x3b, 0x9c, 0xbe, 0x54, 0x63, 0x87, 0x69, 0xcd, 0x30, 0xdd, 0xd2, 0xf0, 0xf1, 0x8c, 0x6d,
	0x30, 0x84, 0xdd, 0x67, 0x54, 0xdd, 0xc8, 0x64, 0xbd, 0xec, 0x43, 0xe7, 0x06, 0x95, 0xf1, 0xfc,
	0x0a, 0x6e, 0xbb, 0xa5, 0x46, 0x69, 0xf0, 0x07, 0x81, 0xf7, 0x4a, 0xa2, 0xde, 0x6a, 0x47, 0xaa,
	0x98, 0x78, 0x55, 0x4c, 0xde, 0xa8, 0x35, 0xdf, 0x12, 0xf0, 0x97, 0x55, 0xd8, 0xbe, 0x7c, 0x06,
	0xb0, 0x78, 0xe1, 0x6c, 0x4f, 0xf6, 0x6e, 0xed, 0x49, 0xe4, 0x1c, 0x58, 0xbb, 0x1b, 0x9f, 0xc0,
	0x3d, 0xb7, 0x1b, 0x33, 0x13, 0x3f, 0x80, 0xb6, 0x3b, 0xef, 0xb6, 0x0d, 0x2d, 0x67, 0xdc, 0x83,
	0x5f, 0x09, 0x74, 0xbe, 0x60, 0x52, 0x0d, 0x63, 0xb9, 0x30, 0x7f, 0xf5, 0x39, 0x7c, 0x07, 0xea,
	0x52, 0xc5, 0xb9, 0xb2, 0x0f, 0x65, 0xb1, 0xc0, 0x0e, 0x78, 0x94, 0x17, 0x26, 0x7b, 0x91, 0xfe,
	0xf9, 0x46, 0xce, 0x0a, 0xb8, 0xeb, 0x50, 0xb3, 0x8e, 0xee, 0x42, 0x43, 0x9c, 0x9d, 0x49, 0xaa,
	0x0c, 0x2b, 0x2f, 0xb2, 0x2b, 0xec, 0xc2, 0xa6, 0xd4, 0xf4, 0x79, 0x42, 0xad, 0x47, 0xf3, 0x75,
	0x95, 0x8d, 0x5e, 0x85, 0x8d, 0x4f, 0xff, 0xae, 0x03, 0x3a, 0x57, 0xda, 0xfc, 0x93, 0x7c, 0x75,
	0x84, 0xbf, 0x11, 0xb8, 0x57, 0x31, 0x7c, 0x38, 0xa8, 0x6a, 0xe4, 0xff, 0xbf, 0x32, 0xdd, 0x70,
	0xed, 0xf8, 0x42, 0x6b, 0xb0, 0xff, 0xdd, 0x9f, 0xff, 0xfc, 0x5c, 0xdb, 0x0b, 0xfc, 0x9b, 0xff,
	0x94, 0x54, 0xc9, 0x50, 0x9a, 0x63, 0x9f, 0x92, 0xc7, 0xf8, 0x03, 0x81, 0x9d, 0xd2, 0x28, 0xe2,
	0xe3, 0xaa, 0x4a, 0xd5, 0xf3, 0xda, 0x5d, 0xf9, 0x44, 0x04, 0x4f, 0x0c, 0x8d, 0x47, 0xf8, 0x70,
	0x99, 0xc6, 0xeb, 0xf2, 0x80, 0x5d, 0xe3, 0x4f, 0x04, 0x3a, 0xe5, 0x79, 0xc0, 0xc3, 0x35, 0xa4,
	0xcf, 0x7d, 0x7a, 0xb2, 0x5e, 0xb0, 0x35, 0xa9, 0x67, 0xd8, 0x75, 0x83, 0x77, 0x6f, 0xb2, 0x73,
	0x1c, 0xba, 0x86, 0xb6, 0xab, 0x1d, 0x3f, 0x5c, 0xe5, 0xce, 0x8c, 0xc8, 0xed, 0x93, 0x1a, 0x1c,
	0x98, 0xca, 0x0f, 0x70, 0xaf, 0x54, 0xf9, 0xb5, 0x3b, 0x3c, 0xd7, 0xf8, 0x3d, 0x81, 0xe6, 0xfc,
	0x1e, 0xe3, 0xc3, 0xaa, 0x9c, 0xe5, 0x09, 0xec, 0x1e, 0xac, 0x88, 0xb2, 0xda, 0x0f, 0x0d, 0x83,
	0x03, 0xdc, 0xbf, 0x95, 0x41, 0x78, 0xaa, 0x0f, 0x0d, 0x07, 0xb0, 0x9b, 0x88, 0x49, 0x45, 0xe2,
	0xe1, 0xce, 0xc2, 0xd6, 0x63, 0xfd, 0x95, 0x74, 0x4c, 0xfe, 0x25, 0xe4, 0xb4, 0x61, 0xbe, 0x98,
	0x3e, 0xfa, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xae, 0x49, 0x82, 0x46, 0xd6, 0x09, 0x00, 0x00,
}
