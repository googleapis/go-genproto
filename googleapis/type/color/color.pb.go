// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/type/color/color.proto
// DO NOT EDIT!

/*
Package google_type is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/genproto/googleapis/type/color/color.proto

It has these top-level messages:
	Color
*/
package google_type // import "google.golang.org/genproto/googleapis/type/color"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Represents a color in the RGBA color space. This representation is designed
// for simplicity of conversion to/from color representations in various
// languages over compactness; for example, the fields of this representation
// can be trivially provided to the constructor of "java.awt.Color" in Java; it
// can also be trivially provided to UIColor's "+colorWithRed:green:blue:alpha"
// method in iOS; and, with just a little work, it can be easily formatted into
// a CSS "rgba()" string in JavaScript, as well. Here are some examples:
//
// Example (Java):
//
//      import com.google.type.Color;
//
//      // ...
//      public static java.awt.Color fromProto(Color protocolor) {
//        float alpha = protocolor.hasAlpha()
//            ? protocolor.getAlpha().getValue()
//            : 1.0;
//
//        return new java.awt.Color(
//            protocolor.getRed(),
//            protocolor.getGreen(),
//            protocolor.getBlue(),
//            alpha);
//      }
//
//      public static Color toProto(java.awt.Color color) {
//        float red = (float) color.getRed();
//        float green = (float) color.getGreen();
//        float blue = (float) color.getBlue();
//        float denominator = 255.0;
//        Color.Builder resultBuilder =
//            Color
//                .newBuilder()
//                .setRed(red / denominator)
//                .setGreen(green / denominator)
//                .setBlue(blue / denominator);
//        int alpha = color.getAlpha();
//        if (alpha != 255) {
//          result.setAlpha(
//              FloatValue
//                  .newBuilder()
//                  .setValue(((float) alpha) / denominator)
//                  .build());
//        }
//        return resultBuilder.build();
//      }
//      // ...
//
// Example (iOS / Obj-C):
//
//      // ...
//      static UIColor* fromProto(Color* protocolor) {
//         float red = [protocolor red];
//         float green = [protocolor green];
//         float blue = [protocolor blue];
//         FloatValue* alpha_wrapper = [protocolor alpha];
//         float alpha = 1.0;
//         if (alpha_wrapper != nil) {
//           alpha = [alpha_wrapper value];
//         }
//         return [UIColor colorWithRed:red green:green blue:blue alpha:alpha];
//      }
//
//      static Color* toProto(UIColor* color) {
//          CGFloat red, green, blue, alpha;
//          if (![color getRed:&red green:&green blue:&blue alpha:&alpha]) {
//            return nil;
//          }
//          Color* result = [Color alloc] init];
//          [result setRed:red];
//          [result setGreen:green];
//          [result setBlue:blue];
//          if (alpha <= 0.9999) {
//            [result setAlpha:floatWrapperWithValue(alpha)];
//          }
//          [result autorelease];
//          return result;
//     }
//     // ...
//
//  Example (JavaScript):
//
//     // ...
//
//     var protoToCssColor = function(rgb_color) {
//        var redFrac = rgb_color.red || 0.0;
//        var greenFrac = rgb_color.green || 0.0;
//        var blueFrac = rgb_color.blue || 0.0;
//        var red = Math.floor(redFrac * 255);
//        var green = Math.floor(greenFrac * 255);
//        var blue = Math.floor(blueFrac * 255);
//
//        if (!('alpha' in rgb_color)) {
//           return rgbToCssColor_(red, green, blue);
//        }
//
//        var alphaFrac = rgb_color.alpha.value || 0.0;
//        var rgbParams = [red, green, blue].join(',');
//        return ['rgba(', rgbParams, ',', alphaFrac, ')'].join('');
//     };
//
//     var rgbToCssColor_ = function(red, green, blue) {
//       var rgbNumber = new Number((red << 16) | (green << 8) | blue);
//       var hexString = rgbNumber.toString(16);
//       var missingZeros = 6 - hexString.length;
//       var resultBuilder = ['#'];
//       for (var i = 0; i < missingZeros; i++) {
//          resultBuilder.push('0');
//       }
//       resultBuilder.push(hexString);
//       return resultBuilder.join('');
//     };
//
//     // ...
type Color struct {
	// The amount of red in the color as a value in the interval [0, 1].
	Red float32 `protobuf:"fixed32,1,opt,name=red" json:"red,omitempty"`
	// The amount of green in the color as a value in the interval [0, 1].
	Green float32 `protobuf:"fixed32,2,opt,name=green" json:"green,omitempty"`
	// The amount of blue in the color as a value in the interval [0, 1].
	Blue float32 `protobuf:"fixed32,3,opt,name=blue" json:"blue,omitempty"`
	// The fraction of this color that should be applied to the pixel. That is,
	// the final pixel color is defined by the equation:
	//
	//   pixel color = alpha * (this color) + (1.0 - alpha) * (background color)
	//
	// This means that a value of 1.0 corresponds to a solid color, whereas
	// a value of 0.0 corresponds to a completely transparent color. This
	// uses a wrapper message rather than a simple float scalar so that it is
	// possible to distinguish between a default value and the value being unset.
	// If omitted, this color object is to be rendered as a solid color
	// (as if the alpha value had been explicitly given with a value of 1.0).
	Alpha *google_protobuf.FloatValue `protobuf:"bytes,4,opt,name=alpha" json:"alpha,omitempty"`
}

