package tree

import (
	"strings"
	"testing"
)

func TestTree_Del(t *testing.T) {
	for _, tc := range []struct {
		name    string
		pattern string
		in      map[string]interface{}
		out     string
	}{
		{
			name:    "unknown",
			pattern: "abc",
			in: map[string]interface{}{
				"supu": 42,
				"tupu": false,
			},
			out: `
├── supu	42
└── tupu	false
`,
		},
		{
			name:    "plain",
			pattern: "supu",
			in: map[string]interface{}{
				"supu": 42,
				"tupu": false,
			},
			out: `
└── tupu	false
`,
		},
		{
			name:    "element_in_struct",
			pattern: "internal.supu",
			in: map[string]interface{}{
				"internal": map[string]interface{}{
					"supu": 42,
					"tupu": false,
				},
				"tupu": false,
			},
			out: `
├── internal
│   └── tupu	false
└── tupu	false
`,
		},
		{
			name:    "element_in_struct_with_wildcard",
			pattern: "a.*.supu",
			in: map[string]interface{}{
				"a": map[string]interface{}{
					"first": map[string]interface{}{
						"supu": 42,
						"tupu": false,
					},
					"last": map[string]interface{}{
						"supu": 42,
						"tupu": false,
					},
				},
				"tupu": false,
			},
			out: `
├── a
│   ├── first
│   │   └── tupu	false
│   └── last
│       └── tupu	false
└── tupu	false
`,
		},
		{
			name:    "struct",
			pattern: "internal",
			in: map[string]interface{}{
				"internal": map[string]interface{}{
					"supu": 42,
					"tupu": false,
				},
				"tupu": false,
			},
			out: `
└── tupu	false
`,
		},
		{
			name:    "element_in_substruct",
			pattern: "internal.internal.supu",
			in: map[string]interface{}{
				"internal": map[string]interface{}{
					"supu": 42,
					"tupu": false,
					"internal": map[string]interface{}{
						"supu": 42,
						"tupu": false,
					},
				},
				"tupu": false,
			},
			out: `
├── internal
│   ├── internal
│   │   └── tupu	false
│   ├── supu	42
│   └── tupu	false
└── tupu	false
`,
		},
		{
			name:    "similar_names",
			pattern: "a.a.a",
			in: map[string]interface{}{
				"a": map[string]interface{}{
					"a": map[string]interface{}{
						"a": map[string]interface{}{
							"a": 1,
						},
						"aa": 1,
					},
					"aa": 1,
				},
				"tupu": false,
			},
			out: `
├── a
│   ├── a
│   │   └── aa	1
│   └── aa	1
└── tupu	false
`,
		},
		{
			name:    "collection_element_atributes",
			pattern: "a.*.a",
			in: map[string]interface{}{
				"a": []interface{}{
					map[string]interface{}{
						"a": map[string]interface{}{
							"a": map[string]interface{}{
								"a": 1,
							},
							"aa": 1,
						},
						"aa": 1,
					},
					map[string]interface{}{
						"a":  42,
						"aa": 1,
					},
				},
				"tupu": false,
			},
			out: `
├── a []
│   ├── 0
│   │   └── aa	1
│   └── 1
│       └── aa	1
└── tupu	false
`,
		},
		{
			name:    "nested_collection_element_atributes",
			pattern: "a.*.b.*.c",
			in: map[string]interface{}{
				"a": []interface{}{
					map[string]interface{}{
						"b": []interface{}{
							map[string]interface{}{
								"c": map[string]interface{}{
									"a": 1,
								},
								"aa": 1,
							},
							map[string]interface{}{
								"c": map[string]interface{}{
									"a": 2,
								},
								"aa": 1,
							},
						},
						"aa": 1,
					},
					map[string]interface{}{
						"b": []interface{}{
							map[string]interface{}{
								"c": map[string]interface{}{
									"a": 1,
								},
								"aa": 1,
							},
						},
						"aa": 1,
					},
				},
				"tupu": false,
			},
			out: `
├── a []
│   ├── 0
│   │   ├── aa	1
│   │   └── b []
│   │       ├── 0
│   │       │   └── aa	1
│   │       └── 1
│   │           └── aa	1
│   └── 1
│       ├── aa	1
│       └── b []
│           └── 0
│               └── aa	1
└── tupu	false
`,
		},
		{
			name:    "large_collection_element_atributes",
			pattern: "a.*.a",
			in: map[string]interface{}{
				"a": []interface{}{
					map[string]interface{}{
						"a":  1,
						"aa": 1,
					},
					map[string]interface{}{
						"a":  2,
						"aa": 1,
					},
					map[string]interface{}{
						"a":  1,
						"aa": 1,
					},
					map[string]interface{}{
						"a":  2,
						"aa": 1,
					},
					map[string]interface{}{
						"a":  1,
						"aa": 1,
					},
					map[string]interface{}{
						"a":  2,
						"aa": 1,
					},
					map[string]interface{}{
						"a":  1,
						"aa": 1,
					},
					map[string]interface{}{
						"a":  2,
						"aa": 1,
					},
					map[string]interface{}{
						"a":  1,
						"aa": 1,
					},
					map[string]interface{}{
						"a":  2,
						"aa": 1,
					},
					map[string]interface{}{
						"a":  1,
						"aa": 1,
					},
					map[string]interface{}{
						"a":  2,
						"aa": 1,
					},
				},
				"tupu": false,
			},
			out: `
├── a []
│   ├── 0
│   │   └── aa	1
│   ├── 1
│   │   └── aa	1
│   ├── 2
│   │   └── aa	1
│   ├── 3
│   │   └── aa	1
│   ├── 4
│   │   └── aa	1
│   ├── 5
│   │   └── aa	1
│   ├── 6
│   │   └── aa	1
│   ├── 7
│   │   └── aa	1
│   ├── 8
│   │   └── aa	1
│   ├── 9
│   │   └── aa	1
│   ├── 10
│   │   └── aa	1
│   └── 11
│       └── aa	1
└── tupu	false
`,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			res, _ := New(tc.in)
			res.Sort()

			res.Del(strings.Split(tc.pattern, "."))
			tree := res.String()
			if tree != tc.out {
				t.Errorf("unexpected result (%s):'%s'\n'%s'", tc.pattern, tree, tc.out)
			}
		})
	}
}

