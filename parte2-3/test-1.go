package main

func test1 {
	cases := []struct{
		name, function, key, value string
		expectedS string
		expectedI int
	}{
		{name: "add element", function: "add", key: "xxx", value: "zzz", expected: 0},
		{name: "add other element", function: "add", key: "zzz", value: "zzz", expected: 0},
		{name: "not add element", function: "add", key: "xxx", value: "zzz", expected: -1},
		{name: "add element", function: "add", key: "xxx", value: "zzz", expected: 0},
	}
	for c, _ := range cases {
		t.Run(c.name, func(tt *testing.T){
			if res := Sum(c.function, c.b); res != c.expected {
				tt.Errorf("expected %d but got %d", c.expected, res)
			}
		})
	}
}