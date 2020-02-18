package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("", func() {

	It("should show error when try create triangle", func() {

		_, err := NewTriangle(1, 2, 3)
		Expect(err).ToNot(BeNil())

		_, err = NewTriangle(1, 2, 4)
		Expect(err).ToNot(BeNil())

		_, err = NewTriangle(12, 15, 30)
		Expect(err).ToNot(BeNil())
	})

	It("should test if triangle is escaleno", func() {
		t, err := NewTriangle(3, 4, 5)

		Expect(err).To(BeNil())
		Expect(t.isEscaleno()).To(BeTrue())
	})

	It("should test if triangle is isosceles", func() {
		t, err := NewTriangle(3, 3, 4)

		Expect(err).To(BeNil())
		Expect(t.isIsosceles()).To(BeTrue())

		t, err = NewTriangle(3, 4, 3)

		Expect(err).To(BeNil())
		Expect(t.isIsosceles()).To(BeTrue())

		t, err = NewTriangle(4, 3, 3)

		Expect(err).To(BeNil())
		Expect(t.isIsosceles()).To(BeTrue())
	})

	It("should test if triangle is equilatero", func() {
		t, err := NewTriangle(3, 3, 3)

		Expect(err).To(BeNil())
		Expect(t.isEquilatero()).To(BeTrue())
	})
})