func TestTree_Get(t *testing.T) {
	in := map[string]interface{}{
		"a": []interface{}{
			map[string]interface{}{
				"a":  1,
				"aa": 1,
			},
			map[string]interface{}{
				"a":  2,
				"aa": 1,
			},
			map[string]interface{}{
				"a":  1,
				"aa": 1,
			},
			map[string]interface{}{
				"a":  2,
				"aa": 1,
			},
			map[string]interface{}{
				"a":  1,
				"aa": 1,
			},
			map[string]interface{}{
				"a":  2,
				"aa": 1,
			},
			map[string]interface{}{
				"a":  1,
				"aa": 1,
			},
			map[string]interface{}{
				"a":  2,
				"aa": 1,
			},
			map[string]interface{}{
				"a":  1,
				"aa": 1,
			},
			map[string]interface{}{
				"a":  2,
				"aa": 1,
			},
			map[string]interface{}{
				"a":  1,
				"aa": 1,
			},
			map[string]interface{}{
				"a":  2,
				"aa": 1,
			},
		},
		"b": map[string]interface{}{
			"c": []interface{}{1, 2, 3, 4},
		},
		"tupu": false,
	}
	tree, _ := New(in)
	v := tree.Get([]string{"a", "1", "a"})
	i, ok := v.(int)
	if !ok {
		t.Errorf("unexpected result type: %v", v)
		return
	}
	if i != 2 {
		t.Errorf("unexpected result %d", i)
	}
}

