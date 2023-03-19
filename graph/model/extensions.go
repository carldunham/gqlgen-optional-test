package model

import "fmt"

func (input *FooInput) String() string {
	if input == nil {
		return "foo is null"
	}

	bar := "foo.bar was not provided"
	if input.Bar.Defined {
		bar = input.Bar.Value.String()
	}

	if !input.Text.Defined {
		return "foo.text was not provided, " + bar
	}

	if input.Text.Value == nil {
		return "foo.text is null, " + bar
	}

	return fmt.Sprintf("foo.text is %q, %s", *input.Text.Value, bar)
}

func (input *BarInput) String() string {
	if input == nil {
		return "foo.bar is null"
	}

	baz := "foo.bar.baz was not provided"
	if input.Baz.Defined {
		baz = input.Baz.Value.String()
	}

	if !input.Count.Defined {
		return "foo.bar.count was not provided, " + baz
	}

	if input.Count.Value == nil {
		return "foo.bar.count is null, " + baz
	}

	return fmt.Sprintf("foo.bar.count is %d, %s", *input.Count.Value, baz)
}

func (input *BazInput) String() string {
	if input == nil {
		return "foo.bar.baz is null"
	}
	if !input.IsSomething.Defined {
		return "foo.bar.baz.isSomething was not provided"
	}

	if input.IsSomething.Value == nil {
		return "foo.bar.baz.isSomething is null"
	}

	return fmt.Sprintf("foo.bar.baz.isSomething is %t", *input.IsSomething.Value)
}
