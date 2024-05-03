package main_test

import (
	"context"
	"os/exec"
	"testing"
	"time"

	protos "github.com/nikimanoledaki/calculator-microservice/protos/calculator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"google.golang.org/grpc"
)

func TestClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Client Suite")
}

var (
	err           error
	serverBinary  string
	clientBinary  string
	serverCommand *exec.Cmd
	serverSession *gexec.Session
	conn          *grpc.ClientConn
	ctx           context.Context
	cancel        context.CancelFunc
	clt           protos.CalculatorClient
)

var _ = BeforeSuite(func() {
	serverBinary, err = gexec.Build("github.com/nikimanoledaki/calculator-microservice/cmd/server", "-mod=vendor")
	Expect(err).NotTo(HaveOccurred())

	clientBinary, err = gexec.Build("github.com/nikimanoledaki/calculator-microservice/cmd/client", "-mod=vendor")
	Expect(err).NotTo(HaveOccurred())

	serverCommand = exec.Command(serverBinary)
	serverSession, err = gexec.Start(serverCommand, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	conn, err = grpc.Dial("localhost:9092", grpc.WithInsecure(), grpc.WithBlock())
	Expect(err).NotTo(HaveOccurred())

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	clt = protos.NewCalculatorClient(conn)
})

var _ = AfterSuite(func() {
	cancel()
	conn.Close()
	gexec.Terminate()
	gexec.CleanupBuildArtifacts()
})
