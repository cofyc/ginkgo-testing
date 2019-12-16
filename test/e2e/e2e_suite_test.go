package e2e

import (
	goflag "flag"
	"fmt"
	"math/rand"
	"os"
	"path"
	"testing"
	"time"

	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/reporters"
	"github.com/onsi/gomega"
	flag "github.com/spf13/pflag"
	"k8s.io/klog"
)

type E2EConfigType struct {
	ReportDir string
}

var E2EConfig = E2EConfigType{}

// A go testing function as main entrypoint of ginkgo test suite.
func TestE2E(t *testing.T) {
	// RegisterFailHandler connects Ginkgo to Gomega. When a matcher fails
	// the fail handler passed into RegisterFailHandler is called.
	gomega.RegisterFailHandler(ginkgo.Fail)
	// Run specs with a test suite name.
	// This will fail the testing.T if any of specs fail.
	// ginkgo.RunSpecs(t, "E2E Suite")
	// Run specs with default and custom reporters.
	var r []ginkgo.Reporter
	if err := os.MkdirAll(E2EConfig.ReportDir, 0755); err != nil {
		klog.Errorf("Failed creating report directory: %v", err)
	} else {
		r = append(r, reporters.NewJUnitReporter(path.Join(E2EConfig.ReportDir, fmt.Sprintf("junit_%02d.xml", config.GinkgoConfig.ParallelNode))))
	}
	ginkgo.RunSpecsWithDefaultAndCustomReporters(t, "E2E Suite", r)
}

func handleFlags() {
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	// change ginkgo default configurations
	// Turn on verbose by default to get spec names
	config.DefaultReporterConfig.Verbose = true
	// Turn on EmitSpecProgress to get spec progress (especially on interrupt)
	config.GinkgoConfig.EmitSpecProgress = true
	// Randomize specs as well as suites
	config.GinkgoConfig.RandomizeAllSpecs = true
}

func TestMain(m *testing.M) {
	handleFlags()
	rand.Seed(time.Now().UnixNano())
	os.Exit(m.Run())
}

// If you want to run some code synchronously, you can use Synchronized version of BeforeSuite/AfterSuite
// var _ = ginkgo.SynchronizedBeforeSuite(func() []byte {
// // run on node 1
// return nil
// }, func(data []byte) {
// // run on all nodes
// return
// })
// var _ = SynchronizedAfterSuite(func() {
// // run on all nodes
// }, func(data []byte) {
// // run on node 1
// })

var _ = ginkgo.BeforeSuite(func() {
	ginkgo.By("run BeforeSuite")
})

var _ = ginkgo.AfterSuite(func() {
	ginkgo.By("run AfterSuite")
})

// You can pass gingko.Done to run with a timeout.
// var _ = ginkgo.AfterSuite(func(done ginkgo.Done) {
// time.Sleep(time.Second)
// ginkgo.By("run AfterSuite")
// close(done)
// }, 2)