func TestTree_Move(t *testing.T) {
	for _, tc := range []struct {
		name string
		src  string
		dst  string
		in   map[string]interface{}
		out  string
	}{
		{
			name: "plain",
			src:  "a",
			dst:  "b",
			in:   map[string]interface{}{"a": 42},
			out: `
└── b	42
`,
		},
		{
			name: "in_struct",
			src:  "b.a",
			dst:  "b.c",
			in: map[string]interface{}{
				"a": 1,
				"b": map[string]interface{}{"a": 42},
			},
			out: `
├── a	1
└── b
    └── c	42
`,
		},
		{
			name: "from_struct",
			src:  "b.a",
			dst:  "c",
			in: map[string]interface{}{
				"a": 1,
				"b": map[string]interface{}{"a": 42},
			},
			out: `
├── a	1
├── b	<nil>
└── c	42
`,
		},
		{
			name: "from_struct_with_wildcard",
			src:  "b.*.c",
			dst:  "b.*.x",
			in: map[string]interface{}{
				"c": 42,
				"b": map[string]interface{}{
					"first": map[string]interface{}{"c": map[string]interface{}{"d": 42}},
					"last":  map[string]interface{}{"m": 42, "c": map[string]interface{}{"d": 42}},
				},
			},
			out: `
├── b
│   ├── first
│   │   └── x
│   │       └── d	42
│   └── last
│       ├── m	42
│       └── x
│           └── d	42
└── c	42
`,
		},
		{
			name: "from_collection",
			src:  "b.*.c",
			dst:  "b.*.x",
			in: map[string]interface{}{
				"a": 42,
				"b": []interface{}{
					map[string]interface{}{"c": 42},
					map[string]interface{}{"c": map[string]interface{}{"d": 42}},
				},
			},
			out: `
├── a	42
└── b []
    ├── 0
    │   └── x	42
    └── 1
        └── x
            └── d	42
`,
		},
		{
			name: "from_struct_nested",
			src:  "b.b",
			dst:  "c",
			in: map[string]interface{}{
				"a": 42,
				"b": map[string]interface{}{
					"a":  42,
					"bb": true,
					"b":  map[string]interface{}{"a": 42},
				},
			},
			out: `
├── a	42
├── b
│   ├── a	42
│   └── bb	true
└── c
    └── a	42
`,
		},
		{
			name: "collection",
			src:  "a.*.b",
			dst:  "a.*.c",
			in: map[string]interface{}{
				"a": []interface{}{
					map[string]interface{}{
						"b": []interface{}{
							map[string]interface{}{
								"c": map[string]interface{}{
									"a": 1,
								},
								"aa": 1,
							},
							map[string]interface{}{
								"c": map[string]interface{}{
									"a": 2,
								},
								"aa": 1,
							},
						},
						"aa": 1,
					},
					map[string]interface{}{
						"b": []interface{}{
							map[string]interface{}{
								"c": map[string]interface{}{
									"a": 1,
								},
								"aa": 1,
							},
						},
						"aa": 1,
					},
				},
				"tupu": false,
			},
			out: `
├── a []
│   ├── 0
│   │   ├── aa	1
│   │   └── c []
│   │       ├── 0
│   │       │   ├── aa	1
│   │       │   └── c
│   │       │       └── a	1
│   │       └── 1
│   │           ├── aa	1
│   │           └── c
│   │               └── a	2
│   └── 1
│       ├── aa	1
│       └── c []
│           └── 0
│               ├── aa	1
│               └── c
│                   └── a	1
└── tupu	false
`,
		},
		{
			name: "recursive_collection",
			src:  "a.*.b.*.c",
			dst:  "a.*.b.*.x",
			in: map[string]interface{}{
				"a": []interface{}{
					map[string]interface{}{
						"b": []interface{}{
							map[string]interface{}{
								"c": map[string]interface{}{
									"a": 1,
								},
								"aa": 1,
							},
							map[string]interface{}{
								"c": map[string]interface{}{
									"a": 2,
								},
								"aa": 1,
							},
						},
						"aa": 1,
					},
					map[string]interface{}{
						"b": []interface{}{
							map[string]interface{}{
								"c": map[string]interface{}{
									"a": 1,
								},
								"aa": 1,
							},
						},
						"aa": 1,
					},
				},
				"tupu": false,
			},
			out: `
├── a []
│   ├── 0
│   │   ├── aa	1
│   │   └── b []
│   │       ├── 0
│   │       │   ├── aa	1
│   │       │   └── x
│   │       │       └── a	1
│   │       └── 1
│   │           ├── aa	1
│   │           └── x
│   │               └── a	2
│   └── 1
│       ├── aa	1
│       └── b []
│           └── 0
│               ├── aa	1
│               └── x
│                   └── a	1
└── tupu	false
`,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			res, _ := New(tc.in)

			res.Move(strings.Split(tc.src, "."), strings.Split(tc.dst, "."))

			res.Sort()

			if tree := res.String(); tree != tc.out {
				t.Errorf("unexpected result (%s -> %s):\n%s\n%s", tc.src, tc.dst, tree, tc.out)
			}
		})
	}
}
