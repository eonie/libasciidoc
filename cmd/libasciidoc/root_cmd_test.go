package main_test

import (
	"bytes"
	"io/ioutil"

	main "github.com/bytesparadise/libasciidoc/cmd/libasciidoc"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("root cmd", func() {
	RegisterFailHandler(Fail)

	It("render with STDOUT output", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		root.SetArgs([]string{"-o", "-", "test/test.adoc"})
		// when
		err := root.Execute()
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(buf.String()).ToNot(BeEmpty())
	})

	It("render with file output", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		root.SetArgs([]string{"test/test.adoc"})
		// when
		err := root.Execute()
		// then
		Expect(err).ToNot(HaveOccurred())
		content, err := ioutil.ReadFile("test/test.html")
		Expect(err).ToNot(HaveOccurred())
		Expect(content).ToNot(BeEmpty())
	})

	It("fail to parse bad log level", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		root.SetArgs([]string{"--log", "debug1", "-s", "test/test.adoc"})
		// when
		err := root.Execute()
		// then
		GinkgoT().Logf("command output: %v", buf.String())
		Expect(err).To(HaveOccurred())
	})

	It("render without header/footer", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		root.SetArgs([]string{"-s", "-o", "-", "test/test.adoc"})
		// when
		err := root.Execute()
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(buf.String()).ToNot(BeEmpty())
		Expect(buf.String()).ToNot(ContainSubstring(`<div id="footer">`))
	})

	It("render multiple files", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		root.SetArgs([]string{"-s", "test/admonition.adoc", "test/test.adoc"})
		// when
		err := root.Execute()
		// then
		Expect(err).ToNot(HaveOccurred())
	})

	It("when rendering multiple files, return last error", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		root.SetArgs([]string{"-s", "test/doesnotexist.adoc", "test/test.adoc"})
		// when
		err := root.Execute()
		// then
		Expect(err).To(HaveOccurred())
	})

	It("show help when executed with no arg", func() {
		// given
		root := main.NewRootCmd()
		buf := new(bytes.Buffer)
		root.SetOutput(buf)
		root.SetArgs([]string{})
		// when
		err := root.Execute()
		// then
		Expect(err).ToNot(HaveOccurred())
	})
})
