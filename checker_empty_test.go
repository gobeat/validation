package validation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type sampleEmptyInputTest1 struct {
	Name string `validate:"empty=false"`
}

var _ = Describe("EmptyChecker", func() {
	It("should return error code emptyValueError", func() {
		err := New().Validate(sampleEmptyInputTest1{})
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("Name", emptyValueError)))
	})

	It("should return nil", func() {
		err := New().Validate(sampleEmptyInputTest1{
			Name: "John",
		})
		Expect(err).To(BeNil())
	})
})
