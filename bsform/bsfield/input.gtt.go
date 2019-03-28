package bsfield

import (
	"context"
	"github.com/go-qbit/template/filter"
	"io"
)

var (
	s7a12fbba1b5990f27101f240bf388f6b = []byte{0xA, 0x20, 0x20, 0x20, 0x20, 0x3C, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6C, 0x61, 0x73, 0x73, 0x3D, 0x22, 0x69, 0x6E, 0x76, 0x61, 0x6C, 0x69, 0x64, 0x2D, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6B, 0x22, 0x3E}
	s1bb4b15041ac78119756a84a2c3cf170 = []byte{0x20}
	sdc90de5788da5ffb1dd5d027fb9a379e = []byte{0x20, 0x69, 0x73, 0x2D, 0x69, 0x6E, 0x76, 0x61, 0x6C, 0x69, 0x64}
	s6f330b8de72a2a4f357ef550e7b98982 = []byte{0x20, 0x70, 0x6C, 0x61, 0x63, 0x65, 0x68, 0x6F, 0x6C, 0x64, 0x65, 0x72, 0x3D, 0x22}
	s3e36bcbcfb880362c7fcd1491fe765a3 = []byte{0x20, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64}
	se15eeac5bfaf4b5366665dbd8d7a2026 = []byte{0x3C, 0x2F, 0x64, 0x69, 0x76, 0x3E}
	s7236c66d53b133f280b8e23ea4d53ac0 = []byte{0x3C, 0x2F, 0x6C, 0x61, 0x62, 0x65, 0x6C, 0x3E}
	s1f57138b33ff7c29dcb169fd12c61d25 = []byte{0x3C, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6C, 0x61, 0x73, 0x73, 0x3D, 0x22, 0x66, 0x6F, 0x72, 0x6D, 0x2D, 0x67, 0x72, 0x6F, 0x75, 0x70, 0x22, 0x3E}
	sd33ded6ff87b1e23f6f9b39ee04340be = []byte{0x3C, 0x69, 0x6E, 0x70, 0x75, 0x74, 0x20, 0x6E, 0x61, 0x6D, 0x65, 0x3D, 0x22}
	s4a5be3b9b0c7d6249d57148667400bff = []byte{0x3C, 0x6C, 0x61, 0x62, 0x65, 0x6C, 0x20, 0x66, 0x6F, 0x72, 0x3D, 0x22, 0x69, 0x6E, 0x70, 0x75, 0x74}
	se55838261bb3f3b3539374aca85fc56f = []byte{0x3E}
	sdcb532f29fa03783aa5364df3aade0c8 = []byte{0x22}
	sb77f264bd4d37561b8494d1563165e19 = []byte{0x22, 0x20}
	sadf37d8651db7061f740a476a2ca4442 = []byte{0x22, 0x3E}
	s05f7b31cf5fa41d6bf1086b25c730a3b = []byte{0x63, 0x6C, 0x61, 0x73, 0x73, 0x3D, 0x22, 0x66, 0x6F, 0x72, 0x6D, 0x2D, 0x63, 0x6F, 0x6E, 0x74, 0x72, 0x6F, 0x6C}
	s6cb72921bf440474747eef8b88c9b981 = []byte{0x69, 0x64, 0x3D, 0x22, 0x69, 0x6E, 0x70, 0x75, 0x74}
	s7da6ef98e6a29375a83fbc0f48cd877a = []byte{0x74, 0x79, 0x70, 0x65, 0x3D, 0x22}
	s60a7e41def7454a2b0b9796e025bf088 = []byte{0x76, 0x61, 0x6C, 0x75, 0x65, 0x3D, 0x22}
)

func ProcessInput(ctx context.Context, w io.Writer, f *Input) {
	if f.Type == "hidden" {
		w.Write(sd33ded6ff87b1e23f6f9b39ee04340be)
		io.WriteString(w, filter.Filterhtml(f.Name))
		w.Write(sb77f264bd4d37561b8494d1563165e19)
		w.Write(s7da6ef98e6a29375a83fbc0f48cd877a)
		io.WriteString(w, filter.Filterhtml(f.Type))
		w.Write(sb77f264bd4d37561b8494d1563165e19)
		w.Write(s60a7e41def7454a2b0b9796e025bf088)
		io.WriteString(w, filter.Filterhtml(f.GetStringValue()))
		w.Write(sdcb532f29fa03783aa5364df3aade0c8)
		w.Write(se55838261bb3f3b3539374aca85fc56f)
	} else {
		w.Write(s1f57138b33ff7c29dcb169fd12c61d25)
		w.Write(s4a5be3b9b0c7d6249d57148667400bff)
		io.WriteString(w, filter.Filterhtml(f.Name))
		w.Write(sadf37d8651db7061f740a476a2ca4442)
		io.WriteString(w, filter.Filterhtml(f.Label))
		w.Write(s7236c66d53b133f280b8e23ea4d53ac0)
		w.Write(sd33ded6ff87b1e23f6f9b39ee04340be)
		io.WriteString(w, filter.Filterhtml(f.Name))
		w.Write(sb77f264bd4d37561b8494d1563165e19)
		w.Write(s7da6ef98e6a29375a83fbc0f48cd877a)
		io.WriteString(w, filter.Filterhtml(f.Type))
		w.Write(sb77f264bd4d37561b8494d1563165e19)
		w.Write(s60a7e41def7454a2b0b9796e025bf088)
		io.WriteString(w, filter.Filterhtml(f.GetStringValue()))
		w.Write(sb77f264bd4d37561b8494d1563165e19)
		w.Write(s05f7b31cf5fa41d6bf1086b25c730a3b)
		if f.err != nil {
			w.Write(sdc90de5788da5ffb1dd5d027fb9a379e)
		}
		for _, c := range f.Class {
			w.Write(s1bb4b15041ac78119756a84a2c3cf170)
			io.WriteString(w, filter.Filterhtml(c))
		}
		w.Write(sb77f264bd4d37561b8494d1563165e19)
		w.Write(s6cb72921bf440474747eef8b88c9b981)
		io.WriteString(w, filter.Filterhtml(f.Name))
		w.Write(sdcb532f29fa03783aa5364df3aade0c8)
		if f.Required {
			w.Write(s3e36bcbcfb880362c7fcd1491fe765a3)
		}
		if f.Placeholder != "" {
			w.Write(s6f330b8de72a2a4f357ef550e7b98982)
			io.WriteString(w, filter.Filterhtml(f.Placeholder))
			w.Write(sdcb532f29fa03783aa5364df3aade0c8)
		}
		w.Write(se55838261bb3f3b3539374aca85fc56f)
		if f.err != nil {
			w.Write(s7a12fbba1b5990f27101f240bf388f6b)
			io.WriteString(w, filter.Filterhtml(f.err.PublicError()))
			w.Write(se15eeac5bfaf4b5366665dbd8d7a2026)
		}
		w.Write(se15eeac5bfaf4b5366665dbd8d7a2026)
	}
}
