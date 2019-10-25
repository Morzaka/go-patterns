package Specification

import "testing"

func TestSpecification(t *testing.T) {
	overDue := NewOverDueSpecification()
	noticeSent := NewNoticeSentSpecification()
	inCollection := NewInCollectionSpecification()

	sendToCollecton := overDue.And(noticeSent).And(inCollection.Not())

	invoice := Invoice{
		Day:    31,    // >= 30
		Notice: 4,     // >= 3
		IsSent: false, // false
	}

	// true!
	result := sendToCollecton.IsSatisfiedBy(invoice)

	if !result {
		t.Errorf("Expect result to equal %v, but %v.\n", false, true)
	}
}