func (m *Color) Reset()                    { *m = Color{} }
func (m *Color) String() string            { return proto.CompactTextString(m) }
func (*Color) ProtoMessage()               {}
func (*Color) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Color) GetAlpha() *google_protobuf.FloatValue {
	if m != nil {
		return m.Alpha
	}
	return nil
}

func init() {
	proto.RegisterType((*Color)(nil), "google.type.Color")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/type/color/color.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x85, 0xd9, 0xdd, 0x5b, 0x8b, 0xb9, 0x42, 0x09, 0x16, 0x8b, 0x36, 0x87, 0x20, 0x5c, 0x95,
	0x41, 0xad, 0x04, 0xab, 0x13, 0xb4, 0x5d, 0x0e, 0xb1, 0xcf, 0xae, 0xe3, 0xdc, 0x41, 0x6e, 0x27,
	0xe4, 0x12, 0xc5, 0xbf, 0xe3, 0x2f, 0x95, 0x24, 0x2b, 0xda, 0x84, 0x97, 0x99, 0xf7, 0x3e, 0xde,
	0xc0, 0x03, 0x8b, 0xb0, 0x25, 0xcd, 0x62, 0xcd, 0xc4, 0x5a, 0x3c, 0x23, 0xd3, 0xe4, 0xbc, 0x04,
	0xc1, 0xb2, 0x32, 0x6e, 0x7f, 0xc4, 0xf0, 0xe5, 0x08, 0x47, 0xb1, 0xe2, 0xcb, 0xab, 0xb3, 0x43,
	0x2d, 0xe7, 0x74, 0x5a, 0x5f, 0xdc, 0xf3, 0x3e, 0xec, 0xe2, 0xa0, 0x47, 0x39, 0x60, 0xc1, 0x61,
	0x76, 0x0d, 0xf1, 0x1d, 0x5d, 0x72, 0x1c, 0xf1, 0xd3, 0x1b, 0xe7, 0xc8, 0xff, 0x89, 0xc2, 0xb9,
	0xfa, 0x80, 0xf6, 0x31, 0x61, 0xd5, 0x19, 0x34, 0x9e, 0xde, 0xba, 0x6a, 0x55, 0xad, 0xeb, 0x6d,
	0x92, 0xea, 0x1c, 0x5a, 0xf6, 0x44, 0x53, 0x57, 0xe7, 0x59, 0xf9, 0x28, 0x05, 0x8b, 0xc1, 0x46,
	0xea, 0x9a, 0x3c, 0xcc, 0x5a, 0xdd, 0x40, 0x6b, 0xac, 0xdb, 0x99, 0x6e, 0xb1, 0xaa, 0xd6, 0xcb,
	0xdb, 0x4b, 0x3d, 0x97, 0xfb, 0x2d, 0xa1, 0x9f, 0xac, 0x98, 0xf0, 0x6a, 0x6c, 0xa4, 0x6d, 0x71,
	0x6e, 0xae, 0xe1, 0x74, 0x94, 0x83, 0xfe, 0x77, 0xc5, 0x06, 0x72, 0x91, 0x3e, 0x65, 0xfa, 0xea,
	0xbb, 0x6e, 0x9e, 0x5f, 0xfa, 0xe1, 0x24, 0x23, 0xee, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6a,
	0xe1, 0xe0, 0x7d, 0x2d, 0x01, 0x00, 0x00,
}
