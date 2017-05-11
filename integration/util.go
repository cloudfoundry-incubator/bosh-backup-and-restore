package integration

import (
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf-experimental/cf-webmock/mockhttp"
)

type helpText struct {
	output []byte
}

func (h helpText) outputString() string {
	return string(h.output)
}

func ShowsTheBackupHelpText(helpText *helpText) {
	Expect(helpText.outputString()).To(ContainSubstring("--target"))
	Expect(helpText.outputString()).To(ContainSubstring("Target BOSH Director URL"))

	Expect(helpText.outputString()).To(ContainSubstring("--username"))
	Expect(helpText.outputString()).To(ContainSubstring("BOSH Director username"))

	Expect(helpText.outputString()).To(ContainSubstring("--password"))
	Expect(helpText.outputString()).To(ContainSubstring("BOSH Director password"))

	Expect(helpText.outputString()).To(ContainSubstring("--deployment"))
	Expect(helpText.outputString()).To(ContainSubstring("Name of BOSH deployment"))
}

func ShowsTheHelpText(helpText *helpText) {
	Expect(helpText.outputString()).To(ContainSubstring(`SUBCOMMANDS:
   backup
   restore
   pre-backup-check`))

	Expect(helpText.outputString()).To(ContainSubstring(`USAGE:
	bbr command [arguments...] [subcommand]`))
}

func mockDirectorWith(director *mockhttp.Server, info mockhttp.MockedResponseBuilder, vmsResponse []mockhttp.MockedResponseBuilder, sshResponse []mockhttp.MockedResponseBuilder, downloadManifestResponse []mockhttp.MockedResponseBuilder, cleanupResponse []mockhttp.MockedResponseBuilder) {
	director.VerifyAndMock(AppendBuilders(
		[]mockhttp.MockedResponseBuilder{info},
		vmsResponse,
		sshResponse,
		downloadManifestResponse,
		cleanupResponse,
	)...)

}
