package validation

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type sampleLengthInput1 struct {
	Password string `validate:"length=8"`
}

var _ = Describe("LengthChecker", func() {
	It("should return error code lengthValueError", func() {
		err := New().Validate(sampleLengthInput1{"aaaa"})
		Expect(err).NotTo(BeNil())
		Expect(err).To(Equal(makeError("Password", fmt.Sprintf(lengthValueError, 8))))
	})

	It("should return nil", func() {
		err := New().Validate(sampleLengthInput1{"aaaaaaaa"})
		Expect(err).To(BeNil())
	})
})
