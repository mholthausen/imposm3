package mapping

import (
	"goposm/element"
	"testing"
)

func TestFilterNodes(t *testing.T) {
	var tags element.Tags

	// test name only
	tags = make(element.Tags)
	tags["name"] = "foo"
	if PointTags.Filter(tags) != false {
		t.Fatal("Filter result not false")
	}
	if len(tags) != 0 {
		t.Fatal("Filter result not empty")
	}

	// test name + unmapped tags
	tags = make(element.Tags)
	tags["name"] = "foo"
	tags["boring"] = "true"

	if PointTags.Filter(tags) != false {
		t.Fatal("Filter result not false")
	}
	if len(tags) != 0 {
		t.Fatal("Filter result not empty")
	}

	// test __any__
	tags = make(element.Tags)
	tags["population"] = "0"
	tags["name"] = "foo"
	tags["boring"] = "true"

	if PointTags.Filter(tags) != true {
		t.Fatal("Filter result true", tags)
	}
	if len(tags) != 2 && tags["population"] == "0" && tags["name"] == "foo" {
		t.Fatal("Filter result not expected", tags)
	}

}

func BenchmarkFilterNodes(b *testing.B) {
	var tags element.Tags

	for i := 0; i < b.N; i++ {
		// test __any__
		tags = make(element.Tags)
		tags["population"] = "0"
		tags["name"] = "foo"
		tags["boring"] = "true"

		if PointTags.Filter(tags) != true {
			b.Fatal("Filter result true", tags)
		}
		if len(tags) != 2 && tags["population"] == "0" && tags["name"] == "foo" {
			b.Fatal("Filter result not expected", tags)
		}
	}

}