package stempl

import "testing"

func TestTemplate(t *testing.T) {
	var (
		result string
		err    error
	)

	result, err = Format("", map[string]string{})
	if err != nil {
		t.Fail()
	}
	if result != "" {
		t.Fail()
	}

	result, err = Format("something strange happening", map[string]string{})
	if err != nil {
		t.Fail()
	}
	if result != "something strange happening" {
		t.Fail()
	}

	result, err = Format(
		"something {strange} happening",
		map[string]string{
			"strange": "good",
		},
	)
	if err != nil {
		t.Error(err)
	}
	if result != "something good happening" {
		t.Error(result)
	}

	result, err = Format(
		"{a}{b}{c}",
		map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	)
	if err != nil {
		t.Error(err)
	}
	if result != "123" {
		t.Error(result)
	}

	result, err = Format(
		"{a} {b} {c}",
		map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	)
	if err != nil {
		t.Error(err)
	}
	if result != "1 2 3" {
		t.Error(result)
	}

	result, err = Format(
		"{a}",
		map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	)
	if err != nil {
		t.Error(err)
	}
	if result != "1" {
		t.Error(result)
	}

	result, err = Format(
		"{",
		map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	)
	if err == nil {
		t.Fail()
	}

	result, err = Format(
		"}",
		map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	)
	if err == nil {
		t.Fail()
	}

	result, err = Format(
		"{a}{b}{",
		map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	)
	if err == nil {
		t.Fail()
	}

	result, err = Format(
		"{a}{b}}",
		map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	)
	if err == nil {
		t.Fail()
	}

	result, err = Format(
		"{a}{b}{}",
		map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	)
	if err == nil {
		t.Fail()
	}

	result, err = Format(
		"{a{b}}",
		map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	)
	if err == nil {
		t.Fail()
	}

	result, err = Format(
		"{a{b}}",
		map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	)
	if err == nil {
		t.Fail()
	}

	result, err = Format(
		"    {    ",
		map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	)
	if err == nil {
		t.Fail()
	}

	result, err = Format(
		"    }    ",
		map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	)
	if err == nil {
		t.Fail()
	}

	result, err = Format(
		"{a}}{",
		map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	)
	if err == nil {
		t.Fail()
	}

	result, err = Format(
		"{{",
		map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	)
	if err != nil {
		t.Error(err)
	}
	if result != "{" {
		t.Error(result)
	}

	result, err = Format(
		"}}",
		map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	)
	if err != nil {
		t.Error(err)
	}
	if result != "}" {
		t.Error(result)
	}

	result, err = Format(
		"{{}}",
		map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	)
	if err != nil {
		t.Error(err)
	}
	if result != "{}" {
		t.Error(result)
	}

	result, err = Format(
		"{a}{{{b}}}{{{c}{{",
		map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	)
	if err != nil {
		t.Error(err)
	}
	if result != "1{2}{3{" {
		t.Error(result)
	}

	result, err = Format(
		"{{{}}}",
		map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	)
	if err == nil {
		t.Fail()
	}

	result, err = Format(
		"    }}}}}    ",
		map[string]string{
			"a": "1",
			"b": "2",
			"c": "3",
		},
	)
	if err == nil {
		t.Fail()
	}
}
