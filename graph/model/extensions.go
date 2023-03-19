package model

import "fmt"

func (input *FooInput) String() string {
	bar := input.Bar.String()

	if !input.Text.Defined {
		return "text was not provided, " + bar
	}

	if input.Text.Value == nil {
		return "text is null, " + bar
	}

	return fmt.Sprintf("text is %q, %s", *input.Text.Value, bar)
}

func (input *BarInput) String() string {
	baz := input.Baz.String()

	if !input.Count.Defined {
		return "count was not provided, " + baz
	}

	if input.Count.Value == nil {
		return "count is null, " + baz
	}

	return fmt.Sprintf("count is %d, %s", *input.Count.Value, baz)
}

func (input *BazInput) String() string {
	if !input.IsSomething.Defined {
		return "isSomething was not provided"
	}

	if input.IsSomething.Value == nil {
		return "isSomething is null"
	}

	return fmt.Sprintf("isSomething is %t", *input.IsSomething.Value)
}
